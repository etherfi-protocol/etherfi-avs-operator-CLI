package symbiotic

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"text/tabwriter"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
	"github.com/fatih/color"
)

type API struct {
	Client *ethclient.Client

	KnownNetworks  map[common.Address]string
	KnownOperators map[common.Address]string
	KnownVaults    map[string]config.SymbioticVaultConfig
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

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
		KnownVaults:    cfg.SymbioticVaults,
	}
}

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

func (a *API) VaultReport(vaultData config.SymbioticVaultConfig) (*VaultReport, error) {

	// Grab global vault data
	vault, _ := NewVault(vaultData.Vault, a.Client)
	activeVaultStakeBig, err := vault.ActiveStake(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch current vault stake: %v", err)
	}

	// fetch stats for each network in the vault
	var networkStats []VaultNetworkStats
	for networkAddr, networkName := range a.KnownNetworks {

		delegator, _ := NewDelegator(vaultData.Delegator, a.Client)
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
		Asset:         vaultData.Asset,
		AssetDecimals: vaultData.Decimals,
		ActiveStake:   activeVaultStakeBig,
		NetworkStats:  networkStats,
	}, nil
}

type OperatorShareUpdate struct {
	Operator common.Address
	Shares   *big.Int
}

func (a *API) SetNetworkLimit(vaultCfg config.SymbioticVaultConfig, limit *big.Int, network common.Address, shareUpdates []OperatorShareUpdate) error {

	// output in gnosis compatible format
	batch := gnosis.GnosisBatch{
		Version: "1.0",
		ChainId: "1",
		Meta:    gnosis.GnosisMetadata{Name: "symbiotic-set-network-limit"},
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	delegatorABI, err := DelegatorMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("invalid delegator abi: %v", err)
	}

	// set the new limit
	subnetwork := toSubnetwork(network)
	limitCalldata, err := delegatorABI.Pack("setNetworkLimit", subnetwork, limit)
	if err != nil {
		return fmt.Errorf("failed to pack setNetworkLimit: %w", err)
	}
	batch.AddTransaction(gnosis.SubTransaction{
		Target: vaultCfg.Delegator,
		Value:  big.NewInt(0),
		Data:   limitCalldata,
	})

	// do the requested share updates
	for _, update := range shareUpdates {
		shareCalldata, err := delegatorABI.Pack("setOperatorNetworkShares", subnetwork, update.Operator, update.Shares)
		if err != nil {
			return fmt.Errorf("failed to pack setOperatorNetworkShares: %w", err)
		}
		batch.AddTransaction(gnosis.SubTransaction{
			Target: vaultCfg.Delegator,
			Value:  big.NewInt(0),
			Data:   shareCalldata,
		})
	}

	return utils.ExportJSON("symbiotic-set-network-limit", 0, batch)
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
