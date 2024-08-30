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

	// AVS specific contracts

	// AltLayer
	AltLayerRegistryCoordinatorAddress common.Address
	AltLayerServiceManagerAddress      common.Address

	//ARPA
	ARPANodeRegistryAddress   common.Address
	ARPAServiceManagerAddress common.Address

	// Automata
	AutomataRegistryCoordinatorAddress common.Address
	AutomataServiceManagerAddress      common.Address

	// Brevis
	BrevisRegistryCoordinatorAddress common.Address
	BrevisServiceManagerAddress      common.Address

	// EigenDA
	EigenDARegistryCoordinatorAddress common.Address
	EigenDAServiceManagerAddress      common.Address

	// EOracle
	EOracleRegistryCoordinatorAddress common.Address
	EOracleServiceManagerAddress      common.Address

	// Hyperlane
	HyperlaneStakeRegistryAddress  common.Address
	HyperlaneServiceManagerAddress common.Address

	// Lagrange State Committee
	LagrangeServiceAddress   common.Address
	LagrangeCommitteeAddress common.Address

	// Lagrange ZK Coprocessor
	LagrangeZKMRServiceManagerAddress common.Address
	LagrangeZKMRStakeRegistryAddress  common.Address

	// Openlayer
	OpenlayerRegistryCoordinatorAddress common.Address
	OpenlayerServiceManagerAddress      common.Address

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

	AvsOperatorManagerAddress: common.HexToAddress("0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a"),

	AltLayerRegistryCoordinatorAddress:  common.HexToAddress("0x561be1AB42170a19f31645F774e6e3862B2139AA"),
	AltLayerServiceManagerAddress:       common.HexToAddress("0x71a77037870169d47aad6c2C9360861A4C0df2bF"),
	ARPANodeRegistryAddress:             common.HexToAddress("0x58e39879374901e17A790af039DC9Ac06baCf25B"),
	ARPAServiceManagerAddress:           common.HexToAddress("0x1DE75EaAb2df55d467494A172652579E6FA4540E"),
	AutomataRegistryCoordinatorAddress:  common.HexToAddress("0x414696E4F7f06273973E89bfD3499e8666D63Bd4"),
	AutomataServiceManagerAddress:       common.HexToAddress("0xe5445838c475a2980e6a88054ff1514230b83aeb"),
	BrevisRegistryCoordinatorAddress:    common.HexToAddress("0x434621cfd8BcDbe8839a33c85aE2B2893a4d596C"),
	BrevisServiceManagerAddress:         common.HexToAddress("0x9FC952BdCbB7Daca7d420fA55b942405B073A89d"),
	EigenDARegistryCoordinatorAddress:   common.HexToAddress("0x0BAAc79acD45A023E19345c352d8a7a83C4e5656"),
	EigenDAServiceManagerAddress:        common.HexToAddress("0x870679E138bCdf293b7Ff14dD44b70FC97e12fc0"),
	EOracleRegistryCoordinatorAddress:   common.HexToAddress("0x757E6f572AfD8E111bD913d35314B5472C051cA8"),
	EOracleServiceManagerAddress:        common.HexToAddress("0x23221c5bB90C7c57ecc1E75513e2E4257673F0ef"),
	HyperlaneStakeRegistryAddress:       common.HexToAddress("0x272CF0BB70D3B4f79414E0823B426d2EaFd48910"),
	HyperlaneServiceManagerAddress:      common.HexToAddress("0xe8E59c6C8B56F2c178f63BCFC4ce5e5e2359c8fc"),
	LagrangeServiceAddress:              common.HexToAddress("0x35F4f28A8d3Ff20EEd10e087e8F96Ea2641E6AA2"),
	LagrangeCommitteeAddress:            common.HexToAddress("0xECc22f3EcD0EFC8aD77A78ad9469eFbc44E746F5"),
	LagrangeZKMRServiceManagerAddress:   common.HexToAddress("0x22CAc0e6A1465F043428e8AeF737b3cb09D0eEDa"),
	LagrangeZKMRStakeRegistryAddress:    common.HexToAddress("0x8dcdCc50Cc00Fe898b037bF61cCf3bf9ba46f15C"),
	OpenlayerRegistryCoordinatorAddress: common.HexToAddress("0x7dd7320044013f7f49B1b6D67aED10726fe6e62b"),
	OpenlayerServiceManagerAddress:      common.HexToAddress("0xF7fcff55d5FDAf2C3Bbeb140Be5e62a2c7D26Db3"),
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

	AvsOperatorManagerAddress: common.HexToAddress("0xdf9679e8bfce22ae503fd2726cb1218a18cd8bf4"),

	AltLayerRegistryCoordinatorAddress:  common.HexToAddress("0x1eA7D160d325B289bF981e0D7aB6Bf3261a0FFf2"),
	AltLayerServiceManagerAddress:       common.HexToAddress("0xae9a4497dee2540daf489beddb0706128a99ec63"),
	ARPANodeRegistryAddress:             common.HexToAddress("0x76416A37BfDCb188a48C01a145D1b8d5794a98bD"),
	ARPAServiceManagerAddress:           common.HexToAddress("0xd36b6e5eee8311d7bffb2f3bb33301a1ab7de101"),
	AutomataRegistryCoordinatorAddress:  common.HexToAddress("0x62c715575cE3Ad7C5a43aA325b881c70564f2215"),
	AutomataServiceManagerAddress:       common.HexToAddress("0x4665Af665df5703445645D243f0FD63eD3b9D132"),
	BrevisRegistryCoordinatorAddress:    common.HexToAddress("0x0dB4ceE042705d47Ef6C0818E82776359c3A80Ca"),
	BrevisServiceManagerAddress:         common.HexToAddress("0x7A46219950d8a9bf2186549552DA35Bf6fb85b1F"),
	EigenDARegistryCoordinatorAddress:   common.HexToAddress("0x53012C69A189cfA2D9d29eb6F19B32e0A2EA3490"),
	EigenDAServiceManagerAddress:        common.HexToAddress("0xD4A7E1Bd8015057293f0D0A557088c286942e84b"),
	WitnessChainOperatorRegistryAddress: common.HexToAddress("0x708CBDDdab358c1fa8efB82c75bB4a116F316Def"),
	WitnessChainWitnessHubAddress:       common.HexToAddress("0xa987EC494b13b21A8a124F8Ac03c9F530648C87D"),
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
