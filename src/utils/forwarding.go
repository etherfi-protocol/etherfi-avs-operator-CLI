package utils

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/common"
)

// PackForwardCallForAdmin encloses the provided calldata in a call to AvsOperatorManager.adminForwardCall()
func PackForwardCallForAdmin(operatorID int64, input []byte, innerTarget common.Address) ([]byte, error) {

	managerABI, err := contracts.AvsOperatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("fetching abi: %w", err)
	}

	subcallTarget := innerTarget
	subcallSelector := [4]byte(input[:4])
	subcallData := input[4:]
	encodedForwardData, err := managerABI.Pack("adminForwardCall", big.NewInt(operatorID), subcallTarget, subcallSelector, subcallData)
	if err != nil {
		return nil, fmt.Errorf("packing forwardCall: %w", err)
	}

	return encodedForwardData, nil
}

// PackCallForOperator encloses the provided calldata in a call to AvsOperatorManager.forwardOperatorCall()
func PackCallForOperator(operatorID int64, input []byte, innerTarget common.Address) ([]byte, error) {

	managerABI, err := contracts.AvsOperatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("fetching abi: %w", err)
	}

	encodedForwardData, err := managerABI.Pack("forwardOperatorCall", big.NewInt(operatorID), innerTarget, input)
	if err != nil {
		return nil, fmt.Errorf("packing forwardCall: %w", err)
	}
	fmt.Printf("forwardOperatorCall: 0x%s\n\n", hex.EncodeToString(encodedForwardData))

	return encodedForwardData, nil
}
