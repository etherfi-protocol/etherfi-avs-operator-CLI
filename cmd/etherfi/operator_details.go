package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var operatorDetailsCmd = &cli.Command{
	Name:   "operator-details",
	Usage:  "Information about the current configuration of target operator",
	Action: handleOperatorDetails,
	//Before: setupCommand,
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

	// connect to RPC node
	// TODO: FIX URL
	rpcClient, err := ethclient.Dial(os.Getenv("HOLESKY_RPC_URL"))
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}

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
	delegationManager, err := contracts.NewDelegationManager(cfg.DelegationManager, rpcClient)
	if err != nil {
		return fmt.Errorf("binding delegationManager: %w", err)
	}
	strategyManager, err := contracts.NewStrategyManager(cfg.StrategyManager, rpcClient)
	if err != nil {
		return fmt.Errorf("binding strategyManager: %w", err)
	}

	// look up address of operator contract
	operatorAddr, err := avsManager.AvsOperators(nil, big.NewInt(operatorID))
	if err != nil {
		return fmt.Errorf("fetching operator address: %w", err)
	}

	signer, err := avsManager.EcdsaSigner(nil, big.NewInt(operatorID))
	if err != nil {
		return fmt.Errorf("fetching ECDSAsigner: %w", err)
	}
	nodeRunner, err := avsManager.AvsNodeRunner(nil, big.NewInt(operatorID))
	if err != nil {
		return fmt.Errorf("fetching avsNodeRunner: %w", err)
	}
	fmt.Printf("OperatorID: %d - %s\n", operatorID, operatorAddr)
	fmt.Printf("ecdsaSigner: %s\n", signer)
	fmt.Printf("runner: %s\n", nodeRunner)

	type AVS struct {
		Name                string
		ServiceManager      common.Address
		RegistryCoordinator common.Address
	}
	activeAVS := []AVS{
		{Name: "EigenDA", ServiceManager: cfg.EigenDAServiceManager, RegistryCoordinator: cfg.EigenDARegistryCoordinator},
		{Name: "Brevis", ServiceManager: cfg.BrevisServiceManager, RegistryCoordinator: cfg.BrevisRegistryCoordinator},
	}
	for _, avs := range activeAVS {
		status, err := avsDirectory.AvsOperatorStatus(nil, cfg.EigenDAServiceManager, operatorAddr)
		if err != nil {
			return fmt.Errorf("fetching operator status: %w", err)
		}
		deprecatedDeets, err := avsManager.GetAvsInfo(nil, big.NewInt(operatorID), cfg.EigenDARegistryCoordinator)
		if err != nil {
			return fmt.Errorf("fetching deprecated avsInfo: %w", err)
		}
		hasUploadedPubkey := deprecatedDeets.Params.PubkeyG1.X != nil && deprecatedDeets.Params.PubkeyG1.X.Cmp(big.NewInt(0)) != 0
		//shares, amounts, err := delegationManager.GetDelegatableShares(&bind.CallOpts{}, operatorAddr)
		shares, err := delegationManager.GetOperatorShares(&bind.CallOpts{}, operatorAddr, []common.Address{cfg.BeaconEthStrategy, cfg.WethStrategy})
		if err != nil {
			return fmt.Errorf("fetching delegated shares: %w", err)
		}
		strats, amts, err := strategyManager.GetDeposits(&bind.CallOpts{}, operatorAddr)
		if err != nil {
			return fmt.Errorf("fetching deposits: %w", err)
		}

		fmt.Printf("shares: %+v\n", shares)
		fmt.Printf("deposits: %+v - %+v\n", strats, amts)

		fmt.Printf("AVS: %s\n", avs.Name)
		fmt.Printf("Registered: %v\n", status == 1)
		fmt.Printf("Legacy pubkey registered: %v\n\n", hasUploadedPubkey)

	}

	return nil
}
