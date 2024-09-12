package eigenda

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/avs/eigenda"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/config"

	"github.com/urfave/cli/v3"
)

var eigendaAPI *eigenda.EigenDA
var EigenDACmd = &cli.Command{
	Name:   "eigenda",
	Usage:  "various actions related to managing eigenDA operators",
	Before: prepareEigendaCmd,
	Commands: []*cli.Command{
		PrepareRegistrationCmd,
		RegisterCmd,
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
	eigendaAPI = eigenda.New(cfg, rpcClient)

	return nil
}
