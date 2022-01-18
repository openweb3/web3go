package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

type TransactionOrHashList struct {
	vtype        string
	Transactions []Transaction
	Hashes       []common.Hash
}

func (l *TransactionOrHashList) UnmarshalJSON(data []byte) error {
	var txs []Transaction
	if e := json.Unmarshal(data, &txs); e == nil {
		l.Transactions = txs
		l.vtype = "Transactions"
		return nil
	}
	var hashes []common.Hash
	if e := json.Unmarshal(data, &hashes); e == nil {
		l.Hashes = hashes
		l.vtype = "Hashes"
		return nil
	}
	l.vtype = "None"
	return nil
}

func (l TransactionOrHashList) MarshalJSON() ([]byte, error) {
	switch l.Type() {
	case "Transactions":
		return json.Marshal(l.Transactions)
	case "Hashes":
		return json.Marshal(l.Hashes)
	}
	return json.Marshal(nil)
}

func (l *TransactionOrHashList) Type() string {
	return l.vtype
}
