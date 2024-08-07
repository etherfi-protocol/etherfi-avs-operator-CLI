package witness

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts/witnesschain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
	Action: handleWitnessPrepareRegistration,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "watchtower-address",
			Usage:    "watchtower address",
			Required: true,
		},
	},
}

func handleWitnessPrepareRegistration(ctx context.Context, cli *cli.Command) error {

	// parse cli input
	operatorID := cli.Int("operator-id")
	watchtowerAddress := common.HexToAddress(cli.String("watchtower-address"))
	rpcURL := cli.String("rpc-url")

	// load watchtower private key and sanity check against provided address
	watchtowerPrivateKey, err := crypto.HexToECDSA(os.Getenv("WATCHTOWER_PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("loading signing key: %w", err)
	}
	derivedWatchtowerAddress := crypto.PubkeyToAddress(watchtowerPrivateKey.PublicKey)
	fmt.Println("derivedWatchtowerAddress:", derivedWatchtowerAddress)
	if derivedWatchtowerAddress != watchtowerAddress {
		return fmt.Errorf("provided watchtower-address does not match address derived from key")
	}

	// load configuration
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing RPC: %w", err)
	}
	cfg, err := bindings.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	// look up operator contract associated with this id and configured ecdsaSigner
	operatorManagerContract, err := contracts.NewAvsOperatorManager(cfg.AvsOperatorManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding operatorManager: %w", err)
	}
	operatorAddr, err := operatorManagerContract.AvsOperators(nil, big.NewInt(operatorID))
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	// compute the watchtower registration digest
	operatorRegistry, err := witnesschain.NewWitnessChainOperatorRegistry(cfg.WitnessChainOperatorRegistryAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding operatorRegistry: %w", err)
	}
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 10).Unix())
	registrationDigest, err := operatorRegistry.CalculateWatchtowerRegistrationMessageHash(nil, operatorAddr, expiry)
	if err != nil {
		return fmt.Errorf("calculating watchtower registration digest: %w", err)
	}

	// sign the digest with the watchtower key
	signed, err := crypto.Sign(registrationDigest[:], watchtowerPrivateKey)
	if err != nil {
		return fmt.Errorf("signing digest: %w", err)
	}

	// account for EIP-155
	signed[64] += 27

	registrationInput := RegistrationInput{
		OperatorID:                operatorID,
		WatchtowerAddress:         watchtowerAddress,
		WatchtowerSignature:       "0x" + hex.EncodeToString(signed),
		WatchtowerSignatureExpiry: expiry,
	}
	out, err := json.MarshalIndent(registrationInput, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))

	return nil
}
