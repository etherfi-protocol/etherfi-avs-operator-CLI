package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bindings"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
)

func TestValidSignature(t *testing.T) {

	// connect to RPC node
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		t.Fatal(err)
	}

	operatorManagerContract, err := etherfi.NewAvsOperatorManager(bindings.Holesky.AvsOperatorManagerAddress, rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	// find operator contract
	operatorID := big.NewInt(1)
	operatorContractAddress, err := operatorManagerContract.AvsOperators(nil, operatorID)
	operatorContract, err := etherfi.NewAvsOperator(operatorContractAddress, rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	// check if is valid signature
	hash, err := hex.DecodeString("587253a5d7251d73e70d28c441e3ce0385cd6e3f5641718fb05a32b5b45fee00")
	if err != nil {
		t.Fatal(err)
	}
	sig, err := hex.DecodeString("79f529ba4257af48865fcec577fd08a506e9aa7f38cc727a674301682c8b40a20fa28b13cde6127b98079366edf5c603aba582fa90a1c3b2847cfb96d8a94df21c")
	if err != nil {
		t.Fatal(err)
	}

	magic, err := operatorContract.IsValidSignature(nil, [32]byte(hash), sig)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("magic: %s\n", hex.EncodeToString(magic[:]))
}

func TestCalculateDigest(t *testing.T) {

	// connect to RPC node
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		t.Fatal(err)
	}

	operatorManagerContract, err := etherfi.NewAvsOperatorManager(bindings.Holesky.AvsOperatorManagerAddress, rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	serviceManagerAddress := common.HexToAddress("0x7A46219950d8a9bf2186549552DA35Bf6fb85b1F")
	salt, err := hex.DecodeString("bd320e070eceae61901e284e73b45355833172dc8fa97cc2fe10a76660aa97a8")
	if err != nil {
		t.Fatal(err)
	}
	expiry := big.NewInt(1744684487)

	operatorID := big.NewInt(1)
	hash, err := operatorManagerContract.CalculateOperatorAVSRegistrationDigestHash(nil, operatorID, serviceManagerAddress, [32]byte(salt), expiry)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("hash: %s\n", hex.EncodeToString(hash[:]))
}
