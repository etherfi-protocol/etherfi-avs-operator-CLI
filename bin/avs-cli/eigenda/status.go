package eigenda

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var StatusCmd = &cli.Command{
	Name:   "status",
	Usage:  "(Operator) info about currently registered operator",
	Action: handleStatus,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
	},
}

func handleStatus(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	operatorID := cli.Int("operator-id")

	if operatorID == 0 {
		return fmt.Errorf("invalid operator-id")
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return eigendaAPI.Status(operator)
}
