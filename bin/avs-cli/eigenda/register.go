package eigenda

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/eigenda"
	"github.com/urfave/cli/v3"
)

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
	var input eigenda.RegistrationInfo
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

	return eigendaAPI.RegisterOperator(operator, input, signingKey)
}

var DeregisterCmd = &cli.Command{
	Name:   "deregister",
	Usage:  "(Admin) Deregister operator from specified quorums",
	Action: handleDeregister,
	Flags: []cli.Flag{
		&cli.IntSliceFlag{
			Name:     "quorums",
			Usage:    "which quorums to deregister from i.e. 0,1",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
	},
}

func handleDeregister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	quorums := cli.IntSlice("quorums")
	operatorID := cli.Int("operator-id")

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return eigendaAPI.DeregisterOperator(operator, quorums)
}
