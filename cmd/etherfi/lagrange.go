package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var lagrangeCmd = &cli.Command{
	Name:  "lagrange",
	Usage: "various actions related to managing lagrange operators",
	Commands: []*cli.Command{
		lagrangeRegisterCmd,
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
			Usage:    "(Required) Registration digest signature",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "salt",
			Usage:    "(Required) Salt used for ECDSA signature",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "expiry",
			Usage:    "(Required) Expiry used for ECDSA signature",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "broadcast",
			Usage: "broadcast signed tx to network",
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

	return lagrangeRegister(operatorID, signAddress, pubkeyBytes, signature, salt, expiry, rpcClient)
}

func lagrangeRegister(
	operatorID int64,
	signAddress common.Address,
	blsPubkey []byte,
	registrationSignature []byte,
	salt []byte,
	expiry int64,
	rpcClient *ethclient.Client,
) error {

	// load configuration
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("querying chainID from RPC: %w", err)
	}
	/*
		cfg, err := configForChain(chainID.Int64())
		if err != nil {
			return err
		}
	*/

	/*
		// bind contracts
		lagrangeService, err := contracts.NewLagrangeService(cfg.LagrangeService, rpcClient)
		if err != nil {
			return fmt.Errorf("binding LagrangeService contract: %w", err)
		}
	*/

	// load key for signing tx
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("loading signing key: %w", err)
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("creating signer from key: %w", err)
	}
	// toggle whether tx is broadcast to network
	// TODO: FIX
	//transactor.NoSend = !cli.Bool("broadcast")
	transactor.NoSend = true

	// convert to format expected by lagrange
	if len(blsPubkey) != 64 {
		return fmt.Errorf("invalid pubkey")
	}
	var blsPubkeys = make([][2]*big.Int, 1)
	blsPubkeys[0][0] = big.NewInt(0).SetBytes(blsPubkey[:32])
	blsPubkeys[0][1] = big.NewInt(0).SetBytes(blsPubkey[32:64])

	sigParams := contracts.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: registrationSignature,
		Salt:      [32]byte(salt),
		Expiry:    big.NewInt(int64(expiry)),
	}

	lagrangeABI, err := contracts.LagrangeServiceMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	registerSelector := lagrangeABI.Methods["register"].ID
	fmt.Println("lagrangeMethod:", hex.EncodeToString(registerSelector))

	input, err := lagrangeABI.Pack("register", signAddress, blsPubkeys, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}
	fmt.Println("register:", hex.EncodeToString(input))
	fmt.Println()

	managerABI, err := contracts.AvsOperatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	// pack forward call for manager
	target := common.HexToAddress("0x35F4f28A8d3Ff20EEd10e087e8F96Ea2641E6AA2")
	//selector := managerABI.Methods["adminForwardCall"].ID
	encodedForwardData, err := managerABI.Pack("adminForwardCall", big.NewInt(operatorID), target, [4]byte(registerSelector), input[4:])
	if err != nil {
		return fmt.Errorf("packing forwardCall: %w", err)
	}
	fmt.Println("adminForwardCall:", hex.EncodeToString(encodedForwardData))

	//	lagrangeService.Register(transactor, signAddress, blsPubkeys, sigParams)

	return nil
}
