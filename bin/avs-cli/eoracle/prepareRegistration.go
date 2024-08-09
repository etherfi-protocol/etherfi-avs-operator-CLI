package eoracle

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/keystore"
	"github.com/urfave/cli/v3"
)

var EOraclePrepareRegistrationCmd = &cli.Command{
	Name:   "prepare-registration",
	Usage:  "(Node Operator) gather all inputs required to register for eOracle",
	Action: handleEOraclePrepareRegistration,
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
			Name:     "alias-address",
			Usage:    "address associated with alias ECDSA key",
			Required: true,
		},
	},
}

func handleEOraclePrepareRegistration(ctx context.Context, cli *cli.Command) error {
	// parse cli input
	operatorID := cli.Int("operator-id")
	blsKeyFile := cli.String("bls-keystore")
	blsKeyPassword := cli.String("bls-password")
	aliasAddress := common.HexToAddress(cli.String("alias-address"))

	// decrypt and load bls key from keystore
	ks := keystore.NewKeystoreV3()
	blsKey, err := ks.LoadBLS(blsKeyFile, blsKeyPassword)
	if err != nil {
		return fmt.Errorf("loading bls keystore: %w", err)
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return eoracleAPI.PrepareRegistration(operator, blsKey, aliasAddress)
}
