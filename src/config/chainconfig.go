package config

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SymbioticVaultType int

const (
	UNKNOWN SymbioticVaultType = iota
	WSTETH_VAULT
	LBTC_VAULT
	ETHFI_VAULT
)

// Relevant contracts to each AVS
type Config struct {

	////////////////////////////////////////////
	//            Eigenlayer                  //
	////////////////////////////////////////////

	// Eigenlayer core contracts
	AvsDirectoryAddress      common.Address
	DelegationManagerAddress common.Address
	StrategyManagerAddress   common.Address
	EigenpodManagerAddress   common.Address

	// additional Eigenlayer contracts
	EigenlayerOperatorRegistryClaimAddress common.Address // for updating stakedrop claims for operator contracts
	EigenlayerRewardsCoordinatorAddress    common.Address // for claiming AVS rewards

	// strategies (how eigenlayer tracks different LRT's + beacon eth)
	BeaconEthStrategyAddress common.Address
	WethStrategyAddress      common.Address

	// Etherfi AVS contracts
	AvsOperatorManagerAddress common.Address

	// AVS specific contracts

	// AlignedLayer
	AlignedLayerRegistryCoordinatorAddress common.Address
	AlignedLayerServiceManagerAddress      common.Address

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

	// CyberMACH
	CyberMachRegistryCoordinatorAddress common.Address
	CyberMachServiceManagerAddress      common.Address

	// EigenDA
	EigenDARegistryCoordinatorAddress common.Address
	EigenDAServiceManagerAddress      common.Address

	// EOracle
	EOracleRegistryCoordinatorAddress common.Address
	EOracleServiceManagerAddress      common.Address

	// Ethgas
	EthgasStakeRegistryAddress  common.Address
	EthgasServiceManagerAddress common.Address

	// Gasp
	GaspRegistryCoordinatorAddress common.Address
	GaspServiceManagerAddress      common.Address

	// Hyperlane
	HyperlaneStakeRegistryAddress  common.Address
	HyperlaneServiceManagerAddress common.Address

	// Lagrange ZK Coprocessor
	LagrangeZKMRServiceManagerAddress common.Address
	LagrangeZKMRStakeRegistryAddress  common.Address

	// Openlayer
	OpenlayerRegistryCoordinatorAddress common.Address
	OpenlayerServiceManagerAddress      common.Address
	OpenlayerStakeRegistryAddress       common.Address

	// Predicate
	PredicateServiceManagerAddress common.Address

	// Ungate
	UngateAVSGovernanceAddress common.Address

	// UniFi
	UniFiAvsManagerAddress common.Address

	// Witnesschain
	WitnessChainOperatorRegistryAddress common.Address
	WitnessChainWitnessHubAddress       common.Address

	////////////////////////////////////////////
	//             Symbiotic                  //
	////////////////////////////////////////////

	SymbioticVaults    map[string]SymbioticVaultConfig
	SymbioticNetworks  map[common.Address]string
	SymbioticOperators map[common.Address]string
}

type SymbioticVaultConfig struct {
	Vault     common.Address
	Delegator common.Address
	Burner    common.Address
	Asset     string
	Decimals  int
}

