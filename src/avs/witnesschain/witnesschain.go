package witnesschain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/avs/signer"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/gnosis"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/etherfi"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// API handle for all core witness chain functionality
type WitnessChain struct {
	Client *ethclient.Client

	OperatorRegistryAddress   common.Address
	OperatorRegistry          *WitnessChainOperatorRegistry
	WitnessHubAddress         common.Address
	WitnessHub                *WitnessChainWitnessHub
	AvsOperatorManagerAddress common.Address
}

// Info that node operator must supply to the ether.fi admin for registration
type RegistrationInfo struct {
	OperatorID                int64
	WatchtowerAddress         common.Address
	WatchtowerSignature       string
	WatchtowerSignatureExpiry *big.Int
}

// PrepareRegistration aggregates all required info from the node operator that
// the ether.fi admin will need to register them to the Witness Chain AVS
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

// RegisterOperator is used by the ether.fi admin to register a node operator's AvsOperator contract
// with the Witness Chain AVS.
func (wc *WitnessChain) RegisterOperator(operator *etherfi.Operator, signingKey *ecdsa.PrivateKey) error {

	// generate and sign registration hash with admin ecdsa key
	signature, err := signer.GenerateAndSignRegistrationDigest(operator, signer.WITNESS_CHAIN, wc.Client, signingKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}

	// convert to types expected by contract call
	sigWithSaltAndExpiry := ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature.Signature,
		Salt:      signature.Salt,
		Expiry:    signature.Expiry,
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	hubABI, err := WitnessChainWitnessHubMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	input, err := hubABI.Pack("registerOperatorToAVS", operator.Address, sigWithSaltAndExpiry)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, input, wc.WitnessHubAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, wc.AvsOperatorManagerAddress, fmt.Sprintf("witness-chain-register-%d", operator.ID))
	return utils.ExportJSON("witness-chain-register-gnosis", operator.ID, batch)
}

func (wc *WitnessChain) RegisterWatchtower(operator *etherfi.Operator, info *RegistrationInfo) error {

	// parse watchtower signature
	if strings.HasPrefix(info.WatchtowerSignature, "0x") {
		info.WatchtowerSignature = info.WatchtowerSignature[2:]
	}
	watchtowerSignature, err := hex.DecodeString(info.WatchtowerSignature)
	if err != nil {
		return fmt.Errorf("invalid watchtower signature")
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	witnessABI, err := WitnessChainOperatorRegistryMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	// pack operatorRegistry.registerWatchtowerAsOperator()
	calldata, err := witnessABI.Pack("registerWatchtowerAsOperator", info.WatchtowerAddress, info.WatchtowerSignatureExpiry, watchtowerSignature)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, calldata, wc.OperatorRegistryAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, wc.AvsOperatorManagerAddress, fmt.Sprintf("witness-chain-register-watchtower-%d", operator.ID))
	return utils.ExportJSON("witness-chain-register-gnosis", operator.ID, batch)
}
