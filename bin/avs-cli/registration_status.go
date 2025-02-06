package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/eigenlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

var RegistrationStatusCmd = &cli.Command{
	Name:   "registration-status",
	Usage:  "check which AVSs a target operator is registered for (from the perspective of eigenlayer)",
	Action: handleRegistrationStatus,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
	},
}

func handleRegistrationStatus(ctx context.Context, cmd *cli.Command) error {

	// parse cli input
	operatorID := cmd.Int("operator-id")

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

	// load all required addresses for this chain
	cfg, err := config.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	etherfiAPI := etherfi.New(cfg, rpcClient)

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	type AVS struct {
		Name           string
		ServiceManager common.Address
	}
	activeAVSs := []AVS{
		{Name: "AlignedLayer", ServiceManager: cfg.AlignedLayerServiceManagerAddress},
		{Name: "AltLayer", ServiceManager: cfg.AltLayerServiceManagerAddress},
		{Name: "ARPA", ServiceManager: cfg.ARPAServiceManagerAddress},
		{Name: "Automata", ServiceManager: cfg.AutomataServiceManagerAddress},
		{Name: "Brevis", ServiceManager: cfg.BrevisServiceManagerAddress},
		{Name: "CyberMACH", ServiceManager: cfg.CyberMachServiceManagerAddress},
		{Name: "EigenDA", ServiceManager: cfg.EigenDAServiceManagerAddress},
		{Name: "eOracle", ServiceManager: cfg.EOracleServiceManagerAddress},
		{Name: "Gasp", ServiceManager: cfg.GaspServiceManagerAddress},
		{Name: "Hyperlane", ServiceManager: cfg.HyperlaneServiceManagerAddress},
		{Name: "LagrangeSC", ServiceManager: cfg.LagrangeServiceAddress},
		{Name: "LagrangeZK", ServiceManager: cfg.LagrangeZKMRServiceManagerAddress},
		{Name: "Openlayer", ServiceManager: cfg.OpenlayerServiceManagerAddress},
		{Name: "Ungate", ServiceManager: cfg.UngateAVSGovernanceAddress},
		{Name: "Unifi", ServiceManager: cfg.UniFiAvsManagerAddress},
		{Name: "Witnesschain", ServiceManager: cfg.WitnessChainWitnessHubAddress},
	}

	avsDirectory, err := eigenlayer.NewAvsDirectory(cfg.AvsDirectoryAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding AVSDirectory contract: %w", err)
	}

	greenText := color.New(color.FgHiGreen)
	redText := color.New(color.FgHiRed)

	for _, avs := range activeAVSs {
		status, err := avsDirectory.AvsOperatorStatus(nil, avs.ServiceManager, operator.Address)
		if err != nil {
			return fmt.Errorf("fetching operator status: %w", err)
		}
		fmt.Printf("%s %s\n", avs.ServiceManager, operator.Address)

		fmt.Printf("%s: ", avs.Name)
		if status == 1 {
			greenText.Printf("Registered\n")
		} else {
			redText.Printf("Unregistered\n")
		}

	}

	return nil
}
