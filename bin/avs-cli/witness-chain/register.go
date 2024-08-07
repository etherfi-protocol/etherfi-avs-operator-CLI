package witness

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/gnosis"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/src/witnesschain"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var WitnessRegisterToAvsCmd = &cli.Command{
	Name:   "register",
	Usage:  "register target operator with avs",
	Action: handleWitnessRegister,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "registration-input",
			Usage:    "path to registration file created by prepare-registration command",
			Required: true,
		},
	},
}

func handleWitnessRegister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
	inputFilepath := cli.String("registration-input")

	// read input file with required witnesschain data
	var input RegistrationInput
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

	// generate and sign registration hash with admin ecdsa key
	signingKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}
	return witnessAPI.RegisterOperator(operator, signingKey)

	/*
		sigWithSaltAndExpiry, err := signer.GenerateAndSignRegistrationDigest(input.OperatorID, signer.WITNESS_CHAIN, rpcClient, privateKey)
		if err != nil {
			return fmt.Errorf("signing registration digest: %w", err)
		}

		return witnessRegister(input.OperatorID, sigWithSaltAndExpiry, cfg, rpcClient)
	*/
}

func witnessRegister(
	operatorID int64,
	signature *types.SignatureWithSaltAndExpiry,
	cfg *bindings.Config,
	rpcClient *ethclient.Client,
) error {

	// TODO(Dave): since I am using raw pack, see if I even need to do this conversion
	//             because underlying types have the same structure
	// convert to types expected by contract call
	sigParams := witnesschain.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature.Signature,
		Salt:      signature.Salt,
		Expiry:    signature.Expiry,
	}

	// look up operator contract associated with this id and configured ecdsaSigner
	operatorManagerContract, err := contracts.NewAvsOperatorManager(cfg.AvsOperatorManagerAddress, rpcClient)
	if err != nil {
		return fmt.Errorf("binding operatorManager: %w", err)
	}
	operatorAddr, err := operatorManagerContract.AvsOperators(nil, big.NewInt(operatorID))
	if err != nil {
		return fmt.Errorf("looking up operator address: %w", err)
	}

	witnessABI, err := witnesschain.WitnessChainWitnessHubMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}
	fmt.Printf("Signature: %s\n", hex.EncodeToString(signature.Signature))
	fmt.Printf("Salt: %s\n", hex.EncodeToString(signature.Salt[:]))
	fmt.Printf("Expiry: %s\n", signature.Expiry)
	fmt.Printf("\n----------------------------------------------\n")

	// pack witnessHub.registerOperatorToAVS()
	input, err := witnessABI.Pack("registerOperatorToAVS", operatorAddr, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}
	fmt.Printf("subcall: 0x%s\n\n", hex.EncodeToString(input))

	adminCall, err := bindings.PackForwardCallForAdmin(operatorID, input, cfg.WitnessChainWitnessHubAddress)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}
	fmt.Printf("admincall: 0x%s\n\n", hex.EncodeToString(adminCall))

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, cfg.AvsOperatorManagerAddress, fmt.Sprintf("witness-chain-register-watchtower-%d", operatorID))
	fmt.Printf("gnosis:\n%s\n", batch.PrettyPrint())

	return nil
}
