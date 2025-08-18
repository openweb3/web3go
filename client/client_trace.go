package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/openweb3/go-rpc-provider/interfaces"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/types"
)

type RpcTraceClient struct {
	BaseClient
}

func NewRpcTraceClient(provider interfaces.Provider) *RpcTraceClient {
	_client := &RpcTraceClient{}
	_client.MiddlewarableProvider = providers.NewMiddlewarableProvider(provider)
	return _client
}

// Returns traces matching given filter.
func (c *RpcTraceClient) Filter(traceFilter types.TraceFilter) (val []types.LocalizedTrace, err error) {
	err = c.CallContext(c.getContext(), &val, "trace_filter", traceFilter)
	return
}

// Returns transaction trace at given index.
func (c *RpcTraceClient) Trace(transactionHash common.Hash, indexes []uint) (val *types.LocalizedTrace, err error) {
	err = c.CallContext(c.getContext(), &val, "trace_get", transactionHash, indexes)
	return
}

// Returns all traces of given transaction.
func (c *RpcTraceClient) Transactions(transactionHash common.Hash) (val []types.LocalizedTrace, err error) {
	err = c.CallContext(c.getContext(), &val, "trace_transaction", transactionHash)
	return
}

// Returns all traces produced at given block.
func (c *RpcTraceClient) Blocks(blockNumber types.BlockNumberOrHash) (val []types.LocalizedTrace, err error) {
	err = c.CallContext(c.getContext(), &val, "trace_block", blockNumber)
	return
}

// Executes the given call and returns a number of possible traces for it.
func (c *RpcTraceClient) Call(request types.CallRequest, options types.TraceOptions, blockNumber *types.BlockNumberOrHash) (val types.TraceResults, err error) {
	err = c.CallContext(c.getContext(), &val, "trace_call", request, options, getRealBlockNumberOrHash(blockNumber))
	return
}

// Executes the given raw transaction and returns a number of possible traces for it.
func (c *RpcTraceClient) RawTransaction(rawTransaction []byte, options types.TraceOptions, blockNumber *types.BlockNumberOrHash) (val types.TraceResults, err error) {
	_rawTransaction := (hexutil.Bytes)(rawTransaction)
	err = c.CallContext(c.getContext(), &val, "trace_rawTransaction", _rawTransaction, options, getRealBlockNumberOrHash(blockNumber))
	return
}

// Executes the transaction with the given hash and returns a number of possible traces for it.
func (c *RpcTraceClient) ReplayTransaction(transactionHash common.Hash, options types.TraceOptions) (val types.TraceResults, err error) {
	err = c.CallContext(c.getContext(), &val, "trace_replayTransaction", transactionHash, options)
	return
}

// Executes all the transactions at the given block and returns a number of possible traces for each transaction.
func (c *RpcTraceClient) ReplayBlockTransactions(blockNumber types.BlockNumberOrHash, options types.TraceOptions) (val []types.TraceResultsWithTransactionHash, err error) {
	err = c.CallContext(c.getContext(), &val, "trace_replayBlockTransactions", blockNumber, options)
	return
}

// Returns all set auth traces produced at the given block.
func (c *RpcTraceClient) BlockSetAuthTraces(blockNumber types.BlockNumberOrHash) (val []types.LocalizedSetAuthTrace, err error) {
	err = c.CallContext(c.getContext(), &val, "trace_blockSetAuth", blockNumber)
	return
}
