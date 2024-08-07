package eigenda

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/avs/signer"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/keystore"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

type RegistrationInput struct {
	OperatorID                  int64
	Socket                      string
	Quorums                     []int64
	BLSPubkeyRegistrationParams *types.BLSPubkeyRegistrationParams
}

var EigenDAPrepareRegistrationCmd = &cli.Command{
	Name:   "prepare-registration",
	Action: handleEigenDAPrepareRegistration,
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
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
	},
}

func handleEigenDAPrepareRegistration(ctx context.Context, cli *cli.Command) error {
	// parse cli input
	operatorID := cli.Int("operator-id")
	blsKeyFile := cli.String("bls-keystore")
	blsKeyPassword := cli.String("bls-password")
	quorums := cli.IntSlice("quorums")
	socket := cli.String("socket")
	rpcURL := cli.String("rpc-url")

	// decrypt and load bls key from keystore
	ks := keystore.NewKeystoreV3()
	keyPair, err := ks.LoadBLS(blsKeyFile, blsKeyPassword)
	if err != nil {
		return fmt.Errorf("loading bls keystore: %w", err)
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

	registryCoordinator, err := bindings.NewRegistryCoordinator(cfg.EigenDARegistryCoordinatorAddress, rpcClient)
	if err != nil {
		return err
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

	// load hash to sign with bls key
	g1MsgToSign, err := registryCoordinator.PubkeyRegistrationMessageHash(operatorAddr)
	if err != nil {
		return fmt.Errorf("fetching pubkeyRegistrationMessageHash: %w", err)
	}
	avsSigner := signer.NewAVSSigner(keyPair)
	g1Sig, err := avsSigner.Sign(g1MsgToSign)
	if err != nil {
		return fmt.Errorf("signing pubkey registration hash: %w", err)
	}
	signedParams := new(types.BLSPubkeyRegistrationParams)
	signedParams.Load(keyPair.GetPubKeyG1().G1Affine, keyPair.GetPubKeyG2().G2Affine, g1Sig)

	isValid, err := avsSigner.Verify(g1MsgToSign, g1Sig)
	if !isValid || err != nil {
		return fmt.Errorf("failed to verify g1 signature: %w", err)
	}

	registrationInput := RegistrationInput{
		OperatorID:                  operatorID,
		Socket:                      socket,
		Quorums:                     quorums,
		BLSPubkeyRegistrationParams: signedParams,
	}

	out, err := json.MarshalIndent(registrationInput, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))

	return nil
}
