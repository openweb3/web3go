package types

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	ethrpctypes "github.com/ethereum/go-ethereum/rpc"
)

type Transaction struct {
	BlockHash        *common.Hash      `json:"blockHash"`
	BlockNumber      *hexutil.Big      `json:"blockNumber"`
	From             common.Address    `json:"from"`
	Gas              hexutil.Uint64    `json:"gas"`
	GasPrice         *hexutil.Big      `json:"gasPrice"`
	GasFeeCap        *hexutil.Big      `json:"maxFeePerGas,omitempty"`
	GasTipCap        *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
	Hash             common.Hash       `json:"hash"`
	Input            hexutil.Bytes     `json:"input"`
	Nonce            hexutil.Uint64    `json:"nonce"`
	To               *common.Address   `json:"to"`
	TransactionIndex *hexutil.Uint64   `json:"transactionIndex"`
	Value            *hexutil.Big      `json:"value"`
	Type             hexutil.Uint64    `json:"type"`
	Accesses         *types.AccessList `json:"accessList,omitempty"`
	ChainID          *hexutil.Big      `json:"chainId,omitempty"`
	V                *hexutil.Big      `json:"v"`
	R                *hexutil.Big      `json:"r"`
	S                *hexutil.Big      `json:"s"`
}

type Receipt struct {
	TransactionType   *hexutil.Uint64 `json:"transactionType"`
	TransactionHash   *common.Hash    `json:"transactionHash"`
	TransactionIndex  *hexutil.Big    `json:"transactionIndex"`
	BlockHash         *common.Hash    `json:"blockHash"`
	From              *common.Address `json:"from"`
	To                *common.Address `json:"to"`
	BlockNumber       *hexutil.Big    `json:"blockNumber"`
	CumulativeGasUsed *hexutil.Big    `json:"cumulativeGasUsed"`
	GasUsed           *hexutil.Big    `json:"gasUsed"`
	ContractAddress   *common.Address `json:"contractAddress"`
	Logs              []Log           `json:"logs"`
	StateRoot         *common.Hash    `json:"stateRoot"`
	LogsBloom         types.Bloom     `json:"logsBloom"`
	StatusCode        *hexutil.Uint64 `json:"statusCode"`
	EffectiveGasPrice *hexutil.Big    `json:"effectiveGasPrice"`
}

type Block struct {
	Hash             *common.Hash    `json:"hash"`
	ParentHash       common.Hash     `json:"parentHash"`
	UnclesHash       common.Hash     `json:"unclesHash"`
	Author           common.Address  `json:"author"`
	Miner            common.Address  `json:"miner"`
	StateRoot        common.Hash     `json:"stateRoot"`
	TransactionsRoot common.Hash     `json:"transactionsRoot"`
	ReceiptsRoot     common.Hash     `json:"receiptsRoot"`
	Number           *hexutil.Big    `json:"number"`
	GasUsed          *hexutil.Big    `json:"gasUsed"`
	GasLimit         *hexutil.Big    `json:"gasLimit"`
	ExtraData        hexutil.Bytes   `json:"extraData"`
	LogsBloom        *types.Bloom    `json:"logsBloom"`
	Timestamp        *hexutil.Big    `json:"timestamp"`
	Difficulty       *hexutil.Big    `json:"difficulty"`
	TotalDifficulty  *hexutil.Big    `json:"totalDifficulty"`
	SealFields       []hexutil.Bytes `json:"sealFields"`
	BaseFeePerGas    *hexutil.Big    `json:"baseFeePerGas"`
	Uncles           []common.Hash   `json:"uncles"`
	Transactions     []Transaction   `json:"transactions"`
	Size             *hexutil.Big    `json:"size"`
}

// TransactionArgs represents the arguments to construct a new transaction
// or a message call.
type TransactionArgs struct {
	From                 *common.Address `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	Value                *hexutil.Big    `json:"value"`
	Nonce                *hexutil.Uint64 `json:"nonce"`

	// We accept "data" and "input" for backwards-compatibility reasons.
	// "input" is the newer name and should be preferred by clients.
	// Issue detail: https://github.com/ethereum/go-ethereum/issues/15628
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`

	// Introduced by AccessListTxType transaction.
	AccessList *types.AccessList `json:"accessList,omitempty"`
	ChainID    *hexutil.Big      `json:"chainId,omitempty"`
}

type Log = types.Log
type EthRpcLogFilter = ethereum.FilterQuery
type BlockNumber = ethrpctypes.BlockNumber
type SyncProgress = ethereum.SyncProgress
