package predicate

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/predicate"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/urfave/cli/v3"
)

var predicateAPI *predicate.API
var etherfiAPI *etherfi.API

var PredicateCmd = &cli.Command{
	Name:   "predicate",
	Usage:  "various actions related to managing Predicate operators",
	Before: prepareCmd,
	Commands: []*cli.Command{
		PrepareRegistrationCmd,
		RegisterCmd,
		RotateSigningKeyCmd,
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
	predicateAPI = predicate.New(cfg, rpcClient)

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
			Name:     "avs-signer",
			Usage:    "Address of predicate specific ecdsa signing key",
			Required: true,
		},
	},
}

func handlePrepareRegistration(ctx context.Context, cli *cli.Command) error {

	// parse cli input
	operatorID := cli.Int("operator-id")
	avsSignerAddress := common.HexToAddress(cli.String("avs-signer"))

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return predicateAPI.PrepareRegistration(operator, avsSignerAddress)
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
	var input predicate.RegistrationInfo
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
	if input.AvsSigner == common.HexToAddress("0x00") {
		return fmt.Errorf("invalid registration input, missing avsSigner")
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

	return predicateAPI.RegisterOperator(operator, input, signingKey)
}

var RotateSigningKeyCmd = &cli.Command{
	Name:   "rotate-signing-key",
	Usage:  "(Admin) Rotate the predicate signing key for an operator",
	Action: handleRotateSigningKey,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "old-signer",
			Usage:    "Address of current predicate specific ecdsa signing key",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "new-signer",
			Usage:    "Address of new predicate specific ecdsa signing key",
			Required: true,
		},
	},
}

func handleRotateSigningKey(ctx context.Context, cli *cli.Command) error {

	// parse cli input
	operatorID := cli.Int("operator-id")
	oldSignerAddress := common.HexToAddress(cli.String("old-signer"))
	newSignerAddress := common.HexToAddress(cli.String("new-signer"))

	if operatorID == 0 {
		return fmt.Errorf("invalid input, missing operator-id")
	}
	if oldSignerAddress == common.HexToAddress("0x00") {
		return fmt.Errorf("invalid input, missing old-signer")
	}
	if newSignerAddress == common.HexToAddress("0x00") {
		return fmt.Errorf("invalid input, missing new-signer")
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return predicateAPI.RotateSigningKey(operator, oldSignerAddress, newSignerAddress)
}
