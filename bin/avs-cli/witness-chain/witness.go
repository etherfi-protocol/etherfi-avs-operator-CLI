package witness

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/gnosis"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/etherfi"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/witnesschain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var witnessAPI *witnesschain.WitnessChain
var etherfiAPI *etherfi.API

var WitnessCmd = &cli.Command{
	Name:   "witness-chain",
	Usage:  "various actions related to managing Witness-Chain operators",
	Before: prepareWitnessChainCmd,
	Commands: []*cli.Command{
		WitnessRegisterToAvsCmd,
		WitnessPrepareRegistrationCmd,
		WitnessRegisterWatchtowerCmd,
	},
}

func prepareWitnessChainCmd(ctx context.Context, cmd *cli.Command) error {
	// try to load RPC_URL from env or flags
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = cmd.String("rpc-url")
	}

	if rpcURL == "" {
		return fmt.Errorf("must set env var $RPC_URL or use --rpc-url flag")
	}

	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing RPC: %w", err)
	}

	cfg, err := bindings.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	witnessHub, err := witnesschain.NewWitnessChainWitnessHub(cfg.WitnessChainWitnessHubAddress, rpcClient)
	if err != nil {
		return err
	}
	operatorRegistry, err := witnesschain.NewWitnessChainOperatorRegistry(cfg.WitnessChainOperatorRegistryAddress, rpcClient)
	if err != nil {
		return err
	}
	avsOperatorManager, err := contracts.NewAvsOperatorManager(cfg.AvsOperatorManagerAddress, rpcClient)
	if err != nil {
		return err
	}

	etherfiAPI = &etherfi.API{
		Client:                    rpcClient,
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		AvsOperatorManager:        avsOperatorManager,
	}

	// make globally accessible by all commands
	witnessAPI = &witnesschain.WitnessChain{
		Client:                  rpcClient,
		OperatorRegistryAddress: cfg.WitnessChainOperatorRegistryAddress,
		OperatorRegistry:        operatorRegistry,
		WitnessHubAddress:       cfg.WitnessChainWitnessHubAddress,
		WitnessHub:              witnessHub,
	}

	return nil
}

var WitnessRegisterWatchtowerCmd = &cli.Command{
	Name:   "register-watchtower",
	Action: handleRegisterWatchtower,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "registration-input",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
	},
}

func handleRegisterWatchtower(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	rpcURL := cli.String("rpc-url")
	inputFilepath := cli.String("registration-input")

	// connect to RPC node
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}
	cfg, err := bindings.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	// read input file with required witnesschain data
	var input RegistrationInput
	buf, err := os.ReadFile(inputFilepath)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}
	if err := json.Unmarshal(buf, &input); err != nil {
		return fmt.Errorf("parsing registration input: %w", err)
	}
	if input.OperatorID == 0 {
		return fmt.Errorf("invalid registration input")
	}
	if strings.HasPrefix(input.WatchtowerSignature, "0x") {
		input.WatchtowerSignature = input.WatchtowerSignature[2:]
	}
	watchtowerSignature, err := hex.DecodeString(input.WatchtowerSignature)
	if err != nil {
		return fmt.Errorf("invalid watchtower signature")
	}

	return witnessRegisterWatchtower(
		input.OperatorID,
		input.WatchtowerAddress,
		watchtowerSignature,
		input.WatchtowerSignatureExpiry,
		cfg,
		rpcClient,
	)
}

func witnessRegisterWatchtower(
	operatorID int64,
	watchtowerAddress common.Address,
	watchtowerSignature []byte,
	watchtowerSignatureExpiry *big.Int,
	cfg *bindings.Config,
	rpcClient *ethclient.Client,
) error {

	witnessABI, err := witnesschain.WitnessChainOperatorRegistryMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	// pack operatorRegistry.registerWatchtowerAsOperator()
	input, err := witnessABI.Pack("registerWatchtowerAsOperator", watchtowerAddress, watchtowerSignatureExpiry, watchtowerSignature)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}
	fmt.Printf("subcall: 0x%s\n\n", hex.EncodeToString(input))

	adminCall, err := bindings.PackForwardCallForAdmin(operatorID, input, cfg.WitnessChainOperatorRegistryAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}
	fmt.Printf("admincall: 0x%s\n\n", hex.EncodeToString(adminCall))

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, cfg.AvsOperatorManagerAddress, fmt.Sprintf("witness-chain-register-watchtower-%d", operatorID))
	fmt.Printf("gnosis:\n%s\n", batch.PrettyPrint())

	return nil
}
