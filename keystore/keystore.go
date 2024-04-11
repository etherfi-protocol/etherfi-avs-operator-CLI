package keystore

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

type Keystore interface {
	Load(keyfile, passwd string) (*bls.KeyPair, error)
}


type eigenlayerKeystore struct {
}

func (k *eigenlayerKeystore) Load(keyfile, passwd string) (*bls.KeyPair, error) {
	return bls.ReadPrivateKeyFromFile(keyfile, passwd)
}

// NewKeystore returns a new instance of Keystore
func NewKeystore() Keystore {
	// TODO: define home directory
	return &eigenlayerKeystore{}
}
