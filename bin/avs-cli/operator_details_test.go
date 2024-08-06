package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/dsrvlabs/etherfi-avs-operator-tool/bindings/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestRegistrationEvents(t *testing.T) {

	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		t.Fatal(err)
	}

	avsDirectory, err := contracts.NewAVSDirectory(common.HexToAddress("0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF"), rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	operatorFilter := []common.Address{common.HexToAddress("0x67943aE8e07bFC9f5C9A90d608F7923D9C21e051")}
	avsFilter := []common.Address{}
	events, err := avsDirectory.FilterOperatorAVSRegistrationStatusUpdated(nil, operatorFilter, avsFilter)
	if err != nil {
		t.Fatal(err)
	}
	for events.Next() {
		event := events.Event
		fmt.Printf("Operator: %s, Avs: %s, Status: %d\n", event.Operator, event.Avs, event.Status)
		fmt.Printf("%s\n", event.Raw.TxHash)
	}

}

func TestOperatorDetails(t *testing.T) {

	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))

	if err != nil {
		t.Fatal(err)
	}

	var testcases = []struct {
		operatorID int64
	}{
		{operatorID: 1},
		{operatorID: 2},
		{operatorID: 3},
		{operatorID: 4},
		{operatorID: 5},
		{operatorID: 6},
		{operatorID: 7},
		{operatorID: 8},
		{operatorID: 9},
		{operatorID: 10},
		{operatorID: 11},
		{operatorID: 12},
		{operatorID: 13},
		{operatorID: 14},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("operator-%d", tc.operatorID), func(t *testing.T) {
			if err := operatorDetails(tc.operatorID, rpcClient); err != nil {
				t.Fatal(err)
			}
		})
	}

}

/*
func TestOperatorDetails(t *testing.T) {

	// connect to RPC node
	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		t.Fatal(err)
	}

	// bind contract
	operatorManagerContract, err := contracts.NewEtherfiAVSOperatorsManager(bindings.Holesky.OperatorManagerAddress, rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	// load signing key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		t.Fatal(err)
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(17000))
	if err != nil {
		t.Fatal(err)
	}

		_, err = operatorManagerContract.InstantiateEtherFiAvsOperator(transactor, big.NewInt(5))
		if err != nil {
			t.Fatal(err)
		}

	operatorID := big.NewInt(1)
	info, err := operatorManagerContract.GetAvsInfo(nil, operatorID, bindings.Holesky.BrevisRegistryCoordinator)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", info)
	return

	deets := contracts.IDelegationManagerOperatorDetails{
		EarningsReceiver: common.HexToAddress("0xD0d7F8a5a86d8271ff87ff24145Cf40CEa9F7A39"),
	}

	operator, err := operatorManagerContract.AvsOperators(nil, operatorID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("operator: %s\n", operator)
	return

	// Manually pack input to support gnosis signing
	managerABI, err := contracts.EtherfiAVSOperatorsManagerMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	input, err := managerABI.Pack("registerAsOperator", operatorID, deets, "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Target: %s\n", bindings.Holesky.OperatorManagerAddress)
	fmt.Printf("Raw Input: 0x%s\n", hex.EncodeToString(input))

	//details := bindings.

	/*
		EarningsReceiver         common.Address
		DelegationApprover       common.Address
		StakerOptOutWindowBlocks uint32

	if _, err := operatorManagerContract.RegisterAsOperator(transactor, operatorID, deets, ""); err != nil {
		t.Fatal(err)
	}

}
*/
