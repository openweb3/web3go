package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	ethrpctypes "github.com/ethereum/go-ethereum/rpc"
)

type Block struct {
	BaseFeePerGas   *hexutil.Big      `json:"baseFeePerGas"`
	Difficulty      *hexutil.Big      `json:"difficulty"`
	ExtraData       hexutil.Bytes     `json:"extraData"`
	GasLimit        hexutil.Uint64    `json:"gasLimit"`
	GasUsed         hexutil.Uint64    `json:"gasUsed"`
	Hash            common.Hash       `json:"hash"`
	LogsBloom       types.Bloom       `json:"logsBloom"`
	Miner           common.Address    `json:"miner"`
	MixHash         *common.Hash      `json:"mixHash,omitempty"` //+ *v return from geth node but not parity
	Nonce           *types.BlockNonce `json:"nonce,omitempty"`   //+ *v return from geth node but not parity
	Number          *hexutil.Big      `json:"number"`
	ParentHash      common.Hash       `json:"parentHash"`
	ReceiptsRoot    common.Hash       `json:"receiptsRoot"`
	Size            hexutil.Uint64    `json:"size"`
	StateRoot       common.Hash       `json:"stateRoot"`
	Timestamp       hexutil.Uint64    `json:"timestamp"`
	TotalDifficulty *hexutil.Big      `json:"totalDifficulty"` //un-include in GetUncleByBlockHashAndIndex and GetUncleByBlockNumberAndIndex

	// Transactions type is []common.Hash when fullTx is false, otherwise []Transaction
	Transactions     TxOrHashList `json:"transactions"`
	TransactionsRoot common.Hash     `json:"transactionsRoot"`
	Uncles           []common.Hash   `json:"uncles"`
	Sha3Uncles       common.Hash     `json:"sha3Uncles"`
	// SealFields       []hexutil.Bytes         `json:"sealFields"` //+ ?
}

type Transaction struct {
	Accesses             *types.AccessList `json:"accessList,omitempty"`
	BlockHash            *common.Hash      `json:"blockHash"`
	BlockNumber          *hexutil.Big      `json:"blockNumber"`
	ChainID              *hexutil.Big      `json:"chainId,omitempty"`
	From                 common.Address    `json:"from"`
	Gas                  hexutil.Uint64    `json:"gas"`
	GasPrice             *hexutil.Big      `json:"gasPrice"`
	Hash                 common.Hash       `json:"hash"`
	Input                hexutil.Bytes     `json:"input"`
	MaxFeePerGas         *hexutil.Big      `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
	Nonce                hexutil.Uint64    `json:"nonce"`
	R                    *hexutil.Big      `json:"r"`
	S                    *hexutil.Big      `json:"s"`
	To                   *common.Address   `json:"to"`
	TransactionIndex     *hexutil.Uint64   `json:"transactionIndex"`
	Type                 hexutil.Uint64    `json:"type"`
	V                    *hexutil.Big      `json:"v"`
	Value                *hexutil.Big      `json:"value"`
}

type Receipt struct {
	BlockHash         common.Hash     `json:"blockHash"`
	BlockNumber       hexutil.Uint64  `json:"blockNumber"`
	ContractAddress   *common.Address `json:"contractAddress"`
	CumulativeGasUsed hexutil.Uint64  `json:"cumulativeGasUsed"`
	EffectiveGasPrice hexutil.Uint64  `json:"effectiveGasPrice"`
	From              common.Address  `json:"from"`
	GasUsed           hexutil.Uint64  `json:"gasUsed"`
	Logs              []*Log          `json:"logs"` //"logs"  [][]*types.Log // when receipt.Logs == nil
	LogsBloom         types.Bloom     `json:"logsBloom"`
	Root              *hexutil.Bytes  `json:"root"`   // when len(receipt.PostState) > 0
	Status            *hexutil.Uint   `json:"status"` // when len(receipt.PostState) = 0
	To                *common.Address `json:"to"`
	TransactionHash   common.Hash     `json:"transactionHash"`
	TransactionIndex  hexutil.Uint64  `json:"transactionIndex"`
	Type              hexutil.Uint    `json:"type"`
}

// CallRequest represents the arguments to construct a new transaction
// or a message call.
type CallRequest struct {
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
	Input *hexutil.Bytes `json:"input"` //+ *v if data!=input throw, else set empty field value by filled field

	// Introduced by AccessListTxType transaction.
	AccessList *types.AccessList `json:"accessList,omitempty"`
	ChainID    *hexutil.Big      `json:"chainId,omitempty"` //+ *v throw if chainId is consensus
}

type Log struct {
	Address             common.Address `json:"address"`
	BlockHash           common.Hash    `json:"blockHash"`
	BlockNumber         hexutil.Uint64 `json:"blockNumber"`
	Data                hexutil.Bytes  `json:"data"`
	Index               hexutil.Uint   `json:"logIndex"`
	LogType             *string        `json:"logType,omitempty"` //+ *v return by parity but not geth
	Removed             bool           `json:"removed"`
	Topics              []common.Hash  `json:"topics"`
	TxHash              common.Hash    `json:"transactionHash"`
	TxIndex             hexutil.Uint   `json:"transactionIndex"`
	TransactionLogIndex *hexutil.Big   `json:"transactionLogIndex,omitempty"` //+ *v return by parity but not geth
}

// type Log = types.Log
type BlockNumber = ethrpctypes.BlockNumber
