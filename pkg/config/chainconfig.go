package config

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Relevant contracts to each AVS
type Config struct {

	// Eigenlayer core contracts
	AvsDirectoryAddress      common.Address
	DelegationManagerAddress common.Address
	StrategyManagerAddress   common.Address
	EigenpodManagerAddress   common.Address

	// EigenDA
	EigenDARegistryCoordinatorAddress common.Address
	EigenDAServiceManagerAddress      common.Address

	SignMessageLibAddress common.Address
}

var Mainnet = Config{
	AvsDirectoryAddress:      common.HexToAddress("0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF"),
	DelegationManagerAddress: common.HexToAddress("0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A"),
	StrategyManagerAddress:   common.HexToAddress("0x858646372CC42E1A627fcE94aa7A7033e7CF075A"),
	EigenpodManagerAddress:   common.HexToAddress("0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338"),

	EigenDARegistryCoordinatorAddress: common.HexToAddress("0x0BAAc79acD45A023E19345c352d8a7a83C4e5656"),
	EigenDAServiceManagerAddress:      common.HexToAddress("0x870679E138bCdf293b7Ff14dD44b70FC97e12fc0"),

	SignMessageLibAddress: common.HexToAddress("0xA65387F16B013cf2Af4605Ad8aA5ec25a2cbA3a2"),
}

var Holesky = Config{
	AvsDirectoryAddress:      common.HexToAddress("0x055733000064333CaDDbC92763c58BF0192fFeBf"),
	DelegationManagerAddress: common.HexToAddress("0xA44151489861Fe9e3055d95adC98FbD462B948e7"),
	StrategyManagerAddress:   common.HexToAddress("0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6"),
	EigenpodManagerAddress:   common.HexToAddress("0x30770d7E3e71112d7A6b7259542D1f680a70e315"),

	EigenDARegistryCoordinatorAddress: common.HexToAddress("0x53012C69A189cfA2D9d29eb6F19B32e0A2EA3490"),
	EigenDAServiceManagerAddress:      common.HexToAddress("0xD4A7E1Bd8015057293f0D0A557088c286942e84b"),

	SignMessageLibAddress: common.HexToAddress("0xA65387F16B013cf2Af4605Ad8aA5ec25a2cbA3a2"),
}

func ConfigForChain(chainID int64) (Config, error) {

	var cfg Config
	switch chainID {
	case 1:
		cfg = Mainnet
	case 17000:
		cfg = Holesky
	default:
		return Config{}, fmt.Errorf("unimplemented chain: %d", chainID)
	}

	return cfg, nil
}

func AutodetectConfig(rpcClient *ethclient.Client) (Config, error) {
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return Config{}, fmt.Errorf("querying chainID from RPC: %w", err)
	}
	return ConfigForChain(chainID.Int64())
}
