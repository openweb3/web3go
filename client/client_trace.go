package client

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/openweb3/go-rpc-provider/interfaces"
	"github.com/openweb3/web3go/types"
)

type RpcTraceClient struct {
	core interfaces.Provider
}

func NewRpcTraceClient(provider interfaces.Provider) *RpcTraceClient {
	return &RpcTraceClient{
		core: provider,
	}
}

/// RPC Metadata
// type Metadata

/// Returns traces matching given filter.
func (c *RpcTraceClient) Filter(ctx context.Context, traceFilter types.TraceFilter) (val []types.LocalizedTrace, err error) {
	err = c.core.CallContext(ctx, &val, "trace_filter", traceFilter)
	return
}

/// Returns transaction trace at given index.
func (c *RpcTraceClient) Trace(ctx context.Context, transactionHash common.Hash, indexes []uint) (val *types.LocalizedTrace, err error) {
	err = c.core.CallContext(ctx, &val, "trace_get", transactionHash, indexes)
	return
}

/// Returns all traces of given transaction.
func (c *RpcTraceClient) Transactions(ctx context.Context, transactionHash common.Hash) (val []types.LocalizedTrace, err error) {
	err = c.core.CallContext(ctx, &val, "trace_transaction", transactionHash)
	return
}

/// Returns all traces produced at given block.
func (c *RpcTraceClient) Blocks(ctx context.Context, blockNumber types.BlockNumberOrHash) (val []types.LocalizedTrace, err error) {
	err = c.core.CallContext(ctx, &val, "trace_block", blockNumber)
	return
}

/// Executes the given call and returns a number of possible traces for it.
func (c *RpcTraceClient) Call(ctx context.Context, request types.CallRequest, options types.TraceOptions, blockNumber *types.BlockNumberOrHash) (val types.TraceResults, err error) {
	err = c.core.CallContext(ctx, &val, "trace_call", request, options, blockNumber)
	return
}

// /// Executes all given calls and returns a number of possible traces for each of it.
// func(c *RpcTraceClient) CallMany(_ types.Vec<(CallRequest, TraceOptions)>, _ *types.BlockNumber)(val []types.TraceResults, err error) {
//     TraceOptions)>

//     err = c.core.CallContext(ctx,&val, "trace_callMany", _, TraceOptions)>, _)

//     return
// }

/// Executes the given raw transaction and returns a number of possible traces for it.
func (c *RpcTraceClient) RawTransaction(ctx context.Context, rawTransaction []byte, options types.TraceOptions, blockNumber *types.BlockNumberOrHash) (val types.TraceResults, err error) {
	_rawTransaction := (hexutil.Bytes)(rawTransaction)
	err = c.core.CallContext(ctx, &val, "trace_rawTransaction", _rawTransaction, options, blockNumber)
	return
}

/// Executes the transaction with the given hash and returns a number of possible traces for it.
func (c *RpcTraceClient) ReplayTransaction(ctx context.Context, transactionHash common.Hash, options types.TraceOptions) (val types.TraceResults, err error) {
	err = c.core.CallContext(ctx, &val, "trace_replayTransaction", transactionHash, options)
	return
}

/// Executes all the transactions at the given block and returns a number of possible traces for each transaction.
func (c *RpcTraceClient) ReplayBlockTransactions(ctx context.Context, blockNumber types.BlockNumberOrHash, options types.TraceOptions) (val []types.TraceResultsWithTransactionHash, err error) {
	err = c.core.CallContext(ctx, &val, "trace_replayBlockTransactions", blockNumber, options)
	return
}
