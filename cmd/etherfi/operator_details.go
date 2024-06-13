package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
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
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
	},
}

func handleOperatorDetails(ctx context.Context, cli *cli.Command) error {

	operatorID := cli.Int("operator-id")
	rpcURL := cli.String("rpc-url")

	// connect to RPC node
	rpcClient, err := ethclient.Dial(rpcURL)
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
	cfg, err := bindings.ConfigForChain(chainID.Int64())
	if err != nil {
		return err
	}

	// bind contracts
	avsManager, err := contracts.NewAvsOperatorManager(cfg.OperatorManagerAddress, rpcClient)
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
	/*
		strategyManager, err := contracts.NewStrategyManager(cfg.StrategyManager, rpcClient)
		if err != nil {
			return fmt.Errorf("binding strategyManager: %w", err)
		}
	*/

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
	shares, err := delegationManager.GetOperatorShares(&bind.CallOpts{}, operatorAddr, []common.Address{cfg.BeaconEthStrategy})
	if err != nil {
		return fmt.Errorf("fetching delegated shares: %w", err)
	}

	fmt.Printf("OperatorID: %d - %s\n", operatorID, operatorAddr)
	fmt.Printf("ecdsaSigner: %s\n", signer)
	fmt.Printf("runner: %s\n", nodeRunner)

	beaconShares, _ := shares[0].Float64()
	fmt.Printf("beaconEthShares: %.2f\n", beaconShares/1e18)

	greenText := color.New(color.FgHiGreen)
	redText := color.New(color.FgHiRed)

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
		status, err := avsDirectory.AvsOperatorStatus(nil, avs.ServiceManager, operatorAddr)
		if err != nil {
			return fmt.Errorf("fetching operator status: %w", err)
		}

		registryCoordinator, err := contracts.NewRegistryCoordinator(avs.RegistryCoordinator, rpcClient)
		if err != nil {
			return fmt.Errorf("binding AVSDirectory contract: %w", err)
		}
		externalOperatorID, err := registryCoordinator.GetOperatorId(nil, operatorAddr)
		if err != nil {
			return fmt.Errorf("fetching external operatorID: %w", err)
		}
		quorumBitmap, err := registryCoordinator.GetCurrentQuorumBitmap(nil, externalOperatorID)
		if err != nil {
			return fmt.Errorf("fetching quorumBitmap: %w", err)
		}
		numQuorums, err := registryCoordinator.QuorumCount(nil)
		if err != nil {
			return fmt.Errorf("fetching QuorumCount: %w", err)
		}

		deprecatedDeets, err := avsManager.GetAvsInfo(nil, big.NewInt(operatorID), avs.RegistryCoordinator)
		if err != nil {
			return fmt.Errorf("fetching deprecated avsInfo: %w", err)
		}
		hasUploadedPubkey := deprecatedDeets.Params.PubkeyG1.X != nil && deprecatedDeets.Params.PubkeyG1.X.Cmp(big.NewInt(0)) != 0

		fmt.Println("----------------------------------------------")
		fmt.Printf("AVS: %s\n", avs.Name)
		fmt.Printf("Registered: %v\n", status == 1)
		fmt.Printf("Legacy pubkey registered: %v\n", hasUploadedPubkey)
		fmt.Printf("Quorums: ")
		for x := 0; x < int(numQuorums); x++ {
			if quorumBitmap.Bit(x) != 0 {
				greenText.Printf("%d ", x)
			} else {
				redText.Printf("%d ", x)
			}
		}
		fmt.Println()
	}

	return nil
}