var Mainnet = Config{
	AvsDirectoryAddress:      common.HexToAddress("0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF"),
	DelegationManagerAddress: common.HexToAddress("0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A"),
	StrategyManagerAddress:   common.HexToAddress("0x858646372CC42E1A627fcE94aa7A7033e7CF075A"),
	EigenpodManagerAddress:   common.HexToAddress("0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338"),
	BeaconEthStrategyAddress: common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0"),
	WethStrategyAddress:      common.HexToAddress(""),

	EigenlayerOperatorRegistryClaimAddress: common.HexToAddress("0x8bB56D1CBA6273478E9B4D79F89857ac8D766eb3"),
	EigenlayerRewardsCoordinatorAddress:    common.HexToAddress("0x7750d328b314EfFa365A0402CcfD489B80B0adda"),

	AvsOperatorManagerAddress: common.HexToAddress("0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a"),

	AlignedLayerRegistryCoordinatorAddress: common.HexToAddress("0xA8CC0749b4409c3c47012323E625aEcBA92f64b9"),
	AlignedLayerServiceManagerAddress:      common.HexToAddress("0xeF2A435e5EE44B2041100EF8cbC8ae035166606c"),
	AltLayerRegistryCoordinatorAddress:     common.HexToAddress("0x561be1AB42170a19f31645F774e6e3862B2139AA"),
	AltLayerServiceManagerAddress:          common.HexToAddress("0x71a77037870169d47aad6c2C9360861A4C0df2bF"),
	ARPANodeRegistryAddress:                common.HexToAddress("0x58e39879374901e17A790af039DC9Ac06baCf25B"),
	ARPAServiceManagerAddress:              common.HexToAddress("0x1DE75EaAb2df55d467494A172652579E6FA4540E"),
	AutomataRegistryCoordinatorAddress:     common.HexToAddress("0x414696E4F7f06273973E89bfD3499e8666D63Bd4"),
	AutomataServiceManagerAddress:          common.HexToAddress("0xe5445838c475a2980e6a88054ff1514230b83aeb"),
	BrevisRegistryCoordinatorAddress:       common.HexToAddress("0x434621cfd8BcDbe8839a33c85aE2B2893a4d596C"),
	BrevisServiceManagerAddress:            common.HexToAddress("0x9FC952BdCbB7Daca7d420fA55b942405B073A89d"),
	CyberMachRegistryCoordinatorAddress:    common.HexToAddress("0x118610D207A32f10F4f7C3a1FEFac5b3327c2bad"),
	CyberMachServiceManagerAddress:         common.HexToAddress("0x1F2c296448f692af840843d993fFC0546619Dcdb"),
	EigenDARegistryCoordinatorAddress:      common.HexToAddress("0x0BAAc79acD45A023E19345c352d8a7a83C4e5656"),
	EigenDAServiceManagerAddress:           common.HexToAddress("0x870679E138bCdf293b7Ff14dD44b70FC97e12fc0"),
	EOracleRegistryCoordinatorAddress:      common.HexToAddress("0x757E6f572AfD8E111bD913d35314B5472C051cA8"),
	EOracleServiceManagerAddress:           common.HexToAddress("0x23221c5bB90C7c57ecc1E75513e2E4257673F0ef"),
	EthgasStakeRegistryAddress:             common.HexToAddress("0xfF94c9859E4b15341c1BA3e80CF80044cA2C4e76"),
	EthgasServiceManagerAddress:            common.HexToAddress("0x6201bc0A699e3b10f324204e6F8EcdD0983De227"),
	GaspRegistryCoordinatorAddress:         common.HexToAddress("0x9A986296d45C327dAa5998519AE1B3757F1e6Ba1"),
	GaspServiceManagerAddress:              common.HexToAddress("0x3aDdEb54ddd43Eb40235eC32DfA7928F28A44bb5"),
	HyperlaneStakeRegistryAddress:          common.HexToAddress("0x272CF0BB70D3B4f79414E0823B426d2EaFd48910"),
	HyperlaneServiceManagerAddress:         common.HexToAddress("0xe8E59c6C8B56F2c178f63BCFC4ce5e5e2359c8fc"),
	LagrangeZKMRServiceManagerAddress:      common.HexToAddress("0x22CAc0e6A1465F043428e8AeF737b3cb09D0eEDa"),
	LagrangeZKMRStakeRegistryAddress:       common.HexToAddress("0x8dcdCc50Cc00Fe898b037bF61cCf3bf9ba46f15C"),
	OpenlayerRegistryCoordinatorAddress:    common.HexToAddress("0x7dd7320044013f7f49B1b6D67aED10726fe6e62b"),
	OpenlayerServiceManagerAddress:         common.HexToAddress("0xF7fcff55d5FDAf2C3Bbeb140Be5e62a2c7D26Db3"),
	OpenlayerStakeRegistryAddress:          common.HexToAddress("0x8702C01EAbC9E5E376ACaB9358383DBDdCdDF018"),
	PredicateServiceManagerAddress:         common.HexToAddress("0xf6f4A30EeF7cf51Ed4Ee1415fB3bFDAf3694B0d2"),
	UngateAVSGovernanceAddress:             common.HexToAddress("0xB3e069FD6dDA251AcBDE09eDa547e0AB207016ee"),
	UniFiAvsManagerAddress:                 common.HexToAddress("0x2d86E90ED40a034C753931eE31b1bD5E1970113d"),
	WitnessChainOperatorRegistryAddress:    common.HexToAddress("0xEf1a89841fd189ba28e780A977ca70eb1A5e985D"),
	WitnessChainWitnessHubAddress:          common.HexToAddress("0xD25c2c5802198CB8541987b73A8db4c9BCaE5cC7"),

	// Symbiotic
	SymbioticVaults: map[string]SymbioticVaultConfig{
		"wstETH": {
			Vault:     common.HexToAddress("0x450a90fdEa8B87a6448Ca1C87c88Ff65676aC45b"),
			Delegator: common.HexToAddress("0xd6c4b4267BFB908BBdf8C9BDa7d0Ae517aA145b0"),
			Burner:    common.HexToAddress("0xD1e54Bf62e089287a9514d581c6b80aA75B81d15"),
			Asset:     "wstETH",
			Decimals:  18,
		},
		"LBTC": {
			Vault:     common.HexToAddress("0xd4E20ECA1f996Dab35883dC0AD5E3428AF888D45"),
			Delegator: common.HexToAddress("0xA32E5868713CBeb1880578F5626ED53cc3E1A2fD"),
			Burner:    common.HexToAddress("0xeb3d85be284d6F83daf1338B22d834ECaF1099fE"),
			Asset:     "LBTC",
			Decimals:  8,
		},
		"ETHFI": {
			Vault:     common.HexToAddress("0x2Bcfa0283C92b7845ECE12cEaDc521414BeF1067"),
			Delegator: common.HexToAddress("0xCcA347C3d3CAFB909A7506eaD727805D38a82EE3"),
			Burner:    common.HexToAddress("0x2242E802b5AAADcc7c4929EF77F0E530BCb5Ce3f"),
			Asset:     "ETHFI",
			Decimals:  18,
		},
	},
	SymbioticOperators: map[common.Address]string{
		common.HexToAddress("0x087c25f83ED20bda587CFA035ED0c96338D4660f"): "P2P",
		common.HexToAddress("0xfF645D02C79141424fb4F8bBB5e494f05067c08B"): "Kiln",
		common.HexToAddress("0x51B6D824bd35AeD4FD1a9E253E41Dc7C9feeFa30"): "Pier 2",
		common.HexToAddress("0x888F7454E65D213C89bA92020e0d716428898f7f"): "Blockdaemon",
	},
	SymbioticNetworks: map[common.Address]string{
		common.HexToAddress("0xAD12e74847d6D1487A6a3A6b75D1f509f3F627e8"): "CapX",
		common.HexToAddress("0x759D4335cb712aa188935C2bD3Aa6D205aC61305"): "Cycle Network",
		common.HexToAddress("0x8560C667Ae72F28D09465B342A480daB28821f6b"): "Ditto",
		common.HexToAddress("0x59cf937Ea9FA9D7398223E3aA33d92F7f5f986A2"): "Hyperlane",
		common.HexToAddress("0xE3a148b25Cca54ECCBD3A4aB01e235D154f03eFa"): "Hyve",
		common.HexToAddress("0x3a7B173124DcFeCff1847FF7f8f56e72ABE02340"): "Marlin",
		common.HexToAddress("0xfCa0128A19A5c06b0148c27ee7623417a11BaAbd"): "Radius",
		common.HexToAddress("0xcf128e88e11507abad12a7624a34e3d22f731abc"): "Router",
		common.HexToAddress("0x5112EbA9bc2468Bb5134CBfbEAb9334EdaE7106a"): "Symbiosis",
	},
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
