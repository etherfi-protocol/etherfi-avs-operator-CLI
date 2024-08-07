package signer

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
)

type AVS int

const (
	UNKNOWN AVS = iota
	EIGEN_DA
	BREVIS
	LAGRANGE
	EORACLE
	WITNESS_CHAIN
)

func GenerateRegistrationDigest(operatorID int64, salt [32]byte, expiry *big.Int, avs AVS, rpcClient *ethclient.Client) ([]byte, error) {

	// load configuration
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("querying chainID from RPC: %w", err)
	}
	cfg, err := bindings.ConfigForChain(chainID.Int64())
	if err != nil {
		return nil, err
	}

	// bind contracts
	avsDirectory, err := contracts.NewAVSDirectory(cfg.AvsDirectoryAddress, rpcClient)
	if err != nil {
		return nil, fmt.Errorf("binding registryCoordinator: %w", err)
	}
	avsManager, err := contracts.NewAvsOperatorManager(cfg.AvsOperatorManagerAddress, rpcClient)
	if err != nil {
		return nil, fmt.Errorf("binding OperatorManager contract: %w", err)
	}

	// look up address of operator contract
	operatorAddr, err := avsManager.AvsOperators(nil, big.NewInt(operatorID))
	if err != nil {
		return nil, fmt.Errorf("fetching operator address: %w", err)
	}

	var serviceManagerAddress common.Address
	switch avs {
	case EIGEN_DA:
		serviceManagerAddress = cfg.EigenDAServiceManagerAddress
	case BREVIS:
		serviceManagerAddress = cfg.BrevisServiceManagerAddress
	case LAGRANGE:
		serviceManagerAddress = cfg.LagrangeServiceAddress
	case EORACLE:
		serviceManagerAddress = cfg.EOracleServiceManagerAddress
	case WITNESS_CHAIN:
		serviceManagerAddress = cfg.WitnessChainWitnessHubAddress
	default:
		panic("unknown avs")
	}

	hash, err := avsDirectory.CalculateOperatorAVSRegistrationDigestHash(nil, operatorAddr, serviceManagerAddress, [32]byte(salt), expiry)
	if err != nil {
		return nil, fmt.Errorf("requesting hash: %w", err)
	}

	return hash[:], nil
}

func GenerateAndSignRegistrationDigest(operatorID int64, avs AVS, rpcClient *ethclient.Client, privKey *ecdsa.PrivateKey) (*types.SignatureWithSaltAndExpiry, error) {

	// load configuration
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("querying chainID from RPC: %w", err)
	}
	cfg, err := bindings.ConfigForChain(chainID.Int64())
	if err != nil {
		return nil, err
	}

	// look up operator contract associated with this id and configured ecdsaSigner
	operatorManagerContract, err := contracts.NewAvsOperatorManager(cfg.AvsOperatorManagerAddress, rpcClient)
	if err != nil {
		return nil, fmt.Errorf("binding operatorManager: %w", err)
	}
	operatorAddr, err := operatorManagerContract.AvsOperators(nil, big.NewInt(operatorID))
	if err != nil {
		return nil, fmt.Errorf("looking up operator address: %w", err)
	}
	operatorContract, err := contracts.NewEtherfiAVSOperator(operatorAddr, rpcClient)
	if err != nil {
		return nil, fmt.Errorf("binding operator contract: %w", err)
	}
	ecdsaSigner, err := operatorContract.EcdsaSigner(nil)
	if err != nil {
		return nil, fmt.Errorf("fetching configured ecdsaSigner: %w", err)
	}

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("gererating random salt: %w", err)
	}
	// TODO: revisit what we want the expiration to be
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 10).Unix())

	hash, err := GenerateRegistrationDigest(operatorID, [32]byte(salt), expiry, avs, rpcClient)
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
	if keyAddress != ecdsaSigner {
		color.HiYellow("WARNING: configured signing key does not match the current ecdsaSigner for operator")
		fmt.Printf("signer: %s, ecdsaSigner: %s\n\n", keyAddress, ecdsaSigner)
	}

	// sign the digest
	signed, err := crypto.Sign(hash[:], privKey)
	if err != nil {
		return nil, fmt.Errorf("signing digest: %w", err)
	}

	// account for EIP-155
	signed[64] += 27

	return &types.SignatureWithSaltAndExpiry{
		Signature: signed,
		Salt:      [32]byte(salt),
		Expiry:    expiry,
	}, nil
}
