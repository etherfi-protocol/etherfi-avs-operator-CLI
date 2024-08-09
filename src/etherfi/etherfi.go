package etherfi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
)

type API struct {
	Client *ethclient.Client

	AvsOperatorManager        *AvsOperatorManager
	AvsOperatorManagerAddress common.Address
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	avsOperatorManager, _ := NewAvsOperatorManager(cfg.AvsOperatorManagerAddress, rpcClient)

	return &API{
		Client:                    rpcClient,
		AvsOperatorManager:        avsOperatorManager,
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
	}
}
