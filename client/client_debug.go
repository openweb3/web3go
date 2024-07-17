package client

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/go-rpc-provider/interfaces"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/types"
	"github.com/openweb3/web3go/types/enums"
)

type RpcDebugClient struct {
	BaseClient
}

func NewRpcDebugClient(provider interfaces.Provider) *RpcDebugClient {
	_client := &RpcDebugClient{}
	_client.MiddlewarableProvider = providers.NewMiddlewarableProvider(provider)
	return _client
}

func (c *RpcDebugClient) TraceTransaction(tx_hash common.Hash, opts ...*types.GethDebugTracingOptions) (val *types.GethTrace, err error) {
	opt := get1stOpt(opts)
	val = &types.GethTrace{Type: getGethTraceTypeByOpt(opt)}
	err = c.CallContext(c.getContext(), &val, "debug_traceTransaction", tx_hash, opt)
	return
}

func (c *RpcDebugClient) TraceBlockByHash(block_hash common.Hash, opts ...*types.GethDebugTracingOptions) (val []*types.GethTraceResult, err error) {
	opt := get1stOpt(opts)

	var tmpVal []any
	err = c.CallContext(c.getContext(), &tmpVal, "debug_traceBlockByHash", block_hash, opt)
	if err != nil {
		return nil, err
	}

	val = make([]*types.GethTraceResult, len(tmpVal))
	tracerType := getGethTraceTypeByOpt(opt)
	for i, v := range tmpVal {
		val[i] = &types.GethTraceResult{TracerType: tracerType}

		b, _ := json.Marshal(v)
		if err := val[i].UnmarshalJSON(b); err != nil {
			return nil, err
		}
	}

	return val, err
}

func (c *RpcDebugClient) TraceBlockByNumber(blockNumber types.BlockNumber, opts ...*types.GethDebugTracingOptions) (val []*types.GethTraceResult, err error) {
	opt := get1stOpt(opts)

	var tmpVal []any
	err = c.CallContext(c.getContext(), &tmpVal, "debug_traceBlockByNumber", blockNumber, opt)
	if err != nil {
		return nil, err
	}

	val = make([]*types.GethTraceResult, len(tmpVal))
	tracerType := getGethTraceTypeByOpt(opt)
	for i, v := range tmpVal {
		val[i] = &types.GethTraceResult{TracerType: tracerType}

		b, _ := json.Marshal(v)
		if err := val[i].UnmarshalJSON(b); err != nil {
			return nil, err
		}
	}

	return val, err
}

func (c *RpcDebugClient) TraceCall(request types.CallRequest, block_number *types.BlockNumber, opts ...*types.GethDebugTracingOptions) (val *types.GethTrace, err error) {
	opt := get1stOpt(opts)
	val = &types.GethTrace{Type: getGethTraceTypeByOpt(opt)}
	err = c.CallContext(c.getContext(), &val, "debug_traceCall", request, block_number, opt)
	return
}

func getGethTraceTypeByOpt(opts *types.GethDebugTracingOptions) enums.GethTraceType {
	t := enums.GETH_TRACE_DEFAULT
	if opts != nil {
		t = enums.ParseGethTraceType(opts.Tracer)
	}
	return t
}

func get1stOpt(opts []*types.GethDebugTracingOptions) *types.GethDebugTracingOptions {
	if len(opts) == 0 {
		return nil
	}
	return opts[0]
}
