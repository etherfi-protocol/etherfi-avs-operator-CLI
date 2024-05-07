package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var operatorDetailsCmd = &cli.Command{
	Name:   "operator-details",
	Usage:  "Information about the current configuration of target operator",
	Action: handleOperatorDetails,
	Before: setupCommand,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.IntFlag{
			Name:  "chain-id",
			Usage: "Chain ID",
			Value: 1, // default to mainnet
		},
	},
}

func handleOperatorDetails(ctx context.Context, cli *cli.Command) error {

	operatorID := cli.Int("operator-id")

	/*
		// connect to RPC node
		rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
		if err != nil {
			return fmt.Errorf("dialing rpc: %w", err)
		}
	*/

	return operatorDetails(operatorID, rpcClient)
}

func operatorDetails(operatorID int64, rpcClient *ethclient.Client) error {

	// load configuration
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("querying chainID from RPC: %w", err)
	}
	cfg, err := configForChain(chainID.Int64())
	if err != nil {
		return err
	}

	// bind contracts
	avsManager, err := contracts.NewEtherfiAVSOperatorsManager(cfg.OperatorManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding OperatorManager contract: %w", err)
	}
	avsDirectory, err := contracts.NewAVSDirectory(cfg.AvsDirectoryAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding AVSDirectory contract: %w", err)
	}

	// look up address of operator contract
	operatorAddr, err := avsManager.AvsOperators(nil, big.NewInt(operatorID))
	if err != nil {
		return fmt.Errorf("fetching operator address: %w", err)
	}

	status, err := avsDirectory.AvsOperatorStatus(nil, cfg.EigenDAServiceManager, operatorAddr)
	if err != nil {
		return fmt.Errorf("fetching operator status: %w", err)
	}
	signer, err := avsManager.EcdsaSigner(nil, big.NewInt(operatorID))
	if err != nil {
		return fmt.Errorf("fetching ECDSAsigner: %w", err)
	}
	nodeRunner, err := avsManager.AvsNodeRunner(nil, big.NewInt(operatorID))
	if err != nil {
		return fmt.Errorf("fetching avsNodeRunner: %w", err)
	}

	/*
		deprecatedDeets, err := avsManager.GetAvsInfo(nil, big.NewInt(operatorID), cfg.EigenDARegistryCoordinator)
		if err != nil {
			return fmt.Errorf("fetching deprecated avsInfo: %w", err)
		}
	*/

	fmt.Printf("OperatorID: %d, Registered: %d\n", operatorID, status)
	fmt.Printf("ecdsaSigner: %s\n", signer)
	fmt.Printf("runner: %s\n", nodeRunner)
	//fmt.Printf("deets %+v\n", deprecatedDeets)
	return nil
}
