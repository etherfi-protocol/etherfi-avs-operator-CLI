package main

import (
	"fmt"
	"math/big"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func OperatorAddressFromID(operatorID int64, cfg *bindings.Config, rpcClient *ethclient.Client) (common.Address, error) {
	operatorManagerContract, err := contracts.NewAvsOperatorManager(cfg.OperatorManagerAddress, rpcClient)
	if err != nil {
		return common.Address{}, fmt.Errorf("binding operatorManager: %w", err)
	}
	operatorAddr, err := operatorManagerContract.AvsOperators(nil, big.NewInt(operatorID))
	if err != nil {
		return common.Address{}, fmt.Errorf("looking up operator address: %w", err)
	}

	return operatorAddr, nil
}
