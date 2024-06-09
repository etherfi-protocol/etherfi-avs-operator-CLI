package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var lagrangeCmd = &cli.Command{
	Name:  "lagrange",
	Usage: "various actions related to managing lagrange operators",
	Commands: []*cli.Command{
		lagrangeRegisterCmd,
	},
}

var lagrangeRegisterCmd = &cli.Command{
	Name:   "register",
	Action: handleLagrangeRegister,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
	},
}

func handleLagrangeRegister(ctx context.Context, cli *cli.Command) error {

	operatorID := cli.Int("operator-id")
	rpcURL := cli.String("rpc-url")

	// connect to RPC node
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}

	return lagrangeRegister(operatorID, rpcClient)
}

func lagrangeRegister(operatorID int64, rpcClient *ethclient.Client) error {

	return nil
}
