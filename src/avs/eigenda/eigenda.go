package eigenda

import (
	"crypto/ecdsa"
	"fmt"

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

// Info that node operator must supply to the ether.fi admin for registration
type RegistrationInfo struct {
	OperatorID                  int64
	Socket                      string
	Quorums                     []int64
	BLSPubkeyRegistrationParams *types.BLSPubkeyRegistrationParams
}

func (a *API) PrepareRegistration(operator *etherfi.Operator, blsKey *bls.KeyPair, socket string, quorums []int64) error {

	// compute hash to sign with bls key
	// the hash is converted to a G1 point on the curve before it is returned
	g1Point, err := a.RegistryCoordinator.PubkeyRegistrationMessageHash(nil, operator.Address)
	if err != nil {
		return fmt.Errorf("fetching pubkeyRegistrationMessageHash: %w", err)
	}

	// TODO: double check that the new setup is legal versus this old setup
	/*
		xElem := fp.NewElement(0)
		xElem.SetBigInt(g1Point.X)
		yElem := fp.NewElement(0)
		yElem.SetBigInt(g1Point.Y)
	*/

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
		Socket:                      socket,
		Quorums:                     quorums,
		BLSPubkeyRegistrationParams: signedParams,
	}
	return utils.ExportJSON("eigenda-prepare-registration", operator.ID, ri)
}

func (a *API) RegisterOperator(operator *etherfi.Operator, info RegistrationInfo, signerKey *ecdsa.PrivateKey) error {

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

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, calldata, a.RegistryCoordinatorAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("eigenda-register-operator-%d", operator.ID))
	return utils.ExportJSON("eigenda-register-gnosis", operator.ID, batch)
}
