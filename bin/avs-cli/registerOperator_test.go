package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: Quick scratch space while developing. Will convert to formal tests
func TestSignerForOperatorID(t *testing.T) {

	// connect to RPC node
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		t.Fatal(err)
	}

	// bind rpc to contract abi
	operatorManagerContract, err := contracts.NewEtherfiAVSOperatorsManager(bindings.Holesky.AvsOperatorManagerAddress, rpcClient)
	if err != nil {
		t.Fatal(err)
	}
	directoryContract, err := contracts.NewAVSDirectory(bindings.Holesky.AvsDirectoryAddress, rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	operatorID := big.NewInt(2)
	operatorAddr, err := operatorManagerContract.AvsOperators(nil, operatorID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("operatorAddress: %s\n", operatorAddr)

	// figure out who needs to sign the message
	signerAddr, err := operatorManagerContract.EcdsaSigner(nil, operatorID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("signerAddr: %s\n", signerAddr)

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Salt: %s\n", hex.EncodeToString(salt[:]))
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 365).Unix())
	fmt.Printf("Expiry: %s\n", expiry)

	hash, err := directoryContract.CalculateOperatorAVSRegistrationDigestHash(nil, operatorAddr, bindings.Holesky.EigenDARegistryCoordinatorAddress, [32]byte(salt), expiry)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("hash: %s\n", hex.EncodeToString(hash[:]))

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		t.Fatal(err)
	}
	signed, err := crypto.Sign(hash[:], privateKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("signed: %s\n", hex.EncodeToString(signed[:]))

	fmt.Printf("len(signed): %d\n", len(signed))

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("fromAddress: %s\n", fromAddress)

	signed[64] += 27

	// see if signature is valid
	//blah := common.HexToAddress("0x6E0d4De93CFA0eEfd4462064634614f81c545101")
	operatorContract, err := contracts.NewEtherfiAVSOperator(operatorAddr, rpcClient)
	//operatorContract, err := contracts.NewEtherfiAVSOperator(blah, rpcClient)
	if err != nil {
		t.Fatal(err)
	}
	signerTest, err := operatorContract.EcdsaSigner(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("signerTest: %s\n", signerTest)

	magic, err := operatorContract.IsValidSignature(nil, hash, signed)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("magic: %s\n", hex.EncodeToString(magic[:]))

	/*
		transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(17000))
		if err != nil {
			t.Fatal(err)
		}
	*/

	/*
		//RegisterOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
		_, err := operatorManagerContract.RegisterOperator(transactor, operatorID, EigenDARegistryCoordinatorHolesky, []byte{1}, "TEST_SOCKET", )

	*/
	/*
		isRegistered, err := operatorManagerContract.IsAvsRegistered(nil, operatorID, common.HexToAddress("0x6E0d4De93CFA0eEfd4462064634614f81c545101"))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("isRegistered: %v\n", isRegistered)
	*/

	/*
		// load key for signing tx
		privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
		if err != nil {
			t.Fatal(err)
		}
		transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(17000))
		if err != nil {
			t.Fatal(err)
		}

		_, err = operatorManagerContract.UpdateEcdsaSigner(transactor, operatorID, common.HexToAddress("0xD0d7F8a5a86d8271ff87ff24145Cf40CEa9F7A39"))
		if err != nil {
			t.Fatal(err)
		}
	*/
}

func TestRegisterOperator(t *testing.T) {

}
