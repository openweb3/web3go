package types

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type TraceOptions []string

type TraceType string

const (
	TRACE_CALL    = "call"
	TRACE_CREATE  = "create"
	TRACE_SUICIDE = "suicide"
	TRACE_REWARD  = "reward"
)

type Trace struct {
	Type         TraceType   `json:"type"`
	Action       interface{} `json:"action"`
	Result       interface{} `json:"result,omitempty"`
	Error        *string     `json:"error,omitempty"`
	TraceAddress []uint      `json:"traceAddress"`
	Subtraces    uint        `json:"subtraces"`
}

type LocalizedTrace struct {
	Type                TraceType    `json:"type"`
	Action              interface{}  `json:"action"`
	Result              interface{}  `json:"result,omitempty"`
	Error               *string      `json:"error,omitempty"`
	TraceAddress        []uint       `json:"traceAddress"`
	Subtraces           uint         `json:"subtraces"`
	TransactionPosition *uint        `json:"transactionPosition"`
	TransactionHash     *common.Hash `json:"transactionHash"`
	BlockNumber         uint64       `json:"blockNumber"`
	BlockHash           common.Hash  `json:"blockHash"`
	Valid               *bool        `json:"valid,omitempty"` // exist in conflux-espace, not in openethereum
}

type StateDiff map[common.Hash]AccountDiff

type TraceFilter struct {
	FromBlock   *BlockNumber     `json:"fromBlock"`
	ToBlock     *BlockNumber     `json:"toBlock"`
	FromAddress []common.Address `json:"fromAddress"`
	ToAddress   []common.Address `json:"toAddress"`
	After       *uint            `json:"after"`
	Count       *uint            `json:"count"`
}

//go:generate gencodec -type TraceResults -field-override traceResultsMarshaling -out gen_trace_results_json.go
type TraceResults struct {
	Output    []byte     `json:"output"`
	Trace     []Trace    `json:"trace"`
	VmTrace   *VMTrace   `json:"vmTrace"`
	StateDiff *StateDiff `json:"stateDiff"`
}

type traceResultsMarshaling struct {
	Output    hexutil.Bytes `json:"output"`
	Trace     []Trace       `json:"trace"`
	VmTrace   *VMTrace      `json:"vmTrace"`
	StateDiff *StateDiff    `json:"stateDiff"`
}

type VMOperation struct {
	Pc   uint                 `json:"pc"`
	Cost uint64               `json:"cost"`
	Ex   *VMExecutedOperation `json:"ex"`
	Sub  *VMTrace             `json:"sub"`
}

// gencodec -type VMExecutedOperation -field-override vMExecutedOperationMarshaling -out gen_vm_executed_operation_json.go
type VMExecutedOperation struct {
	Used  uint64       `json:"used"`
	Push  []*big.Int   `json:"push"`
	Mem   *MemoryDiff  `json:"mem"`
	Store *StorageDiff `json:"store"`
}

type vMExecutedOperationMarshaling struct {
	Used  uint64         `json:"used"`
	Push  []*hexutil.Big `json:"push"`
	Mem   *MemoryDiff    `json:"mem"`
	Store *StorageDiff   `json:"store"`
}

//go:generate gencodec -type VMTrace -field-override vMTraceMarshaling -out gen_vm_trace_json.go
type VMTrace struct {
	Code []byte        `json:"code"`
	Ops  []VMOperation `json:"ops"`
}

type vMTraceMarshaling struct {
	Code hexutil.Bytes `json:"code"`
	Ops  []VMOperation `json:"ops"`
}

//go:generate gencodec -type TraceResultsWithTransactionHash -field-override traceResultsWithTransactionHashMarshaling -out gen_trace_results_with_transaction_hash_json.go
type TraceResultsWithTransactionHash struct {
	Output          []byte      `json:"output"`
	Trace           []Trace     `json:"trace"`
	VmTrace         *VMTrace    `json:"vmTrace"`
	StateDiff       *StateDiff  `json:"stateDiff"`
	TransactionHash common.Hash `json:"transactionHash"`
}

