package etherfi

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Operator struct {
	Client   *ethclient.Client
	Contract *AvsOperator
	Address  common.Address
	ID       int64
}

func (e API) LookupOperatorByID(operatorID int64) (*Operator, error) {

	// look up operator contract associated with this id
	operatorAddr, err := e.AvsOperatorManager.AvsOperators(nil, big.NewInt(operatorID))
	if err != nil {
		return nil, fmt.Errorf("looking up operator address: %w", err)
	}
	contract, err := NewAvsOperator(operatorAddr, e.Client)
	if err != nil {
		return nil, fmt.Errorf("binding operator contract: %w", err)
	}

	return &Operator{
		Client:   e.Client,
		Contract: contract,
		Address:  operatorAddr,
		ID:       operatorID,
	}, nil
}
