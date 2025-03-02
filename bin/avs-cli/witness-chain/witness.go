package witness

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/witnesschain"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/urfave/cli/v3"
)

var witnessAPI *witnesschain.API
var etherfiAPI *etherfi.API

var WitnessCmd = &cli.Command{
	Name:   "witness-chain",
	Usage:  "various actions related to managing Witness-Chain operators",
	Before: prepareWitnessChainCmd,
	Commands: []*cli.Command{
		WitnessPrepareRegistrationCmd,
		WitnessRegisterToAvsCmd,
		WitnessRegisterWatchtowerCmd,
	},
}

// run before any subcommand executes
func prepareWitnessChainCmd(ctx context.Context, cmd *cli.Command) error {
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
	witnessAPI = witnesschain.New(cfg, rpcClient)

	return nil
}
