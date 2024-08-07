package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bin/avs-cli/eigenda"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bin/avs-cli/eoracle"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bin/avs-cli/witness-chain"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

// global state accessible by all commands
type cmdContext struct {
	rpcClient *ethclient.Client

	AVSDirectory               *contracts.AVSDirectory
	EigenDARegistryCoordinator *contracts.RegistryCoordinator
	//EigenDA
}

var rpcClient *ethclient.Client

var (
	// TODO: refactoring
	createBlsCmd = &cli.Command{
		Name:   "create-bls-signature",
		Usage:  "Create BLS signature for AVS registration",
		Action: createBLS,
		// TODO: Check flags details
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "bls-key-file",
				Usage: "BLS key file",
			},
			&cli.StringFlag{
				Name:  "bls-key-password",
				Usage: "BLS key password",
			},
			&cli.StringFlag{
				Name:  "rpc",
				Usage: "RPC Endpoint",
			},
			&cli.StringFlag{
				Name:  "service-manager",
				Usage: "Contract Address of Service Manager",
			},
			&cli.StringFlag{
				Name:  "eigenlayer-operator",
				Usage: "Address of Registered Eigenlayer Operator",
			},
		},
	}
)

// type BeforeFunc func(context.Context, *Command) error
func before(ctx context.Context, cmd *cli.Command) error {
	// try to load RPC_URL from env or flags
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = cmd.String("rpc-url")
	}

	if rpcURL == "" {
		return fmt.Errorf("must set env var $RPC_URL or use --rpc-url flag")
	}

	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return fmt.Errorf("dialing RPC: %w", err)
	}

	// make globally accessible by all commands
	rpcClient = client

	return nil
}

func main() {
	cmd := &cli.Command{
		// global flags
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "rpc-url",
			},
		},

		// operations we want to run for every single command
		Before: before,

		Commands: []*cli.Command{
			createBlsCmd,
			//		registerBLSCmd,
			registrationDigestCmd,
			registerOperatorCmd,
			//updateWhitelistCmd,
			operatorDetailsCmd,
			updateEcdsaSignerCmd,
			lagrangeCmd,
			eoracle.EOracleCmd,
			eigenda.EigenDACmd,
			witness.WitnessCmd,
		},
	}

	ctx := context.Background()
	if err := cmd.Run(ctx, os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
