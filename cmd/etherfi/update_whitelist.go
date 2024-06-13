package main

/*
var updateWhitelistCmd = &cli.Command{
	Name:   "update-whitelist",
	Usage:  "(Ether.fi Admin) Update which AVS's, an operator is approved for",
	Action: updateWhitelist,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "operator-id",
			Usage:    "Operator ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "registry-coordinator",
			Usage:    "address of the registry-coordinator associated with the avs your are registering to",
			Required: true,
		},
		&cli.IntFlag{
			Name:  "chain-id",
			Usage: "Chain ID",
			Value: 1, // default to mainnet
		},
		&cli.BoolFlag{
			Name:  "is-whitelisted",
			Usage: "true / false",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "broadcast",
			Usage: "broadcast signed tx to network",
		},
	},
}

func updateWhitelist(ctx context.Context, cli *cli.Command) error {

	chainID := cli.Int("chain-id")
	cfg, err := configForChain(chainID)
	if err != nil {
		return err
	}

	// inputs
	operatorID := big.NewInt(cli.Int("operator-id"))
	registryCoordinator := common.HexToAddress(cli.String("registry-coordinator"))
	isWhitelisted := cli.Bool("is-whitelisted")

	// Manually pack input to support gnosis signing
	managerABI, err := contracts.AvsOperatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("retrieving abi: %w", err)
	}
	input, err := managerABI.Pack("updateAvsWhitelist", operatorID, registryCoordinator, isWhitelisted)
	if err != nil {
		return fmt.Errorf("packing input data from abi: %w", err)
	}
	fmt.Printf("Target: %s\n", cfg.OperatorManagerAddress)
	fmt.Printf("Raw Input: 0x%s\n", hex.EncodeToString(input))

	if cli.Bool("broadcast") {
		fmt.Printf("broadcasting tx...\n")

		// load signing key
		privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
		if err != nil {
			return fmt.Errorf("loading signing key: %w", err)
		}
		transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
		if err != nil {
			return fmt.Errorf("creating signer from key: %w", err)
		}

		// connect to RPC node
		rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
		if err != nil {
			return fmt.Errorf("dialing rpc: %w", err)
		}

		// bind contract
		operatorManagerContract, err := contracts.NewAvsOperatorManager(cfg.OperatorManagerAddress, rpcClient)
		if err != nil {
			return fmt.Errorf("binding contract: %w", err)
		}

		// execute transaction
		_, err = operatorManagerContract.UpdateAvsWhitelist(transactor, operatorID, registryCoordinator, isWhitelisted)
		if err != nil {
			return fmt.Errorf("updating whitelist tx: %w", err)
		}

	}

	return nil
}
*/

// TODO: move somewhere else
