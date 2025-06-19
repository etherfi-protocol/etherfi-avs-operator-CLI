package witness

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
	"github.com/urfave/cli/v3"
)

type RegistrationInput struct {
	OperatorID                int64
	WatchtowerAddress         common.Address
	WatchtowerSignature       string
	WatchtowerSignatureExpiry *big.Int
}

var WitnessPrepareRegistrationCmd = &cli.Command{
	Name:   "prepare-registration",
	Usage:  "(Node Operator) gather all inputs required to register for avs",
	Action: handleWitnessPrepareRegistration,
	Flags: append([]cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
	}, utils.GetECDSAKeystoreFlags()...),
}

func handleWitnessPrepareRegistration(ctx context.Context, cli *cli.Command) error {

	// parse cli input
	operatorID := cli.Int("operator-id")

	// load watchtower private key
	watchtowerPrivateKey, err := utils.LoadECDSAKey(cli, "WATCHTOWER_PRIVATE_KEY")
	if err != nil {
		return fmt.Errorf("loading watchtower private key: %w", err)
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	if err := witnessAPI.PrepareRegistration(operator, watchtowerPrivateKey); err != nil {
		return fmt.Errorf("failed to prepare witnesschain registration: %w", err)
	}

	return nil
}
