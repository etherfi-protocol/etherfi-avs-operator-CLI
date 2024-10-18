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

var ClaimRewardsCmd = &cli.Command{
	Name:   "claim-rewards",
	Usage:  "(Admin) claim AVS rewards for target operator",
	Action: handleClaimRewards,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "recipient",
			Usage:    "address to recieve claimed rewards",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     "tokens",
			Usage:    "which tokens to claim as rewards",
			Value:    []string{"0xec53bF9167f50cDEB3Ae105f56099aaaB9061F83"}, // default to EIGEN
			Required: false,
		},
	},
}

func handleClaimRewards(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	operatorID := cli.Int("operator-id")
	recipient := common.HexToAddress(cli.String("recipient"))
	tokenStrs := cli.StringSlice("tokens")

	var tokens []common.Address
	for _, token := range tokenStrs {
		tokens = append(tokens, common.HexToAddress(token))
	}
	if operatorID == 0 {
		return fmt.Errorf("invalid operator-id")
	}

	// try to load RPC_URL from env or flags
	rpcURL := os.Getenv("RPC_URL")
	if cli.String("rpc-url") != "" {
		rpcURL = cli.String("rpc-url")
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

	// look up operator contract associated with this id
	etherfiAPI := etherfi.New(cfg, rpcClient)
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	eigenlayerAPI := eigenlayer.New(cfg, rpcClient)
	return eigenlayerAPI.ClaimAvsOperatorRewards(operator, recipient, tokens)
}
