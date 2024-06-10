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

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

var registrationDigestCmd = &cli.Command{
	Name:   "registration-digest",
	Usage:  "compute and optionally sign registration digest for registering an operator",
	Action: handleRegistrationDigest,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "registry-coordinator",
			Usage:    "address of the registry-coordinator associated with the avs your are registering to",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "sign",
			Usage: "sign the created digest with configured private key",
		},
	},
}

func handleRegistrationDigest(ctx context.Context, cli *cli.Command) error {

	operatorID := cli.Int("operator-id")
	registryCoordinator := common.HexToAddress(cli.String("registry-coordinator"))
	shouldSign := cli.Bool("sign")
	rpcURL := cli.String("rpc-url")

	// connect to RPC node
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}

	return registrationDigest(operatorID, registryCoordinator, shouldSign, rpcClient)
}

func registrationDigest(operatorID int64, registryCoordinator common.Address, shouldSign bool, rpcClient *ethclient.Client) error {

	// load configuration
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("querying chainID from RPC: %w", err)
	}
	cfg, err := configForChain(chainID.Int64())
	if err != nil {
		return err
	}

	// bind rpc to contract abi
	operatorManagerContract, err := contracts.NewAvsOperatorManager(cfg.OperatorManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding operatorManager: %w", err)
	}
	registryCoordinatorContract, err := contracts.NewRegistryCoordinator(registryCoordinator, rpcClient)
	if err != nil {
		return fmt.Errorf("binding registryCoordinator: %w", err)
	}

	// TODO: better system
	var serviceManagerAddress common.Address
	if registryCoordinator == common.HexToAddress("0x35F4f28A8d3Ff20EEd10e087e8F96Ea2641E6AA2") { // lagrange registryCoordinator equivalent
		serviceManagerAddress = common.HexToAddress("0x35F4f28A8d3Ff20EEd10e087e8F96Ea2641E6AA2")
	} else {
		// need to find associated service manager
		serviceManagerAddress, err = registryCoordinatorContract.ServiceManager(nil)
		if err != nil {
			return fmt.Errorf("finding associated service manager: %w", err)
		}
	}

	// TODO: option in pass salt as a param
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("gererating random salt: %w", err)
	}
	// TODO: option for expiry as a param
	expiry := new(big.Int).SetInt64(time.Now().Add(24 * time.Hour * 365).Unix())

	hash, err := operatorManagerContract.CalculateOperatorAVSRegistrationDigestHash(nil, big.NewInt(operatorID), serviceManagerAddress, [32]byte(salt), expiry)
	if err != nil {
		return fmt.Errorf("requesting hash: %w", err)
	}

	// sign if requested
	if shouldSign {

		// look up operator contract associated with this id and configured ecdsaSigner
		operatorAddr, err := operatorManagerContract.AvsOperators(nil, big.NewInt(operatorID))
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
			color.HiYellow("WARNING: configured signing key does not match the current ecdsaSigner for operator")
			fmt.Printf("signer: %s, ecdsaSigner: %s\n\n", keyAddress, ecdsaSigner)
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
