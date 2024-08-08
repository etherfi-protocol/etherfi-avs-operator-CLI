package eoracle

import (
	"context"
	"fmt"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/eoracle"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/etherfi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var eoracleAPI *eoracle.API
var etherfiAPI *etherfi.API

var EOracleCmd = &cli.Command{
	Name:   "eoracle",
	Usage:  "various actions related to managing eOracle operators",
	Before: prepareEoracleCmd,
	Commands: []*cli.Command{
		EOracleRegisterCmd,
		EOraclePrepareRegistrationCmd,
	},
}

// run before any subcommand executes
func prepareEoracleCmd(ctx context.Context, cmd *cli.Command) error {
	// try to load RPC_URL from env or flags
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
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
	cfg, err := bindings.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}
	registryCoordinator, _ := eoracle.NewRegistryCoordinator(cfg.EOracleRegistryCoordinatorAddress, rpcClient)
	serviceManager, _ := eoracle.NewServiceManager(cfg.EOracleServiceManagerAddress, rpcClient)
	avsOperatorManager, _ := etherfi.NewAvsOperatorManager(cfg.AvsOperatorManagerAddress, rpcClient)

	// make globally accessible by all sub commands
	etherfiAPI = &etherfi.API{
		Client:                    rpcClient,
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		AvsOperatorManager:        avsOperatorManager,
	}
	eoracleAPI = &eoracle.API{
		Client:                     rpcClient,
		RegistryCoordinatorAddress: cfg.EOracleRegistryCoordinatorAddress,
		RegistryCoordinator:        registryCoordinator,
		ServiceManagerAddress:      cfg.EOracleServiceManagerAddress,
		ServiceManager:             serviceManager,
		AvsOperatorManagerAddress:  cfg.AvsOperatorManagerAddress,
	}

	return nil
}
