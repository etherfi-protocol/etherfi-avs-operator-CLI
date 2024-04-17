package signer

import (
	"crypto/sha256"
	"testing"

	eigenbn254 "github.com/Layr-Labs/eigensdk-go/crypto/bn254"
	"github.com/stretchr/testify/assert"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/keystore"
)

func TestNewAVSSigner(t *testing.T) {
	// Prepare keystores
	blsKeyFile := "../../keystore/fixtures/fixture.bls.key.json"
	blsKeyPassword := "fixture@test1234"

	ks := keystore.NewKeystore()

	blsKeyPair, err := ks.LoadBLS(blsKeyFile, blsKeyPassword)
	assert.Nil(t, err)

	signer := NewAVSSigner(blsKeyPair)

	// Prepare message
	msg := "hello world"
	hash := sha256.Sum256([]byte(msg))
	g1MsgHash := eigenbn254.MapToCurve(hash)

	// Sign and verify
	signature, err := signer.Sign(g1MsgHash)
	assert.Nil(t, err)

	isValid, err := signer.Verify(g1MsgHash, signature)
	assert.Nil(t, err)

	assert.True(t, isValid)

	wrongMsg := "another message"
	hash = sha256.Sum256([]byte(wrongMsg))
	g1MsgHash = eigenbn254.MapToCurve(hash)

	isValid, err = signer.Verify(g1MsgHash, signature)
	assert.Nil(t, err)

	assert.False(t, isValid)
}
