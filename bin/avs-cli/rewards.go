package main

import (
	"context"
	"fmt"
	"math/big"
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
		&cli.IntSliceFlag{
			Name:     "operator-ids",
			Usage:    "Operator ID",
			Value:    []int64{}, // default to EIGEN
			Required: false,
		},
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "Claim rewards for all operators",
			Value:   false,
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
	operatorIDs := cli.IntSlice("operator-ids")
	recipient := common.HexToAddress(cli.String("recipient"))
	tokenStrs := cli.StringSlice("tokens")
	isAllOperators := cli.Bool("all")
	var tokens []common.Address
	for _, token := range tokenStrs {
		tokens = append(tokens, common.HexToAddress(token))
	}
	if len(operatorIDs) == 0 && !isAllOperators {
		return fmt.Errorf("no operators provided")
	}
	if len(operatorIDs) > 0 && isAllOperators {
		return fmt.Errorf("must use --all or operator-ids")
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
	etherfiAPI := etherfi.New(cfg, rpcClient)
	var lastAvsId *big.Int
	if isAllOperators {
		lastAvsId, err = etherfiAPI.AvsOperatorManager.NextAvsOperatorId(nil)
		if err != nil {
			return err
		}
		lastAvsIdInt := int(lastAvsId.Int64())
		for i := 0; i < lastAvsIdInt; i++ {
			operatorIDs = append(operatorIDs, int64(i))
		}
	}

	// look up operator contracts associated with these ids
	var operators []*etherfi.Operator
	for _, id := range operatorIDs {
		operator, err := etherfiAPI.LookupOperatorByID(id)
		if err != nil {
			return fmt.Errorf("looking up operator address: %w", err)
		}
		operators = append(operators, operator)
	}

	eigenlayerAPI := eigenlayer.New(cfg, rpcClient)
	return eigenlayerAPI.ClaimAvsOperatorRewards(operators, recipient, tokens)
}
