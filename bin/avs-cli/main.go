package main

import (
	"context"
	"fmt"
	"os"

	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/alignedlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/altlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/arpa"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/automata"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/brevis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/cybermach"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/eigenda"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/eoracle"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/ethgas"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/gasp"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/hyperlane"
	lagrangezk "github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/lagrangeZK"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/openlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/predicate"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/redstone"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/symbiotic"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/ungate"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bin/avs-cli/unifi"
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
			RegistrationStatusCmd,
			StakedropAddressCmd,
			operatorDetailsCmd,
			updateEcdsaSignerCmd,
			ClaimRewardsCmd,
			symbiotic.SymbioticCmd,

			// AVS specific commands
			alignedlayer.AlignedLayerCmd,
			altlayer.AltLayerCmd,
			arpa.ARPACmd,
			automata.AutomataCmd,
			brevis.BrevisCmd,
			cybermach.CyberMachCmd,
			eigenda.EigenDACmd,
			eoracle.EOracleCmd,
			ethgas.EthgasCmd,
			gasp.GaspCmd,
			hyperlane.HyperlaneCmd,
			lagrangezk.LagrangeZKCmd,
			openlayer.OpenlayerCmd,
			predicate.PredicateCmd,
			redstone.RedstoneCmd,
			ungate.UngateCmd,
			unifi.UniFiCmd,
			witness.WitnessCmd,
		},
	}

	ctx := context.Background()
	if err := cmd.Run(ctx, os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
