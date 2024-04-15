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

	/*
		debugSigningFunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			fmt.Printf("--------------------------------------\n")

			signer := types.LatestSignerForChainID(big.NewInt(17000))
			keyAddr := common.HexToAddress("0xD0d7F8a5a86d8271ff87ff24145Cf40CEa9F7A39")

			if address != keyAddr {
				return nil, fmt.Errorf("wrong signing key")
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privateKey)
			if err != nil {
				return nil, err
			}

			test, err := tx.WithSignature(signer, signature)
			if err != nil {
				return nil, fmt.Errorf("unable to sign: %w", err)
			}
			buf := bytes.Buffer{}
			if err := test.EncodeRLP(&buf); err != nil {
				panic(err)
			}

			fmt.Printf("signed tx: %s\n", string(buf.Bytes()))

			return tx.WithSignature(signer, signature)
		}
		transactor.Signer = debugSigningFunc
	*/

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

	_, err = operatorManagerContract.RegisterOperator(transactor, operatorID, registryCoordinator, quorumNumbers, socket, *blsParams, sigParams)
	if err != nil {
		return fmt.Errorf("registering operator: %w", err)
	}

	//return _EtherfiAVSOperatorsManager.contract.Transact(opts, "registerOperator", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
	//input, err := c.abi.Pack(method, params...)

	return nil
}
