package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/cmd/etherfi/eoracle"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var rpcClient *ethclient.Client

var (
	/*
		registerBLSCmd = &cli.Command{
			Name:   "register-bls",
			Usage:  "register a bls key with the avsManager contract",
			Action: registerBLS,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "bls-signature-file",
					Usage:    "BLS signature json file",
					Required: true,
				},
				&cli.IntFlag{
					Name:     "operator-id",
					Usage:    "Operator ID",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "registry-coordinator",
					Usage:    "Registry Coordinator address",
					Required: true,
				},
				&cli.IntSliceFlag{
					Name:  "quorum-numbers",
					Usage: "Quorum Numbers",
				},
				&cli.StringFlag{
					Name:  "socket",
					Usage: "Socket",
				},
				&cli.BoolFlag{
					Name:  "broadcast",
					Usage: "broadcast signed tx to network",
				},
				&cli.IntFlag{
					Name:  "chain-id",
					Usage: "Chain ID",
					Value: 1, // default to mainnet
				},
			},
		}
	*/

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
		},
	}

	ctx := context.Background()
	if err := cmd.Run(ctx, os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
