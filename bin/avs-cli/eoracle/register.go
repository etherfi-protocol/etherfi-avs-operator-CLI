package eoracle

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/eoracle"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v3"
)

var EOracleRegisterCmd = &cli.Command{
	Name:   "register",
	Action: handleEOracleRegister,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "registration-input",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
	},
}

func handleEOracleRegister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	inputFilepath := cli.String("registration-input")

	// read input file with required eOracle data
	var input eoracle.RegistrationInfo
	buf, err := os.ReadFile(inputFilepath)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}
	if err := json.Unmarshal(buf, &input); err != nil {
		return fmt.Errorf("parsing registration input: %w", err)
	}
	if input.OperatorID == 0 {
		return fmt.Errorf("invalid registration input")
	}

	// look up operator contract associated with this id
	operator, err := etherfiAPI.LookupOperatorByID(input.OperatorID)
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	// TODO: quorums as parameter. Their CLI just hardcodes it as quorum 0 for now
	quorums := []byte{0}

	// generate and sign registration hash to be signed by admin ecdsa key
	signerKey, err := crypto.HexToECDSA(os.Getenv("ADMIN_1271_SIGNING_KEY"))
	if err != nil {
		return fmt.Errorf("---invalid private key: %w", err)
	}

	return eoracleAPI.RegisterOperator(operator, input, signerKey, quorums)
}

// helper to parse the string format output of the eOracle cli
func parseEOracleG1Point(g1Str string) (*eoracle.BN254G1Point, error) {
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

	return &eoracle.BN254G1Point{
		X: x,
		Y: y,
	}, nil
}

// helper to parse the string format output of the eOracle cli
func parseEOracleG2Point(g2Str string) (*eoracle.BN254G2Point, error) {
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
	return &eoracle.BN254G2Point{
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
