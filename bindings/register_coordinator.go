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
	PubkeyRegistrationMessageHash(operatorAddress common.Address) (*bn254.G1Affine, error)
}

type registryCoordinator struct {
	//client ethclient.Client
	contractRegistryCoordinator *regcoordinator.ContractRegistryCoordinator
}

func (r *registryCoordinator) PubkeyRegistrationMessageHash(operatorAddress common.Address) (*bn254.G1Affine, error) {

	g1Point, err := r.contractRegistryCoordinator.PubkeyRegistrationMessageHash(nil, operatorAddress)
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
func NewRegistryCoordinator(contractAddr common.Address, client *ethclient.Client) (RegistryCoordinator, error) {

	coord, err := regcoordinator.NewContractRegistryCoordinator(contractAddr, client)
	if err != nil {
		logging.Error("RegistryCoordinator", err)
		return nil, err
	}

	return &registryCoordinator{
		contractRegistryCoordinator: coord,
	}, nil
}
