package lagrangesc

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	lagrangeutils "github.com/Lagrange-Labs/client-cli/utils"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/eigenlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
)

// API handle for all core Brevis functionality
type API struct {
	Client *ethclient.Client

	LagrangeServiceAddress    common.Address
	LagrangeService           *LagrangeService
	LagrangeCommitteeAddress  common.Address
	LagrangeCommittee         *LagrangeCommittee
	AvsOperatorManagerAddress common.Address

	EigenlayerAPI *eigenlayer.API
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	lagrangeService, _ := NewLagrangeService(cfg.LagrangeServiceAddress, rpcClient)
	lagrangeCommittee, _ := NewLagrangeCommittee(cfg.LagrangeCommitteeAddress, rpcClient)

	return &API{
		Client:                    rpcClient,
		LagrangeService:           lagrangeService,
		LagrangeServiceAddress:    cfg.LagrangeServiceAddress,
		LagrangeCommittee:         lagrangeCommittee,
		LagrangeCommitteeAddress:  cfg.LagrangeCommitteeAddress,
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		EigenlayerAPI:             eigenlayer.New(cfg, rpcClient),
	}
}

type RegistrationInfo struct {
	OperatorID      int64
	BLSKeyWithProof *IBLSKeyCheckerBLSKeyWithProof
	SignerAddress   common.Address
}

func (a *API) PrepareRegistration(operator *etherfi.Operator, signerAddr common.Address, blsKey *bls.KeyPair) error {

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("generating random salt: %w", err)
	}
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 7).Unix())

	keyWithProofDigest, err := a.LagrangeCommittee.CalculateKeyWithProofHash(nil, operator.Address, [32]byte(salt), expiry)
	if err != nil {
		return fmt.Errorf("calculating committee proof hash: %w", err)
	}

	// Lagrange library wants the private keys as a hex string
	blsPrivKey := blsKey.PrivKey.Text(2)
	blsPrivKeys := []string{blsPrivKey}

	g1Pubkeys, g2Pubkey, sig, err := lagrangeutils.GenerateBLSSignature(keyWithProofDigest[:], blsPrivKeys...)
	if err != nil {
		fmt.Errorf("generating aggregate BLS signature: %w", err)
	}

	ri := RegistrationInfo{
		OperatorID:    operator.ID,
		SignerAddress: signerAddr,
		BLSKeyWithProof: &IBLSKeyCheckerBLSKeyWithProof{
			BlsG1PublicKeys: g1Pubkeys,
			AggG2PublicKey:  g2Pubkey,
			Signature:       sig,
			Salt:            [32]byte(salt),
			Expiry:          expiry,
		},
	}
	return utils.ExportJSON("lagrangeSC-prepare-registration", operator.ID, ri)
}

func (a *API) RegisterOperator(operator *etherfi.Operator, info RegistrationInfo, signerKey *ecdsa.PrivateKey) error {

	// generate and sign registration hash to be signed by admin ecdsa key
	sigWithSaltAndExpiry, err := a.EigenlayerAPI.GenerateAndSignRegistrationDigest(operator, a.LagrangeServiceAddress, signerKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}

	sigParams := ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sigWithSaltAndExpiry.Signature,
		Salt:      sigWithSaltAndExpiry.Salt,
		Expiry:    sigWithSaltAndExpiry.Expiry,
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	lagrangeServiceABI, err := LagrangeServiceMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	calldata, err := lagrangeServiceABI.Pack("register", info.SignerAddress, info.BLSKeyWithProof, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, calldata, a.LagrangeServiceAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("lagrangeSC-register-operator-%d", operator.ID))
	return utils.ExportJSON("lagrangeSC-register-gnosis", operator.ID, batch)
}
