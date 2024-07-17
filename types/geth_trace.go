package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/openweb3/web3go/types/enums"
	"github.com/pkg/errors"
)

type GethTrace struct {
	Type           enums.GethTraceType
	Default        *DefaultFrame
	CallTracer     *CallFrame
	FourByteTracer *FourByteFrame
	PreStateTracer *PreStateFrame
	NoopTracer     *NoopFrame
	MuxTracer      *MuxFrame
	Js             interface{}
}

func (g GethTrace) MarshalJSON() ([]byte, error) {
	switch g.Type {
	case enums.GETH_TRACE_DEFAULT:
		return json.Marshal(g.Default)
	case enums.GETH_TRACE_CALL:
		return json.Marshal(g.CallTracer)
	case enums.GETH_TRACE_FOUR_BYTE:
		return json.Marshal(g.FourByteTracer)
	case enums.GETH_TRACE_PRE_STATE:
		return json.Marshal(g.PreStateTracer)
	case enums.GETH_TRACE_NOOP:
		return json.Marshal(g.NoopTracer)
	case enums.GETH_TRACE_MUX:
		return json.Marshal(g.MuxTracer)
	case enums.GETH_TRACE_JS:
		return json.Marshal(g.Js)
	default:
		return nil, errors.New("unknown trace type")
	}
}

func (g *GethTrace) UnmarshalJSON(data []byte) error {
	if g == nil {
		return errors.New("must specify the tracer type")
	}
	switch g.Type {
	case enums.GETH_TRACE_DEFAULT:
		return json.Unmarshal(data, &g.Default)
	case enums.GETH_TRACE_CALL:
		return json.Unmarshal(data, &g.CallTracer)
	case enums.GETH_TRACE_FOUR_BYTE:
		return json.Unmarshal(data, &g.FourByteTracer)
	case enums.GETH_TRACE_PRE_STATE:
		return json.Unmarshal(data, &g.PreStateTracer)
	case enums.GETH_TRACE_NOOP:
		return json.Unmarshal(data, &g.NoopTracer)
	case enums.GETH_TRACE_MUX:
		return json.Unmarshal(data, &g.MuxTracer)
	case enums.GETH_TRACE_JS:
		return json.Unmarshal(data, &g.Js)
	default:
		return errors.New("unknown tracer type")
	}
}

type DefaultFrame struct {
	Failed      bool          `json:"failed"`
	Gas         uint64        `json:"gas"`
	ReturnValue hexutil.Bytes `json:"returnValue"`
	StructLogs  []*StructLog  `json:"structLogs"`
}

type StructLog struct {
	Pc            uint64             `json:"pc"`
	Op            string             `json:"op"`
	Gas           uint64             `json:"gas"`
	GasCost       uint64             `json:"gasCost"`
	Depth         uint64             `json:"depth"`
	Error         *string            `json:"error,omitempty"`
	Stack         []*hexutil.Big     `json:"stack,omitempty"`
	ReturnData    *hexutil.Bytes     `json:"returnData,omitempty"`
	Memory        []string           `json:"memory,omitempty"`
	MemSize       *uint64            `json:"memSize,omitempty"`
	Storage       *map[string]string `json:"storage,omitempty"`
	RefundCounter *uint64            `json:"refund,omitempty"`
}

type CallFrame struct {
	From         common.Address  `json:"from"`
	Gas          *hexutil.Big    `json:"gas,omitempty"`
	GasUsed      *hexutil.Big    `json:"gasUsed,omitempty"`
	To           *common.Address `json:"to,omitempty"`
	Input        *hexutil.Bytes  `json:"input"`
	Output       *hexutil.Bytes  `json:"output,omitempty"`
	Error        *string         `json:"error,omitempty"`
	RevertReason *string         `json:"revertReason,omitempty"`
	Calls        []CallFrame     `json:"calls,omitempty"`
	Logs         []CallLogFrame  `json:"logs,omitempty"`
	Value        *hexutil.Big    `json:"value,omitempty"`
	Type         string          `json:"type"`
}

type CallLogFrame struct {
	Address *common.Address `json:"address,omitempty"`
	Topics  []common.Hash   `json:"topics,omitempty"`
	Data    *hexutil.Bytes  `json:"data,omitempty"`
}

type FourByteFrame = map[string]uint64

type PreStateFrame struct {
	// The default mode returns the accounts necessary to execute a given transaction.
	//
	// It re-executes the given transaction and tracks every part of state that is touched.
	Default *PreStateMode `json:"Default,omitempty"`
	// Diff mode returns the differences between the transaction's pre and post-state (i.e. what
	// changed because the transaction happened).
	Diff *DiffMode `json:"Diff,omitempty"`
}

// Includes all the account states necessary to execute a given transaction.
//
// This corresponds to the default mode of the [PreStateConfig].
//
// The [AccountState]'s storage will include all touched slots of an account.
type PreStateMode struct {
	Accounts map[string]AccountState `json:"Accounts"`
}

// Represents the account states before and after the transaction is executed.
//
// This corresponds to the [DiffMode] of the [PreStateConfig].
//
// This will only contain changed [AccountState]s, created accounts will not be included in the pre
// state and selfdestructed accounts will not be included in the post state.
type DiffMode struct {
	// The account states after the transaction is executed.
	Post map[string]AccountState `json:"Post"`
	// The account states before the transaction is executed.
	Pre map[string]AccountState `json:"Pre"`
}

type AccountState struct {
	Balance *hexutil.Big                `json:"balance,omitempty"`
	Code    hexutil.Bytes               `json:"code,omitempty"`
	Nonce   *uint64                     `json:"nonce,omitempty"`
	Storage map[common.Hash]common.Hash `json:"storage,omitempty"`
}

type NoopFrame struct {
}

type MuxFrame map[enums.GethDebugBuiltInTracerType]GethTrace
