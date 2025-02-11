package enums

import (
	"fmt"
)

type GethDebugBuiltInTracerType int

const (
	GETH_BUILDIN_TRACER_FOUR_BYTE GethDebugBuiltInTracerType = iota
	GETH_BUILDIN_TRACER_CALL
	GETH_BUILDIN_TRACER_PRE_STATE
	GETH_BUILDIN_TRACER_NOOP
	GETH_BUILDIN_TRACER_MUX
)

var (
	GethDebugBuiltInTracerTypeV2S map[GethDebugBuiltInTracerType]string
	GethDebugBuiltInTracerTypeS2V map[string]GethDebugBuiltInTracerType
)

func init() {
	GethDebugBuiltInTracerTypeV2S = map[GethDebugBuiltInTracerType]string{
		GETH_BUILDIN_TRACER_FOUR_BYTE: "4byteTracer",
		GETH_BUILDIN_TRACER_CALL:      "callTracer",
		GETH_BUILDIN_TRACER_PRE_STATE: "prestateTracer",
		GETH_BUILDIN_TRACER_NOOP:      "noopTracer",
		GETH_BUILDIN_TRACER_MUX:       "muxTracer",
	}

	GethDebugBuiltInTracerTypeS2V = make(map[string]GethDebugBuiltInTracerType)
	for k, v := range GethDebugBuiltInTracerTypeV2S {
		GethDebugBuiltInTracerTypeS2V[v] = k
	}
}

func (t GethDebugBuiltInTracerType) String() string {
	if t == 0 {
		return ""
	}

	v, ok := GethDebugBuiltInTracerTypeV2S[t]
	if ok {
		return v
	}
	return "unknown"
}

func (t GethDebugBuiltInTracerType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *GethDebugBuiltInTracerType) UnmarshalText(data []byte) error {
	v, ok := GethDebugBuiltInTracerTypeS2V[string(data)]
	if ok {
		*t = v
		return nil
	}
	return fmt.Errorf("unknown tracer type %v", string(data))
}

func ParseGethDebugBuiltInTracerType(str string) (*GethDebugBuiltInTracerType, error) {
	v, ok := GethDebugBuiltInTracerTypeS2V[str]
	if !ok {
		return nil, fmt.Errorf("unknown tracer type %v", str)
	}
	return &v, nil
}

func (t GethDebugBuiltInTracerType) ToGethTraceType() GethTraceType {
	switch t {
	case GETH_BUILDIN_TRACER_FOUR_BYTE:
		return GETH_TRACE_FOUR_BYTE
	case GETH_BUILDIN_TRACER_CALL:
		return GETH_TRACE_CALL
	case GETH_BUILDIN_TRACER_PRE_STATE:
		return GETH_TRACE_PRE_STATE
	case GETH_BUILDIN_TRACER_NOOP:
		return GETH_TRACE_NOOP
	case GETH_BUILDIN_TRACER_MUX:
		return GETH_TRACE_MUX
	default:
		return GETH_TRACE_JS
	}
}
