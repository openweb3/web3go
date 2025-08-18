package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// TxpoolStatus represents the status of the transaction pool
//
//go:generate gencodec -type TxpoolStatus -field-override txpoolStatusMarshaling -out gen_txpool_status_json.go
type TxpoolStatus struct {
	Pending uint64 `json:"pending"`
	Queued  uint64 `json:"queued"`
}

type txpoolStatusMarshaling struct {
	Pending hexutil.Uint64 `json:"pending"`
	Queued  hexutil.Uint64 `json:"queued"`
}

// TxpoolInspectSummary represents a summary of a transaction in the pool
// This is a string type that matches the geth format:
// "0x1234...5678: 1000000000000000000 wei + 21000 gas × 20000000000 wei"
// or "contract creation: 0 wei + 500000 gas × 20000000000 wei"
type TxpoolInspectSummary string

// TxpoolInspect represents the inspection result of the transaction pool
type TxpoolInspect struct {
	Pending map[common.Address]map[string]TxpoolInspectSummary `json:"pending"`
	Queued  map[common.Address]map[string]TxpoolInspectSummary `json:"queued"`
}

// TxpoolContentFrom represents the content of transactions from a specific address
type TxpoolContentFrom struct {
	Pending map[string]types.Transaction `json:"pending"`
	Queued  map[string]types.Transaction `json:"queued"`
}

// TxpoolContent represents the content of all transactions in the pool
type TxpoolContent struct {
	Pending map[common.Address]map[string]types.Transaction `json:"pending"`
	Queued  map[common.Address]map[string]types.Transaction `json:"queued"`
}
