package eoracle

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/eigenlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils/signer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/types"
)

// API handle for all core eOracle functionality
type API struct {
	Client *ethclient.Client

	RegistryCoordinator        *RegistryCoordinator
	RegistryCoordinatorAddress common.Address
	ServiceManager             *ServiceManager
	ServiceManagerAddress      common.Address
	AvsOperatorManagerAddress  common.Address

	EigenlayerAPI *eigenlayer.API
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	registryCoordinator, _ := NewRegistryCoordinator(cfg.EOracleRegistryCoordinatorAddress, rpcClient)
	serviceManager, _ := NewServiceManager(cfg.EOracleServiceManagerAddress, rpcClient)

	return &API{
		Client:                     rpcClient,
		RegistryCoordinator:        registryCoordinator,
		RegistryCoordinatorAddress: cfg.EOracleRegistryCoordinatorAddress,
		ServiceManager:             serviceManager,
		ServiceManagerAddress:      cfg.EOracleServiceManagerAddress,
		AvsOperatorManagerAddress:  cfg.AvsOperatorManagerAddress,
		EigenlayerAPI:              eigenlayer.New(cfg, rpcClient),
	}
}

// Info that node operator must supply to the ether.fi admin for registration
type RegistrationInfo struct {
	OperatorID                  int64
	BLSPubkeyRegistrationParams *types.BLSPubkeyRegistrationParams
	AliasAddress                common.Address
}

func (a *API) PrepareRegistration(operator *etherfi.Operator, blsKey *bls.KeyPair, alias common.Address) error {

	// compute hash to sign with bls key
	// the hash is converted to a G1 point on the curve before it is returned
	g1Point, err := a.RegistryCoordinator.PubkeyRegistrationMessageHash(nil, operator.Address)
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
		OperatorID:                  operator.ID,
		BLSPubkeyRegistrationParams: signedParams,
		AliasAddress:                alias,
	}
	return utils.ExportJSON("eoracle-prepare-registration", operator.ID, ri)
}

func (a *API) RegisterOperator(operator *etherfi.Operator, info RegistrationInfo, signerKey *ecdsa.PrivateKey, quorums []byte) error {

	// generate and sign registration hash to be signed by admin ecdsa key
	sigWithSaltAndExpiry, err := a.EigenlayerAPI.GenerateAndSignRegistrationDigest(operator, a.ServiceManagerAddress, signerKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}

	// convert to types expected by contract call
	sigParams := ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sigWithSaltAndExpiry.Signature,
		Salt:      sigWithSaltAndExpiry.Salt,
		Expiry:    sigWithSaltAndExpiry.Expiry,
	}
	pubkeyParams := IEOBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: BN254G1Point(info.BLSPubkeyRegistrationParams.Signature),
		ChainValidatorSignature:     BN254G1Point{X: big.NewInt(0), Y: big.NewInt(0)}, // not currently used by protocol
		PubkeyG1:                    BN254G1Point(info.BLSPubkeyRegistrationParams.G1),
		PubkeyG2:                    BN254G2Point(info.BLSPubkeyRegistrationParams.G2),
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	registryCoordinatorABI, err := RegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	calldata, err := registryCoordinatorABI.Pack("registerOperator", quorums, pubkeyParams, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, calldata, a.RegistryCoordinatorAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("witness-chain-register-watchtower-%d", operator.ID))
	return utils.ExportJSON("eoracle-register-gnosis", operator.ID, batch)
}
