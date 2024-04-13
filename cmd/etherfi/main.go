package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/keystore"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/logging"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
)

var (
	registerBLSCmd = &cli.Command{
		Name:   "register-bls",
		Usage:  "register a bls key with the avsManager contract",
		Action: registerBLS,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "bls-signature-file",
				Usage: "BLS signature json file",
			},
			&cli.IntFlag{
				Name:  "operator-id",
				Usage: "Operator ID",
			},
			&cli.IntSliceFlag{
				Name:  "quorum-numbers",
				Usage: "Quorum Numbers",
			},
			&cli.StringFlag{
				Name:  "socket",
				Usage: "Socket",
			},
			&cli.BoolFlag{
				Name:  "broadcast",
				Usage: "broadcast signed tx to network",
			},
			&cli.IntFlag{
				Name:  "chain-id",
				Usage: "Chain ID",
				Value: 1, // default to mainnet
			},
		},
	}

	// TODO: refactoring
	avsCmd = &cli.Command{
		Name:  "avs",
		Usage: "AVS operator command",
		Action: func(ctx context.Context, cli *cli.Command) error {
			blsKeyFile := cli.String("bls-key-file")
			blsKeyPassword := cli.String("bls-key-password")

			serviceManager := cli.String("service-manager")
			eigenlayerOperator := cli.String("eigenlayer-operator")

			rpc := cli.String("rpc")

			logging.Info("main", "rpc", rpc)
			logging.Info("main", "servicemanager addr", serviceManager)

			// TODO: Get password by prompt
			ks := keystore.NewKeystore()
			keyPair, err := ks.LoadBLS(blsKeyFile, blsKeyPassword)
			if err != nil {
				logging.Error("main", err)
				return err
			}

			//
			client, err := ethclient.Dial(rpc)

			daServiceManager, err := bindings.NewEigenDAServiceManager(serviceManager, client)
			if err != nil {
				logging.Error("main", err)
				return err
			}

			coordinator, err := daServiceManager.RegistryCoordinator()
			if err != nil {
				logging.Error("main", err)
				return err
			}

			g1MsgToSign, err := coordinator.PubkeyRegistrationMessageHash(eigenlayerOperator)
			if err != nil {
				logging.Error("main", err)
				return err
			}

			signature := keyPair.SignHashedToCurveMessage(g1MsgToSign)

			sig := new(types.AVSBLSSignature)
			sig.G1.X = keyPair.GetPubKeyG1().X.String()
			sig.G1.Y = keyPair.GetPubKeyG1().Y.String()
			sig.G2.X[0] = keyPair.GetPubKeyG2().X.A0.String()
			sig.G2.X[1] = keyPair.GetPubKeyG2().X.A1.String()
			sig.G2.Y[0] = keyPair.GetPubKeyG2().Y.A0.String()
			sig.G2.Y[1] = keyPair.GetPubKeyG2().Y.A1.String()
			sig.Signature.X = signature.X.String()
			sig.Signature.Y = signature.Y.String()

			d, err := json.Marshal(sig)
			if err != nil {
				logging.Error("main", err)
				return err
			}

			// Write into json
			// TODO: or send contract transaction
			sha256sum := sha256.Sum256(d)
			filename := hex.EncodeToString(sha256sum[:]) + ".json"
			err = os.WriteFile(filename, d, 0644)
			if err != nil {
				logging.Error("main", err)
				return err
			}

			logging.Info("main", "Signature file: ", filename)

			return nil
		},
		// TODO: Check flags details
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "bls-key-file",
				Usage: "BLS key file",
			},
			&cli.StringFlag{
				Name:  "bls-key-password",
				Usage: "BLS key password",
			},
			&cli.StringFlag{
				Name:  "reg-msg",
				Usage: "PubkeyRegistragionMessageHash",
			},

			&cli.StringFlag{
				Name:  "rpc",
				Usage: "RPC Endpoint",
			},
			&cli.StringFlag{
				Name:  "service-manager",
				Usage: "Contract Address of Service Manager",
			},
			&cli.StringFlag{
				Name:  "eigenlayer-operator",
				Usage: "Address of Registered Eigenlayer Operator",
			},
		},
	}
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			avsCmd,
			registerBLSCmd,
		},
	}

	ctx := context.Background()
	if err := cmd.Run(ctx, os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
