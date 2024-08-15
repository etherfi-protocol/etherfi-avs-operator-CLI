package lagrangezk

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

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
			Name:     "pubkey",
			Usage:    "hex pubkey of lagrange ECDSA key",
			Required: true,
		},
	},
}

func handlePrepareRegistration(ctx context.Context, cli *cli.Command) error {

	// parse cli input
	operatorID := cli.Int("operator-id")
	pubkeyHex := cli.String("pubkey")

	// validate pubkey
	if strings.HasPrefix(pubkeyHex, "0x") {
		pubkeyHex = pubkeyHex[2:]
	}
	buf, err := hex.DecodeString(pubkeyHex)
	if err != nil {
		return fmt.Errorf("invalid pubkey: %w", err)
	}
	if len(buf) != 64 {
		return fmt.Errorf("invalid pubkey length")
	}
	pubkeyX := new(big.Int).SetBytes(buf[0:32])
	pubkeyY := new(big.Int).SetBytes(buf[32:64])

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	return lagrangeZKAPI.PrepareRegistration(operator, pubkeyX, pubkeyY)
}
