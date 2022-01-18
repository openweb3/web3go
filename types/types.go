package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	ethrpctypes "github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
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
	Transactions     TransactionOrHashList `json:"transactions"`
	TransactionsRoot common.Hash           `json:"transactionsRoot"`
	Uncles           []common.Hash         `json:"uncles"`
	Sha3Uncles       common.Hash           `json:"sha3Uncles"`
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
	Logs              []*types.Log    `json:"logs"` //"logs"  [][]*types.Log // when receipt.Logs == nil
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

// FilterQuery contains options for contract log filtering.
type FilterQuery struct {
	BlockHash *common.Hash     // used by eth_getLogs, return logs only from block with this hash
	FromBlock *BlockNumber     // beginning of the queried range, nil means latest block
	ToBlock   *BlockNumber     // end of the range, nil means latest block
	Addresses []common.Address // restricts matches to events created by specific contracts

	// The Topic list restricts matches to particular event topics. Each event has a list
	// of topics. Topics matches a prefix of that list. An empty element slice matches any
	// topic. Non-empty elements represent an alternative that matches any of the
	// contained topics.
	//
	// Examples:
	// {} or nil          matches any topic list
	// {{A}}              matches topic A in first position
	// {{}, {B}}          matches any topic in first position AND B in second position
	// {{A}, {B}}         matches topic A in first position AND B in second position
	// {{A, B}, {C, D}}   matches topic (A OR B) in first position AND (C OR D) in second position
	Topics [][]common.Hash
}

func (args *FilterQuery) UnmarshalJSON(data []byte) error {
	var fc filters.FilterCriteria
	if err := json.Unmarshal(data, &fc); err != nil {
		return errors.Wrapf(err, "failed to unmarshal filter criteria")
	}

	args.BlockHash = fc.BlockHash
	args.FromBlock = BigIntToBlockNumber(fc.FromBlock)
	args.ToBlock = BigIntToBlockNumber(fc.ToBlock)
	args.Addresses = fc.Addresses
	args.Topics = fc.Topics
	return nil
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
