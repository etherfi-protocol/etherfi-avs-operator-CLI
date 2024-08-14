package lagrangezk

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/eigenlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
)

// API handle for all core lagrangeZK functionality
type API struct {
	Client *ethclient.Client

	ServiceManager            *ZKMRServiceManager
	ServiceManagerAddress     common.Address
	StakeRegistry             *ZKMRStakeRegistry
	StakeRegistryAddress      common.Address
	AvsOperatorManagerAddress common.Address

	EigenlayerAPI *eigenlayer.API
}

// Info that node operator must supply to the ether.fi admin for registration
type RegistrationInfo struct {
	OperatorID     int64
	LagrangePubkey ECDSAPubkey
}

type ECDSAPubkey struct {
	X *big.Int
	Y *big.Int
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	serviceManager, _ := NewZKMRServiceManager(cfg.LagrangeZKMRServiceManagerAddress, rpcClient)
	stakeRegistry, _ := NewZKMRStakeRegistry(cfg.LagrangeZKMRStakeRegistryAddress, rpcClient)

	return &API{
		Client:                    rpcClient,
		ServiceManager:            serviceManager,
		ServiceManagerAddress:     cfg.LagrangeZKMRServiceManagerAddress,
		StakeRegistry:             stakeRegistry,
		StakeRegistryAddress:      cfg.LagrangeZKMRStakeRegistryAddress,
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		EigenlayerAPI:             eigenlayer.New(cfg, rpcClient),
	}
}

// PrepareRegistration aggregates all required info from the node operator that
// the ether.fi admin will need to register them to the Witness Chain AVS
func (a *API) PrepareRegistration(operator *etherfi.Operator, pubkeyX *big.Int, pubkeyY *big.Int) error {

	ri := RegistrationInfo{
		OperatorID:     operator.ID,
		LagrangePubkey: ECDSAPubkey{X: pubkeyX, Y: pubkeyY},
	}

	return utils.ExportJSON("lagrangeZK-prepare-registration", operator.ID, ri)
}

func (a *API) RegisterOperator(operator *etherfi.Operator, info RegistrationInfo, signerKey *ecdsa.PrivateKey) error {

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
	pubkey := PublicKey{
		X: info.LagrangePubkey.X,
		Y: info.LagrangePubkey.Y,
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	stakeRegistryABI, err := ZKMRStakeRegistryMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	calldata, err := stakeRegistryABI.Pack("registerOperator", pubkey, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, calldata, a.StakeRegistryAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}
	fmt.Println("stakeRegistry:", a.StakeRegistryAddress)
	fmt.Println("manager:", a.AvsOperatorManagerAddress)
	fmt.Println("operator:", operator.Address)
	fmt.Println("serviceManager:", a.ServiceManagerAddress)

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("lagrangeZK-register-operator-%d", operator.ID))
	return utils.ExportJSON("lagrangeZK-register-gnosis", operator.ID, batch)
}
