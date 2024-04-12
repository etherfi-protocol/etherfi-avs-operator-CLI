package bindings

import (
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"

	regcoordinator "github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/logging"
)

// RegisterCoordinator is a wrapper around the RegistryCoordinator contract
type RegistryCoordinator interface {
	PubkeyRegistrationMessageHash(operatorAddress string) (*bn254.G1Affine, error)
}

type registryCoordinator struct {
	//client ethclient.Client
	contractRegistryCoordinator *regcoordinator.ContractRegistryCoordinator
}

func (r *registryCoordinator) PubkeyRegistrationMessageHash(operatorAddress string) (*bn254.G1Affine, error) {
	logging.Info("RegistryCoordinator", "call NewEigenDAServiceManager")
	logging.Info("RegistryCoordinator", "operator address", operatorAddress)

	addr := common.HexToAddress(operatorAddress)
	g1Point, err := r.contractRegistryCoordinator.PubkeyRegistrationMessageHash(nil, addr)
	if err != nil {
		logging.Error("RegistryCoordinator", err)
		return nil, err
	}

	xElem := fp.NewElement(0)
	xElem.SetBigInt(g1Point.X)
	yElem := fp.NewElement(0)
	yElem.SetBigInt(g1Point.Y)

	g1Affine := bn254.G1Affine{
		X: xElem,
		Y: yElem,
	}

	return &g1Affine, nil
}

// NewRegisterCoordinator creates a new RegisterCoordinator
func NewRegistryCoordinator(contractAddr string, client *ethclient.Client) (RegistryCoordinator, error) {
	logging.Info("RegistryCoordinator", "call NewEigenDAServiceManager")
	logging.Info("RegistryCoordinator", "contract address", contractAddr)

	regCoordContract := common.HexToAddress(contractAddr)
	coord, err := regcoordinator.NewContractRegistryCoordinator(regCoordContract, client)
	if err != nil {
		logging.Error("RegistryCoordinator", err)
		return nil, err
	}

	return &registryCoordinator{
		contractRegistryCoordinator: coord,
	}, nil
}
