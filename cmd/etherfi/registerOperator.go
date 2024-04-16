package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

// Needs a bunch of refactoring

// TODO: This function is still a WIP as I debug an issue related to the BLS keys
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
	fmt.Println(operatorManagerContract)

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return fmt.Errorf("creating signer from key: %w", err)
	}

	// toggle whether tx is broadcast to network
	transactor.NoSend = !cli.Bool("broadcast")

	// hardcoded test
	operatorID := big.NewInt(1)
	registryCoordinator := cfg.BrevisRegistryCoordinator
	quorumNumbers := []byte{2}
	socket := "no need"
	blsParams, err := BLSJsonToRegistrationParams("keystore/fixtures/brevis.json")
	if err != nil {
		return fmt.Errorf("loading bls json: %w", err)
	}

	sig, err := hex.DecodeString("79f529ba4257af48865fcec577fd08a506e9aa7f38cc727a674301682c8b40a20fa28b13cde6127b98079366edf5c603aba582fa90a1c3b2847cfb96d8a94df21c")
	if err != nil {
		return fmt.Errorf("invalid signature value: %w", err)
	}
	salt, err := hex.DecodeString("bd320e070eceae61901e284e73b45355833172dc8fa97cc2fe10a76660aa97a8")
	if err != nil {
		return fmt.Errorf("invalid salt value: %w", err)
	}
	expiry := 1744684487

	fmt.Printf("operatorID: %s\n", operatorID)
	fmt.Printf("registryCoordinator: %s\n", registryCoordinator)
	fmt.Printf("quorumNumbers: %v\n", quorumNumbers)
	fmt.Printf("socket: %s\n", socket)
	fmt.Printf("BLS: %+v\n", blsParams)

	sigParams := contracts.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sig,
		Salt:      [32]byte(salt),
		Expiry:    big.NewInt(int64(expiry)),
	}
	fmt.Printf("SigParams: %+v\n", sigParams)

	_, err = operatorManagerContract.RegisterOperator(transactor, operatorID, registryCoordinator, sigParams)
	if err != nil {
		return fmt.Errorf("registering operator: %w", err)
	}
	return nil
	/*
		abi, err := contracts.EtherfiAVSOperatorsManagerMetaData.GetAbi()
		if err != nil {
			panic(err)
		}
		for name, method := range abi.Methods {
			fmt.Printf("manager: %s: %s\n", name, hex.EncodeToString(method.ID))
		}

		operatorAbi, err := contracts.EtherfiAVSOperatorMetaData.GetAbi()
		if err != nil {
			panic(err)
		}
		for name, method := range operatorAbi.Methods {
			fmt.Printf("operator: %s: %s\n", name, hex.EncodeToString(method.ID))
		}
	*/

	/*

		fmt.Printf("\nmanager-registerOperator: %s\n", hex.EncodeToString(abi.Methods["registerOperator"].ID))
		fmt.Printf("operator-registerOperator: %s\n", hex.EncodeToString(operatorAbi.Methods["registerOperator"].ID))

		fmt.Printf("getAvsInfo: %s\n", hex.EncodeToString(operatorAbi.Methods["getAvsInfo"].ID))
	*/

	//	return nil
	managerABI, err := contracts.EtherfiAVSOperatorsManagerMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	//input, err := managerABI.Pack("registerOperator", operatorID, registryCoordinator, quorumNumbers, socket, *blsParams, sigParams)
	input, err := managerABI.Pack("registerOperator", operatorID, registryCoordinator, sigParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("input: %s\n", hex.EncodeToString(input))

	//return _EtherfiAVSOperatorsManager.contract.Transact(opts, "registerOperator", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
	//input, err := c.abi.Pack(method, params...)

	return nil
}
