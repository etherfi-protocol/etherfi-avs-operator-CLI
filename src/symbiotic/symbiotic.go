package symbiotic

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/fatih/color"
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

	/*
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
	*/

	return &API{
		Client:         rpcClient,
		KnownNetworks:  cfg.SymbioticNetworks,
		KnownOperators: cfg.SymbioticOperators,
	}
}

type VaultReport struct {
}

func (a *API) PrintVaultTable(vaultData config.SymbioticVaultConfig, name string) error {

	//delegator, _ := NewDelegator(delegatorAddr, a.Client)
	vault, _ := NewVault(vaultData.Vault, a.Client)

	currentVaultStakeBig, err := vault.ActiveStake(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch current vault stake: %v", err)
	}

	totalDelegation, _ := currentVaultStakeBig.Float64()
	totalDelegation /= math.Pow(10, float64(vaultData.Decimals))

	// TODO: redo name
	color.HiCyan("%s\n", name)
	fmt.Printf("VaultStake: %0.2f %s\n", totalDelegation, vaultData.Asset)

	fmt.Printf("Networks: \n")
	for networkAddr, networkName := range a.KnownNetworks {

		delegator, _ := NewDelegator(vaultData.Delegator, a.Client)
		subnetwork := toSubnetwork(networkAddr)
		maxLimitBig, err := delegator.MaxNetworkLimit(nil, subnetwork)
		if err != nil {
			return fmt.Errorf("fectching MaxNetworkLimit: %w", err)
		}
		currentLimitBig, err := delegator.NetworkLimit(nil, subnetwork)
		if err != nil {
			return fmt.Errorf("fectching NetworkLimit: %w", err)
		}

		// okay to be lossy, this is just for metrics
		maxLimit, _ := maxLimitBig.Float64()
		currentLimit, _ := currentLimitBig.Float64()
		maxLimit /= math.Pow(10, float64(vaultData.Decimals))
		currentLimit /= math.Pow(10, float64(vaultData.Decimals))

		color.HiYellow("%s\n", networkName)
		padding := 4
		w := tabwriter.NewWriter(os.Stdout, 5, 0, padding, ' ', tabwriter.AlignRight)
		additionalAllocation := min(maxLimit-currentLimit, totalDelegation-currentLimit)
		fmt.Fprintf(w, "NetworkLimit (set by us)     \t%0.2f %s\t\n", currentLimit, vaultData.Asset)
		fmt.Fprintf(w, "NetworkMax (set by network)  \t%0.2f %s\t\n", maxLimit, vaultData.Asset)
		fmt.Fprintf(w, "PossibleAdditionalAllocation \t%0.2f %s\t\n", additionalAllocation, vaultData.Asset)
		w.Flush()
	}

	return nil
}

func (a *API) NetworkStats(networkAddr common.Address, delegatorAddr common.Address) error {
	/*
		What data do I need for table?
		Vault Name:
		Total Delegation:

		// Networks
		Network: "Symbiosis"
		CurrentDelegation (set by us):
		NetworkMax   (set by network):
		PossibleAdditionalAllocation:

		// TABLE
		OperatorName address stake

	*/

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
