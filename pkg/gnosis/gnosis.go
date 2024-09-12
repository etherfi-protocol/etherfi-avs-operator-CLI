package gnosis

import (
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type SubTransaction struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

type GnosisMetadata struct {
	Name string `json:"name"`
}

type GnosisBatch struct {
	Version      string           `json:"version"`
	ChainId      string           `json:"chainId"`
	Meta         GnosisMetadata   `json:"meta"`
	Transactions []SubTransaction `json:"transactions"`
}

func (s *SubTransaction) MarshalJSON() ([]byte, error) {
	type out struct {
		To    string `json:"to"`
		Value string `json:"value"`
		Data  string `json:"data"`
	}

	return json.Marshal(&out{
		To:    s.Target.Hex(),
		Value: s.Value.String(),
		Data:  "0x" + hex.EncodeToString(s.Data),
	})
}

func (b *GnosisBatch) AddTransaction(tx SubTransaction) {
	b.Transactions = append(b.Transactions, tx)
}

func (b *GnosisBatch) AddTransactions(txs []SubTransaction) {
	b.Transactions = append(b.Transactions, txs...)
}

func (b *GnosisBatch) PrettyPrint() string {
	buf, _ := json.MarshalIndent(b, "", "    ")
	return string(buf)
}

func NewSingleTxBatch(data []byte, chainID *big.Int, name string) *GnosisBatch {

	batch := GnosisBatch{
		Version: "1.0",
		ChainId: chainID.String(),
		Meta:    GnosisMetadata{Name: name},
	}

	return &batch
}
