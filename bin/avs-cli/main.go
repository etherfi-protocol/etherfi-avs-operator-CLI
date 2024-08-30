package main

import (
	"context"
	"fmt"
	"os"

	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/altlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/arpa"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/automata"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/brevis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/eigenda"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/eoracle"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/hyperlane"
	lagrangezk "github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/lagrangeZK"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/openlayer"
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
			BlsPubkeyRegistrationHashCmd,
			operatorDetailsCmd,
			updateEcdsaSignerCmd,

			// AVS specific commands
			altlayer.AltLayerCmd,
			arpa.ARPACmd,
			automata.AutomataCmd,
			brevis.BrevisCmd,
			eigenda.EigenDACmd,
			eoracle.EOracleCmd,
			hyperlane.HyperlaneCmd,
			//lagrangesc.LagrangeSCCmd,
			lagrangezk.LagrangeZKCmd,
			openlayer.OpenlayerCmd,
			witness.WitnessCmd,
		},
	}

	ctx := context.Background()
	if err := cmd.Run(ctx, os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