type traceResultsWithTransactionHashMarshaling struct {
	Output          hexutil.Bytes `json:"output"`
	Trace           []Trace       `json:"trace"`
	VmTrace         *VMTrace      `json:"vmTrace"`
	StateDiff       *StateDiff    `json:"stateDiff"`
	TransactionHash common.Hash   `json:"transactionHash"`
}

//go:generate gencodec -type MemoryDiff -field-override memoryDiffMarshaling -out gen_memory_diff_json.go
type MemoryDiff struct {
	Off  uint   `json:"off"`
	Data []byte `json:"data"`
}

type memoryDiffMarshaling struct {
	Off  uint          `json:"off"`
	Data hexutil.Bytes `json:"data"`
}

//go:generate gencodec -type StorageDiff -field-override StorageDiffMarshaling -out gen_storage_diff_json.go
type StorageDiff struct {
	Key *big.Int `json:"key"`
	Val *big.Int `json:"val"`
}

type StorageDiffMarshaling struct {
	Key *hexutil.Big `json:"key"`
	Val *hexutil.Big `json:"val"`
}

type AccountDiff struct {
	Balance string                 `json:"balance"`
	Nonce   string                 `json:"nonce"`
	Code    string                 `json:"code"`
	Storage map[common.Hash]string `json:"storage"`
}

// UnmarshalJSON unmarshals Input and Init type from []byte to hexutil.Bytes
func (l *Trace) UnmarshalJSON(data []byte) error {

	type alias Trace

	a := alias{}
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}
	*l = Trace(a)

	l.Action, l.Result, err = getActionAndResult(data)
	return err
}

func (l *LocalizedTrace) UnmarshalJSON(data []byte) error {
	type alias LocalizedTrace

	a := alias{}
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}
	*l = LocalizedTrace(a)

	l.Action, l.Result, err = getActionAndResult(data)
	return err
}

func getActionAndResult(data []byte) (interface{}, interface{}, error) {
	tmp := struct {
		Type   TraceType              `json:"type"`
		Action map[string]interface{} `json:"action"`
		Result map[string]interface{} `json:"result"`
		Error  string                 `json:"error"`
	}{}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, nil, err
	}

	var action, result interface{}
	if action, err = parseAction(tmp.Action, tmp.Type); err != nil {
		return nil, nil, err
	}
	if result, err = parseActionResult(tmp.Result, tmp.Error, tmp.Type); err != nil {
		return nil, nil, err
	}

	return action, result, nil
}

func parseAction(actionInMap map[string]interface{}, actionType TraceType) (interface{}, error) {
	actionJson, err := json.Marshal(actionInMap)
	if err != nil {
		return nil, err
	}

	var result interface{}
	switch actionType {
	case TRACE_CALL:
		action := Call{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case TRACE_CREATE:
		action := Create{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case TRACE_SUICIDE:
		action := Suicide{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case TRACE_REWARD:
		action := Reward{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	default:
		return nil, fmt.Errorf("unknown action type %v", actionType)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %v to %v ", string(actionJson), actionType)
	}
	return result, nil
}

// parseActionResult parses action result, result will be nil if actionError not empty
// one-to-one mapping between action result type and action type
//	action type TRACE_CALL => CallResult
//	action type TRACE_CREATE => CreateResult
//	action type TRACE_SUICIDE => uint8(0)
//	action type TRACE_REWARD => uint8(0)
func parseActionResult(actionResInMap map[string]interface{}, actionError string, actionType TraceType) (interface{}, error) {
	if actionError != "" {
		return nil, nil
	}

	actionResJson, err := json.Marshal(actionResInMap)
	if err != nil {
		return nil, err
	}

	var result interface{}
	switch actionType {
	case TRACE_CALL:
		action := CallResult{}
		err = json.Unmarshal(actionResJson, &action)
		result = action
	case TRACE_CREATE:
		action := CreateResult{}
		err = json.Unmarshal(actionResJson, &action)
		result = action
	case TRACE_SUICIDE:
		fallthrough
	case TRACE_REWARD:
		var action uint8
		err = json.Unmarshal(actionResJson, &action)
		result = action

	default:
		return nil, fmt.Errorf("unknown action type %v", actionType)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %v to %v ", string(actionResJson), actionType)
	}
	return result, nil
}
