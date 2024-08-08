package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bin/avs-cli/eigenda"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bin/avs-cli/eoracle"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bin/avs-cli/witness-chain"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{

		Commands: []*cli.Command{
			//		registerBLSCmd,
			//registrationDigestCmd,
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
