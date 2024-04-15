package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

func registrationDigest(ctx context.Context, cli *cli.Command) error {

	// parse flags
	registryCoordinatorAddress := common.HexToAddress(cli.String("registry-coordinator"))
	operatorID := big.NewInt(cli.Int("operator-id"))
	chainID := cli.Int("chain-id")
	shouldSign := cli.Bool("sign")

	// connect to RPC node
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return fmt.Errorf("dialing RPC: %w", err)
	}

	var cfg bindings.Config
	switch chainID {
	case 1:
		cfg = bindings.Mainnet
	case 17000:
		cfg = bindings.Holesky
	default:
		return fmt.Errorf("unimplemented chain: %d", chainID)
	}

	// bind rpc to contract abi
	operatorManagerContract, err := contracts.NewEtherfiAVSOperatorsManager(cfg.OperatorManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding operatorManager: %w", err)
	}
	registryCoordinatorContract, err := contracts.NewRegistryCoordinator(registryCoordinatorAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding registryCoordinator: %w", err)
	}

	// need to find associated service manager
	serviceManagerAddress, err := registryCoordinatorContract.ServiceManager(nil)
	if err != nil {
		return fmt.Errorf("finding associated service manager: %w", err)
	}

	// TODO: option in pass salt as a param
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("gererating random salt: %w", err)
	}
	// TODO: option for expiry as a param
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 365).Unix())

	hash, err := operatorManagerContract.CalculateOperatorAVSRegistrationDigestHash(nil, operatorID, serviceManagerAddress, [32]byte(salt), expiry)
	if err != nil {
		return fmt.Errorf("requesting hash: %w", err)
	}

	// sign if requested
	if shouldSign {

		// look up operator contract associated with this id and configured ecdsaSigner
		operatorAddr, err := operatorManagerContract.AvsOperators(nil, operatorID)
		if err != nil {
			return fmt.Errorf("looking up operator address: %w", err)
		}
		operatorContract, err := contracts.NewEtherfiAVSOperator(operatorAddr, rpcClient)
		if err != nil {
			return fmt.Errorf("binding operator contract: %w", err)
		}
		ecdsaSigner, err := operatorContract.EcdsaSigner(nil)
		if err != nil {
			return fmt.Errorf("fetching configured ecdsaSigner: %w", err)
		}

		// ensure the provided key matches the ecdsaSigner the operator contract is expecting
		privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
		if err != nil {
			return fmt.Errorf("invalid private key: %w", err)
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("private key is unexpected type: %w", err)
		}
		keyAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		if keyAddress != ecdsaSigner {
			return fmt.Errorf("configured signing key does not match expected ecdsaSigner for operatorContract (%s | %s)", keyAddress, ecdsaSigner)
		}

		// sign the digest
		signed, err := crypto.Sign(hash[:], privateKey)
		if err != nil {
			return fmt.Errorf("signing digest: %w", err)
		}

		// account for EIP-155
		signed[64] += 27
		fmt.Printf("signature: %s\n", hex.EncodeToString(signed))

	}

	fmt.Printf("hash: %s\n", hex.EncodeToString(hash[:]))
	fmt.Printf("salt: %s\n", hex.EncodeToString(salt))
	fmt.Printf("Expiry: %s\n", expiry)

	return nil
}
