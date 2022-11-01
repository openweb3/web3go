package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/go-rpc-provider/interfaces"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/types"
)

// /// Eth filters rpc api (polling).
// // TODO: do filters api properly
// #[rpc(server)]
// pub trait EthFilter {
//     /// Returns id of new filter.
//     #[rpc(name = "eth_newFilter")]
//     fn new_filter(&self, _: EthRpcLogFilter) -> Result<U256>;

//     /// Returns id of new block filter.
//     #[rpc(name = "eth_newBlockFilter")]
//     fn new_block_filter(&self) -> Result<U256>;

//     /// Returns id of new block filter.
//     #[rpc(name = "eth_newPendingTransactionFilter")]
//     fn new_pending_transaction_filter(&self) -> Result<U256>;

//     /// Returns filter changes since last poll.
//     #[rpc(name = "eth_getFilterChanges")]
//     fn filter_changes(&self, _: Index) -> Result<FilterChanges>;

//     /// Returns all logs matching given filter (in a range 'from' - 'to').
//     #[rpc(name = "eth_getFilterLogs")]
//     fn filter_logs(&self, _: Index) -> Result<Vec<Log>>;

//     /// Uninstalls filter.
//     #[rpc(name = "eth_uninstallFilter")]
//     fn uninstall_filter(&self, _: Index) -> Result<bool>;
// }

type RpcFilterClient struct {
	*providers.MiddlewarableProvider
}

func NewRpcFilterClient(provider interfaces.Provider) *RpcFilterClient {
	return &RpcFilterClient{
		MiddlewarableProvider: providers.NewMiddlewarableProvider(provider),
	}
}

// Returns id of new filter.
func (c *RpcFilterClient) NewLogFilter(filter *types.FilterQuery) (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_newFilter", filter)
	val = (*big.Int)(_val)
	return
}

// Returns id of new block filter.
func (c *RpcFilterClient) NewBlockFilter() (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_newBlockFilter")
	val = (*big.Int)(_val)
	return
}

// Returns id of new block filter.
func (c *RpcFilterClient) NewPendingTransactionFilter() (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_newPendingTransactionFilter")
	val = (*big.Int)(_val)
	return
}

// Returns filter changes since last poll.
func (c *RpcFilterClient) GetFilterChanges(filterID *big.Int) (val *types.FilterChanges, err error) {
	_filterID := (*hexutil.Big)(filterID)
	// var _val interface{}
	err = c.CallContext(context.Background(), &val, "eth_getFilterChanges", _filterID)
	return
}

// Returns all logs matching given filter (in a range 'from' - 'to').
func (c *RpcFilterClient) GetFilterLogs(filterID *big.Int) (val []types.Log, err error) {
	_filterID := (*hexutil.Big)(filterID)
	err = c.CallContext(context.Background(), &val, "eth_getFilterLogs", _filterID)
	return
}

// Uninstalls filter.
func (c *RpcFilterClient) UninstallFilter(filterID *big.Int) (val bool, err error) {
	_filterID := (*hexutil.Big)(filterID)
	err = c.CallContext(context.Background(), &val, "eth_uninstallFilter", _filterID)
	return
}
