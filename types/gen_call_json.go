// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*callMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (c Call) MarshalJSON() ([]byte, error) {
	type Call struct {
		From     common.Address `json:"from"`
		To       common.Address `json:"to"`
		Value    *hexutil.Big   `json:"value"`
		Gas      *hexutil.Big   `json:"gas"`
		Input    hexutil.Bytes  `json:"input"`
		CallType CallType       `json:"callType"`
	}
	var enc Call
	enc.From = c.From
	enc.To = c.To
	enc.Value = (*hexutil.Big)(c.Value)
	enc.Gas = (*hexutil.Big)(c.Gas)
	enc.Input = c.Input
	enc.CallType = c.CallType
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (c *Call) UnmarshalJSON(input []byte) error {
	type Call struct {
		From     *common.Address `json:"from"`
		To       *common.Address `json:"to"`
		Value    *hexutil.Big    `json:"value"`
		Gas      *hexutil.Big    `json:"gas"`
		Input    *hexutil.Bytes  `json:"input"`
		CallType *CallType       `json:"callType"`
	}
	var dec Call
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.From != nil {
		c.From = *dec.From
	}
	if dec.To != nil {
		c.To = *dec.To
	}
	if dec.Value != nil {
		c.Value = (*big.Int)(dec.Value)
	}
	if dec.Gas != nil {
		c.Gas = (*big.Int)(dec.Gas)
	}
	if dec.Input != nil {
		c.Input = *dec.Input
	}
	if dec.CallType != nil {
		c.CallType = *dec.CallType
	}
	return nil
}
