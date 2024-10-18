package unifi

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

// API handle for all core UniFi functionality
type API struct {
	Client *ethclient.Client

	ServiceManagerAddress     common.Address
	AvsOperatorManagerAddress common.Address

	EigenlayerAPI *eigenlayer.API
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	return &API{
		Client: rpcClient,

		ServiceManagerAddress:     cfg.UniFiAvsManagerAddress, // AvsManager serves as the Service Manager
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		EigenlayerAPI:             eigenlayer.New(cfg, rpcClient),
	}
}

// RegisterOperator is used by the ether.fi admin to register a node operator's AvsOperator
// contract with the AVS.
func (a *API) RegisterOperator(operator *etherfi.Operator, signingKey *ecdsa.PrivateKey) error {

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
	avsManagerABI, err := AvsManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	input, err := avsManagerABI.Pack("registerOperator", sigWithSaltAndExpiry)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, input, a.ServiceManagerAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("unifi-register-%d", operator.ID))
	return utils.ExportJSON("unifi-register-gnosis", operator.ID, batch)
}
