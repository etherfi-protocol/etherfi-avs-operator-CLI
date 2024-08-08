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
