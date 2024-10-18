package eigenda

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v3"
)

var UpdateSocketCmd = &cli.Command{
	Name:   "update-socket",
	Usage:  "(Operator) Update the socket",
	Action: handleUpdateSocket,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "socket",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
	},
}

func handleUpdateSocket(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	socket := cli.String("socket")
	operatorID := cli.Int("operator-id")

	if operatorID == 0 {
		return fmt.Errorf("invalid operator-id")
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	// load eip-1271 admin signing key
	signingKey, err := crypto.HexToECDSA(os.Getenv("ADMIN_1271_SIGNING_KEY"))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	return eigendaAPI.UpdateSocket(operator, socket, signingKey)
}
