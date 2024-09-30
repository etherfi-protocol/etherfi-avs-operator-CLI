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
	"github.com/urfave/cli/v3"
)

var StakedropAddressCmd = &cli.Command{
	Name:   "stakedrop-address",
	Usage:  "update the claims address for an operator in the season 2 stakedrop",
	Action: handleStakedropAddress,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "claim-address",
			Usage:    "Address the operator wishes to claim with",
			Required: true,
		},
	},
}

func handleStakedropAddress(ctx context.Context, cmd *cli.Command) error {

	// parse cli input
	operatorID := cmd.Int("operator-id")
	claimAddress := common.HexToAddress(cmd.String("claim-address"))

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

	return eigenlayer.RegisterClaimAddressForOperator(operator, claimAddress)
}
