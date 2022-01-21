package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	ethrpctypes "github.com/ethereum/go-ethereum/rpc"
)

//go:generate gencodec -type Block -field-override blockMarshaling -out gen_block_json.go
type Block struct {
	BaseFeePerGas   *big.Int          `json:"baseFeePerGas"`
	Difficulty      *big.Int          `json:"difficulty"     gencodec:"required"`
	ExtraData       []byte            `json:"extraData"`
	GasLimit        uint64            `json:"gasLimit"`
	GasUsed         uint64            `json:"gasUsed"`
	Hash            common.Hash       `json:"hash"`
	LogsBloom       types.Bloom       `json:"logsBloom"`
	Miner           common.Address    `json:"miner"`
	MixHash         *common.Hash      `json:"mixHash,omitempty"` //+ *v return from geth node but not parity
	Nonce           *types.BlockNonce `json:"nonce,omitempty"`   //+ *v return from geth node but not parity
	Number          *big.Int          `json:"number"         gencodec:"required"`
	ParentHash      common.Hash       `json:"parentHash"`
	ReceiptsRoot    common.Hash       `json:"receiptsRoot"`
	Size            uint64            `json:"size"`
	StateRoot       common.Hash       `json:"stateRoot"`
	Timestamp       uint64            `json:"timestamp"`
	TotalDifficulty *big.Int          `json:"totalDifficulty"` //un-include in GetUncleByBlockHashAndIndex and GetUncleByBlockNumberAndIndex

	// Transactions type is []common.Hash when fullTx is false, otherwise []Transaction
	Transactions     TxOrHashList  `json:"transactions"`
	TransactionsRoot common.Hash   `json:"transactionsRoot"`
	Uncles           []common.Hash `json:"uncles"`
	Sha3Uncles       common.Hash   `json:"sha3Uncles"`
	// SealFields       [][]byte         `json:"sealFields"` //+ ?
}

type blockMarshaling struct {
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
	Transactions     TxOrHashList  `json:"transactions"`
	TransactionsRoot common.Hash   `json:"transactionsRoot"`
	Uncles           []common.Hash `json:"uncles"`
	Sha3Uncles       common.Hash   `json:"sha3Uncles"`
	// SealFields       []hexutil.Bytes         `json:"sealFields"` //+ ?
}

//go:generate gencodec -type Transaction -field-override transactionMarshaling -out gen_transaction_json.go
type Transaction struct {
	Accesses    types.AccessList `json:"accessList,omitempty"`
	BlockHash   *common.Hash     `json:"blockHash"`
	BlockNumber *big.Int         `json:"blockNumber"`
	ChainID     *big.Int         `json:"chainId,omitempty"`
	// Creates not guarantee to be valid, it's valid for parity node but not geth node
	Creates              *common.Address `json:"creates,omitempty"`
	From                 common.Address  `json:"from"`
	Gas                  uint64          `json:"gas"                            gencodec:"required"`
	GasPrice             *big.Int        `json:"gasPrice"`
	Hash                 common.Hash     `json:"hash"`
	Input                hexutil.Bytes   `json:"input"                          gencodec:"required"`
	MaxFeePerGas         *big.Int        `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *big.Int        `json:"maxPriorityFeePerGas,omitempty"`
	Nonce                uint64          `json:"nonce"                          gencodec:"required"`
	R                    *big.Int        `json:"r"                              gencodec:"required"`
	S                    *big.Int        `json:"s"                              gencodec:"required"`
	// Status not guarantee to be valid, it's valid for some evm compatiable chains, such as conflux chain
	Status           *uint64         `json:"status,omitempty"`
	To               *common.Address `json:"to" rlp:"nil"`
	TransactionIndex *uint64         `json:"transactionIndex"`
	Type             uint64          `json:"type"`
	V                *big.Int        `json:"v"                              gencodec:"required"`
	Value            *big.Int        `json:"value"                          gencodec:"required"`
}

type transactionMarshaling struct {
	Accesses             types.AccessList `json:"accessList,omitempty"`
	BlockHash            *common.Hash     `json:"blockHash"`
	BlockNumber          *hexutil.Big     `json:"blockNumber"`
	ChainID              *hexutil.Big     `json:"chainId,omitempty"`
	From                 common.Address   `json:"from"`
	Gas                  hexutil.Uint64   `json:"gas"`
	GasPrice             *hexutil.Big     `json:"gasPrice"`
	Hash                 common.Hash      `json:"hash"`
	Input                hexutil.Bytes    `json:"input"`
	MaxFeePerGas         *hexutil.Big     `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *hexutil.Big     `json:"maxPriorityFeePerGas,omitempty"`
	Nonce                hexutil.Uint64   `json:"nonce"`
	R                    *hexutil.Big     `json:"r"`
	S                    *hexutil.Big     `json:"s"`
	To                   *common.Address  `json:"to"`
	TransactionIndex     *hexutil.Uint64  `json:"transactionIndex"`
	Type                 hexutil.Uint64   `json:"type"`
	V                    *hexutil.Big     `json:"v"`
	Value                *hexutil.Big     `json:"value"`
}

