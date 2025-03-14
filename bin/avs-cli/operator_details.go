package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/eigenlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/urfave/cli/v3"
)

// TODO(Dave): dynamically scan via events
const TotalOperators = 15

var operatorDetailsCmd = &cli.Command{
	Name:   "operator-details",
	Usage:  "Information about the current configuration of target operator",
	Action: handleOperatorDetails,
	//Before: setupCommand,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "all",
			Usage:    "list details of all operators",
			Required: false,
		},
	},
}

func handleOperatorDetails(ctx context.Context, cmd *cli.Command) error {

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

	delegationManager, err := eigenlayer.NewDelegationManager(cfg.DelegationManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding delegationManager: %w", err)
	}
	shares, err := delegationManager.GetOperatorShares(&bind.CallOpts{}, operator.Address, []common.Address{cfg.BeaconEthStrategyAddress, cfg.EigenStrategyAddress})
	if err != nil {
		return fmt.Errorf("fetching delegated shares: %w", err)
	}
	for _, s := range shares {
		v, _ := weiToEther(s).Float64()
		fmt.Printf("%.04f\n", v)
	}
	return nil

}

func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}

func operatorDetails(operatorID int64, rpcClient *ethclient.Client) error {

	return nil
	// TODO(Dave): redo after refactor

	/*
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
		avsManager, err := etherfi.NewAvsOperatorManager(cfg.AvsOperatorManagerAddress, rpcClient)
		if err != nil {
			return fmt.Errorf("binding OperatorManager contract: %w", err)
		}
		avsDirectory, err := eigenlayer.NewAvsDirectory(cfg.AvsDirectoryAddress, rpcClient)
		if err != nil {
			return fmt.Errorf("binding AVSDirectory contract: %w", err)
		}
		delegationManager, err := eigenlayer.NewDelegationManager(cfg.DelegationManagerAddress, rpcClient)
		if err != nil {
			return fmt.Errorf("binding delegationManager: %w", err)
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
		shares, err := delegationManager.GetOperatorShares(&bind.CallOpts{}, operatorAddr, []common.Address{cfg.BeaconEthStrategyAddress})
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
			{Name: "EigenDA", ServiceManager: cfg.EigenDAServiceManagerAddress, RegistryCoordinator: cfg.EigenDARegistryCoordinatorAddress},
			{Name: "Brevis", ServiceManager: cfg.BrevisServiceManagerAddress, RegistryCoordinator: cfg.BrevisRegistryCoordinatorAddress},
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
	*/
}
