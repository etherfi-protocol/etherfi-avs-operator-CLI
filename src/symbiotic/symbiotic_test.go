package symbiotic

import (
	"fmt"
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

	for name, vault := range cfg.SymbioticVaults {
		fmt.Printf("Vault: %s\n", name)
		for networkAddr, networkName := range cfg.SymbioticNetworks {
			fmt.Printf("Network: %s\n", networkName)
			api.NetworkStats(networkAddr, vault.Delegator)
		}
	}

	//api.NetworkStats(cfg.CapXNetwork, cfg.SymbioticVaults[config.WSTETH_VAULT].Delegator)
}

func TestVaultTable(t *testing.T) {
	cfg := config.Mainnet
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))

	if err != nil {
		t.Fatalf("failed to dial RPC: %v", err)
	}

	api := New(&cfg, rpcClient)

	api.PrintVaultTable(cfg.SymbioticVaults["wstETH"])

}
