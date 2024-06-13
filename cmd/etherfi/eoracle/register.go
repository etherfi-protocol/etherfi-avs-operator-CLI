package eoracle

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
)

/*
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
			Name:     "bls-pubkey-g1",
			Usage:    "e.g. E([303...,218...])",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "bls-pubkey-g2",
			Usage:    "e.g. E([303...+123...*u,218...+987...*u])",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "rpc-url",
			Usage:    "rpc url",
			Required: true,
		},
	},
}
*/

/*
func handleEOracleRegister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	operatorID := cli.Int("operator-id")
	rpcURL := cli.String("rpc-url")
	signAddress := common.HexToAddress(cli.String("sign-address"))
	g1Str := cli.String("bls-pubkey-g1")
	g2Str := cli.String("bls-pubkey-g2")

	g1, err := parseEOracleG1Point(g1Str)
	if err != nil {
		return fmt.Errorf("parsing g1 point: %w", err)
	}
	g2, err := parseEOracleG2Point(g2Str)
	if err != nil {
		return fmt.Errorf("parsing g2 point: %w", err)
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

	return nil
	// return eoracleRegister(operatorID, signAddress, pubkeyBytes, signature, salt, expiry, rpcClient, outputGnosis)
}
*/

/*
func eoracleRegister(
	operatorID int64,
	g1 *contracts.BN254G1Point,
	g2 *contracts.BN254G2Point,
*/

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
