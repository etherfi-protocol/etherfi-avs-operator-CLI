package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var (
	avsDirectoryMainnet = common.HexToAddress("0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF")
	avsDirectoryHolesky = common.HexToAddress("0x055733000064333CaDDbC92763c58BF0192fFeBf")
)

// TODO: Inject
//var etherfi

// Needs a bunch of refactoring
func registerOperator(ctx context.Context, cli *cli.Command) error {

	var (
		registryCoordinatorAddress common.Address
		operatorManagerAddress     common.Address
		avsDirectoryAddress        common.Address
	)

	chainID := cli.Int("chain-id")
	switch chainID {
	case 1:
		registryCoordinatorAddress = EigenDARegistryCoordinatorMainnet
		operatorManagerAddress = operatorManagerMainnet
		avsDirectoryAddress = avsDirectoryMainnet
	case 17000:
		registryCoordinatorAddress = EigenDARegistryCoordinatorHolesky
		operatorManagerAddress = operatorManagerHolesky
		avsDirectoryAddress = avsDirectoryHolesky
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
	operatorManagerContract, err := contracts.NewEtherfiAVSOperatorsManager(operatorManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding contract: %w", err)
	}
	directoryContract, err := contracts.NewAVSDirectory(avsDirectoryAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding contract: %w", err)
	}

	// get hash to sign

	/*
		// convert to ethereum address
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("publicKey is not of type *ecdsa.PublicKey")
		}
		operatorAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	*/

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return fmt.Errorf("creating signer from key: %w", err)
	}

	// toggle whether tx is broadcast to network
	transactor.NoSend = !cli.Bool("broadcast")

	// calculate the hash
	//directoryContract.CalculateOperatorAVSRegistrationDigestHash(nil, ope
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("unable to generate salt: %w", err)
	}

	// TODO: Inject
	//operatorID := 1
	operatorID := big.NewInt(cli.Int("operator-id"))

	// grab the operator
	operatorAddr, err := operatorManagerContract.AvsOperators(nil, operatorID)
	if err != nil {
		return fmt.Errorf("failed to look up operator address: %w", err)
	}

	// figure out who needs to sign the message
	signerAddr, err := operatorManagerContract.EcdsaSigner(nil, operatorID)
	if err != nil {
		return fmt.Errorf("failed to look up operator address: %w", err)
	}
	fmt.Printf("signerAddr: %s\n", signerAddr)
	return nil

	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 365).Unix())

	// sign

	// pass along

	// operator, avs, salt, expiry

	directoryContract.CalculateOperatorAVSRegistrationDigestHash(nil, operatorAddr, EigenDARegistryCoordinatorHolesky, [32]byte(salt), expiry)

	// TODO: validate params
	registryCoordinator := registryCoordinatorAddress
	socket := cli.String("socket")
	var quorumNumbers []uint8
	for _, v := range cli.IntSlice("quorum-numbers") {
		quorumNumbers = append(quorumNumbers, uint8(v))
	}

	// load bls signature and convert to format expected by contract
	params, err := BLSJsonToRegistrationParams(cli.String("bls-signature-file"))
	if err != nil {
		return fmt.Errorf("barsing bls signature file: %w", err)
	}

	// Sign transaction and broadcast if requested
	tx, err := operatorManagerContract.RegisterBlsKeyAsDelegatedNodeOperator(transactor, operatorID, registryCoordinator, quorumNumbers, socket, *params)
	if err != nil {
		return fmt.Errorf("failed to sign and/or broadcast tx: %w", err)
	}
	var buf bytes.Buffer
	tx.EncodeRLP(&buf)
	fmt.Printf("raw tx: %s\n", hex.EncodeToString(buf.Bytes()))

	return nil
}
