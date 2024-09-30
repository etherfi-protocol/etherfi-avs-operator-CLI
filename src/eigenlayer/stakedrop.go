package eigenlayer

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
)

func RegisterClaimAddressForOperator(operator *etherfi.Operator, claimAddress common.Address) error {

	// must send tx with this exact message
	msg := "I agree to have my operator EIGEN allocation be claimed from the registered wallet in this transaction."

	// manually pack tx data since we are submitting via gnosis instead of directly
	operatorRegistryABI, err := OperatorRegistryMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	calldata, err := operatorRegistryABI.Pack("Register", claimAddress, msg)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, calldata, config.Mainnet.EigenlayerOperatorRegistryClaimAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, config.Mainnet.AvsOperatorManagerAddress, fmt.Sprintf("stakedrop-claim-address-%d", operator.ID))
	return utils.ExportJSON("stakedrop-claim-address", operator.ID, batch)
}
