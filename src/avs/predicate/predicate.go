package predicate

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/eigenlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/types"
)

// API handle for all core Automata functionality
type API struct {
	Client *ethclient.Client

	ServiceManager            *ServiceManager
	ServiceManagerAddress     common.Address
	AvsOperatorManagerAddress common.Address

	EigenlayerAPI *eigenlayer.API
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	serviceManager, _ := NewServiceManager(cfg.PredicateServiceManagerAddress, rpcClient)

	return &API{
		Client:                    rpcClient,
		ServiceManager:            serviceManager,
		ServiceManagerAddress:     cfg.PredicateServiceManagerAddress,
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		EigenlayerAPI:             eigenlayer.New(cfg, rpcClient),
	}
}

// Info that node operator must supply to the ether.fi admin for registration
type RegistrationInfo struct {
	OperatorID                  int64
	AvsSigner                   common.Address
	BLSPubkeyRegistrationParams *types.BLSPubkeyRegistrationParams
}

func (a *API) PrepareRegistration(operator *etherfi.Operator, avsSigner common.Address) error {

	ri := RegistrationInfo{
		OperatorID: operator.ID,
		AvsSigner:  avsSigner,
	}
	return utils.ExportJSON("predicate-prepare-registration", operator.ID, ri)
}

func (a *API) RegisterOperator(operator *etherfi.Operator, info RegistrationInfo, signerKey *ecdsa.PrivateKey) error {

	// generate and sign registration hash to be signed by admin ecdsa key
	sigWithSaltAndExpiry, err := a.EigenlayerAPI.GenerateAndSignRegistrationDigest(operator, a.ServiceManagerAddress, signerKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}
	sigParams := ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sigWithSaltAndExpiry.Signature,
		Salt:      sigWithSaltAndExpiry.Salt,
		Expiry:    sigWithSaltAndExpiry.Expiry,
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	serviceManagerABI, err := ServiceManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	calldata, err := serviceManagerABI.Pack("registerOperatorToAVS", info.AvsSigner, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, calldata, a.ServiceManagerAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("predicate-register-operator-%d", operator.ID))
	return utils.ExportJSON("predicate-register-gnosis", operator.ID, batch)
}
