package arpa

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	// TODO: replace with the custom signer once we build it
	SignerFn := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.LatestSignerForChainID(big.NewInt(1))
		signedTx, err := types.SignTx(tx, signer, signingKey)
		if err != nil {
			return nil, err
		}
		return signedTx, nil
	}

	transactOpts := &bind.TransactOpts{
		From:    crypto.PubkeyToAddress(signingKey.PublicKey),
		Signer:  SignerFn,
		Context: context.Background(),
	}

	transaction, err := a.NodeRegistry.NodeRegister(transactOpts, dkgPublicKey, true, operator.Address, inputSignature)
	if err != nil {
		fmt.Printf("Error details: %v\n", err)
		return fmt.Errorf("registering node: %w", err)
	}

	txHash := transaction.Hash().Hex()
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
