package eigenda

import (
	"context"
	"fmt"
	"os"

	// need to alias because eigenlayer has a package name that doesn't match the filepath

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/eigenda"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/urfave/cli/v3"
)

var eigendaAPI *eigenda.API
var etherfiAPI *etherfi.API

var EigenDACmd = &cli.Command{
	Name:   "eigenda",
	Usage:  "various actions related to managing eigenDA operators",
	Before: prepareEigendaCmd,
	Commands: []*cli.Command{
		PrepareRegistrationCmd,
		RegisterCmd,
		UpdateSocketCmd,
	},
}

// run before any subcommand executes
func prepareEigendaCmd(ctx context.Context, cmd *cli.Command) error {
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
	etherfiAPI = etherfi.New(cfg, rpcClient)
	eigendaAPI = eigenda.New(cfg, rpcClient)

	return nil
}
