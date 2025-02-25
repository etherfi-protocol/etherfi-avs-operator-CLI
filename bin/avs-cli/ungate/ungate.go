package ungate

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/ungate"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/othentic"
	"github.com/urfave/cli/v3"
)

var ungateAPI *ungate.API
var etherfiAPI *etherfi.API

var UngateCmd = &cli.Command{
	Name:   "ungate",
	Usage:  "various actions related to managing Ungate operators",
	Before: prepareCmd,
	Commands: []*cli.Command{
		RegisterCmd,
		UnregisterCmd,
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
	ungateAPI = ungate.New(cfg, rpcClient)

	return nil
}

var RegisterCmd = &cli.Command{
	Name:   "register",
	Usage:  "(Admin) Register target operator to the AVS",
	Action: handleRegister,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "othentic-output",
			Usage:    "path to output from the othentic CLI",
			Required: true,
		},
	},
}

func handleRegister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	operatorID := cli.Int("operator-id")
	othenticOutputPath := cli.String("othentic-output")

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	// read the data the operator produced with the othentic cli
	buf, err := ioutil.ReadFile(othenticOutputPath)
	if err != nil {
		return fmt.Errorf("failed to load othentic cli file: %w", err)
	}
	var othenticOutput othentic.CLIOperatorData
	if err := json.Unmarshal(buf, &othenticOutput); err != nil {
		return fmt.Errorf("failed to parse othentic cli data: %w", err)
	}

	// sanity check since the first step is using othentics external tooling
	if operator.Address != common.HexToAddress(othenticOutput.Operator) {
		return fmt.Errorf("operator mismatch: %s - %s", operator.Address, othenticOutput.Operator)
	}

	// load eip-1271 admin signing key
	signingKey, err := crypto.HexToECDSA(os.Getenv("ADMIN_1271_SIGNING_KEY"))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	whitelistEnabled := true // TODO: this could change

	return ungateAPI.RegisterOperator(operator, &othenticOutput, signingKey, whitelistEnabled)
}

var UnregisterCmd = &cli.Command{
	Name:   "unregister",
	Usage:  "(Admin) Unregister target operator from the AVS",
	Action: handleUnregister,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
	},
}

func handleUnregister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	operatorID := cli.Int("operator-id")

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return ungateAPI.UnregisterOperator(operator)
}
