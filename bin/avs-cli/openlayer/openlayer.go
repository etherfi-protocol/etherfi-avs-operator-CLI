package openlayer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/openlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/keystore"
	"github.com/urfave/cli/v3"
)

var openlayerAPI *openlayer.API
var etherfiAPI *etherfi.API

var OpenlayerCmd = &cli.Command{
	Name:   "openlayer",
	Usage:  "various actions related to managing Openlayer operators",
	Before: prepareOpenlayerCmd,
	Commands: []*cli.Command{
		PrepareRegistrationCmd,
		RegisterCmd,
		UpdateSignerCmd,
	},
}

// run before any subcommand executes
func prepareOpenlayerCmd(ctx context.Context, cmd *cli.Command) error {
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

	// make globally accessible by all sub commands
	etherfiAPI = etherfi.New(cfg, rpcClient)
	openlayerAPI = openlayer.New(cfg, rpcClient)

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
			Name:     "bls-keystore",
			Usage:    "path to bls keystore file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "bls-password",
			Usage:    "password for encrypted keystore file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "signer-address",
			Usage:    "address of the generated signer ecdsa key",
			Required: true,
		},
		&cli.IntSliceFlag{
			Name:     "quorums",
			Usage:    "which quorums to register for i.e. 0,1",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "socket",
			Usage:    "socket",
			Required: true,
		},
	},
}

func handlePrepareRegistration(ctx context.Context, cli *cli.Command) error {

	// parse cli input
	operatorID := cli.Int("operator-id")
	blsKeyFile := cli.String("bls-keystore")
	blsKeyPassword := cli.String("bls-password")
	signerAddress := common.HexToAddress(cli.String("signer-address"))
	quorums := cli.IntSlice("quorums")
	socket := cli.String("socket")

	// decrypt and load bls key from keystore
	ks := keystore.NewKeystoreV3()
	keyPair, err := ks.LoadBLS(blsKeyFile, blsKeyPassword)
	if err != nil {
		return fmt.Errorf("loading bls keystore: %w", err)
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return openlayerAPI.PrepareRegistration(operator, keyPair, signerAddress, socket, quorums)
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
	var input openlayer.RegistrationInfo
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
	if input.Socket == "" {
		return fmt.Errorf("invalid registration input, missing socket")
	}
	if input.SignerAddress == common.HexToAddress("0x0") {
		return fmt.Errorf("invalid registration input, missing signerAddress")
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

	return openlayerAPI.RegisterOperator(operator, input, signingKey)
}

var UpdateSignerCmd = &cli.Command{
	Name:   "update-signer",
	Usage:  "(Admin) Update the signer address (this is within openlayer, unerelated to the EIP-1271 signer)",
	Action: handleUpdateSigner,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "signer",
			Usage:    "signer address",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "operator ID",
			Required: true,
		},
	},
}

func handleUpdateSigner(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	signerAddr := common.HexToAddress(cli.String("signer"))
	operatorID := cli.Int("operator-id")

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return openlayerAPI.UpdateSignerAddress(operator, signerAddr)
}
