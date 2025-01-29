package symbiotic

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
)

func TestNetworkStats(t *testing.T) {

	cfg := config.Mainnet
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))

	if err != nil {
		t.Fatalf("failed to dial RPC: %v", err)
	}

	api := New(&cfg, rpcClient)
	api.NetworkStats(cfg.CapXNetwork, cfg.SymbioticVaults[config.WSTETH_VAULT].Delegator)

}
