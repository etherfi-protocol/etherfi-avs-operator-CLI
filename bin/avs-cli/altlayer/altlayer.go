package altlayer

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/altlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/urfave/cli/v3"
)

var altlayerAPI *altlayer.API
var etherfiAPI *etherfi.API

var AltLayerCmd = &cli.Command{
	Name:   "altlayer",
	Usage:  "various actions related to managing AltLayer operators",
	Before: prepareAltLayerCmd,
	Commands: []*cli.Command{
		AltLayerPrepareRegistrationCmd,
		AltLayerRegisterCmd,
	},
}

// run before any subcommand executes
func prepareAltLayerCmd(ctx context.Context, cmd *cli.Command) error {
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
	altlayerAPI = altlayer.New(cfg, rpcClient)

	return nil
}
