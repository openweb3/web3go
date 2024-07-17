package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types/enums"
)

type GethTraceResult struct {
	TracerType enums.GethTraceType `json:"-"`
	Result     *GethTrace          `json:"result,omitempty"`
	Error      *string             `json:"error,omitempty"`
	TxHash     *common.Hash        `json:"txHash,omitempty"`
}

func (r *GethTraceResult) UnmarshalJSON(data []byte) error {

	type Alias GethTraceResult
	temp := struct {
		Alias
		Result any `json:"result,omitempty"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if temp.Error != nil {
		r.Error = temp.Error
		r.TxHash = temp.TxHash
		return nil
	} else {
		r.Result = &GethTrace{Type: r.TracerType}
		b, _ := json.Marshal(temp.Result)
		return json.Unmarshal(b, &r.Result)
	}
}
