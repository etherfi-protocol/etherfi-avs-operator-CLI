package eigenlayer

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/etherfi-protocol/eigenlayer-rewards-proofs/pkg/claimgen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/eigenlayer-rewards-proofs/pkg/proofDataFetcher/httpProofDataFetcher"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/types"
	"github.com/fatih/color"
)

type API struct {
	Client *ethclient.Client

	AvsDirectory              *AvsDirectory
	AvsDirectoryAddress       common.Address
	AvsOperatorManagerAddress common.Address

	RewardsCoordinator        *RewardsCoordinator
	RewardsCoordinatorAddress common.Address

	DelegationManager        *DelegationManager
	DelegationManagerAddress common.Address
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	avsDirectory, _ := NewAvsDirectory(cfg.AvsDirectoryAddress, rpcClient)
	delegationManager, _ := NewDelegationManager(cfg.DelegationManagerAddress, rpcClient)
	rewardsCoordinator, _ := NewRewardsCoordinator(cfg.EigenlayerRewardsCoordinatorAddress, rpcClient)

	return &API{
		Client:                    rpcClient,
		AvsDirectory:              avsDirectory,
		AvsDirectoryAddress:       cfg.AvsDirectoryAddress,
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		RewardsCoordinator:        rewardsCoordinator,
		RewardsCoordinatorAddress: cfg.EigenlayerRewardsCoordinatorAddress,
		DelegationManager:         delegationManager,
		DelegationManagerAddress:  cfg.DelegationManagerAddress,
	}
}

func (a *API) GenerateRegistrationDigest(operator *etherfi.Operator, salt [32]byte, expiry *big.Int, serviceManager common.Address) ([]byte, error) {

	hash, err := a.AvsDirectory.CalculateOperatorAVSRegistrationDigestHash(nil, operator.Address, serviceManager, [32]byte(salt), expiry)
	if err != nil {
		return nil, fmt.Errorf("requesting hash: %w", err)
	}

	return hash[:], nil
}

func (a *API) GenerateAndSignRegistrationDigest(operator *etherfi.Operator, serviceManager common.Address, privKey *ecdsa.PrivateKey) (*types.SignatureWithSaltAndExpiry, error) {

	// look up configured ecdsaSigner for this operator
	ecdsaSigner, err := operator.Contract.EcdsaSigner(nil)
	if err != nil {
		return nil, fmt.Errorf("fetching configured ecdsaSigner: %w", err)
	}

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("generating random salt: %w", err)
	}
	// TODO: revisit what we want the expiration to be (currently 7 days)
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 7).Unix())

	hash, err := a.GenerateRegistrationDigest(operator, [32]byte(salt), expiry, serviceManager)
	if err != nil {
		return nil, fmt.Errorf("generating registration digest: %w", err)
	}

	// ensure the provided key matches the ecdsaSigner the operator contract is expecting
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("private key is unexpected type: %w", err)
	}
	keyAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if keyAddress != ecdsaSigner {
		color.HiYellow("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		color.HiYellow("WARNING: configured signing key does not match the current ecdsaSigner for operator")
		color.HiYellow("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Printf("signer: %s, ecdsaSigner: %s\n\n", keyAddress, ecdsaSigner)
	}

	// sign the digest
	signed, err := utils.SignDigestECDSA(hash[:], privKey)
	if err != nil {
		return nil, fmt.Errorf("signing digest: %w", err)
	}

	return &types.SignatureWithSaltAndExpiry{
		Signature: signed,
		Salt:      [32]byte(salt),
		Expiry:    expiry,
	}, nil
}

func (a *API) ClaimAvsOperatorRewards(operators []*etherfi.Operator, rewardsRecipient common.Address, rewardTokens []common.Address) error {

	// find the latest claimable root
	latestClaimableRoot, err := a.RewardsCoordinator.GetCurrentClaimableDistributionRoot(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get latest claimable root: %w", err)
	}
	rootIndex, err := a.RewardsCoordinator.GetRootIndexFromHash(&bind.CallOpts{}, latestClaimableRoot.Root)
	if err != nil {
		return fmt.Errorf("failed to get root index from hash: %w", err)
	}

	// TODO: hardcoded to mainnet ethereum for now
	// fetch merkle tree using eigenSDK
	df := httpProofDataFetcher.NewHttpProofDataFetcher(
		"https://eigenlabs-rewards-mainnet-ethereum.s3.amazonaws.com", // ProofStoreBaseURL
		"mainnet",  // Environment
		"ethereum", // Network
		http.DefaultClient,
	)

	claimDate := time.Unix(int64(latestClaimableRoot.RewardsCalculationEndTimestamp), 0).UTC().Format(time.DateOnly)
	proofData, err := df.FetchClaimAmountsForDate(context.Background(), claimDate)
	if err != nil {
		return fmt.Errorf("failed to fetch claim amounts for date: %w", err)
	}

	// generate proofs
	cg := claimgen.NewClaimgen(proofData.Distribution)
	rewardsCoordinatorABI, _ := RewardsCoordinatorMetaData.GetAbi()
	batch := gnosis.GnosisBatch{
		Version: "1.0",
		ChainId: "1",
		Meta:    gnosis.GnosisMetadata{Name: "operator-rewards-claim-batch"},
	}

	for _, operator := range operators {
		_, claim, err := cg.GenerateClaimProofForEarner(operator.Address, rewardTokens, rootIndex)
		if err != nil {
			if strings.HasPrefix(err.Error(), "earner index not found") {
				continue
			}
			return fmt.Errorf("failed to generate claim proof for earner: %w", err)
		}

		// manually pack tx data since we are forwarding the call via the etherfiNodesManager
		claimCalldata, err := rewardsCoordinatorABI.Pack("processClaim", claim, rewardsRecipient)
		if err != nil {
			return fmt.Errorf("packing processClaim: %w", err)
		}

		// wrap the inner call to be forwarded via AvsOperatorManager
		adminCall, err := utils.PackForwardCallForAdmin(operator.ID, claimCalldata, a.RewardsCoordinatorAddress)
		if err != nil {
			return fmt.Errorf("wrapping call for admin: %w", err)
		}

		// output in gnosis compatible format
		batch.AddTransaction(gnosis.SubTransaction{
			Target: a.AvsOperatorManagerAddress,
			Value:  big.NewInt(0),
			Data:   adminCall,
		})
	}

	return utils.ExportJSON("avs-operator-bulk-reward-claim", operators[0].ID, batch)
}
