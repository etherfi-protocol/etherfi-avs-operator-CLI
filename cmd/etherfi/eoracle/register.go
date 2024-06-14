package eoracle

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/avs/signer"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var EOracleRegisterCmd = &cli.Command{
	Name:   "register",
	Action: handleEOracleRegister,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "registration-input",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
	},
}

func handleEOracleRegister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	operatorID := cli.Int("operator-id")
	rpcURL := cli.String("rpc-url")
	inputFilepath := cli.String("registration-input")

	// connect to RPC node
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dialing rpc: %w", err)
	}
	cfg, err := bindings.AutodetectConfig(rpcClient)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	// generate and sign registration hash to be signed by admin ecdsa key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}
	sigWithSaltAndExpiry, err := signer.GenerateAndSignRegistrationDigest(operatorID, signer.EORACLE, rpcClient, privateKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}

	// read input file with required eOracle data
	var input RegistrationInput
	buf, err := os.ReadFile(inputFilepath)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}
	if err := json.Unmarshal(buf, &input); err != nil {
		return fmt.Errorf("parsing registration input: %w", err)
	}

	// TODO: quorums as parameter. Their CLI just sets it as quorum 0 for now
	quorums := []byte{0}

	return eoracleRegister(operatorID, input.BLSPubkeyRegistrationParams, sigWithSaltAndExpiry, input.AliasAddress, quorums, cfg)
}

func eoracleRegister(
	operatorID int64,
	pubkeyRegistrationParams *types.BLSPubkeyRegistrationParams,
	signature *types.SignatureWithSaltAndExpiry,
	aliasAddress common.Address,
	quorums []byte,
	cfg *bindings.Config,
) error {

	// convert to types expected by contract call
	sigParams := contracts.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature.Signature,
		Salt:      signature.Salt,
		Expiry:    signature.Expiry,
	}
	pubkeyParams := contracts.IEOBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: contracts.BN254G1Point(pubkeyRegistrationParams.Signature),
		ChainValidatorSignature:     contracts.BN254G1Point{X: big.NewInt(0), Y: big.NewInt(0)}, // not currently used by protocol
		PubkeyG1:                    contracts.BN254G1Point(pubkeyRegistrationParams.G1),
		PubkeyG2:                    contracts.BN254G2Point(pubkeyRegistrationParams.G2),
	}

	eOracleABI, err := contracts.EOracleRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	fmt.Printf("Signature: %s\n", hex.EncodeToString(signature.Signature))
	fmt.Printf("Salt: %s\n", hex.EncodeToString(signature.Salt[:]))
	fmt.Printf("Expiry: %s\n", signature.Expiry)
	fmt.Printf("\n----------------------------------------------\n")

	// pack registryCoordinator.registerOperator()
	input, err := eOracleABI.Pack("registerOperator", quorums, pubkeyParams, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}
	fmt.Printf("subcall: 0x%s\n\n", hex.EncodeToString(input))

	adminCall, err := bindings.PackForwardCallForAdmin(operatorID, input, cfg.EOracleRegistryCoordinator)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}
	fmt.Printf("admincall: 0x%s\n\n", hex.EncodeToString(adminCall))

	return nil
}

func parseEOracleG1Point(g1Str string) (*contracts.BN254G1Point, error) {
	// E([3033487302225738788775015552649894032347580654423716360411568660579702112705,2187413669180197747690573834337371262619472523581965284255666855512773801492])
	trimmed := strings.Trim(g1Str, "E([])")
	arr := strings.Split(trimmed, ",")
	if len(arr) != 2 {
		return nil, fmt.Errorf("invalid G1 point")
	}

	x, ok := big.NewInt(0).SetString(arr[0], 10)
	if !ok {
		return nil, fmt.Errorf("invalid a0 for G1 point")
	}
	y, ok := big.NewInt(0).SetString(arr[1], 10)
	if !ok {
		return nil, fmt.Errorf("invalid a1 for G1 point")
	}

	return &contracts.BN254G1Point{
		X: x,
		Y: y,
	}, nil
}

func parseEOracleG2Point(g2Str string) (*contracts.BN254G2Point, error) {
	// E([19315800099032002908983964818159634958299588208342896291709704755313327796065+1179626283257211881751865173779937098418653714448541503945687956982725670988*u,21798117894530687048377190724608441104430993460154237236704829347245529201922+4867520522353599842689725945491968901103451548555484878124286336346270391416*u])
	trimmed := removeAll(g2Str, "E([])*u")
	fmt.Println("trimmed:", trimmed)
	arr := strings.Split(trimmed, ",")
	if len(arr) != 2 {
		return nil, fmt.Errorf("invalid G2 point")
	}

	x := strings.Split(arr[0], "+")
	y := strings.Split(arr[1], "+")

	xA0, ok := big.NewInt(0).SetString(x[0], 10)
	if !ok {
		return nil, fmt.Errorf("invalid xa0 for G2 point")
	}
	xA1, ok := big.NewInt(0).SetString(x[1], 10)
	if !ok {
		fmt.Println("Xa1", x[1])
		return nil, fmt.Errorf("invalid xa1 for G2 point")
	}
	yA0, ok := big.NewInt(0).SetString(y[0], 10)
	if !ok {
		return nil, fmt.Errorf("invalid ya0 for G2 point")
	}
	yA1, ok := big.NewInt(0).SetString(y[1], 10)
	if !ok {
		return nil, fmt.Errorf("invalid ya1 for G2 point")
	}

	// TODO: Do we need to flip the a0 and a1 points
	return &contracts.BN254G2Point{
		X: [2]*big.Int{xA0, xA1},
		Y: [2]*big.Int{yA0, yA1},
	}, nil
}

func removeAll(s, cutset string) string {
	var b strings.Builder
	for _, ch := range s {
		if !strings.ContainsRune(cutset, ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
