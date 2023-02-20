package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	rpc "github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
)

//go:generate gencodec -type Block -field-override blockMarshaling -out gen_block_json.go
type Block struct {
	Author          *common.Address   `json:"author,omitempty"`
	BaseFeePerGas   *big.Int          `json:"baseFeePerGas,omitempty"`
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
	TotalDifficulty *big.Int          `json:"totalDifficulty,omitempty"` //un-include in GetUncleByBlockHashAndIndex and GetUncleByBlockNumberAndIndex

	// Transactions type is []common.Hash when fullTx is false, otherwise []Transaction
	Transactions     TxOrHashList  `json:"transactions"`
	TransactionsRoot common.Hash   `json:"transactionsRoot"`
	Uncles           []common.Hash `json:"uncles"`
	Sha3Uncles       common.Hash   `json:"sha3Uncles"`
	// SealFields       [][]byte         `json:"sealFields"` //+ ?
}

func (b *Block) Header() (*Header, error) {

	if b.MixHash == nil {
		return nil, errors.New("MixHash is nil")
	}

	if b.Nonce == nil {
		return nil, errors.New("Nonce is nil")
	}

	h := Header{
		HeaderExtra: HeaderExtra{
			Author: b.Author,
			Hash:   &b.Hash,
		},

		Header: ethtypes.Header{
			ParentHash:  b.ParentHash,
			UncleHash:   b.Sha3Uncles,
			Coinbase:    b.Miner,
			Root:        b.StateRoot,
			TxHash:      b.TransactionsRoot,
			ReceiptHash: b.ReceiptsRoot,
			Bloom:       b.LogsBloom,
			Difficulty:  b.Difficulty,
			Number:      b.Number,
			GasLimit:    b.GasLimit,
			GasUsed:     b.GasUsed,
			Time:        b.Timestamp,
			Extra:       b.ExtraData,
			MixDigest:   *b.MixHash,
			Nonce:       *b.Nonce,
			BaseFee:     b.BaseFeePerGas,
		},
	}
	return &h, nil
}

type blockMarshaling struct {
	Author          *common.Address   `json:"author,omitempty"`
	BaseFeePerGas   *hexutil.Big      `json:"baseFeePerGas,omitempty"`
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
	TotalDifficulty *hexutil.Big      `json:"totalDifficulty,omitempty"` //un-include in GetUncleByBlockHashAndIndex and GetUncleByBlockNumberAndIndex

	// Transactions type is []common.Hash when fullTx is false, otherwise []Transaction
	Transactions     TxOrHashList  `json:"transactions"`
	TransactionsRoot common.Hash   `json:"transactionsRoot"`
	Uncles           []common.Hash `json:"uncles"`
	Sha3Uncles       common.Hash   `json:"sha3Uncles"`
	// SealFields       []hexutil.Bytes         `json:"sealFields"` //+ ?
}

//go:generate gencodec -type PlainTransaction -field-override plainTransactionMarshaling -out gen_plain_transaction_json.go
// "testomit" tag is used to omit the field in rpc test, omit when testomit is true and un-omit when testomit is false.
type TransactionDetail struct {
	Accesses    types.AccessList `json:"accessList,omitempty"`
	BlockHash   *common.Hash     `json:"blockHash"`
	BlockNumber *big.Int         `json:"blockNumber"`
	ChainID     *big.Int         `json:"chainId,omitempty"`
	// Creates not guarantee to be valid, it's valid for parity node but not geth node
	Creates              *common.Address `json:"creates,omitempty"                                   testomit:"false"`
	From                 common.Address  `json:"from"`
	Gas                  uint64          `json:"gas"                            gencodec:"required"`
	GasPrice             *big.Int        `json:"gasPrice"`
	Hash                 common.Hash     `json:"hash"`
	Input                hexutil.Bytes   `json:"input"                          gencodec:"required"`
	MaxFeePerGas         *big.Int        `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *big.Int        `json:"maxPriorityFeePerGas,omitempty"`
	Nonce                uint64          `json:"nonce"                          gencodec:"required"`
	// Creates not guarantee to be valid, it's valid for parity node but not geth node
	PublicKey *hexutil.Bytes `json:"publicKey,omitempty"                                 testomit:"false"` //+ x
	R         *big.Int       `json:"r"                              gencodec:"required"`
	// Creates not guarantee to be valid, it's valid for parity node but not geth node
	Raw       *hexutil.Bytes `json:"raw,omitempty"                                       testomit:"false"` //+ x
	S         *big.Int       `json:"s"                              gencodec:"required"`
	StandardV *big.Int       `json:"standardV,omitempty"`
	// Status not guarantee to be valid, it's valid for some evm compatiable chains, such as conflux chain
	Status           *uint64         `json:"status,omitempty"`
	To               *common.Address `json:"to" rlp:"nil"`
	TransactionIndex *uint64         `json:"transactionIndex"`
	Type             *uint64         `json:"type,omitempty"`
	V                *big.Int        `json:"v"                              gencodec:"required"`
	Value            *big.Int        `json:"value"                          gencodec:"required"`
}

type plainTransactionMarshaling struct {
	Accesses             types.AccessList `json:"accessList,omitempty"`
	BlockHash            *common.Hash     `json:"blockHash"`
	BlockNumber          *hexutil.Big     `json:"blockNumber"`
	ChainID              *hexutil.Big     `json:"chainId,omitempty"`
	Creates              *common.Address  `json:"creates,omitempty"`
	From                 common.Address   `json:"from"`
	Gas                  hexutil.Uint64   `json:"gas"`
	GasPrice             *hexutil.Big     `json:"gasPrice"`
	Hash                 common.Hash      `json:"hash"`
	Input                hexutil.Bytes    `json:"input"`
	MaxFeePerGas         *hexutil.Big     `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *hexutil.Big     `json:"maxPriorityFeePerGas,omitempty"`
	Nonce                hexutil.Uint64   `json:"nonce"`
	PublicKey            *hexutil.Bytes   `json:"publicKey,omitempty"`
	R                    *hexutil.Big     `json:"r"`
	Raw                  *hexutil.Bytes   `json:"raw,omitempty"`
	S                    *hexutil.Big     `json:"s"`
	StandardV            *hexutil.Big     `json:"standardV,omitempty"`
	Status               *hexutil.Uint64  `json:"status,omitempty"`
	To                   *common.Address  `json:"to"`
	TransactionIndex     *hexutil.Uint64  `json:"transactionIndex"`
	Type                 *hexutil.Uint64  `json:"type"`
	V                    *hexutil.Big     `json:"v"`
	Value                *hexutil.Big     `json:"value"`
}

