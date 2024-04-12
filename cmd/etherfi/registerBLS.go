package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var (
	registryCoordinatorMainnet = common.HexToAddress("0x0BAAc79acD45A023E19345c352d8a7a83C4e5656")
	registryCoordinatorHolesky = common.HexToAddress("0x53012C69A189cfA2D9d29eb6F19B32e0A2EA3490")
	operatorManagerMainnet     = common.HexToAddress("0x2093Bbb221f1d8C7c932c32ee28Be6dEe4a37A6a")
	operatorManagerHolesky     = common.HexToAddress("0xdf9679e8bfce22ae503fd2726cb1218a18cd8bf4")
)

// Needs a bunch of refactoring
func registerBLS(ctx context.Context, cli *cli.Command) error {

	// TODO: network selection
	chainID := 17000

	// load key for signing tx
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("loading signing key: %w", err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainID)))
	if err != nil {
		return fmt.Errorf("creating signer from key: %w", err)
	}

	// toggle whether tx is broadcast to network
	transactor.NoSend = !cli.Bool("broadcast")

	// connect to RPC node
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}

	// bind rpc to contract abi
	operatorContract, err := bindings.NewAVSOperator(operatorManagerHolesky, rpcClient)
	if err != nil {
		return fmt.Errorf("binding contract: %w", err)
	}

	// TODO: validate params
	operatorID := big.NewInt(cli.Int("operator-id"))
	registryCoordinator := registryCoordinatorHolesky // TODO: network selection
	socket := cli.String("socket")
	var quorumNumbers []uint8
	for _, v := range cli.IntSlice("quorum-numbers") {
		quorumNumbers = append(quorumNumbers, uint8(v))
	}

	// load bls signature and convert to format expected by contract
	params, err := BLSJsonToRegistrationParams(cli.String("bls-signature-file"))
	if err != nil {
		return fmt.Errorf("barsing bls signature file: %w", err)
	}

	// Sign transaction and broadcast if requested
	tx, err := operatorContract.RegisterBlsKeyAsDelegatedNodeOperator(transactor, operatorID, registryCoordinator, quorumNumbers, socket, *params)
	if err != nil {
		return fmt.Errorf("failed to sign and/or broadcast tx: %w", err)
	}
	var buf bytes.Buffer
	tx.EncodeRLP(&buf)
	fmt.Printf("raw tx: %s\n", hex.EncodeToString(buf.Bytes()))

	return nil
}

func BLSJsonToRegistrationParams(filepath string) (*bindings.IBLSApkRegistryPubkeyRegistrationParams, error) {

	buf, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading signature file: %w", err)
	}

	var input types.AVSBLSSignature
	if err := json.Unmarshal(buf, &input); err != nil {
		return nil, fmt.Errorf("unmarshalling signature file: %w", err)
	}

	g1x, _ := big.NewInt(0).SetString(input.G1.X, 10)
	g1y, _ := big.NewInt(0).SetString(input.G1.Y, 10)

	g2x0, _ := big.NewInt(0).SetString(input.G2.X[0], 10)
	g2x1, _ := big.NewInt(0).SetString(input.G2.X[1], 10)
	g2y0, _ := big.NewInt(0).SetString(input.G2.Y[0], 10)
	g2y1, _ := big.NewInt(0).SetString(input.G2.Y[1], 10)

	sigX, _ := big.NewInt(0).SetString(input.Signature.X, 10)
	sigY, _ := big.NewInt(0).SetString(input.Signature.Y, 10)

	contractParams := bindings.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: bindings.BN254G1Point{
			X: sigX,
			Y: sigY,
		},
		PubkeyG1: bindings.BN254G1Point{
			X: g1x,
			Y: g1y,
		},
		PubkeyG2: bindings.BN254G2Point{
			X: [2]*big.Int{g2x0, g2x1},
			Y: [2]*big.Int{g2y0, g2y1},
		},
	}

	return &contractParams, nil
}
