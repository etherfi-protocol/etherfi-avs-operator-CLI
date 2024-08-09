package brevis

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/brevis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/urfave/cli/v3"
)

var brevisAPI *brevis.API
var etherfiAPI *etherfi.API

var BrevisCmd = &cli.Command{
	Name:   "brevis",
	Usage:  "various actions related to managing Brevis operators",
	Before: prepareBrevisCmd,
	Commands: []*cli.Command{
		BrevisPrepareRegistrationCmd,
		BrevisRegisterCmd,
	},
}

// run before any subcommand executes
func prepareBrevisCmd(ctx context.Context, cmd *cli.Command) error {
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

	// load all required addresses for this chain
	cfg, err := config.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	// make globally accessible by all sub commands
	etherfiAPI = etherfi.New(cfg, rpcClient)
	brevisAPI = brevis.New(cfg, rpcClient)

	return nil
}
