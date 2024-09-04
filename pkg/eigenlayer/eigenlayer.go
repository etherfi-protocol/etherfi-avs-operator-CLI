package eigenlayer

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/a41-official/mantle-eigenlayer-operator-cli/pkg/config"
	"github.com/a41-official/mantle-eigenlayer-operator-cli/pkg/types"
	"github.com/a41-official/mantle-eigenlayer-operator-cli/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
)

type API struct {
	Client *ethclient.Client

	AvsDirectory        *AvsDirectory
	AvsDirectoryAddress common.Address

	DelegationManager        *DelegationManager
	DelegationManagerAddress common.Address
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	avsDirectory, _ := NewAvsDirectory(cfg.AvsDirectoryAddress, rpcClient)
	delegationManager, _ := NewDelegationManager(cfg.DelegationManagerAddress, rpcClient)

	return &API{
		Client:                   rpcClient,
		AvsDirectory:             avsDirectory,
		AvsDirectoryAddress:      cfg.AvsDirectoryAddress,
		DelegationManager:        delegationManager,
		DelegationManagerAddress: cfg.DelegationManagerAddress,
	}
}

func (a *API) GenerateRegistrationDigest(operator common.Address, salt [32]byte, expiry *big.Int, serviceManager common.Address) ([]byte, error) {

	hash, err := a.AvsDirectory.CalculateOperatorAVSRegistrationDigestHash(nil, operator, serviceManager, [32]byte(salt), expiry)
	if err != nil {
		return nil, fmt.Errorf("requesting hash: %w", err)
	}

	return hash[:], nil
}

func (a *API) GenerateAndSignRegistrationDigest(operator common.Address, serviceManager common.Address, privKey *ecdsa.PrivateKey) (*types.SignatureWithSaltAndExpiry, error) {

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("generating random salt: %w", err)
	}
	
	// TODO: revisit what we want the expiration to be
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 7).Unix())

	hash, err := a.GenerateRegistrationDigest(operator, [32]byte(salt), expiry, serviceManager)
	if err != nil {
		return nil, fmt.Errorf("generating registration digest: %w", err)
	}

	// ensure the provided key matches the ecdsaSigner the operator contract is expecting
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("private key is unexpected type: %w", err)
	}
	keyAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if keyAddress != operator {
		color.HiYellow("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		color.HiYellow("WARNING: configured signing key does not match the current ecdsaSigner for operator")
		color.HiYellow("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Printf("signer: %s, ecdsaSigner: %s\n\n", keyAddress, operator)
	}

	// sign the digest
	signed, err := utils.SignDigestECDSA(hash[:], privKey)
	if err != nil {
		return nil, fmt.Errorf("signing digest: %w", err)
	}

	return &types.SignatureWithSaltAndExpiry{
		Signature: signed,
		Salt:      [32]byte(salt),
		Expiry:    expiry,
	}, nil
}
