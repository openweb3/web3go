package client

import (
	"github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/go-rpc-provider/interfaces"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/types"
)

type RpcFilterClient struct {
	BaseClient
}

func NewRpcFilterClient(provider interfaces.Provider) *RpcFilterClient {
	_client := &RpcFilterClient{}
	_client.MiddlewarableProvider = providers.NewMiddlewarableProvider(provider)
	return _client
}

// Returns id of new filter.
func (c *RpcFilterClient) NewLogFilter(filter *types.FilterQuery) (val *rpc.ID, err error) {
	err = c.CallContext(c.getContext(), &val, "eth_newFilter", filter)
	return
}

// Returns id of new block filter.
func (c *RpcFilterClient) NewBlockFilter() (val *rpc.ID, err error) {
	err = c.CallContext(c.getContext(), &val, "eth_newBlockFilter")
	return
}

// Returns id of new block filter.
func (c *RpcFilterClient) NewPendingTransactionFilter() (val *rpc.ID, err error) {
	err = c.CallContext(c.getContext(), &val, "eth_newPendingTransactionFilter")
	return
}

// Returns filter changes since last poll.
func (c *RpcFilterClient) GetFilterChanges(filterID rpc.ID) (val *types.FilterChanges, err error) {
	err = c.CallContext(c.getContext(), &val, "eth_getFilterChanges", filterID)
	return
}

// Returns all logs matching given filter (in a range 'from' - 'to').
func (c *RpcFilterClient) GetFilterLogs(filterID rpc.ID) (val []types.Log, err error) {
	err = c.CallContext(c.getContext(), &val, "eth_getFilterLogs", filterID)
	return
}

// Uninstalls filter.
func (c *RpcFilterClient) UninstallFilter(filterID rpc.ID) (val bool, err error) {
	err = c.CallContext(c.getContext(), &val, "eth_uninstallFilter", filterID)
	return
}
