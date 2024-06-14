package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

type cmdContext struct {
	rpcClient *ethclient.Client

	AVSDirectory               *contracts.AVSDirectory
	EigenDARegistryCoordinator *contracts.RegistryCoordinator
	//EigenDA
}

// playing with helper that will initialize all RPC and contract bindings before command runs
func setupCommand(ctx context.Context, cmd *cli.Command) error {
	//rpcClient
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return fmt.Errorf("dialing RPC: %w", err)
	}

	// eigenDAStakeRegistry
	// 0x006124Ae7976137266feeBFb3F4D2BE4C073139D
	// indexRegistry
	// 0xBd35a7a1CDeF403a6a99e4E8BA0974D198455030
	// blsApkRegistry
	// 0x00A5Fd09F6CeE6AE9C8b0E5e33287F7c82880505

	rpcClient = client
	return nil
}
