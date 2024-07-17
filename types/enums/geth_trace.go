package enums

type GethTraceType int

const (
	GETH_TRACE_DEFAULT GethTraceType = iota
	GETH_TRACE_CALL
	GETH_TRACE_FOUR_BYTE
	GETH_TRACE_PRE_STATE
	GETH_TRACE_NOOP
	GETH_TRACE_MUX
	GETH_TRACE_JS
)

func ParseGethTraceType(tracer string) GethTraceType {
	debugTracerType, err := ParseGethDebugBuiltInTracerType(tracer)
	if err != nil {
		return GETH_TRACE_JS
	} else {
		return debugTracerType.ToGethTraceType()
	}
}
