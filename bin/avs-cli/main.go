package main

import (
	"context"
	"fmt"
	"os"

	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/eigenda"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/eoracle"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/witness-chain"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{

		// global flags
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "rpc-url",
				Usage:    "rpc url",
				Required: false,
			},
		},
		Commands: []*cli.Command{
			//registrationDigestCmd,
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
