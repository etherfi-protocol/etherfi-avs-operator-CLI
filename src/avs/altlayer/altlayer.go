package altlayer

import (
	"crypto/ecdsa"
	"fmt"

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

// API handle for all core AltLayer functionality
type API struct {
	Client *ethclient.Client

	RegistryCoordinatorAddress common.Address
	RegistryCoordinator        *RegistryCoordinator
	ServiceManagerAddress      common.Address
	AvsOperatorManagerAddress  common.Address

	EigenlayerAPI *eigenlayer.API
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	registryCoordinator, _ := NewRegistryCoordinator(cfg.AltLayerRegistryCoordinatorAddress, rpcClient)

	return &API{
		Client:                     rpcClient,
		RegistryCoordinator:        registryCoordinator,
		RegistryCoordinatorAddress: cfg.AltLayerRegistryCoordinatorAddress,
		ServiceManagerAddress:      cfg.AltLayerServiceManagerAddress,
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
	return utils.ExportJSON("altlayer-prepare-registration", operator.ID, ri)
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
	sigParams := ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sigWithSaltAndExpiry.Signature,
		Salt:      sigWithSaltAndExpiry.Salt,
		Expiry:    sigWithSaltAndExpiry.Expiry,
	}
	pubkeyParams := IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: BN254G1Point(info.BLSPubkeyRegistrationParams.Signature),
		PubkeyG1:                    BN254G1Point(info.BLSPubkeyRegistrationParams.G1),
		PubkeyG2:                    BN254G2Point(info.BLSPubkeyRegistrationParams.G2),
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	coordinatorABI, err := RegistryCoordinatorMetaData.GetAbi()
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
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("altlayer-register-operator-%d", operator.ID))
	return utils.ExportJSON("altlayer-register-gnosis", operator.ID, batch)
}
