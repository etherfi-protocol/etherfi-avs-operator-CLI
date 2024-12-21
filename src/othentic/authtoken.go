package othentic

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
)

type API struct {
	rpcClient *ethclient.Client
}

// Output produced by an operator using the Othentic CLI
type CLIOperatorData struct {
	Operator                 string   `json:"operator"`
	ChainID                  string   `json:"chainId"`
	Avs                      string   `json:"avs"`
	BlsKey                   []string `json:"blsKey"`
	BlsRegistrationSignature []string `json:"blsRegistrationSignature"`
}

func AcquireSignedAuthToken(signerKey *ecdsa.PrivateKey, operator *etherfi.Operator, governanceAddr common.Address) ([]byte, error) {

	// Wallet and AvsGovernanceAddress fields must be in EIP-55 checksummed case
	// By default common.Address does not use that casing when marshalling to json
	type AuthTokenRequest struct {
		Wallet               string `json:"wallet"`
		AvsGovernanceAddress string `json:"avsGovernanceAddress"`
		Signature            string `json:"signature"`
		IsSmartWallet        bool   `json:"isSmartWallet"`
	}
	type AuthTokenResponse struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
		Error   bool `json:"error"`
		Message any  `json:"message"`
	}

	// need to sign `operatorAddress || governanceAddress`
	message := fmt.Sprintf("%s%s", operator.Address, governanceAddr)
	hash := crypto.Keccak256([]byte(message))

	signed, err := utils.SignDigestECDSA(hash[:], signerKey)
	if err != nil {
		return nil, fmt.Errorf("signing digest: %w", err)
	}

	// build and request a signed token from the official othentic signer
	request := AuthTokenRequest{
		Wallet:               operator.Address.String(),
		AvsGovernanceAddress: governanceAddr.String(),
		Signature:            "0x" + hex.EncodeToString(signed),
		IsSmartWallet:        true,
	}
	buf, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal auth token request")
	}
	resp, err := http.Post("http://allowlist-signer.othentic.xyz/signer/sign", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return nil, fmt.Errorf("failed to acquire token from othentic signer: %w", err)
	}
	defer resp.Body.Close()

	// parse response
	var signedResponse AuthTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&signedResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal othentic response: %w", err)
	}
	if signedResponse.Error == true {
		return nil, fmt.Errorf("othentic signer returned error: %s", signedResponse.Message)
	}

	token, err := hexutil.Decode(signedResponse.Data.Token)
	if err != nil {
		return nil, fmt.Errorf("invalidToken: %w", err)
	}

	return token, nil
}
