package eigenda

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	// need to alias because eigenlayer has a package name that doesn't match the filepath
	registryCoordinator "github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/avs/signer"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/gnosis"
	"github.com/dsrvlabs/etherfi-avs-operator-tool/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v3"
)

var EigenDARegisterCmd = &cli.Command{
	Name:   "register",
	Action: handleEigenDARegister,
	Flags: []cli.Flag{
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

func handleEigenDARegister(ctx context.Context, cli *cli.Command) error {

	// parse cli params
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

	// read input file with required eigenDA data
	var input RegistrationInput
	buf, err := os.ReadFile(inputFilepath)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}
	if err := json.Unmarshal(buf, &input); err != nil {
		return fmt.Errorf("parsing registration input: %w", err)
	}

	if input.OperatorID == 0 {
		return fmt.Errorf("invalid registration input, missing operatorID")
	}
	if input.Socket == "" {
		return fmt.Errorf("invalid registration input, missing socket")
	}

	// generate and sign registration hash to be signed by admin ecdsa key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}
	sigWithSaltAndExpiry, err := signer.GenerateAndSignRegistrationDigest(input.OperatorID, signer.EIGEN_DA, rpcClient, privateKey)
	if err != nil {
		return fmt.Errorf("signing registration digest: %w", err)
	}

	quorums := make([]uint8, len(input.Quorums))
	for idx, q := range input.Quorums {
		quorums[idx] = uint8(q)
	}

	return eigenDARegister(input.OperatorID, input.BLSPubkeyRegistrationParams, sigWithSaltAndExpiry, quorums, input.Socket, cfg)
}

func eigenDARegister(
	operatorID int64,
	pubkeyRegistrationParams *types.BLSPubkeyRegistrationParams,
	signature *types.SignatureWithSaltAndExpiry,
	quorums []byte,
	socket string,
	cfg *bindings.Config,
) error {

	// convert to types expected by contract call
	sigParams := contracts.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature.Signature,
		Salt:      signature.Salt,
		Expiry:    signature.Expiry,
	}
	pubkeyParams := contracts.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: contracts.BN254G1Point(pubkeyRegistrationParams.Signature),
		PubkeyG1:                    contracts.BN254G1Point(pubkeyRegistrationParams.G1),
		PubkeyG2:                    contracts.BN254G2Point(pubkeyRegistrationParams.G2),
	}

	// (Dave): eigenlayer has a vendored version of their contract bindings. We should do this too...
	eigenABI, err := registryCoordinator.ContractRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("fetching abi: %w", err)
	}

	fmt.Printf("Signature: %s\n", hex.EncodeToString(signature.Signature))
	fmt.Printf("Salt: %s\n", hex.EncodeToString(signature.Salt[:]))
	fmt.Printf("Expiry: %s\n", signature.Expiry)
	fmt.Printf("\n----------------------------------------------\n")

	// pack registryCoordinator.registerOperator()
	input, err := eigenABI.Pack("registerOperator", quorums, socket, pubkeyParams, sigParams)
	if err != nil {
		return fmt.Errorf("packing input: %w", err)
	}
	fmt.Printf("subcall: 0x%s\n\n", hex.EncodeToString(input))

	adminCall, err := bindings.PackForwardCallForAdmin(operatorID, input, cfg.EigenDARegistryCoordinator)
	if err != nil {
		return fmt.Errorf("wrapping call for admin: %w", err)
	}
	fmt.Printf("admincall: 0x%s\n\n", hex.EncodeToString(adminCall))

	// output in gnosis compatible format
	batch := gnosis.NewSingleTxBatch(adminCall, cfg.OperatorManagerAddress, fmt.Sprintf("eigenda-register-%d", operatorID))
	fmt.Printf("gnosis:\n%s\n", batch.PrettyPrint())

	return nil
}
