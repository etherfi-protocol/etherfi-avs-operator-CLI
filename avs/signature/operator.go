package signature

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func CreateOperatorSignature(key *ecdsa.PrivateKey, msg, salt [32]byte, expiry *big.Int) ([]byte, error) {
	// TODO: Call AVSDirectory.calculateOperatorAVSRegistrationDigestHash
	return crypto.Sign(msg[:], key)
}
