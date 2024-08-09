package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	// need to alias because eigenlayer has a package name that doesn't match the filepath
	registryCoordinator "github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/config"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/etherfi"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/keystore"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/src/utils/signer"
	"github.com/etherfi-protocol/etherfi-avs-operator-tool/types"
	"github.com/urfave/cli/v3"
)

var BlsPubkeyRegistrationHashCmd = &cli.Command{
	Name:        "bls-pubkey-registration-hash",
	Usage:       "compute and sign pubkey registration hash with a bls-key",
	Description: "This command is mostly useful for testing. It only makes sense for AVS's implementing the same interface as EigenDA",
	Action:      handleBlsPubkeyRegistrationHashCmd,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "bls-keystore",
			Usage:    "path to bls keystore file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "bls-password",
			Usage:    "password for encrypted keystore file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "registry-coordinator",
			Usage:    "address of contract that conforms to eigenDA's interface",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: false,
		},
	},
}

// try to load RPC_URL from env or flags
func handleBlsPubkeyRegistrationHashCmd(ctx context.Context, cmd *cli.Command) error {

	// parse cli input
	operatorID := cmd.Int("operator-id")
	blsKeyFile := cmd.String("bls-keystore")
	registryCoordinatorAddress := common.HexToAddress(cmd.String("registry-coordinator"))
	blsKeyPassword := cmd.String("bls-password")

	// load rpc from flag or env
	rpcURL := os.Getenv("RPC_URL")
	if cmd.String("rpc-url") != "" {
		rpcURL = cmd.String("rpc-url")
	}
	if rpcURL == "" {
		return fmt.Errorf("must set env var $RPC_URL or use --rpc-url flag")
	}
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing RPC: %w", err)
	}

	// load all required addresses for this chain
	cfg, err := config.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}
	etherfiAPI := etherfi.New(cfg, rpcClient)

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(operatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	// decrypt and load bls key from keystore
	ks := keystore.NewKeystoreV3()
	blsKey, err := ks.LoadBLS(blsKeyFile, blsKeyPassword)
	if err != nil {
		return fmt.Errorf("loading bls keystore: %w", err)
	}

	// bind contract
	registryCoordinator, _ := registryCoordinator.NewContractRegistryCoordinator(registryCoordinatorAddress, rpcClient)

	// compute hash to sign with bls key
	// the hash is converted to a G1 point on the curve before it is returned
	g1Point, err := registryCoordinator.PubkeyRegistrationMessageHash(nil, operator.Address)
	if err != nil {
		return fmt.Errorf("fetching pubkeyRegistrationMessageHash: %w", err)
	}

	fmt.Printf("hash point: %v\n", g1Point)

	// map from contract type to type expected by signing algorithm
	g1MsgToSign := &bn254.G1Affine{
		X: *new(fp.Element).SetBigInt(g1Point.X),
		Y: *new(fp.Element).SetBigInt(g1Point.Y),
	}

	// sign the registration hash (G1 Point) with the bls key to prove ownership of key
	blsSigner := signer.NewBLSSigner(blsKey)
	g1Sig, err := blsSigner.Sign(g1MsgToSign)
	if err != nil {
		return fmt.Errorf("signing pubkey registration hash: %w", err)
	}
	signedParams := new(types.BLSPubkeyRegistrationParams)
	signedParams.Load(blsKey.GetPubKeyG1().G1Affine, blsKey.GetPubKeyG2().G2Affine, g1Sig)

	// sanity check on the bls signature
	isValid, err := blsSigner.Verify(g1MsgToSign, g1Sig)
	if !isValid || err != nil {
		return fmt.Errorf("failed to verify g1 signature: %w", err)
	}

	// output
	buf, err := json.MarshalIndent(signedParams, "", "    ")
	if err != nil {
		return fmt.Errorf("marshalling signature: %w", err)
	}
	fmt.Printf("\n%s\n", string(buf))

	return nil
}
