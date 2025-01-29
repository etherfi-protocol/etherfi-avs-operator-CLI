package symbiotic

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
)

type API struct {
	Client *ethclient.Client

	Vault        *Vault
	VaultAddress common.Address

	Delegator        *Delegator
	DelegatorAddress common.Address

	BurnerRouter        *BurnerRouter
	BurnerRouterAddress common.Address

	KnownNetworks  map[common.Address]string
	KnownOperators map[common.Address]string
}

func New(cfg *config.Config, rpcClient *ethclient.Client) *API {

	knownNetworks := map[common.Address]string{
		cfg.CapXNetwork:      "CapX Cloud",
		cfg.CycleNetwork:     "Cycle Network",
		cfg.DittoNetwork:     "Ditto",
		cfg.HyperlaneNetwork: "Hyperlane",
		cfg.HyveNetwork:      "Hyve",
		cfg.MarlinNetwork:    "Marlin",
		cfg.RadiusNetwork:    "Radius",
		cfg.RouterNetwork:    "Router",
		cfg.SymbiosisNetwork: "Symbiosis",
	}
	knownOperators := map[common.Address]string{
		cfg.BlockdaemonSymbioticOperator: "Blockdaemon",
		cfg.KilnSymbioticOperator:        "Kiln",
		cfg.P2PSymbioticOperator:         "P2P",
		cfg.Pier2SymbioticOperator:       "Pier2",
	}

	return &API{
		Client:         rpcClient,
		KnownNetworks:  knownNetworks,
		KnownOperators: knownOperators,
	}
}

func (a *API) NetworkStats(networkAddr common.Address, delegatorAddr common.Address) error {

	subnetwork := toSubnetwork(networkAddr)

	// MaxNetworkLimit
	delegator, _ := NewDelegator(delegatorAddr, a.Client)
	fmt.Printf("delegator: %s\n", delegatorAddr)

	maxLimit, err := delegator.MaxNetworkLimit(nil, subnetwork)
	if err != nil {
		return fmt.Errorf("fectching MaxNetworkLimit: %w", err)
	}

	currentLimit, err := delegator.NetworkLimit(nil, subnetwork)
	if err != nil {
		return fmt.Errorf("fectching NetworkLimit: %w", err)
	}

	fmt.Printf("maxLimit %s\n", maxLimit)
	fmt.Printf("currentLimit %s\n", currentLimit)
	return nil
}

// convert from network address to 32byte subnetwork ID
// This effectively just appends 12 bytes of zeroes.
// Will need to update once a network actually uses the subnetwork feature
func toSubnetwork(networkAddr common.Address) [32]byte {
	var subnetwork [32]byte
	copy(subnetwork[:], networkAddr[:])
	return subnetwork
}
