package ungate

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/eigenlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/othentic"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
)

// API handle for all core Ungate functionality
type API struct {
	Client *ethclient.Client

	AVSGovernanceAddress      common.Address
	AVSGovernance             *AVSGovernance
	ServiceManagerAddress     common.Address
	AvsOperatorManagerAddress common.Address

	EigenlayerAPI *eigenlayer.API
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	avsGovernance, _ := NewAVSGovernance(cfg.UngateAVSGovernanceAddress, rpcClient)

	return &API{
		Client:                    rpcClient,
		AVSGovernance:             avsGovernance,
		AVSGovernanceAddress:      cfg.UngateAVSGovernanceAddress,
		ServiceManagerAddress:     cfg.UngateAVSGovernanceAddress, // AVSGovernance contract serves as the service manager for this AVS
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		EigenlayerAPI:             eigenlayer.New(cfg, rpcClient),
	}
}

func (a *API) RegisterOperator(operator *etherfi.Operator, othenticOperatorData *othentic.CLIOperatorData, signerKey *ecdsa.PrivateKey, whitelistEnabled bool) error {

	// othentic cli stores pubkey as 4 "0x" prefixed hex strings representing a G2 point
	var blsPubkey [4]*big.Int
	blsPubkey[0], _ = new(big.Int).SetString(othenticOperatorData.BlsKey[0][2:], 16)
	blsPubkey[1], _ = new(big.Int).SetString(othenticOperatorData.BlsKey[1][2:], 16)
	blsPubkey[2], _ = new(big.Int).SetString(othenticOperatorData.BlsKey[2][2:], 16)
	blsPubkey[3], _ = new(big.Int).SetString(othenticOperatorData.BlsKey[3][2:], 16)

	// othentic cli stores pubkey as 2 "0x" prefixed hex strings representing a G1 point
	var blsSignature BLSAuthLibrarySignature // I'm not sure why they put this in a struct... but not the pubkey
	blsSignature.Signature[0], _ = new(big.Int).SetString(othenticOperatorData.BlsRegistrationSignature[0][2:], 16)
	blsSignature.Signature[1], _ = new(big.Int).SetString(othenticOperatorData.BlsRegistrationSignature[1][2:], 16)

	// generate and sign registration hash to be signed by admin ecdsa key
	sigWithSaltAndExpiry, err := a.EigenlayerAPI.GenerateAndSignRegistrationDigest(operator, a.ServiceManagerAddress, signerKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}
	sigParams := ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sigWithSaltAndExpiry.Signature,
		Salt:      sigWithSaltAndExpiry.Salt,
		Expiry:    sigWithSaltAndExpiry.Expiry,
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	governanceABI, err := AVSGovernanceMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	// need to call different method depending on if whitelist enabled or not
	var calldata []byte
	if whitelistEnabled {

		// need to get a signed auth token from the othentic signing server
		token, err := othentic.AcquireSignedAuthToken(signerKey, operator, a.AVSGovernanceAddress)
		if err != nil {
			return fmt.Errorf("failed to acquire allowlist auth token: %w", err)
		}
		calldata, err = governanceABI.Pack("registerAsAllowedOperator", blsPubkey, token, operator.Address, sigParams, blsSignature)
		if err != nil {
			return fmt.Errorf("packing input: %w", err)
		}
	} else {

		// non whitelist flow
		calldata, err = governanceABI.Pack("registerAsOperator", blsPubkey, operator.Address, sigParams, blsSignature)
		if err != nil {
			return fmt.Errorf("packing input: %w", err)
		}
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, calldata, a.AVSGovernanceAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("ungate-register-operator-%d", operator.ID))
	return utils.ExportJSON("ungate-register-gnosis", operator.ID, batch)
}

func (a *API) UnregisterOperator(operator *etherfi.Operator) error {

	// manually pack tx data since we are submitting via gnosis instead of directly
	governanceABI, err := AVSGovernanceMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	calldata, err := governanceABI.Pack("unregisterAsOperator")
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// wrap the inner call to be forwarded via AvsOperatorManager
	adminCall, err := utils.PackForwardCallForAdmin(operator.ID, calldata, a.AVSGovernanceAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, a.AvsOperatorManagerAddress, fmt.Sprintf("ungate-unregister-operator-%d", operator.ID))
	return utils.ExportJSON("ungate-unregister-gnosis", operator.ID, batch)
}

/*

keeping this in case we want to entirely replace the othentic typescript cli

// Info that node operator must supply to the ether.fi admin for registration
type RegistrationInfo struct {
	OperatorID   int64
	BLSPubkey    [4]*big.Int
	BLSSignature [2]*big.Int
}

// Ungate is an Othentic based AVS
func (a *API) PrepareRegistration(operator *etherfi.Operator, operatorData *othentic.CLIOperatorData) error {

	// othentic cli stores pubkey as 4 "0x" prefixed hex strings representing a G2 point
	var blsPubkey [4]*big.Int
	blsPubkey[0], _ = new(big.Int).SetString(operatorData.BlsKey[0][2:], 16)
	blsPubkey[1], _ = new(big.Int).SetString(operatorData.BlsKey[1][2:], 16)
	blsPubkey[2], _ = new(big.Int).SetString(operatorData.BlsKey[2][2:], 16)
	blsPubkey[3], _ = new(big.Int).SetString(operatorData.BlsKey[3][2:], 16)

	// othentic cli stores pubkey as 2 "0x" prefixed hex strings representing a G1 point
	var blsSignature [2]*big.Int
	blsSignature[0], _ = new(big.Int).SetString(operatorData.BlsRegistrationSignature[0][2:], 16)
	blsSignature[1], _ = new(big.Int).SetString(operatorData.BlsRegistrationSignature[1][2:], 16)

	ri := RegistrationInfo{
		OperatorID:   operator.ID,
		BLSPubkey:    blsPubkey,
		BLSSignature: blsSignature,
	}
	return utils.ExportJSON("ungate-prepare-registration", operator.ID, ri)
}
*/