//go:generate gencodec -type Receipt -field-override receiptMarshaling -out gen_receipt_json.go
// "testomit" tag is used to omit the field in rpc test, omit when testomit is true and un-omit when testomit is false.
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
	Root              []byte          `json:"root,omitempty"`   // when len(receipt.PostState) > 0
	Status            *uint64         `json:"status,omitempty"` // when len(receipt.PostState) = 0
	To                *common.Address `json:"to"`
	TransactionHash   common.Hash     `json:"transactionHash"`
	TransactionIndex  uint64          `json:"transactionIndex"`
	// Not guarantee to be valid, it's valid for some evm compatiable chains, such as conflux chain
	TxExecErrorMsg *string `json:"txExecErrorMsg,omitempty"        testomit:"false"`
	Type           *uint   `json:"type,omitempty"`
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
	Root              hexutil.Bytes   `json:"root,omitempty"`   // when len(receipt.PostState) > 0
	Status            *hexutil.Uint64 `json:"status,omitempty"` // when len(receipt.PostState) = 0
	To                *common.Address `json:"to"`
	TransactionHash   common.Hash     `json:"transactionHash"`
	TransactionIndex  hexutil.Uint64  `json:"transactionIndex"`
	TxExecErrorMsg    *string         `json:"txExecErrorMsg"`
	Type              *hexutil.Uint   `json:"type,omitempty"`
}

//go:generate gencodec -type CallRequest -field-override callRequestMarshaling -out gen_call_request_json.go
// CallRequest represents the arguments to construct a new transaction
// or a message call.
type CallRequest struct {
	From                 *common.Address `json:"from,omitempty"`
	To                   *common.Address `json:"to,omitempty"`
	Gas                  *uint64         `json:"gas,omitempty"`
	GasPrice             *big.Int        `json:"gasPrice,omitempty"`
	MaxFeePerGas         *big.Int        `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *big.Int        `json:"maxPriorityFeePerGas,omitempty"`
	Value                *big.Int        `json:"value,omitempty"`
	Nonce                *uint64         `json:"nonce,omitempty"`

	Data []byte `json:"data,omitempty"`

	// Introduced by AccessListTxType transaction.
	AccessList *types.AccessList `json:"accessList,omitempty"`
	ChainID    *big.Int          `json:"chainId,omitempty"` //+ *v throw if chainId is consensus
}

type callRequestMarshaling struct {
	From                 *common.Address `json:"from,omitempty"`
	To                   *common.Address `json:"to,omitempty"`
	Gas                  *hexutil.Uint64 `json:"gas,omitempty"`
	GasPrice             *hexutil.Big    `json:"gasPrice,omitempty"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas,omitempty"`
	Value                *hexutil.Big    `json:"value,omitempty"`
	Nonce                *hexutil.Uint64 `json:"nonce,omitempty"`

	Data hexutil.Bytes `json:"data,omitempty"`

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

type BlockNumber = rpc.BlockNumber
type BlockNumberOrHash rpc.BlockNumberOrHash
type Transaction = ethtypes.Transaction
type Subscription = ethereum.Subscription

const (
	SafeBlockNumber      = rpc.SafeBlockNumber
	FinalizedBlockNumber = rpc.FinalizedBlockNumber
	PendingBlockNumber   = rpc.PendingBlockNumber
	LatestBlockNumber    = rpc.LatestBlockNumber
	EarliestBlockNumber  = rpc.EarliestBlockNumber
)

func NewBlockNumber(blockNumber int64) BlockNumber {
	return BlockNumber(blockNumber)
}

func (bnh *BlockNumberOrHash) UnmarshalJSON(data []byte) error {
	return (*rpc.BlockNumberOrHash)(bnh).UnmarshalJSON(data)
}

func (bnh BlockNumberOrHash) MarshalJSON() ([]byte, error) {
	if bnh.BlockNumber != nil {
		return json.Marshal(bnh.BlockNumber)
	}
	return json.Marshal(rpc.BlockNumberOrHash(bnh))
}

func (bnh *BlockNumberOrHash) Number() (BlockNumber, bool) {
	return (*rpc.BlockNumberOrHash)(bnh).Number()
}

func (bnh *BlockNumberOrHash) String() string {
	return (*rpc.BlockNumberOrHash)(bnh).String()
}

func (bnh *BlockNumberOrHash) Hash() (common.Hash, bool) {
	return (*rpc.BlockNumberOrHash)(bnh).Hash()
}

func BlockNumberOrHashWithNumber(blockNr BlockNumber) BlockNumberOrHash {
	return (BlockNumberOrHash)(rpc.BlockNumberOrHashWithNumber(blockNr))
}

func BlockNumberOrHashWithHash(hash common.Hash, canonical bool) BlockNumberOrHash {
	return (BlockNumberOrHash)(rpc.BlockNumberOrHashWithHash(hash, canonical))
}
