package bindings

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	servicemanager "github.com/Layr-Labs/eigenda/contracts/bindings/EigenDAServiceManager"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/logging"
)

// EigenDAServiceManager is a wrapper around the EigenDAServiceManager contract
type EigenDAServiceManager interface {
	RegistryCoordinator() (RegistryCoordinator, error)
}

type eigenDAServiceManager struct {
	contractEigenDAServiceManager *servicemanager.ContractEigenDAServiceManager
	ethClient                     *ethclient.Client
}

func (m *eigenDAServiceManager) RegistryCoordinator() (RegistryCoordinator, error) {
	logging.Info("eigenDAServiceManager", "call RegisterCoordinator")

	addr, err := m.contractEigenDAServiceManager.RegistryCoordinator(nil)
	if err != nil {
		logging.Error("eigenDAServiceManager", err)
		return nil, err
	}

	logging.Info("eigenDAServiceManager", "addr "+addr.String())

	return NewRegistryCoordinator(addr.String(), m.ethClient)
}

// NewEigenDAServiceManager creates a new EigenDAServiceManager
func NewEigenDAServiceManager(contractAddr string, client *ethclient.Client) (EigenDAServiceManager, error) {
	logging.Info("eigenDAServiceManager", "call NewEigenDAServiceManager")
	logging.Info("eigenDAServiceManager", "address", contractAddr)

	eigenDAServiceManagerContract := common.HexToAddress(contractAddr)
	contract, err := servicemanager.NewContractEigenDAServiceManager(
		eigenDAServiceManagerContract, client)
	if err != nil {
		logging.Error("eigenDAServiceManager", err)
		return nil, err
	}

	return &eigenDAServiceManager{
		ethClient:                     client,
		contractEigenDAServiceManager: contract,
	}, nil
}
