// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*logMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (l Log) MarshalJSON() ([]byte, error) {
	type Log struct {
		Address             common.Address `json:"address"`
		BlockHash           common.Hash    `json:"blockHash"`
		BlockNumber         hexutil.Uint64 `json:"blockNumber"`
		Data                hexutil.Bytes  `json:"data"`
		Index               hexutil.Uint   `json:"logIndex"`
		LogType             *string        `json:"logType,omitempty"`
		Removed             bool           `json:"removed"`
		Topics              []common.Hash  `json:"topics"`
		TxHash              common.Hash    `json:"transactionHash"`
		TxIndex             hexutil.Uint   `json:"transactionIndex"`
		TransactionLogIndex *hexutil.Big   `json:"transactionLogIndex,omitempty"`
	}
	var enc Log
	enc.Address = l.Address
	enc.BlockHash = l.BlockHash
	enc.BlockNumber = hexutil.Uint64(l.BlockNumber)
	enc.Data = l.Data
	enc.Index = hexutil.Uint(l.Index)
	enc.LogType = l.LogType
	enc.Removed = l.Removed
	enc.Topics = l.Topics
	enc.TxHash = l.TxHash
	enc.TxIndex = hexutil.Uint(l.TxIndex)
	enc.TransactionLogIndex = (*hexutil.Big)(l.TransactionLogIndex)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (l *Log) UnmarshalJSON(input []byte) error {
	type Log struct {
		Address             *common.Address `json:"address"`
		BlockHash           *common.Hash    `json:"blockHash"`
		BlockNumber         *hexutil.Uint64 `json:"blockNumber"`
		Data                *hexutil.Bytes  `json:"data"`
		Index               *hexutil.Uint   `json:"logIndex"`
		LogType             *string         `json:"logType,omitempty"`
		Removed             *bool           `json:"removed"`
		Topics              []common.Hash   `json:"topics"`
		TxHash              *common.Hash    `json:"transactionHash"`
		TxIndex             *hexutil.Uint   `json:"transactionIndex"`
		TransactionLogIndex *hexutil.Big    `json:"transactionLogIndex,omitempty"`
	}
	var dec Log
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Address != nil {
		l.Address = *dec.Address
	}
	if dec.BlockHash != nil {
		l.BlockHash = *dec.BlockHash
	}
	if dec.BlockNumber != nil {
		l.BlockNumber = uint64(*dec.BlockNumber)
	}
	if dec.Data != nil {
		l.Data = *dec.Data
	}
	if dec.Index != nil {
		l.Index = uint(*dec.Index)
	}
	if dec.LogType != nil {
		l.LogType = dec.LogType
	}
	if dec.Removed != nil {
		l.Removed = *dec.Removed
	}
	if dec.Topics != nil {
		l.Topics = dec.Topics
	}
	if dec.TxHash != nil {
		l.TxHash = *dec.TxHash
	}
	if dec.TxIndex != nil {
		l.TxIndex = uint(*dec.TxIndex)
	}
	if dec.TransactionLogIndex != nil {
		l.TransactionLogIndex = (*big.Int)(dec.TransactionLogIndex)
	}
	return nil
}