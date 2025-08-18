package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/go-rpc-provider/interfaces"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/types"
)

type RpcTxPoolClient struct {
	BaseClient
}

func NewRpcTxPoolClient(provider interfaces.Provider) *RpcTxPoolClient {
	_client := &RpcTxPoolClient{}
	_client.MiddlewarableProvider = providers.NewMiddlewarableProvider(provider)
	return _client
}

// TxpoolStatus returns the number of transactions currently pending for inclusion in
// the next block(s), as well as the ones that are being scheduled for
// future execution only.
//
// See https://geth.ethereum.org/docs/rpc/ns-txpool#txpool_status for more details
func (c *RpcTxPoolClient) TxpoolStatus() (val *types.TxpoolStatus, err error) {
	err = c.CallContext(c.getContext(), &val, "txpool_status")
	return
}

// TxpoolInspect returns a summary of all the transactions currently pending for
// inclusion in the next block(s), as well as the ones that are being
// scheduled for future execution only.
//
// See https://geth.ethereum.org/docs/rpc/ns-txpool#txpool_inspect for more details
func (c *RpcTxPoolClient) TxpoolInspect() (val *types.TxpoolInspect, err error) {
	err = c.CallContext(c.getContext(), &val, "txpool_inspect")
	return
}

// TxpoolContentFrom retrieves the transactions contained within the txpool, returning
// pending as well as queued transactions of this address, grouped by
// nonce.
//
// See https://geth.ethereum.org/docs/rpc/ns-txpool#txpool_contentFrom for more details
func (c *RpcTxPoolClient) TxpoolContentFrom(from common.Address) (val *types.TxpoolContentFrom, err error) {
	err = c.CallContext(c.getContext(), &val, "txpool_contentFrom", from)
	return
}

// TxpoolContent returns the details of all transactions currently pending for inclusion
// in the next block(s), as well as the ones that are being scheduled
// for future execution only.
//
// See https://geth.ethereum.org/docs/rpc/ns-txpool#txpool_content for more details
func (c *RpcTxPoolClient) TxpoolContent() (val *types.TxpoolContent, err error) {
	err = c.CallContext(c.getContext(), &val, "txpool_content")
	return
}
