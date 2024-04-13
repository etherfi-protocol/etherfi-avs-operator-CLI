package keystore

import (
	"crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigenecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
)

type Keystore interface {
	LoadBLS(keyfile, passwd string) (*bls.KeyPair, error)
	LoadECDSA(keyfile, passwd string) (*ecdsa.PrivateKey, error)
}

type eigenlayerKeystore struct {
}

func (k *eigenlayerKeystore) LoadBLS(keyfile, passwd string) (*bls.KeyPair, error) {
	return bls.ReadPrivateKeyFromFile(keyfile, passwd)
}

func (k *eigenlayerKeystore) LoadECDSA(keyfile, passwd string) (*ecdsa.PrivateKey, error) {
	return eigenecdsa.ReadKey(keyfile, passwd)
}

// NewKeystore returns a new instance of Keystore
func NewKeystore() Keystore {
	// TODO: define home directory
	return &eigenlayerKeystore{}
}
