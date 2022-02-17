package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type TransactionFilter struct {
	From     *SenderArgument      `json:"from"`
	To       *ActionArgument      `json:"to"`
	Gas      *ValueFilterArgument `json:"gas"`
	GasPrice *ValueFilterArgument `json:"gasPrice"`
	Value    *ValueFilterArgument `json:"value"`
	Nonce    *ValueFilterArgument `json:"nonce"`
}

type SenderArgument struct {
	Eq common.Address `json:"eq"`
}

type ActionArgument struct {
	Eq     common.Address `json:"eq,omitempty"`
	Action string         `json:"action,omitempty"`
}

//go:generate gencodec -type ValueFilterArgument -field-override valueFilterArgumentMarshaling -out gen_value_filter_argument_json.go
type ValueFilterArgument struct {
	Eq *big.Int `json:"eq,omitempty"`
	Lt *big.Int `json:"lt,omitempty"`
	Gt *big.Int `json:"gt,omitempty"`
}

type valueFilterArgumentMarshaling struct {
	Eq *hexutil.Big `json:"eq"`
	Lt *hexutil.Big `json:"lt"`
	Gt *hexutil.Big `json:"gt"`
}
