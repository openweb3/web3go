package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/go-rpc-provider/interfaces"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/types"
)

/*
#[derive(Default, Serialize, Deserialize, Clone)]
pub struct TxpoolStatus {
    pub pending: U64,
    pub queued: U64,
}

#[derive(Clone, Debug, Default, PartialEq, Eq, Serialize, Deserialize)]
pub struct TxpoolInspect {
    /// pending tx
    pub pending: BTreeMap<Address, BTreeMap<String, TxpoolInspectSummary>>,
    /// queued tx
    pub queued: BTreeMap<Address, BTreeMap<String, TxpoolInspectSummary>>,
}

#[derive(Clone, Copy, Debug, PartialEq, Eq, Default)]
pub struct TxpoolInspectSummary {
    /// Recipient (None when contract creation)
    pub to: Option<Address>,
    /// Transferred value
    pub value: U256,
    /// Gas amount
    pub gas: u64,
    /// Gas Price
    pub gas_price: u128,
}

/// Implement the `Deserialize` trait for `TxpoolInspectSummary` struct.
impl<'de> Deserialize<'de> for TxpoolInspectSummary {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where D: Deserializer<'de> {
        deserializer.deserialize_str(TxpoolInspectSummaryVisitor)
    }
}

/// Implement the `Serialize` trait for `TxpoolInspectSummary` struct so that
/// the format matches the one from geth.
impl Serialize for TxpoolInspectSummary {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where S: serde::Serializer {
        let formatted_to = self.to.map_or_else(
            || "contract creation".to_string(),
            |to| format!("{to:?}"),
        );
        let formatted = format!(
            "{}: {} wei + {} gas × {} wei",
            formatted_to, self.value, self.gas, self.gas_price
        );
        serializer.serialize_str(&formatted)
    }
}

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct TxpoolContentFrom<T = Transaction> {
    /// pending tx
    pub pending: BTreeMap<String, T>,
    /// queued tx
    pub queued: BTreeMap<String, T>,
}

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct TxpoolContent<T = Transaction> {
    /// pending tx
    pub pending: BTreeMap<Address, BTreeMap<String, T>>,
    /// queued tx
    pub queued: BTreeMap<Address, BTreeMap<String, T>>,
}


#[rpc(server, namespace = "txpool")]
pub trait TxPoolApi {
    /// Returns the number of transactions currently pending for inclusion in
    /// the next block(s), as well as the ones that are being scheduled for
    /// future execution only.
    ///
    /// See [here](https://geth.ethereum.org/docs/rpc/ns-txpool#txpool_status) for more details
    #[method(name = "status")]
    async fn txpool_status(&self) -> RpcResult<TxpoolStatus>;

    /// Returns a summary of all the transactions currently pending for
    /// inclusion in the next block(s), as well as the ones that are being
    /// scheduled for future execution only.
    ///
    /// See [here](https://geth.ethereum.org/docs/rpc/ns-txpool#txpool_inspect) for more details
    #[method(name = "inspect")]
    async fn txpool_inspect(&self) -> RpcResult<TxpoolInspect>;

    /// Retrieves the transactions contained within the txpool, returning
    /// pending as well as queued transactions of this address, grouped by
    /// nonce.
    ///
    /// See [here](https://geth.ethereum.org/docs/rpc/ns-txpool#txpool_contentFrom) for more details
    #[method(name = "contentFrom")]
    async fn txpool_content_from(
        &self, from: Address,
    ) -> RpcResult<TxpoolContentFrom>;

    /// Returns the details of all transactions currently pending for inclusion
    /// in the next block(s), as well as the ones that are being scheduled
    /// for future execution only.
    ///
    /// See [here](https://geth.ethereum.org/docs/rpc/ns-txpool#txpool_content) for more details
    #[method(name = "content")]
    async fn txpool_content(&self) -> RpcResult<TxpoolContent>;
}
*/

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