//go:generate gencodec -type Receipt -field-override receiptMarshaling -out gen_receipt_json.go
type Receipt struct {
	BlockHash         common.Hash     `json:"blockHash"`
	BlockNumber       uint64          `json:"blockNumber"`
	ContractAddress   *common.Address `json:"contractAddress"`
	CumulativeGasUsed uint64          `json:"cumulativeGasUsed"`
	EffectiveGasPrice uint64          `json:"effectiveGasPrice"`
	From              common.Address  `json:"from"`
	GasUsed           uint64          `json:"gasUsed"`
	Logs              []*Log          `json:"logs"` //"logs"  [][]*types.Log // when receipt.Logs == nil
	LogsBloom         types.Bloom     `json:"logsBloom"`
	Root              []byte          `json:"root"`   // when len(receipt.PostState) > 0
	Status            uint            `json:"status"` // when len(receipt.PostState) = 0
	To                *common.Address `json:"to"`
	TransactionHash   common.Hash     `json:"transactionHash"`
	TransactionIndex  uint64          `json:"transactionIndex"`
	// Not guarantee to be valid, it's valid for some evm compatiable chains, such as conflux chain
	TxExecErrorMsg *string `json:"txExecErrorMsg"`
	Type           uint    `json:"type"`
}
type receiptMarshaling struct {
	BlockHash         common.Hash     `json:"blockHash"`
	BlockNumber       hexutil.Uint64  `json:"blockNumber"`
	ContractAddress   *common.Address `json:"contractAddress"`
	CumulativeGasUsed hexutil.Uint64  `json:"cumulativeGasUsed"`
	EffectiveGasPrice hexutil.Uint64  `json:"effectiveGasPrice"`
	From              common.Address  `json:"from"`
	GasUsed           hexutil.Uint64  `json:"gasUsed"`
	Logs              []*Log          `json:"logs"` //"logs"  [][]*types.Log // when receipt.Logs == nil
	LogsBloom         types.Bloom     `json:"logsBloom"`
	Root              hexutil.Bytes   `json:"root"`   // when len(receipt.PostState) > 0
	Status            hexutil.Uint    `json:"status"` // when len(receipt.PostState) = 0
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
	Gas                  *uint64         `json:"gas"`
	GasPrice             *big.Int        `json:"gasPrice"`
	MaxFeePerGas         *big.Int        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int        `json:"maxPriorityFeePerGas"`
	Value                *big.Int        `json:"value"`
	Nonce                *uint64         `json:"nonce"`

	// We accept "data" and "input" for backwards-compatibility reasons.
	// "input" is the newer name and should be preferred by clients.
	// Issue detail: https://github.com/ethereum/go-ethereum/issues/15628
	Data  []byte `json:"data"`
	Input []byte `json:"input"` //+ *v if data!=input throw, else set empty field value by filled field

	// Introduced by AccessListTxType transaction.
	AccessList *types.AccessList `json:"accessList,omitempty"`
	ChainID    *big.Int          `json:"chainId,omitempty"` //+ *v throw if chainId is consensus
}

type callRequestMarshaling struct {
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
	Input *hexutil.Bytes `json:"input,omitempty"` //+ *v if data!=input throw, else set empty field value by filled field

	// Introduced by AccessListTxType transaction.
	AccessList *types.AccessList `json:"accessList,omitempty"`
	ChainID    *hexutil.Big      `json:"chainId,omitempty"` //+ *v throw if chainId is consensus
}

//go:generate gencodec -type Log -field-override logMarshaling -out gen_log_json.go
type Log struct {
	Address             common.Address `json:"address"`
	BlockHash           common.Hash    `json:"blockHash"`
	BlockNumber         uint64         `json:"blockNumber"`
	Data                []byte         `json:"data"`
	Index               uint           `json:"logIndex"`
	LogType             *string        `json:"logType,omitempty"` //+ *v return by parity but not geth
	Removed             bool           `json:"removed"`
	Topics              []common.Hash  `json:"topics"`
	TxHash              common.Hash    `json:"transactionHash"`
	TxIndex             uint           `json:"transactionIndex"`
	TransactionLogIndex *uint          `json:"transactionLogIndex,omitempty"` //+ *v return by parity but not geth
}

type logMarshaling struct {
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
	TransactionLogIndex *hexutil.Uint  `json:"transactionLogIndex,omitempty"` //+ *v return by parity but not geth
}

// type Log = types.Log
type BlockNumber = ethrpctypes.BlockNumber
