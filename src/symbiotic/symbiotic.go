package symbiotic

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils"
)

type API struct {
	Client *ethclient.Client

	KnownNetworks  map[common.Address]string
	KnownOperators map[common.Address]string
	KnownVaults    map[string]config.SymbioticVaultConfig
}

func New(cfg config.Config, rpcClient *ethclient.Client) *API {

	return &API{
		Client:         rpcClient,
		KnownNetworks:  cfg.SymbioticNetworks,
		KnownOperators: cfg.SymbioticOperators,
		KnownVaults:    cfg.SymbioticVaults,
	}
}

type OperatorShareUpdate struct {
	Operator common.Address
	Shares   *big.Int
}

func (a *API) SetNetworkLimit(vaultCfg config.SymbioticVaultConfig, limit *big.Int, network common.Address, shareUpdates []OperatorShareUpdate) error {

	// output in gnosis compatible format
	batch := gnosis.GnosisBatch{
		Version: "1.0",
		ChainId: "1",
		Meta:    gnosis.GnosisMetadata{Name: "symbiotic-set-network-limit"},
	}

	// manually pack tx data since we are submitting via gnosis instead of directly
	delegatorABI, err := DelegatorMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("invalid delegator abi: %v", err)
	}

	// set the new limit
	subnetwork := toSubnetwork(network)
	limitCalldata, err := delegatorABI.Pack("setNetworkLimit", subnetwork, limit)
	if err != nil {
		return fmt.Errorf("failed to pack setNetworkLimit: %w", err)
	}
	batch.AddTransaction(gnosis.SubTransaction{
		Target: vaultCfg.Delegator,
		Value:  big.NewInt(0),
		Data:   limitCalldata,
	})

	// do the requested share updates
	for _, update := range shareUpdates {
		shareCalldata, err := delegatorABI.Pack("setOperatorNetworkShares", subnetwork, update.Operator, update.Shares)
		if err != nil {
			return fmt.Errorf("failed to pack setOperatorNetworkShares: %w", err)
		}
		batch.AddTransaction(gnosis.SubTransaction{
			Target: vaultCfg.Delegator,
			Value:  big.NewInt(0),
			Data:   shareCalldata,
		})
	}

	return utils.ExportJSON("symbiotic-set-network-limit", 0, batch)
}

// convert from network address to 32byte subnetwork ID
// This effectively just appends 12 bytes of zeroes.
// Will need to update once a network actually uses the subnetwork feature
func toSubnetwork(networkAddr common.Address) [32]byte {
	var subnetwork [32]byte
	copy(subnetwork[:], networkAddr[:])
	return subnetwork
}
