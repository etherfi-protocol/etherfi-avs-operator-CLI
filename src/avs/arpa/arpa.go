package arpa

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
)

// API handle for all core arpa functionality
type API struct {
	Client *ethclient.Client

	NodeRegistry              *NodeRegistry
	NodeRegistryAddress       common.Address
	ServiceManagerAddress     common.Address
	AvsOperatorManagerAddress common.Address

	EigenlayerAPI *eigenlayer.API
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	nodeRegistry, _ := NewNodeRegistry(cfg.ARPANodeRegistryAddress, rpcClient)

	return &API{
		Client: rpcClient,

		NodeRegistry:              nodeRegistry,
		NodeRegistryAddress:       cfg.ARPANodeRegistryAddress,
		ServiceManagerAddress:     cfg.ARPAServiceManagerAddress,
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		EigenlayerAPI:             eigenlayer.New(cfg, rpcClient),
	}
}

// Info that node operator must supply to the ether.fi admin for registration
type RegistrationInfo struct {
	OperatorID int64
	AvsSigner  common.Address
}

// PrepareRegistration aggregates all required info from the node operator that
// the ether.fi admin will need to register them to the AVS
func (a *API) PrepareRegistration(operator *etherfi.Operator, avsSigner common.Address) error {

	ri := RegistrationInfo{
		OperatorID: operator.ID,
		AvsSigner:  avsSigner,
	}

	return utils.ExportJSON("arpa-prepare-registration", operator.ID, ri)
}

// RegisterOperator is used by the ether.fi admin to register a node operator's AvsOperator
// contract with the AVS.
func (a *API) RegisterOperator(operator *etherfi.Operator, info RegistrationInfo, signingKey *ecdsa.PrivateKey) error {

	// generate and sign registration hash with admin ecdsa key
	signature, err := a.EigenlayerAPI.GenerateAndSignRegistrationDigest(operator, a.ServiceManagerAddress, signingKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}

	// convert to types expected by contract call
	sigWithSaltAndExpiry := ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature.Signature,
		Salt:      signature.Salt,
		Expiry:    signature.Expiry,
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	nodeRegistryABI, err := NodeRegistryMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	input, err := nodeRegistryABI.Pack("nodeRegister", , operator.Address, sigWithSaltAndExpiry)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, input, a.NodeRegistryAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("arpa-register-%d", operator.ID))
	return utils.ExportJSON("arpa-register-gnosis", operator.ID, batch)
}
