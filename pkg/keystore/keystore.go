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

type eigenlayerKeystoreV3 struct {
}

func (k *eigenlayerKeystoreV3) LoadBLS(keyfile, passwd string) (*bls.KeyPair, error) {
	return bls.ReadPrivateKeyFromFile(keyfile, passwd)
}

func (k *eigenlayerKeystoreV3) LoadECDSA(keyfile, passwd string) (*ecdsa.PrivateKey, error) {
	return eigenecdsa.ReadKey(keyfile, passwd)
}

// NewKeystoreV3 returns a new instance of Keystore
func NewKeystoreV3() Keystore {
	// TODO: define home directory
	return &eigenlayerKeystoreV3{}
}
