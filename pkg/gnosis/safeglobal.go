package gnosis

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/contracts/safeglobal"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/config"
)

type SafeGlobal struct {
	Client *ethclient.Client

	SignMessageLib        *safeglobal.SignMessageLib
	SignMessageLibAddress common.Address
}

func New(cfg config.Config, rpcClient *ethclient.Client) *SafeGlobal {

	signMessageLib, _ := safeglobal.NewSignMessageLib(cfg.SignMessageLibAddress, rpcClient)

	return &SafeGlobal{
		Client:                rpcClient,
		SignMessageLib:        signMessageLib,
		SignMessageLibAddress: cfg.SignMessageLibAddress,
	}
}
