// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*traceResultsWithTransactionHashMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (t TraceResultsWithTransactionHash) MarshalJSON() ([]byte, error) {
	type TraceResultsWithTransactionHash struct {
		Output          hexutil.Bytes `json:"output"`
		Trace           []Trace       `json:"trace"`
		VmTrace         *VMTrace      `json:"vmTrace"`
		StateDiff       *StateDiff    `json:"stateDiff"`
		TransactionHash common.Hash   `json:"transactionHash"`
	}
	var enc TraceResultsWithTransactionHash
	enc.Output = t.Output
	enc.Trace = t.Trace
	enc.VmTrace = t.VmTrace
	enc.StateDiff = t.StateDiff
	enc.TransactionHash = t.TransactionHash
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *TraceResultsWithTransactionHash) UnmarshalJSON(input []byte) error {
	type TraceResultsWithTransactionHash struct {
		Output          *hexutil.Bytes `json:"output"`
		Trace           []Trace        `json:"trace"`
		VmTrace         *VMTrace       `json:"vmTrace"`
		StateDiff       *StateDiff     `json:"stateDiff"`
		TransactionHash *common.Hash   `json:"transactionHash"`
	}
	var dec TraceResultsWithTransactionHash
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Output != nil {
		t.Output = *dec.Output
	}
	if dec.Trace != nil {
		t.Trace = dec.Trace
	}
	if dec.VmTrace != nil {
		t.VmTrace = dec.VmTrace
	}
	if dec.StateDiff != nil {
		t.StateDiff = dec.StateDiff
	}
	if dec.TransactionHash != nil {
		t.TransactionHash = *dec.TransactionHash
	}
	return nil
}
