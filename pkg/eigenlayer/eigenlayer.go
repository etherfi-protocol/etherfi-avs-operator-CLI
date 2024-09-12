package eigenlayer

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigenlayer-contracts/pkg/bindings/AVSDirectory"
	"github.com/Layr-Labs/eigenlayer-contracts/pkg/bindings/DelegationManager"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/config"
	"github.com/mantle-lsp/mantle-avs-operator-CLI/pkg/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EigenLayer struct {
	Client *ethclient.Client

	AvsDirectory        *AVSDirectory.AVSDirectory
	AvsDirectoryAddress common.Address

	DelegationManager        *DelegationManager.DelegationManager
	DelegationManagerAddress common.Address
}

func New(cfg config.Config, rpcClient *ethclient.Client) *EigenLayer {

	avsDirectory, _ := AVSDirectory.NewAVSDirectory(cfg.AvsDirectoryAddress, rpcClient)
	delegationManager, _ := DelegationManager.NewDelegationManager(cfg.DelegationManagerAddress, rpcClient)

	return &EigenLayer{
		Client:                   rpcClient,
		AvsDirectory:             avsDirectory,
		AvsDirectoryAddress:      cfg.AvsDirectoryAddress,
		DelegationManager:        delegationManager,
		DelegationManagerAddress: cfg.DelegationManagerAddress,
	}
}

func (a *EigenLayer) GenerateAndSignRegistrationDigest(operator common.Address, serviceManager common.Address) (*types.SignatureWithSaltAndExpiry, error) {

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("generating random salt: %w", err)
	}

	// TODO: revisit what we want the expiration to be
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 7).Unix())

	hash, err := a.AvsDirectory.CalculateOperatorAVSRegistrationDigestHash(nil, operator, serviceManager, [32]byte(salt), expiry)
	if err != nil {
		return nil, fmt.Errorf("requesting hash: %w", err)
	}

	return &types.SignatureWithSaltAndExpiry{
		Signature: hash[:],
		Salt:      [32]byte(salt),
		Expiry:    expiry,
	}, nil
}
