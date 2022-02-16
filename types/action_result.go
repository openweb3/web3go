package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//go:generate gencodec -type CallResult -field-override callResultMarshaling -out gen_call_result_json.go
type CallResult struct {
	GasUsed *big.Int `json:"gasUsed"`
	Output  []byte   `json:"output"`
}

//go:generate gencodec -type CreateResult -field-override createResultMarshaling -out gen_create_result_json.go
type CreateResult struct {
	GasUsed *big.Int       `json:"gasUsed"`
	Code    []byte         `json:"code"`
	Address common.Address `json:"address"`
}

type callResultMarshaling struct {
	GasUsed *hexutil.Big  `json:"gasUsed"`
	Output  hexutil.Bytes `json:"output"`
}
type createResultMarshaling struct {
	GasUsed *hexutil.Big   `json:"gasUsed"`
	Code    hexutil.Bytes  `json:"code"`
	Address common.Address `json:"address"`
}
