package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/bindings"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/avs/lagrange"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/gnosis"
	"github.com/urfave/cli/v3"
)

var lagrangeCmd = &cli.Command{
	Name:  "lagrange",
	Usage: "various actions related to managing lagrange operators",
	Commands: []*cli.Command{
		lagrangeRegisterCmd,
		lagrangeSubscribeCmd,
	},
}

var lagrangeRegisterCmd = &cli.Command{
	Name:   "register",
	Action: handleLagrangeRegister,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "sign-address",
			Usage:    "separate ecdsa key address for lagrange operation. See their docs for details",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "bls-pubkey",
			Usage:    "BN254 bls pubkey in hex format",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "signature",
			Usage:    "Registration digest signature",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "salt",
			Usage:    "Salt used for ECDSA signature",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "expiry",
			Usage:    "Expiry used for ECDSA signature",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "gnosis",
			Usage: "output the transaction as gnosis compatible json",
		},
	},
}

func handleLagrangeRegister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	operatorID := cli.Int("operator-id")
	rpcURL := cli.String("rpc-url")
	signAddress := common.HexToAddress(cli.String("sign-address"))
	pubkeyHex := cli.String("bls-pubkey")
	expiry := cli.Int("expiry")
	outputGnosis := cli.Bool("gnosis")
	signature, err := hex.DecodeString(strings.TrimPrefix(cli.String("signature"), "0x"))
	if err != nil {
		return fmt.Errorf("invalid signature: %w", err)
	}
	salt, err := hex.DecodeString(strings.TrimPrefix(cli.String("salt"), "0x"))
	if err != nil {
		return fmt.Errorf("invalid salt: %w", err)
	}

	// connect to RPC node
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}

	pubkeyBytes, err := hex.DecodeString(pubkeyHex)
	if err != nil {
		return fmt.Errorf("parsing pubkey: %w", err)
	}

	return lagrangeRegister(operatorID, signAddress, pubkeyBytes, signature, salt, expiry, rpcClient, outputGnosis)
}

func lagrangeRegister(
	operatorID int64,
	signAddress common.Address,
	blsPubkey []byte,
	registrationSignature []byte,
	salt []byte,
	expiry int64,
	rpcClient *ethclient.Client,
	outputGnosis bool,
) error {

	// load configuration
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("querying chainID from RPC: %w", err)
	}
	cfg, err := bindings.ConfigForChain(chainID.Int64())
	if err != nil {
		return err
	}

	// convert to format expected by lagrange
	if len(blsPubkey) != 64 {
		return fmt.Errorf("invalid pubkey")
	}
	var blsPubkeys = make([][2]*big.Int, 1)
	blsPubkeys[0][0] = big.NewInt(0).SetBytes(blsPubkey[:32])
	blsPubkeys[0][1] = big.NewInt(0).SetBytes(blsPubkey[32:64])

	sigParams := lagrange.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: registrationSignature,
		Salt:      [32]byte(salt),
		Expiry:    big.NewInt(int64(expiry)),
	}

	lagrangeABI, err := lagrange.LagrangeServiceMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	// pack LagrangeService.register()
	input, err := lagrangeABI.Pack("register", signAddress, blsPubkeys, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}
	fmt.Printf("subcall: 0x%s\n\n", hex.EncodeToString(input))

	managerABI, err := etherfi.AvsOperatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	// pack AvsOperatorManager.adminForwardCall()
	subcallTarget := cfg.LagrangeServiceAddress
	subcallSelector := [4]byte(input[:4])
	subcallData := input[4:]
	encodedForwardData, err := managerABI.Pack("adminForwardCall", big.NewInt(operatorID), subcallTarget, subcallSelector, subcallData)
	if err != nil {
		return fmt.Errorf("packing forwardCall: %w", err)
	}
	fmt.Printf("adminForwardCall: 0x%s\n\n", hex.EncodeToString(encodedForwardData))

	if outputGnosis {

		batch := gnosis.NewSingleTxBatch(encodedForwardData, cfg.AvsOperatorManagerAddress, fmt.Sprintf("lagrange-register-%d", operatorID))
		fmt.Printf("gnosis:\n%s\n", batch.PrettyPrint())
	}

	return nil
}

var lagrangeSubscribeCmd = &cli.Command{
	Name:   "subscribe",
	Action: handleLagrangeSubscribe,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "rollup-chain-id",
			Usage:    "chainID of rollup chain you are subscribing to",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "gnosis",
			Usage: "output the transaction as gnosis compatible json",
		},
	},
}

func handleLagrangeSubscribe(ctx context.Context, cli *cli.Command) error {
	// parse cli params
	operatorID := cli.Int("operator-id")
	rollupChainID := cli.Int("rollup-chain-id")
	rpcURL := cli.String("rpc-url")

	// load configuration
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("querying chainID from RPC: %w", err)
	}
	cfg, err := bindings.ConfigForChain(chainID.Int64())
	if err != nil {
		return err
	}

	return lagrangeSubscribe(operatorID, uint32(rollupChainID), cfg)
}

func lagrangeSubscribe(operatorID int64, rollupChainID uint32, cfg *bindings.Config) error {

	lagrangeABI, err := lagrange.LagrangeServiceMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	managerABI, err := etherfi.AvsOperatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	// pack LagrangeService.subscribe()
	input, err := lagrangeABI.Pack("subscribe", rollupChainID)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}
	fmt.Printf("subcall: 0x%s\n\n", hex.EncodeToString(input))

	// pack AvsOperatorManager.adminForwardCall()
	subcallTarget := cfg.LagrangeServiceAddress
	subcallSelector := [4]byte(input[:4])
	subcallData := input[4:]
	encodedForwardData, err := managerABI.Pack("adminForwardCall", big.NewInt(operatorID), subcallTarget, subcallSelector, subcallData)
	if err != nil {
		return fmt.Errorf("packing forwardCall: %w", err)
	}
	fmt.Printf("adminForwardCall: 0x%s\n\n", hex.EncodeToString(encodedForwardData))

	batch := gnosis.NewSingleTxBatch(encodedForwardData, cfg.AvsOperatorManagerAddress, fmt.Sprintf("lagrange-subscribe-%d", operatorID))
	fmt.Printf("gnosis:\n%s\n", batch.PrettyPrint())

	return nil
}
