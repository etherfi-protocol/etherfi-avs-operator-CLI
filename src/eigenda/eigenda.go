package eigenda

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/avs/signer"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/etherfi"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/utils"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// need to alias because eigenlayer has a package name that doesn't match the filepath
	registryCoordinator "github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

// API handle for all core witness chain functionality
type API struct {
	Client *ethclient.Client

	RegistryCoordinatorAddress common.Address
	RegistryCoordinator        *registryCoordinator.ContractRegistryCoordinator
	AvsOperatorManagerAddress  common.Address
}

// Info that node operator must supply to the ether.fi admin for registration
type RegistrationInfo struct {
	OperatorID                  int64
	Socket                      string
	Quorums                     []int64
	BLSPubkeyRegistrationParams *types.BLSPubkeyRegistrationParams
}

func (a *API) PrepareRegistration(operator *etherfi.Operator, blsKey *bls.KeyPair, socket string, quorums []int64) error {

	// compute hash to sign with bls key
	// the hash is converted to a G1 point on the curve before it is returned
	g1Point, err := a.RegistryCoordinator.PubkeyRegistrationMessageHash(nil, operator.Address)
	if err != nil {
		return fmt.Errorf("fetching pubkeyRegistrationMessageHash: %w", err)
	}

	// TODO: double check that the new setup is legal versus this old setup
	/*
		xElem := fp.NewElement(0)
		xElem.SetBigInt(g1Point.X)
		yElem := fp.NewElement(0)
		yElem.SetBigInt(g1Point.Y)
	*/

	// map from contract type to type expected by signing algorithm
	g1MsgToSign := &bn254.G1Affine{
		X: *new(fp.Element).SetBigInt(g1Point.X),
		Y: *new(fp.Element).SetBigInt(g1Point.Y),
	}

	// sign the registration hash (G1 Point) with the bls key to prove ownership of key
	avsSigner := signer.NewBLSSigner(blsKey)
	g1Sig, err := avsSigner.Sign(g1MsgToSign)
	if err != nil {
		return fmt.Errorf("signing pubkey registration hash: %w", err)
	}
	signedParams := new(types.BLSPubkeyRegistrationParams)
	signedParams.Load(blsKey.GetPubKeyG1().G1Affine, blsKey.GetPubKeyG2().G2Affine, g1Sig)

	// sanity check on the bls signature
	isValid, err := avsSigner.Verify(g1MsgToSign, g1Sig)
	if !isValid || err != nil {
		return fmt.Errorf("failed to verify g1 signature: %w", err)
	}

	ri := RegistrationInfo{
		OperatorID:                  operator.ID,
		Socket:                      socket,
		Quorums:                     quorums,
		BLSPubkeyRegistrationParams: signedParams,
	}
	return utils.ExportJSON("eigenda-prepare-registration", operator.ID, ri)
}
