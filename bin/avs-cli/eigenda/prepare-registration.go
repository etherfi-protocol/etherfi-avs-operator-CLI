package eigenda

import (
	"context"
	"fmt"

	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/keystore"
	"github.com/urfave/cli/v3"
)

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
		&cli.IntSliceFlag{
			Name:     "quorums",
			Usage:    "which quorums to register for i.e. 0,1",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "socket",
			Usage:    "eigenda socket",
			Required: true,
		},
	},
}

func handlePrepareRegistration(ctx context.Context, cli *cli.Command) error {

	// parse cli input
	operatorID := cli.Int("operator-id")
	blsKeyFile := cli.String("bls-keystore")
	blsKeyPassword := cli.String("bls-password")
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

	return eigendaAPI.PrepareRegistration(operator, keyPair, socket, quorums)
}
