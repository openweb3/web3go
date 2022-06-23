// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = (*plainTransactionMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (p PlainTransaction) MarshalJSON() ([]byte, error) {
	type PlainTransaction struct {
		Accesses             types.AccessList `json:"accessList,omitempty"`
		BlockHash            *common.Hash     `json:"blockHash"`
		BlockNumber          *hexutil.Big     `json:"blockNumber"`
		ChainID              *hexutil.Big     `json:"chainId,omitempty"`
		Creates              *common.Address  `json:"creates,omitempty"                                   testomit:"false"`
		From                 common.Address   `json:"from"`
		Gas                  hexutil.Uint64   `json:"gas"                            gencodec:"required"`
		GasPrice             *hexutil.Big     `json:"gasPrice"`
		Hash                 common.Hash      `json:"hash"`
		Input                hexutil.Bytes    `json:"input"                          gencodec:"required"`
		MaxFeePerGas         *hexutil.Big     `json:"maxFeePerGas,omitempty"`
		MaxPriorityFeePerGas *hexutil.Big     `json:"maxPriorityFeePerGas,omitempty"`
		Nonce                hexutil.Uint64   `json:"nonce"                          gencodec:"required"`
		PublicKey            *hexutil.Bytes   `json:"publicKey,omitempty"                                 testomit:"false"`
		R                    *hexutil.Big     `json:"r"                              gencodec:"required"`
		Raw                  *hexutil.Bytes   `json:"raw,omitempty"                                       testomit:"false"`
		S                    *hexutil.Big     `json:"s"                              gencodec:"required"`
		StandardV            *hexutil.Big     `json:"standardV,omitempty"`
		Status               *hexutil.Uint64  `json:"status,omitempty"`
		To                   *common.Address  `json:"to" rlp:"nil"`
		TransactionIndex     *hexutil.Uint64  `json:"transactionIndex"`
		Type                 *hexutil.Uint64  `json:"type,omitempty"`
		V                    *hexutil.Big     `json:"v"                              gencodec:"required"`
		Value                *hexutil.Big     `json:"value"                          gencodec:"required"`
	}
	var enc PlainTransaction
	enc.Accesses = p.Accesses
	enc.BlockHash = p.BlockHash
	enc.BlockNumber = (*hexutil.Big)(p.BlockNumber)
	enc.ChainID = (*hexutil.Big)(p.ChainID)
	enc.Creates = p.Creates
	enc.From = p.From
	enc.Gas = hexutil.Uint64(p.Gas)
	enc.GasPrice = (*hexutil.Big)(p.GasPrice)
	enc.Hash = p.Hash
	enc.Input = p.Input
	enc.MaxFeePerGas = (*hexutil.Big)(p.MaxFeePerGas)
	enc.MaxPriorityFeePerGas = (*hexutil.Big)(p.MaxPriorityFeePerGas)
	enc.Nonce = hexutil.Uint64(p.Nonce)
	enc.PublicKey = p.PublicKey
	enc.R = (*hexutil.Big)(p.R)
	enc.Raw = p.Raw
	enc.S = (*hexutil.Big)(p.S)
	enc.StandardV = (*hexutil.Big)(p.StandardV)
	enc.Status = (*hexutil.Uint64)(p.Status)
	enc.To = p.To
	enc.TransactionIndex = (*hexutil.Uint64)(p.TransactionIndex)
	enc.Type = (*hexutil.Uint64)(p.Type)
	enc.V = (*hexutil.Big)(p.V)
	enc.Value = (*hexutil.Big)(p.Value)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (p *PlainTransaction) UnmarshalJSON(input []byte) error {
	type PlainTransaction struct {
		Accesses             *types.AccessList `json:"accessList,omitempty"`
		BlockHash            *common.Hash      `json:"blockHash"`
		BlockNumber          *hexutil.Big      `json:"blockNumber"`
		ChainID              *hexutil.Big      `json:"chainId,omitempty"`
		Creates              *common.Address   `json:"creates,omitempty"                                   testomit:"false"`
		From                 *common.Address   `json:"from"`
		Gas                  *hexutil.Uint64   `json:"gas"                            gencodec:"required"`
		GasPrice             *hexutil.Big      `json:"gasPrice"`
		Hash                 *common.Hash      `json:"hash"`
		Input                *hexutil.Bytes    `json:"input"                          gencodec:"required"`
		MaxFeePerGas         *hexutil.Big      `json:"maxFeePerGas,omitempty"`
		MaxPriorityFeePerGas *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
		Nonce                *hexutil.Uint64   `json:"nonce"                          gencodec:"required"`
		PublicKey            *hexutil.Bytes    `json:"publicKey,omitempty"                                 testomit:"false"`
		R                    *hexutil.Big      `json:"r"                              gencodec:"required"`
		Raw                  *hexutil.Bytes    `json:"raw,omitempty"                                       testomit:"false"`
		S                    *hexutil.Big      `json:"s"                              gencodec:"required"`
		StandardV            *hexutil.Big      `json:"standardV,omitempty"`
		Status               *hexutil.Uint64   `json:"status,omitempty"`
		To                   *common.Address   `json:"to" rlp:"nil"`
		TransactionIndex     *hexutil.Uint64   `json:"transactionIndex"`
		Type                 *hexutil.Uint64   `json:"type,omitempty"`
		V                    *hexutil.Big      `json:"v"                              gencodec:"required"`
		Value                *hexutil.Big      `json:"value"                          gencodec:"required"`
	}
	var dec PlainTransaction
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Accesses != nil {
		p.Accesses = *dec.Accesses
	}
	if dec.BlockHash != nil {
		p.BlockHash = dec.BlockHash
	}
	if dec.BlockNumber != nil {
		p.BlockNumber = (*big.Int)(dec.BlockNumber)
	}
	if dec.ChainID != nil {
		p.ChainID = (*big.Int)(dec.ChainID)
	}
	if dec.Creates != nil {
		p.Creates = dec.Creates
	}
	if dec.From != nil {
		p.From = *dec.From
	}
	if dec.Gas == nil {
		return errors.New("missing required field 'gas' for PlainTransaction")
	}
	p.Gas = uint64(*dec.Gas)
	if dec.GasPrice != nil {
		p.GasPrice = (*big.Int)(dec.GasPrice)
	}
	if dec.Hash != nil {
		p.Hash = *dec.Hash
	}
	if dec.Input == nil {
		return errors.New("missing required field 'input' for PlainTransaction")
	}
	p.Input = *dec.Input
	if dec.MaxFeePerGas != nil {
		p.MaxFeePerGas = (*big.Int)(dec.MaxFeePerGas)
	}
	if dec.MaxPriorityFeePerGas != nil {
		p.MaxPriorityFeePerGas = (*big.Int)(dec.MaxPriorityFeePerGas)
	}
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' for PlainTransaction")
	}
	p.Nonce = uint64(*dec.Nonce)
	if dec.PublicKey != nil {
		p.PublicKey = dec.PublicKey
	}
	if dec.R == nil {
		return errors.New("missing required field 'r' for PlainTransaction")
	}
	p.R = (*big.Int)(dec.R)
	if dec.Raw != nil {
		p.Raw = dec.Raw
	}
	if dec.S == nil {
		return errors.New("missing required field 's' for PlainTransaction")
	}
	p.S = (*big.Int)(dec.S)
	if dec.StandardV != nil {
		p.StandardV = (*big.Int)(dec.StandardV)
	}
	if dec.Status != nil {
		p.Status = (*uint64)(dec.Status)
	}
	if dec.To != nil {
		p.To = dec.To
	}
	if dec.TransactionIndex != nil {
		p.TransactionIndex = (*uint64)(dec.TransactionIndex)
	}
	if dec.Type != nil {
		p.Type = (*uint64)(dec.Type)
	}
	if dec.V == nil {
		return errors.New("missing required field 'v' for PlainTransaction")
	}
	p.V = (*big.Int)(dec.V)
	if dec.Value == nil {
		return errors.New("missing required field 'value' for PlainTransaction")
	}
	p.Value = (*big.Int)(dec.Value)
	return nil
}
