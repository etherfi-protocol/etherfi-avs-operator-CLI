package witness

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/witnesschain"
	"github.com/urfave/cli/v3"
)

var WitnessRegisterWatchtowerCmd = &cli.Command{
	Name:   "register-watchtower",
	Usage:  "(Admin) Register a watchtower to target operator",
	Action: handleRegisterWatchtower,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "registration-input",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
	},
}

func handleRegisterWatchtower(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	inputFilepath := cli.String("registration-input")

	// read input file with required witnesschain data
	var input witnesschain.RegistrationInfo
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

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(input.OperatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return witnessAPI.RegisterWatchtower(operator, &input)
}
