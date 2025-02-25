package symbiotic

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"text/tabwriter"

	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/fatih/color"
)

type VaultReport struct {
	Asset         string
	AssetDecimals int
	ActiveStake   *big.Int

	NetworkStats []VaultNetworkStats
}

type VaultNetworkStats struct {
	Name            string
	NetworkLimit    *big.Int
	MaxNetworkLimit *big.Int
}

func (a *API) VaultReport(vaultCfg config.SymbioticVaultConfig) (*VaultReport, error) {

	// Grab global vault data
	vault, _ := NewVault(vaultCfg.Vault, a.Client)
	activeVaultStakeBig, err := vault.ActiveStake(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch current vault stake: %v", err)
	}

	// fetch stats for each network in the vault
	var networkStats []VaultNetworkStats
	for networkAddr, networkName := range a.KnownNetworks {

		delegator, _ := NewDelegator(vaultCfg.Delegator, a.Client)
		subnetwork := toSubnetwork(networkAddr)

		maxNetworkLimit, err := delegator.MaxNetworkLimit(nil, subnetwork)
		if err != nil {
			return nil, fmt.Errorf("fectching MaxNetworkLimit: %w", err)
		}
		networkLimit, err := delegator.NetworkLimit(nil, subnetwork)
		if err != nil {
			return nil, fmt.Errorf("fectching NetworkLimit: %w", err)
		}

		networkStats = append(networkStats, VaultNetworkStats{
			Name:            networkName,
			NetworkLimit:    networkLimit,
			MaxNetworkLimit: maxNetworkLimit,
		})
	}

	return &VaultReport{
		Asset:         vaultCfg.Asset,
		AssetDecimals: vaultCfg.Decimals,
		ActiveStake:   activeVaultStakeBig,
		NetworkStats:  networkStats,
	}, nil
}

func (v *VaultReport) PrettyPrint() {

	activeVaultStake, _ := v.ActiveStake.Float64()
	activeVaultStake /= math.Pow(10, float64(v.AssetDecimals))

	color.HiCyan("%s Vault\n", v.Asset)
	fmt.Printf("VaultStake: %0.2f %s\n", activeVaultStake, v.Asset)
	for _, network := range v.NetworkStats {

		maxNetworkLimit, _ := network.MaxNetworkLimit.Float64()
		maxNetworkLimit /= math.Pow(10, float64(v.AssetDecimals))
		networkLimit, _ := network.NetworkLimit.Float64()
		networkLimit /= math.Pow(10, float64(v.AssetDecimals))

		color.HiYellow("%s\n", network.Name)
		padding := 4
		w := tabwriter.NewWriter(os.Stdout, 5, 0, padding, ' ', tabwriter.AlignRight)
		additionalAllocation := min(maxNetworkLimit-networkLimit, activeVaultStake-networkLimit)
		fmt.Fprintf(w, "NetworkLimit (set by us)     \t%0.2f %s\t\n", networkLimit, v.Asset)
		fmt.Fprintf(w, "NetworkMax (set by network)  \t%0.2f %s\t\n", maxNetworkLimit, v.Asset)
		fmt.Fprintf(w, "PossibleAdditionalAllocation \t%0.2f %s\t\n", additionalAllocation, v.Asset)
		w.Flush()
	}
}
