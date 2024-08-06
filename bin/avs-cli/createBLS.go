package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/avs/signer"
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
	ks := keystore.NewKeystoreV3()
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

	g1MsgToSign, err := coordinator.PubkeyRegistrationMessageHash(common.HexToAddress(eigenlayerOperatorAddr))
	if err != nil {
		logging.Error("main", err)
		return err
	}

	avsSigner := signer.NewAVSSigner(keyPair)
	g1Sig, err := avsSigner.Sign(g1MsgToSign)
	if err != nil {
		return err
	}

	sig := new(types.BLSPubkeyRegistrationParams)
	sig.Load(keyPair.GetPubKeyG1().G1Affine, keyPair.GetPubKeyG2().G2Affine, g1Sig)

	logging.Info("main", "Verify Created Signature")
	isValid, err := avsSigner.Verify(g1MsgToSign, g1Sig)
	if err != nil {
		return err
	}

	logging.Info("main", fmt.Sprintf("Signature is valid: %v", isValid))

	d, err := json.Marshal(sig)
	if err != nil {
		logging.Error("main", err)
		return err
	}

	// Write into json
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
