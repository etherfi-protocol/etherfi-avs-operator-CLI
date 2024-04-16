package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/keystore"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/logging"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
)

func createBLS(ctx context.Context, cli *cli.Command) error {
	blsKeyFile := cli.String("bls-key-file")
	blsKeyPassword := cli.String("bls-key-password")

	serviceManagerAddr := cli.String("service-manager")
	eigenlayerOperatorAddr := cli.String("eigenlayer-operator")

	rpc := cli.String("rpc")

	logging.Info("main", "rpc", rpc)
	logging.Info("main", "servicemanager addr", serviceManagerAddr)

	// TODO: Get password by prompt
	ks := keystore.NewKeystore()
	keyPair, err := ks.LoadBLS(blsKeyFile, blsKeyPassword)
	if err != nil {
		logging.Error("main", err)
		return err
	}

	//
	client, err := ethclient.Dial(rpc)

	serviceManager, err := bindings.NewEigenDAServiceManager(serviceManagerAddr, client)
	if err != nil {
		logging.Error("main", err)
		return err
	}

	coordinator, err := serviceManager.RegistryCoordinator()
	if err != nil {
		logging.Error("main", err)
		return err
	}

	g1MsgToSign, err := coordinator.PubkeyRegistrationMessageHash(eigenlayerOperatorAddr)
	if err != nil {
		logging.Error("main", err)
		return err
	}

	signature := keyPair.SignHashedToCurveMessage(g1MsgToSign)

	sig := new(types.AVSBLSSignature)
	sig.G1.X = keyPair.GetPubKeyG1().X.String()
	sig.G1.Y = keyPair.GetPubKeyG1().Y.String()
	
	sig.G2.X[1] = keyPair.GetPubKeyG2().X.A0.String()
	sig.G2.X[0] = keyPair.GetPubKeyG2().X.A1.String()
	sig.G2.Y[1] = keyPair.GetPubKeyG2().Y.A0.String()
	sig.G2.Y[0] = keyPair.GetPubKeyG2().Y.A1.String()
	
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

}
