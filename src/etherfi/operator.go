package etherfi

import (
	"fmt"
	"math/big"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type API struct {
	Client                    *ethclient.Client
	AvsOperatorManagerAddress common.Address
	AvsOperatorManager        *contracts.AvsOperatorManager
}

type Operator struct {
	Client  *ethclient.Client
	Address common.Address
	ID      int64
}

func (e API) LookupOperatorByID(operatorID int64) (*Operator, error) {

	// look up operator contract associated with this id
	operatorAddr, err := e.AvsOperatorManager.AvsOperators(nil, big.NewInt(operatorID))
	if err != nil {
		return nil, fmt.Errorf("looking up operator address: %w", err)
	}

	return &Operator{
		Client:  e.Client,
		Address: operatorAddr,
		ID:      operatorID,
	}, nil
}
