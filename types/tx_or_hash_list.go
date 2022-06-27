package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/go-sdk-common/rpctest"
)

type txListType string

const (
	TXLIST_TRANSACTION txListType = "transaction"
	TXLIST_HASH        txListType = "hash"
)

type TxOrHashList struct {
	vtype        txListType
	transactions []TransactionDetail
	hashes       []common.Hash
}

func NewTxOrHashList(isFull bool) *TxOrHashList {
	l := TxOrHashList{}
	l.vtype = TxListType(isFull)
	return &l
}

func NewTxOrHashListByTxs(txs []TransactionDetail) *TxOrHashList {
	return &TxOrHashList{
		vtype:        TXLIST_TRANSACTION,
		transactions: txs,
	}
}

func NewTxOrHashListByHashes(hashes []common.Hash) *TxOrHashList {
	return &TxOrHashList{
		vtype:  TXLIST_HASH,
		hashes: hashes,
	}
}

func TxListType(isFull bool) txListType {
	if isFull {
		return TXLIST_TRANSACTION
	}
	return TXLIST_HASH
}

func (l *TxOrHashList) Transactions() []TransactionDetail {
	return l.transactions
}

func (l *TxOrHashList) Hashes() []common.Hash {
	return l.hashes
}

func (l *TxOrHashList) UnmarshalJSON(data []byte) error {

	if l.vtype == TXLIST_TRANSACTION {
		var txs []TransactionDetail
		e := json.Unmarshal(data, &txs)
		l.transactions = txs
		return e
	}

	if l.vtype == TXLIST_HASH {
		var hashes []common.Hash
		e := json.Unmarshal(data, &hashes)
		l.hashes = hashes
		return e
	}

	var txs []TransactionDetail
	var e error
	if e = json.Unmarshal(data, &txs); e == nil {
		l.vtype = TXLIST_TRANSACTION
		l.transactions = txs
		return nil
	}

	var hashes []common.Hash
	if e = json.Unmarshal(data, &hashes); e == nil {
		l.vtype = TXLIST_HASH
		l.hashes = hashes
		return nil
	}

	return e
}

func (l TxOrHashList) MarshalJSON() ([]byte, error) {
	switch l.Type() {
	case TXLIST_TRANSACTION:
		return json.Marshal(l.transactions)
	case TXLIST_HASH:
		return json.Marshal(l.hashes)
	}
	return json.Marshal(nil)
}

func (l TxOrHashList) MarshalJSONForRPCTest(indent ...bool) ([]byte, error) {
	switch l.Type() {
	case TXLIST_TRANSACTION:
		return rpctest.JsonMarshalForRpcTest(l.transactions, indent...)
	case TXLIST_HASH:
		return rpctest.JsonMarshalForRpcTest(l.hashes, indent...)
	}
	return json.Marshal(nil)
}

func (l *TxOrHashList) Type() txListType {
	return l.vtype
}
