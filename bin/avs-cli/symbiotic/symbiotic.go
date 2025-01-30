package symbiotic

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/symbiotic"
	"github.com/urfave/cli/v3"
)

var symbioticAPI *symbiotic.API

var SymbioticCmd = &cli.Command{
	Name:   "symbiotic",
	Usage:  "various actions related to managing symbiotic vaults",
	Before: prepareCmd,
	Commands: []*cli.Command{
		VaultReportCmd,
		SetNetworkLimitCmd,
	},
}

// run before any subcommand executes
func prepareCmd(ctx context.Context, cmd *cli.Command) error {
	// try to load RPC_URL from env or flags
	rpcURL := os.Getenv("RPC_URL")
	if cmd.String("rpc-url") != "" {
		rpcURL = cmd.String("rpc-url")
	}
	if rpcURL == "" {
		return fmt.Errorf("must set env var $RPC_URL or use --rpc-url flag")
	}
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing RPC: %w", err)
	}

	// load all required addresses for this chain and bind applicable contracts
	cfg, err := config.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	// make globally accessible by all sub commands
	symbioticAPI = symbiotic.New(cfg, rpcClient)

	return nil
}

var SetNetworkLimitCmd = &cli.Command{
	Name:   "set-network-limit",
	Usage:  "(Admin) set-network-limit alongside optional operator share updates",
	Action: handleSetNetworkLimit,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "asset",
			Usage:    "which asset vault to update i.e. [wstETH, LBTC, ETHFI]",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "network",
			Usage:    "address of network to update limit for",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "limit",
			Usage:    "new network limit in wei",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:  "share-updates",
			Usage: "comma separated list of {operator}:{shares} entries. i.e '0x51B6D824bd35AeD4FD1a9E253E41Dc7C9feeFa30:3333",
		},
	},
}

func handleSetNetworkLimit(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	asset := cli.String("asset")
	network := common.HexToAddress(cli.String("network"))
	rawShareUpdates := cli.StringSlice("share-updates")

	// validate
	networkLimit, ok := big.NewInt(0).SetString(cli.String("limit"), 10)
	if !ok {
		return fmt.Errorf("invalid limit")
	}
	vaultCfg, exists := symbioticAPI.KnownVaults[asset]
	if !exists {
		return fmt.Errorf("Unknown vault")
	}
	_, networkExists := symbioticAPI.KnownNetworks[network]
	if !networkExists {
		return fmt.Errorf("Unknown network")
	}

	// parse operator share updates
	var shareUpdates []symbiotic.OperatorShareUpdate
	for _, update := range rawShareUpdates {

		raw := strings.Split(update, ":")
		if len(raw) != 2 {
			return fmt.Errorf("invalid operator share update format: %s", update)
		}

		operator := common.HexToAddress(raw[0])
		shares, ok := big.NewInt(0).SetString(raw[1], 10)
		if !ok {
			return fmt.Errorf("invalid shares value: %s", raw[1])
		}
		shareUpdates = append(shareUpdates, symbiotic.OperatorShareUpdate{
			Operator: operator,
			Shares:   shares,
		})
	}

	return symbioticAPI.SetNetworkLimit(vaultCfg, networkLimit, network, shareUpdates)
}

var VaultReportCmd = &cli.Command{
	Name:   "vault-report",
	Usage:  "(Admin) see vault allocation data",
	Action: handleVaultReport,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "asset",
			Usage:    "which vault asset to pull data for",
			Required: true,
		},
	},
}

func handleVaultReport(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	asset := cli.String("asset")

	vault, exists := symbioticAPI.KnownVaults[asset]
	if !exists {
		return fmt.Errorf("Unknown vault")
	}

	report, err := symbioticAPI.VaultReport(vault)
	if err != nil {
		return fmt.Errorf("failed to create vault report: %w", err)
	}
	report.PrettyPrint()

	return nil
}
