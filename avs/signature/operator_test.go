package signature

import (
	"fmt"
	"crypto/sha256"
	"math/big"
	"testing"
	"time"
	"encoding/hex"

	"github.com/stretchr/testify/assert"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/keystore"
)


func TestOperator_Sign(t *testing.T) {
	keyFile := "../../keystore/fixtures/fixture.ecdsa.key.json"
	passwd := "fixture@test1234"

	ks := keystore.NewKeystore()
	key, err := ks.LoadECDSA(keyFile, passwd)

	assert.Nil(t, err)

	// TODO: Rand salt
	salt := sha256.Sum256([]byte(time.Now().String()))
	expiry := big.NewInt(time.Now().Add(24 * time.Hour * 365).Unix())

	fmt.Println(hex.EncodeToString(salt[:]), expiry)
	fmt.Println(len(salt))

	msg := [32]byte{}
	CreateOperatorSignature(key, msg, salt, expiry)

	assert.Nil(t, nil)
}
