package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var updateEcdsaSignerCmd = &cli.Command{
	Name:   "update-ecdsa-signer",
	Usage:  "(Ether.fi Admin) the signer associated with this operator",
	Action: handleUpdateEcdsaSigner,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "ecdsa-signer",
			Usage:    "0x123456789...",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "broadcast",
			Usage: "broadcast signed tx to network",
		},
		&cli.StringFlag{
			Name:  "rpc-url",
			Usage: "rpc url",
		},
	},
}

func handleUpdateEcdsaSigner(ctx context.Context, cli *cli.Command) error {

	operatorID := cli.Int("operator-id")
	ecdsaSigner := common.HexToAddress(cli.String("ecdsa-signer"))
	broadcast := cli.Bool("broadcast")
	rpcURL := cli.String("rpc-url")

	// connect to RPC node
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}

	return updateEcdsaSigner(rpcClient, operatorID, ecdsaSigner, broadcast)
}

func updateEcdsaSigner(rpcClient *ethclient.Client, operatorID int64, ecdsaSigner common.Address, broadcast bool) error {

	// load configuration
	chainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("querying chainID from RPC: %w", err)
	}
	cfg, err := configForChain(chainID.Int64())
	if err != nil {
		return err
	}

	// bind contracts
	avsManager, err := contracts.NewAvsOperatorManager(cfg.OperatorManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding OperatorManager contract: %w", err)
	}

	// load signing key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("loading signing key: %w", err)
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("creating signer from key: %w", err)
	}

	if !broadcast {
		transactor.NoSend = true
	}

	tx, err := avsManager.UpdateEcdsaSigner(transactor, big.NewInt(operatorID), ecdsaSigner)
	if err != nil {
		return fmt.Errorf("executing tx: %w", err)
	}

	var buf bytes.Buffer
	tx.EncodeRLP(&buf)
	fmt.Println("TxHash:", tx.Hash().Hex())
	fmt.Println("Tx:", hex.EncodeToString(buf.Bytes()))

	return nil
}
