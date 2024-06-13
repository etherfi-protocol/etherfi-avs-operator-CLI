package eoracle

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

var EOraclePrepareRegistrationCmd = &cli.Command{
	Name:   "prepare-registration",
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
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
	},
}

func handleEOraclePrepareRegistration(ctx context.Context, cli *cli.Command) error {
	// parse cli input
	operatorID := cli.Int("operator-id")
	blsKeyFile := cli.String("bls-keystore")
	blsKeyPassword := cli.String("bls-password")
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
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("querying chainID from RPC: %w", err)
	}
	cfg, err := bindings.ConfigForChain(chainID.Int64())
	if err != nil {
		return err
	}
	registryCoordinator, err := bindings.NewRegistryCoordinator(cfg.EOracleRegistryCoordinator, rpcClient)
	if err != nil {
		return err
	}

	// look up operator contract associated with this id and configured ecdsaSigner
	operatorManagerContract, err := contracts.NewAvsOperatorManager(cfg.OperatorManagerAddress, rpcClient)
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
	sig := new(types.AVSBLSSignature)
	sig.Load(keyPair.GetPubKeyG1().G1Affine, keyPair.GetPubKeyG2().G2Affine, g1Sig)

	isValid, err := avsSigner.Verify(g1MsgToSign, g1Sig)
	if !isValid || err != nil {
		return fmt.Errorf("failed to verify g1 signature: %w", err)
	}

	out, err := json.MarshalIndent(sig, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))

	return nil
}
