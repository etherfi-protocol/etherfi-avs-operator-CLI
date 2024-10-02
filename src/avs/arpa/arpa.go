package arpa

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/eigenlayer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
)

// API handle for all core arpa functionality
type API struct {
	Client *ethclient.Client

	NodeRegistry              *NodeRegistry
	NodeRegistryAddress       common.Address
	ServiceManagerAddress     common.Address
	AvsOperatorManagerAddress common.Address

	EigenlayerAPI *eigenlayer.API
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	nodeRegistry, _ := NewNodeRegistry(cfg.ARPANodeRegistryAddress, rpcClient)

	return &API{
		Client: rpcClient,

		NodeRegistry:              nodeRegistry,
		NodeRegistryAddress:       cfg.ARPANodeRegistryAddress,
		ServiceManagerAddress:     cfg.ARPAServiceManagerAddress,
		AvsOperatorManagerAddress: cfg.AvsOperatorManagerAddress,
		EigenlayerAPI:             eigenlayer.New(cfg, rpcClient),
	}
}

// Info that node operator must supply to the ether.fi admin for registration
type RegistrationInfo struct {
	OperatorID   int64
	DKGPublicKey []byte
}

// Register is called by the node operator with their `Node Account` to register with ARPA
func (a *API) Register(operator *etherfi.Operator, dkgPublicKey []byte, inputSignature ISignatureUtilsSignatureWithSaltAndExpiry) error {

	// registration will be signed the node operator with their `Node Account`
	signingKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	// pack input for `nodeRegister` call
	nodeRegistryABI, err := NodeRegistryMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	input, err := nodeRegistryABI.Pack("nodeRegister", dkgPublicKey, true, operator.Address, inputSignature)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}

	// get the nonce and the gas values for the transaction
	fromAddress := crypto.PubkeyToAddress(signingKey.PublicKey)
	nonce, err := a.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("fetching nonce: %w", err)
	}
	gasPrice, err := a.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("fetching gas price: %w", err)
	}
	gasLimit, err := a.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &a.NodeRegistryAddress,
		Data: input,
	})
	if err != nil {
		return fmt.Errorf("estimating gas: %w", err)
	}
	// add 10% buffer to the gas limit
	finalGasLimit := gasLimit + (gasLimit / 10)

	// sign and send the transaction
	tx := types.NewTransaction(nonce, a.NodeRegistryAddress, big.NewInt(0), finalGasLimit, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), signingKey)
	if err != nil {
		return fmt.Errorf("signing transaction: %w", err)
	}
	err = a.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("sending transaction: %w", err)
	}

	txHash := signedTx.Hash().Hex()
	etherscanURL := fmt.Sprintf("https://etherscan.io/tx/%s", txHash)
	fmt.Printf("Registration sent successfully.\nView it on Etherscan: %s\n", etherscanURL)

	return nil
}

// GenerateAVSRegistrationSignature is used by the ether.fi admin to generate an AVS registration signature for a node operator
func (a *API) GenerateAVSRegistrationSignature(operator *etherfi.Operator, signingKey *ecdsa.PrivateKey) error {

	// generate and sign registration hash with admin ecdsa key
	signature, err := a.EigenlayerAPI.GenerateAndSignRegistrationDigest(operator, a.ServiceManagerAddress, signingKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}

	// convert to types expected by contract call
	sigWithSaltAndExpiry := ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature.Signature,
		Salt:      signature.Salt,
		Expiry:    signature.Expiry,
	}

	return utils.ExportJSON("arpa-registration-signature", operator.ID, sigWithSaltAndExpiry)
}
