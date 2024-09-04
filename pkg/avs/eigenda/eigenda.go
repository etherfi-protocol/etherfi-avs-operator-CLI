package eigenda

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/a41-official/mantle-eigenlayer-operator-cli/pkg/config"
	"github.com/a41-official/mantle-eigenlayer-operator-cli/pkg/eigenlayer"
	"github.com/a41-official/mantle-eigenlayer-operator-cli/pkg/gnosis"
	"github.com/a41-official/mantle-eigenlayer-operator-cli/pkg/types"
	"github.com/a41-official/mantle-eigenlayer-operator-cli/pkg/utils"
	"github.com/a41-official/mantle-eigenlayer-operator-cli/pkg/utils/signer"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// need to alias because eigenlayer has a package name that doesn't match the filepath
	registryCoordinator "github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

// API handle for all core EigenDA functionality
type API struct {
	Client *ethclient.Client

	RegistryCoordinatorAddress common.Address
	RegistryCoordinator        *registryCoordinator.ContractRegistryCoordinator
	ServiceManagerAddress      common.Address
	AvsOperatorManagerAddress  common.Address

	EigenlayerAPI *eigenlayer.API
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	registryCoordinator, _ := registryCoordinator.NewContractRegistryCoordinator(cfg.EigenDARegistryCoordinatorAddress, rpcClient)

	return &API{
		Client:                     rpcClient,
		RegistryCoordinator:        registryCoordinator,
		RegistryCoordinatorAddress: cfg.EigenDARegistryCoordinatorAddress,
		ServiceManagerAddress:      cfg.EigenDAServiceManagerAddress,
		AvsOperatorManagerAddress:  cfg.AvsOperatorManagerAddress,
		EigenlayerAPI:              eigenlayer.New(cfg, rpcClient),
	}
}

// Info that node operator must supply to the mantle admin for registration
type RegistrationInfo struct {
	Operator                    common.Address
	Socket                      string
	Quorums                     []int64
	BLSPubkeyRegistrationParams *types.BLSPubkeyRegistrationParams
}

func (a *API) PrepareRegistration(operator common.Address, blsKey *bls.KeyPair, socket string, quorums []int64) error {
	// compute hash to sign with bls key
	// the hash is converted to a G1 point on the curve before it is returned
	g1Point, err := a.RegistryCoordinator.PubkeyRegistrationMessageHash(nil, operator)
	if err != nil {
		return fmt.Errorf("fetching pubkeyRegistrationMessageHash: %w", err)
	}

	// map from contract type to type expected by signing algorithm
	g1MsgToSign := &bn254.G1Affine{
		X: *new(fp.Element).SetBigInt(g1Point.X),
		Y: *new(fp.Element).SetBigInt(g1Point.Y),
	}

	// sign the registration hash (G1 Point) with the bls key to prove ownership of key
	blsSigner := signer.NewBLSSigner(blsKey)
	g1Sig, err := blsSigner.Sign(g1MsgToSign)
	if err != nil {
		return fmt.Errorf("signing pubkey registration hash: %w", err)
	}
	signedParams := new(types.BLSPubkeyRegistrationParams)
	signedParams.Load(blsKey.GetPubKeyG1().G1Affine, blsKey.GetPubKeyG2().G2Affine, g1Sig)

	// sanity check on the bls signature
	isValid, err := blsSigner.Verify(g1MsgToSign, g1Sig)
	if !isValid || err != nil {
		return fmt.Errorf("failed to verify g1 signature: %w", err)
	}

	ri := RegistrationInfo{
		Operator:                    operator,
		Socket:                      socket,
		Quorums:                     quorums,
		BLSPubkeyRegistrationParams: signedParams,
	}
	return utils.ExportJSON("eigenda-prepare-registration", operator, ri)
}

func (a *API) RegisterOperator(operator common.Address, info RegistrationInfo, signerKey *ecdsa.PrivateKey) error {

	// generate and sign registration hash to be signed by admin ecdsa key
	sigWithSaltAndExpiry, err := a.EigenlayerAPI.GenerateAndSignRegistrationDigest(operator, a.ServiceManagerAddress, signerKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}

	// convert to types expected by contract call
	quorums := make([]byte, len(info.Quorums))
	for i, v := range info.Quorums {
		quorums[i] = byte(v)
	}
	sigParams := registryCoordinator.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sigWithSaltAndExpiry.Signature,
		Salt:      sigWithSaltAndExpiry.Salt,
		Expiry:    sigWithSaltAndExpiry.Expiry,
	}
	pubkeyParams := registryCoordinator.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: registryCoordinator.BN254G1Point(info.BLSPubkeyRegistrationParams.Signature),
		PubkeyG1:                    registryCoordinator.BN254G1Point(info.BLSPubkeyRegistrationParams.G1),
		PubkeyG2:                    registryCoordinator.BN254G2Point(info.BLSPubkeyRegistrationParams.G2),
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	coordinatorABI, err := registryCoordinator.ContractRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	calldata, err := coordinatorABI.Pack("registerOperator", quorums, info.Socket, pubkeyParams, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(calldata, a.AvsOperatorManagerAddress, fmt.Sprintf("eigenda-register-operator-%s", operator))
	return utils.ExportJSON("eigenda-register-gnosis", operator, batch)
}
