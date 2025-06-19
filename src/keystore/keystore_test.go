package keystore

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestKeystore_Load(t *testing.T) {
	// fixtures
	blsKeyFile := "./fixtures/fixture.bls.key.json"
	privKey := "16590985200001101433457980282110542214238205519476686128549391575925906182766"
	passwd := "fixture@test1234"

	ks := NewKeystoreV3()
	blsKeyPair, err := ks.LoadBLS(blsKeyFile, passwd)

	assert.Nil(t, err)
	assert.Equal(t, privKey, blsKeyPair.PrivKey.String())

	fmt.Println("=== G1 ===")
	fmt.Printf("g1.x: %s\n", blsKeyPair.GetPubKeyG1().X.String())
	fmt.Printf("g1.y: %s\n", blsKeyPair.GetPubKeyG1().Y.String())

	fmt.Println("=== G2 ===")
	fmt.Printf("g2.x1: %s\n", blsKeyPair.GetPubKeyG2().X.A0.String())
	fmt.Printf("g2.x2: %s\n", blsKeyPair.GetPubKeyG2().X.A1.String())
	fmt.Printf("g2.y1: %s\n", blsKeyPair.GetPubKeyG2().Y.A0.String())
	fmt.Printf("g2.y2: %s\n", blsKeyPair.GetPubKeyG2().Y.A1.String())

	d, err := json.Marshal(blsKeyPair.GetPubKeyG1())
	fmt.Println(string(d), err)

	d, err = json.Marshal(blsKeyPair.GetPubKeyG2())
	fmt.Println(string(d), err)
}

func TestKeystore_BLSSign(t *testing.T) {
	// fixtures
	blsKeyFile := "./fixtures/fixture.bls.key.json"
	passwd := "fixture@test1234"

	ks := NewKeystoreV3()
	blsKeyPair, err := ks.LoadBLS(blsKeyFile, passwd)

	assert.Nil(t, err)

	// TODO: Get message
	// Retrieve from RegistryCoordinator.PubkeyRegistrationMessageHash
	// 19046105841993330018967249168635223163277136948200676530072132334676557545907,9769332298681190511594115723881244838126887677320011858505970076150623739668
	msgG1X := new(big.Int)
	msgG1X.SetString("19046105841993330018967249168635223163277136948200676530072132334676557545907", 10)

	msgG1Y := new(big.Int)
	msgG1Y.SetString("9769332298681190511594115723881244838126887677320011858505970076150623739668", 10)

	xElem := fp.NewElement(0)
	xElem.SetBigInt(msgG1X)

	yElem := fp.NewElement(0)
	yElem.SetBigInt(msgG1Y)

	g1PointMsg := bn254.G1Affine{
		X: xElem,
		Y: yElem,
	}

	signature := blsKeyPair.SignHashedToCurveMessage(&g1PointMsg)

	fmt.Println("signature", signature.String())

	// TODO: Verify signature
	// input should be [32]byte format but how get I get this?
	//signature.Verify(blsKeyPair.GetPubKeyG2(), )
}

func TestKeystore_LoadECDSA(t *testing.T) {
	// fixtures
	ecdsaKeyFile := "./fixtures/fixture.ecdsa.key.json"
	expectedAddress := "0x287B703F25CE707D7974D26DBE5B78121f70794f"
	passwd := "fixture@test1234"

	ks := NewKeystoreV3()
	ecdsaKey, err := ks.LoadECDSA(ecdsaKeyFile, passwd)

	assert.Nil(t, err)
	assert.NotNil(t, ecdsaKey)

	// Verify the loaded key generates the expected address
	address := crypto.PubkeyToAddress(ecdsaKey.PublicKey)
	assert.Equal(t, strings.ToLower(expectedAddress), strings.ToLower(address.Hex()))

	fmt.Printf("Loaded ECDSA key for address: %s\n", address.Hex())
}

func TestKeystore_LoadECDSA_InvalidPassword(t *testing.T) {
	ecdsaKeyFile := "./fixtures/fixture.ecdsa.key.json"
	wrongPasswd := "wrongpassword"

	ks := NewKeystoreV3()
	ecdsaKey, err := ks.LoadECDSA(ecdsaKeyFile, wrongPasswd)

	assert.NotNil(t, err)
	assert.Nil(t, ecdsaKey)
	assert.Contains(t, err.Error(), "could not decrypt key")
}

func TestKeystore_LoadECDSA_NonexistentFile(t *testing.T) {
	nonexistentFile := "./fixtures/nonexistent.json"
	passwd := "fixture@test1234"

	ks := NewKeystoreV3()
	ecdsaKey, err := ks.LoadECDSA(nonexistentFile, passwd)

	assert.NotNil(t, err)
	assert.Nil(t, ecdsaKey)
}
