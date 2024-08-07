package witnesschain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/etherfi"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// API handle for all core witness chain functionality
type WitnessChain struct {
	Client *ethclient.Client

	OperatorRegistryAddress common.Address
	OperatorRegistry        *WitnessChainOperatorRegistry
	WitnessHubAddress       common.Address
	WitnessHub              *WitnessChainWitnessHub
}

type RegistrationInfo struct {
	OperatorID                int64
	WatchtowerAddress         common.Address
	WatchtowerSignature       string
	WatchtowerSignatureExpiry *big.Int
}

func (wc *WitnessChain) PrepareRegistration(operator *etherfi.Operator, watchtowerKey *ecdsa.PrivateKey) error {

	// compute the watchtower registration digest
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 10).Unix())
	registrationDigest, err := wc.OperatorRegistry.CalculateWatchtowerRegistrationMessageHash(nil, operator.Address, expiry)
	if err != nil {
		return fmt.Errorf("calculating watchtower registration digest: %w", err)
	}

	// sign the digest with the watchtower key
	signed, err := utils.SignDigestECDSA(registrationDigest[:], watchtowerKey)
	if err != nil {
		return fmt.Errorf("signing digest: %w", err)
	}

	ri := RegistrationInfo{
		OperatorID:                operator.ID,
		WatchtowerAddress:         crypto.PubkeyToAddress(watchtowerKey.PublicKey),
		WatchtowerSignature:       "0x" + hex.EncodeToString(signed),
		WatchtowerSignatureExpiry: expiry,
	}

	return utils.ExportJSON("witnesschain-prepare-registration", operator.ID, ri)
}
