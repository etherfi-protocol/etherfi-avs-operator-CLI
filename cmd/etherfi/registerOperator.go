package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

// Needs a bunch of refactoring

func registerOperator(ctx context.Context, cli *cli.Command) error {

	var cfg bindings.Config
	chainID := cli.Int("chain-id")
	switch chainID {
	case 1:
		cfg = bindings.Mainnet
	case 17000:
		cfg = bindings.Holesky
	default:
		return fmt.Errorf("unimplemented chain: %d", chainID)
	}

	// TODO: validate params
	operatorID := big.NewInt(cli.Int("operator-id"))
	registryCoordinator := common.HexToAddress(cli.String("registry-coordinator"))
	expiry := cli.Int("expiry")
	signature, err := hex.DecodeString(strings.TrimPrefix(cli.String("signature"), "0x"))
	if err != nil {
		return fmt.Errorf("invalid signature: %w", err)
	}
	salt, err := hex.DecodeString(strings.TrimPrefix(cli.String("salt"), "0x"))
	if err != nil {
		return fmt.Errorf("invalid salt: %w", err)
	}

	// load key for signing tx
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("loading signing key: %w", err)
	}

	// connect to RPC node
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}

	// bind rpc to contract abi
	operatorManagerContract, err := contracts.NewEtherfiAVSOperatorsManager(cfg.OperatorManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding contract: %w", err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return fmt.Errorf("creating signer from key: %w", err)
	}

	// toggle whether tx is broadcast to network
	transactor.NoSend = !cli.Bool("broadcast")

	sigParams := contracts.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature,
		Salt:      [32]byte(salt),
		Expiry:    big.NewInt(int64(expiry)),
	}

	_, err = operatorManagerContract.RegisterOperator(transactor, operatorID, registryCoordinator, sigParams)
	if err != nil {
		return fmt.Errorf("registering operator: %w", err)
	}

	// TODO: enable dumping of raw tx without signing to enable cli to be used with multisigs
	/*
		managerABI, err := contracts.EtherfiAVSOperatorsManagerMetaData.GetAbi()
		if err != nil {
			panic(err)
		}
		input, err := managerABI.Pack("registerOperator", operatorID, registryCoordinator, sigParams)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Raw Input: %s\n", hex.EncodeToString(input))
	*/

	return nil
}
