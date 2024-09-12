package eigenda

import (
	"context"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/contracts/safeglobal"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/config"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/eigenlayer"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/gnosis"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/types"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/utils"

	// need to alias because eigenlayer has a package name that doesn't match the filepath
	registryCoordinator "github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

// API handle for all core EigenDA functionality
type EigenDA struct {
	Client                     *ethclient.Client
	RegistryCoordinatorAddress common.Address
	RegistryCoordinator        *registryCoordinator.ContractRegistryCoordinator
	ServiceManagerAddress      common.Address

	EigenLayer *eigenlayer.EigenLayer
	SafeGlobal *gnosis.SafeGlobal
}

func New(cfg config.Config, rpcClient *ethclient.Client) *EigenDA {

	registryCoordinator, _ := registryCoordinator.NewContractRegistryCoordinator(cfg.EigenDARegistryCoordinatorAddress, rpcClient)

	return &EigenDA{
		Client:                     rpcClient,
		RegistryCoordinator:        registryCoordinator,
		RegistryCoordinatorAddress: cfg.EigenDARegistryCoordinatorAddress,
		ServiceManagerAddress:      cfg.EigenDAServiceManagerAddress,
		EigenLayer:                 eigenlayer.New(cfg, rpcClient),
		SafeGlobal:                 gnosis.New(cfg, rpcClient),
	}
}

// Info that node operator must supply to the mantle admin for registration
type RegistrationInfo struct {
	Operator                    common.Address
	Socket                      string
	Quorums                     []int64
	BLSPubkeyRegistrationParams *types.BLSPubkeyRegistrationParams
}

func (a *EigenDA) PrepareRegistration(operator common.Address, blsKey *bls.KeyPair, socket string, quorums []int64) error {
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

	g1Sig := blsKey.SignHashedToCurveMessage(g1MsgToSign)

	signedParams := new(types.BLSPubkeyRegistrationParams)
	signedParams.Load(blsKey.GetPubKeyG1().G1Affine, blsKey.GetPubKeyG2().G2Affine, g1Sig.G1Affine)

	ri := RegistrationInfo{
		Operator:                    operator,
		Socket:                      socket,
		Quorums:                     quorums,
		BLSPubkeyRegistrationParams: signedParams,
	}
	return utils.ExportJSON("eigenda-prepare-registration", operator, ri)
}

func (a *EigenDA) RegisterOperator(operator common.Address, info RegistrationInfo) error {

	// generate and sign registration hash to be signed by admin ecdsa key
	sigWithSaltAndExpiry, err := a.EigenLayer.GenerateAndSignRegistrationDigest(operator, a.ServiceManagerAddress)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}

	// convert to types expected by contract call
	quorums := make([]byte, len(info.Quorums))
	for i, v := range info.Quorums {
		quorums[i] = byte(v)
	}

	sigParams := registryCoordinator.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: nil, //signature is signed in safe wallet
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

	signMessageLibABI, err := safeglobal.SignMessageLibMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	signCalldata, err := signMessageLibABI.Pack("signMessage", sigWithSaltAndExpiry.Signature)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	chainID, _ := a.Client.ChainID(context.Background())

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(calldata, chainID, fmt.Sprintf("eigenda-register-operator-%s", operator))

	batch.AddTransaction(gnosis.SubTransaction{
		Target: a.SafeGlobal.SignMessageLibAddress,
		Value:  big.NewInt(0),
		Data:   signCalldata,
	})

	batch.AddTransaction(gnosis.SubTransaction{
		Target: a.RegistryCoordinatorAddress,
		Value:  big.NewInt(0),
		Data:   calldata,
	})

	return utils.ExportJSON("eigenda-register-gnosis", operator, batch)
}
