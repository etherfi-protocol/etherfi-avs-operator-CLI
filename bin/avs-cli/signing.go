package main

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

// signEIP155 performs an ECDSA signature over the provided key but increments the V value
// of the signature to match signatures expected by most ethereum applications
func SignDigest(digest []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {

	signed, err := crypto.Sign(digest, privateKey)
	if err != nil {
		return nil, err
	}

	// account for EIP-155 by incrementing V if necessary
	if signed[64] < 27 {
		signed[64] += 27
	}

	return signed, nil
}
