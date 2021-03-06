// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*valueFilterArgumentMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (v ValueFilterArgument) MarshalJSON() ([]byte, error) {
	type ValueFilterArgument struct {
		Eq *hexutil.Big `json:"eq,omitempty"`
		Lt *hexutil.Big `json:"lt,omitempty"`
		Gt *hexutil.Big `json:"gt,omitempty"`
	}
	var enc ValueFilterArgument
	enc.Eq = (*hexutil.Big)(v.Eq)
	enc.Lt = (*hexutil.Big)(v.Lt)
	enc.Gt = (*hexutil.Big)(v.Gt)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (v *ValueFilterArgument) UnmarshalJSON(input []byte) error {
	type ValueFilterArgument struct {
		Eq *hexutil.Big `json:"eq,omitempty"`
		Lt *hexutil.Big `json:"lt,omitempty"`
		Gt *hexutil.Big `json:"gt,omitempty"`
	}
	var dec ValueFilterArgument
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Eq != nil {
		v.Eq = (*big.Int)(dec.Eq)
	}
	if dec.Lt != nil {
		v.Lt = (*big.Int)(dec.Lt)
	}
	if dec.Gt != nil {
		v.Gt = (*big.Int)(dec.Gt)
	}
	return nil
}
