package lagrangesc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	lagrangesc "github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/lagrangeSC"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/keystore"
	"github.com/urfave/cli/v3"
)

var lagrangeAPI *lagrangesc.API
var etherfiAPI *etherfi.API

const (
	OptimismChainID = 10
	BaseChainID     = 8453
	ArbitrumChainID = 42161
)

var LagrangeSCCmd = &cli.Command{
	Name:   "lagrangeSC",
	Usage:  "various actions related to managing Lagrange operators",
	Before: prepareCmd,
	Commands: []*cli.Command{
		PrepareRegistrationCmd,
		RegisterCmd,
		SubscribeCmd,
		UpdateBLSPubkeyCmd,
	},
}

// run before any subcommand executes
func prepareCmd(ctx context.Context, cmd *cli.Command) error {
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

	// load all required addresses for this chain and bind applicable contracts
	cfg, err := config.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	// make globally accessible by all sub commands
	etherfiAPI = etherfi.New(cfg, rpcClient)
	lagrangeAPI = lagrangesc.New(cfg, rpcClient)

	return nil
}

var PrepareRegistrationCmd = &cli.Command{
	Name:   "prepare-registration",
	Usage:  "(Node Operator) gather all inputs required to register for avs",
	Action: handlePrepareRegistration,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "signer-address",
			Usage:    "Address of lagrange specific ecdsa signing key",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "bls-keystore",
			Usage:    "path to bls keystore file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "bls-password",
			Usage:    "password for encrypted keystore file",
			Required: true,
		},
	},
}

func handlePrepareRegistration(ctx context.Context, cli *cli.Command) error {

	// parse cli input
	operatorID := cli.Int("operator-id")
	signerAddress := common.HexToAddress(cli.String("signer-address"))
	blsKeyFile := cli.String("bls-keystore")
	blsKeyPassword := cli.String("bls-password")

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	// decrypt and load bls key from keystore
	ks := keystore.NewKeystoreV3()
	keyPair, err := ks.LoadBLS(blsKeyFile, blsKeyPassword)
	if err != nil {
		return fmt.Errorf("loading bls keystore: %w", err)
	}

	return lagrangeAPI.PrepareRegistration(operator, signerAddress, keyPair)
}

var RegisterCmd = &cli.Command{
	Name:   "register",
	Usage:  "(Admin) Register target operator to the AVS",
	Action: handleRegister,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "registration-input",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
	},
}

func handleRegister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	inputFilepath := cli.String("registration-input")

	// read input file with required registration data
	var input lagrangesc.RegistrationInfo
	buf, err := os.ReadFile(inputFilepath)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}
	if err := json.Unmarshal(buf, &input); err != nil {
		return fmt.Errorf("parsing registration input: %w", err)
	}
	if input.OperatorID == 0 {
		return fmt.Errorf("invalid registration input, missing operatorID")
	}
	if input.SignerAddress == common.HexToAddress("0x00") {
		return fmt.Errorf("invalid registration input, missing SignerAddress")
	}
	if input.BLSKeyWithProof == nil {
		return fmt.Errorf("invalid registration input, missing BLSKeyWithProof")
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(input.OperatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	// load eip-1271 admin signing key
	signingKey, err := crypto.HexToECDSA(os.Getenv("ADMIN_1271_SIGNING_KEY"))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	return lagrangeAPI.RegisterOperator(operator, input, signingKey)
}

var SubscribeCmd = &cli.Command{
	Name:   "subscribe",
	Usage:  "(Admin) Subscribe the target operator to specified chains",
	Action: handleSubscribe,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "registration-input",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
		&cli.IntSliceFlag{
			Name:     "chain-ids",
			Usage:    "which chainIDs to subscribe to. Defaults to all supported chains",
			Value:    []int64{OptimismChainID, BaseChainID, ArbitrumChainID},
			Required: false,
		},
	},
}

func handleSubscribe(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	inputFilepath := cli.String("registration-input")
	chainIDs := cli.IntSlice("chain-ids")

	// read input file with required registration data
	var input lagrangesc.RegistrationInfo
	buf, err := os.ReadFile(inputFilepath)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}
	if err := json.Unmarshal(buf, &input); err != nil {
		return fmt.Errorf("parsing registration input: %w", err)
	}
	if input.OperatorID == 0 {
		return fmt.Errorf("invalid registration input, missing operatorID")
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(input.OperatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return lagrangeAPI.SubscribeToChains(operator, chainIDs)
}

var UpdateBLSPubkeyCmd = &cli.Command{
	Name:   "update-bls-pubkey",
	Usage:  "(Admin) Change the registered BLS pubkey for the operator",
	Action: handleUpdateBLSPubkey,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "registration-input",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
	},
}

func handleUpdateBLSPubkey(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	inputFilepath := cli.String("registration-input")

	// read input file with required registration data
	var input lagrangesc.RegistrationInfo
	buf, err := os.ReadFile(inputFilepath)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}
	if err := json.Unmarshal(buf, &input); err != nil {
		return fmt.Errorf("parsing registration input: %w", err)
	}
	if input.OperatorID == 0 {
		return fmt.Errorf("invalid registration input, missing operatorID")
	}
	if input.SignerAddress == common.HexToAddress("0x00") {
		return fmt.Errorf("invalid registration input, missing SignerAddress")
	}
	if input.BLSKeyWithProof == nil {
		return fmt.Errorf("invalid registration input, missing BLSKeyWithProof")
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(input.OperatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return lagrangeAPI.UpdateBLSPubkey(operator, input)
}
