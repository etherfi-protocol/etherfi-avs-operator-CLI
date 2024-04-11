package keystore

import (
	"fmt"
	"testing"
	"math/big"

	//"github.com/Layr-Labs/eigenda/core"
	//"github.com/Layr-Labs/eigenda/core/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/stretchr/testify/assert"
)

func TestKeystore_Load(t *testing.T) {
	// fixtures
	blsKeyFile := "./fixtures/fixture.bls.key.json"
	privKey := "16590985200001101433457980282110542214238205519476686128549391575925906182766"
	passwd := "fixture@test1234"

	ks := NewKeystore()
	blsKeyPair, err := ks.Load(blsKeyFile, passwd)

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
}

func TestKeystore_BLSSign(t *testing.T) {
	// fixtures
	blsKeyFile := "./fixtures/fixture.bls.key.json"
	passwd := "fixture@test1234"

	ks := NewKeystore()
	blsKeyPair, err := ks.Load(blsKeyFile, passwd)

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
