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

	// strategies (how eigenlayer tracks different LRT's + beacon eth)
	BeaconEthStrategyAddress common.Address
	WethStrategyAddress      common.Address

	// Etherfi AVS contracts
	AvsOperatorManagerAddress common.Address

	// EigenDA
	EigenDARegistryCoordinatorAddress common.Address
	EigenDAServiceManagerAddress      common.Address

	// EOracle
	EOracleRegistryCoordinatorAddress common.Address
	EOracleServiceManagerAddress      common.Address

	// Brevis
	BrevisRegistryCoordinatorAddress common.Address
	BrevisServiceManagerAddress      common.Address

	// Lagrange
	LagrangeServiceAddress common.Address

	// Witnesschain
	WitnessChainOperatorRegistryAddress common.Address
	WitnessChainWitnessHubAddress       common.Address
}

var Mainnet = Config{
	AvsDirectoryAddress:      common.HexToAddress("0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF"),
	DelegationManagerAddress: common.HexToAddress("0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A"),
	StrategyManagerAddress:   common.HexToAddress("0x858646372CC42E1A627fcE94aa7A7033e7CF075A"),
	EigenpodManagerAddress:   common.HexToAddress("0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338"),
	BeaconEthStrategyAddress: common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0"),
	WethStrategyAddress:      common.HexToAddress(""),

	AvsOperatorManagerAddress:           common.HexToAddress("0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a"),
	EigenDARegistryCoordinatorAddress:   common.HexToAddress("0x0BAAc79acD45A023E19345c352d8a7a83C4e5656"),
	EigenDAServiceManagerAddress:        common.HexToAddress("0x870679E138bCdf293b7Ff14dD44b70FC97e12fc0"),
	EOracleRegistryCoordinatorAddress:   common.HexToAddress("0x757E6f572AfD8E111bD913d35314B5472C051cA8"),
	EOracleServiceManagerAddress:        common.HexToAddress("0x23221c5bB90C7c57ecc1E75513e2E4257673F0ef"),
	BrevisRegistryCoordinatorAddress:    common.HexToAddress("0x434621cfd8BcDbe8839a33c85aE2B2893a4d596C"),
	BrevisServiceManagerAddress:         common.HexToAddress("0x9FC952BdCbB7Daca7d420fA55b942405B073A89d"),
	LagrangeServiceAddress:              common.HexToAddress("0x35F4f28A8d3Ff20EEd10e087e8F96Ea2641E6AA2"),
	WitnessChainOperatorRegistryAddress: common.HexToAddress("0xEf1a89841fd189ba28e780A977ca70eb1A5e985D"),
	WitnessChainWitnessHubAddress:       common.HexToAddress("0xD25c2c5802198CB8541987b73A8db4c9BCaE5cC7"),
}

var Holesky = Config{
	AvsDirectoryAddress:      common.HexToAddress("0x055733000064333CaDDbC92763c58BF0192fFeBf"),
	DelegationManagerAddress: common.HexToAddress("0xA44151489861Fe9e3055d95adC98FbD462B948e7"),
	StrategyManagerAddress:   common.HexToAddress("0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6"),
	EigenpodManagerAddress:   common.HexToAddress("0x30770d7E3e71112d7A6b7259542D1f680a70e315"),
	BeaconEthStrategyAddress: common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0"),
	WethStrategyAddress:      common.HexToAddress("0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9"),

	AvsOperatorManagerAddress:           common.HexToAddress("0xdf9679e8bfce22ae503fd2726cb1218a18cd8bf4"),
	EigenDARegistryCoordinatorAddress:   common.HexToAddress("0x53012C69A189cfA2D9d29eb6F19B32e0A2EA3490"),
	EigenDAServiceManagerAddress:        common.HexToAddress("0xD4A7E1Bd8015057293f0D0A557088c286942e84b"),
	BrevisRegistryCoordinatorAddress:    common.HexToAddress("0x0dB4ceE042705d47Ef6C0818E82776359c3A80Ca"),
	BrevisServiceManagerAddress:         common.HexToAddress("0x7A46219950d8a9bf2186549552DA35Bf6fb85b1F"),
	WitnessChainOperatorRegistryAddress: common.HexToAddress("0x708CBDDdab358c1fa8efB82c75bB4a116F316Def"),
	WitnessChainWitnessHubAddress:       common.HexToAddress("0xa987EC494b13b21A8a124F8Ac03c9F530648C87D"),
}

func ConfigForChain(chainID int64) (*Config, error) {

	var cfg Config
	switch chainID {
	case 1:
		cfg = Mainnet
	case 17000:
		cfg = Holesky
	default:
		return nil, fmt.Errorf("unimplemented chain: %d", chainID)
	}

	return &cfg, nil
}

func AutodetectConfig(rpcClient *ethclient.Client) (*Config, error) {
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("querying chainID from RPC: %w", err)
	}
	return ConfigForChain(chainID.Int64())
}
