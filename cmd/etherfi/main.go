package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/keystore"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/urfave/cli/v3"
)

var (
	// TODO: refactoring
	avsCmd = &cli.Command{
		Name:  "avs",
		Usage: "AVS operator command",
		Action: func(ctx context.Context, cli *cli.Command) error {
			blsKeyFile := cli.String("bls-key-file")
			blsKeyPassword := cli.String("bls-key-password")
			regMsg := cli.String("reg-msg")

			// TODO: Get password by prompt

			ks := keystore.NewKeystore()
			keyPair, err := ks.Load(blsKeyFile, blsKeyPassword)
			if err != nil {
				return err
			}

			msgTokens := strings.Split(regMsg, ",")
			if len(msgTokens) != 2 {
				return fmt.Errorf("invalid registration message")
			}

			msgG1X := new(big.Int)
			msgG1X.SetString(msgTokens[0], 10)

			msgG1Y := new(big.Int)
			msgG1Y.SetString(msgTokens[1], 10)

			xElem := fp.NewElement(0)
			xElem.SetBigInt(msgG1X)

			yElem := fp.NewElement(0)
			yElem.SetBigInt(msgG1Y)

			g1PointMsg := bn254.G1Affine{
				X: xElem,
				Y: yElem,
			}

			// Sign
			signature := keyPair.SignHashedToCurveMessage(&g1PointMsg)

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

			// Write into json
			// TODO: or send contract transaction
			sha256sum := sha256.Sum256(d)
			filename := hex.EncodeToString(sha256sum[:]) + ".json"
			err = os.WriteFile(filename, d, 0644)
			if err != nil {
				return err
			}

			fmt.Println("Signature file: ", filename)

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
		},
	}
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			avsCmd,
		},
	}

	ctx := context.Background()
	cmd.Run(ctx, os.Args)
}
