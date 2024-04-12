package bindings

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"time"

	avsdir "github.com/Layr-Labs/eigenda/contracts/bindings/AVSDirectory"
	regcoordinator "github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/keystore"
)

const (
	etherfiOperator            = "0x53F69255A16E6a924665EB839f2730bFF01BE7D8"
	eigenlayerAVSDirectory     = "0x055733000064333CaDDbC92763c58BF0192fFeBf"
	eigendaServiceManager      = "0xd4a7e1bd8015057293f0d0a557088c286942e84b"
	eigendaRegistryCoordinator = "0x53012C69A189cfA2D9d29eb6F19B32e0A2EA3490"
)

func TestBinding(t *testing.T) {
	ks := keystore.NewKeystore()
	keyPair, err := ks.LoadBLS("../keystore/fixtures/fixture.bls.key.json", "fixture@test1234")

	assert.Nil(t, err)

	ecdsaKey, err := ks.LoadECDSA("../keystore/fixtures/fixture.ecdsa.key.json", "fixture@test1234")

	assert.Nil(t, err)

	url := "https://ethereum-holesky-rpc.publicnode.com"
	client, err := ethclient.Dial(url)

	assert.Nil(t, err)

	regCoordContract := common.HexToAddress(eigendaRegistryCoordinator)
	coord, err := regcoordinator.NewContractRegistryCoordinator(regCoordContract, client)
	assert.Nil(t, err)

	operatorAddr := common.HexToAddress(etherfiOperator)
	g1Point, err := coord.PubkeyRegistrationMessageHash(nil, operatorAddr)
	fmt.Println(g1Point.X.String(), g1Point.Y.String())

	xElem := fp.NewElement(0)
	xElem.SetBigInt(g1Point.X)
	yElem := fp.NewElement(0)
	yElem.SetBigInt(g1Point.Y)

	g1PointMsg := bn254.G1Affine{
		X: xElem,
		Y: yElem,
	}

	blsMsgSig := keyPair.SignHashedToCurveMessage(&g1PointMsg)

	fmt.Println(blsMsgSig)
	fmt.Println(keyPair.GetPubKeyG1().X.String())
	fmt.Println(keyPair.GetPubKeyG1().Y.String())
	fmt.Println(keyPair.GetPubKeyG1().X.BigInt(big.NewInt(0)))

	blsG1Point := regcoordinator.BN254G1Point{
		X: keyPair.GetPubKeyG1().X.BigInt(big.NewInt(0)),
		Y: keyPair.GetPubKeyG1().Y.BigInt(big.NewInt(0)),
	}

	blsG2Point := regcoordinator.BN254G2Point{
		X: [2]*big.Int{keyPair.GetPubKeyG2().X.A0.BigInt(big.NewInt(0)), keyPair.GetPubKeyG2().X.A1.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{keyPair.GetPubKeyG2().Y.A0.BigInt(big.NewInt(0)), keyPair.GetPubKeyG2().Y.A1.BigInt(big.NewInt(0))},
	}

	sigG1Point := regcoordinator.BN254G1Point{
		X: blsMsgSig.X.BigInt(big.NewInt(0)),
		Y: blsMsgSig.Y.BigInt(big.NewInt(0)),
	}

	params := regcoordinator.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyG1:                    blsG1Point,
		PubkeyG2:                    blsG2Point,
		PubkeyRegistrationSignature: sigG1Point,
	}

	_ = blsG1Point
	_ = blsG2Point
	_ = params

	salt := [32]byte{}
	_salt := crypto.Keccak256([]byte("test")) // FIXME: random salt
	copy(salt[:], _salt)

	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 365).Unix())

	avsDirectory, err := avsdir.NewContractAVSDirectory(common.HexToAddress(eigenlayerAVSDirectory), client)

	eigendaServiceManager := common.HexToAddress(eigendaServiceManager)
	msgToSign, err := avsDirectory.CalculateOperatorAVSRegistrationDigestHash(nil, operatorAddr, eigendaServiceManager, salt, expiry)

	assert.Nil(t, err)

	fmt.Println("MsgToSign", hex.EncodeToString(msgToSign[:]), err)

	sign, err := crypto.Sign(msgToSign[:], ecdsaKey)

	assert.Nil(t, err)

	fmt.Println("Sign", hex.EncodeToString(sign), err)

	sign[64] += 27
	oprSig := regcoordinator.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sign,
		Salt:      salt,
		Expiry:    expiry,
	}

	socket := "50.18.13.221:32005;32004"
	quorums := []byte{0x0}

	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaKey, big.NewInt(17000))

	assert.Nil(t, err)

	auth.NoSend = false

	fmt.Printf("Auth: %+v\n", auth)
	fmt.Printf("quorums: %+v\n", quorums)
	fmt.Printf("socket: %+v\n", socket)
	fmt.Printf("params: %+v\n", params)
	fmt.Printf("oprSig: %+v\n", oprSig)

	//tx, err := coord.RegisterOperator(auth, quorums, socket, params, oprSig)

	tx, err := coord.UpdateSocket(auth, socket)

	fmt.Println(tx, err)
}

func TestBinding_RegisterCoordinator(t *testing.T) {
	rpc := "https://ethereum-holesky-rpc.publicnode.com"

	client, err := ethclient.Dial(rpc)

	assert.Nil(t, err)

	holeskyEigenDAServiceManger := "0xd4a7e1bd8015057293f0d0a557088c286942e84b"
	daServiceManager, err := NewEigenDAServiceManager(holeskyEigenDAServiceManger, client)

	assert.Nil(t, err)

	regCoord, err := daServiceManager.RegistryCoordinator()

	assert.Nil(t, err)

	g1Msg, err := regCoord.PubkeyRegistrationMessageHash(etherfiOperator)

	assert.Nil(t, err)

	fmt.Println(g1Msg)
}
