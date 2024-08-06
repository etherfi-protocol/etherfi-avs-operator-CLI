package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/cmd/etherfi/eigenda"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/cmd/etherfi/eoracle"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/cmd/etherfi/witness-chain"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

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

func main() {
	cmd := &cli.Command{
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
