package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/go-rpc-provider/interfaces"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/types"
)

type RpcParityClient struct {
	*providers.MiddlewarableProvider
}

func NewRpcParityClient(provider interfaces.Provider) *RpcParityClient {
	return &RpcParityClient{
		MiddlewarableProvider: providers.NewMiddlewarableProvider(provider),
	}
}

/// Returns current transactions limit.
func (c *RpcParityClient) TransactionsLimit() (val uint, err error) {
	err = c.CallContext(context.Background(), &val, "parity_transactionsLimit")
	return
}

/// Returns mining extra data.
func (c *RpcParityClient) ExtraData() (val []byte, err error) {
	var _val hexutil.Bytes
	err = c.CallContext(context.Background(), &_val, "parity_extraData")
	val = ([]byte)(_val)
	return
}

/// Returns mining gas floor target.
func (c *RpcParityClient) GasFloorTarget() (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "parity_gasFloorTarget")
	val = (*big.Int)(_val)
	return
}

/// Returns mining gas floor cap.
func (c *RpcParityClient) GasCeilTarget() (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "parity_gasCeilTarget")
	val = (*big.Int)(_val)
	return
}

/// Returns minimal gas price for transaction to be included in queue.
func (c *RpcParityClient) MinGasPrice() (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "parity_minGasPrice")
	val = (*big.Int)(_val)
	return
}

/// Returns latest logs
func (c *RpcParityClient) DevLogs() (val []string, err error) {
	err = c.CallContext(context.Background(), &val, "parity_devLogs")
	return
}

/// Returns logs levels
func (c *RpcParityClient) DevLogsLevels() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "parity_devLogsLevels")
	return
}

/// Returns chain name - DEPRECATED. Use `parity_chainName` instead.
func (c *RpcParityClient) NetChain() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "parity_netChain")
	return
}

/// Returns peers details
func (c *RpcParityClient) NetPeers() (val types.Peers, err error) {
	err = c.CallContext(context.Background(), &val, "parity_netPeers")
	return
}

/// Returns network port
func (c *RpcParityClient) NetPort() (val uint16, err error) {
	err = c.CallContext(context.Background(), &val, "parity_netPort")
	return
}

/// Returns rpc settings
func (c *RpcParityClient) RpcSettings() (val types.RpcSettings, err error) {
	err = c.CallContext(context.Background(), &val, "parity_rpcSettings")
	return
}

/// Returns node name
func (c *RpcParityClient) NodeName() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "parity_nodeName")
	return
}

/// Returns default extra data
func (c *RpcParityClient) DefaultExtraData() (val []byte, err error) {
	var _val hexutil.Bytes
	err = c.CallContext(context.Background(), &_val, "parity_defaultExtraData")
	val = ([]byte)(_val)
	return
}

/// Returns distribution of gas price in latest blocks.
func (c *RpcParityClient) GasPriceHistogram() (val types.Histogram, err error) {
	err = c.CallContext(context.Background(), &val, "parity_gasPriceHistogram")
	return
}

/// Returns number of unsigned transactions waiting in the signer queue (if signer enabled)
/// Returns error when signer is disabled
func (c *RpcParityClient) UnsignedTransactionsCount() (val uint, err error) {
	err = c.CallContext(context.Background(), &val, "parity_unsignedTransactionsCount")
	return
}

/// Returns a cryptographically random phrase sufficient for securely seeding a secret key.
func (c *RpcParityClient) GenerateSecretPhrase() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "parity_generateSecretPhrase")
	return
}

/// Returns whatever address would be derived from the given phrase if it were to seed a brainwallet.
func (c *RpcParityClient) PhraseToAddress(phrase string) (val common.Address, err error) {
	err = c.CallContext(context.Background(), &val, "parity_phraseToAddress", phrase)
	return
}

/// Returns the value of the registrar for this network.
func (c *RpcParityClient) RegistryAddress() (val *common.Address, err error) {
	err = c.CallContext(context.Background(), &val, "parity_registryAddress")
	return
}

/// Returns all addresses if Fat DB is enabled (`--fat-db`), or null if not.
func (c *RpcParityClient) ListAccounts(count uint64, after *common.Address, blockNumber *types.BlockNumberOrHash) (val []common.Address, err error) {
	err = c.CallContext(context.Background(), &val, "parity_listAccounts", count, after, getRealBlockNumberOrHash(blockNumber))
	return
}

/// Returns all storage keys of the given address (first parameter) if Fat DB is enabled (`--fat-db`),
/// or null if not.
func (c *RpcParityClient) ListStorageKeys(address common.Address, count uint64, after *common.Hash, blockNumber *types.BlockNumberOrHash) (val []common.Hash, err error) {
	err = c.CallContext(context.Background(), &val, "parity_listStorageKeys", address, count, after, getRealBlockNumberOrHash(blockNumber))
	return
}

/// Encrypt some data with a public key under ECIES.
/// First parameter is the 512-byte destination public key, second is the message.
/// FIXME: key shoule be H512 public key
func (c *RpcParityClient) EncryptMessage(key string, phrase []byte) (val []byte, err error) {
	_phrase := (hexutil.Bytes)(phrase)
	var _val hexutil.Bytes
	err = c.CallContext(context.Background(), &_val, "parity_encryptMessage", key, _phrase)
	val = ([]byte)(_val)
	return
}

/// Returns all pending transactions from transaction queue.
func (c *RpcParityClient) PendingTransactions(limit *uint, filter *types.TransactionFilter) (val []types.TransactionDetail, err error) {
	err = c.CallContext(context.Background(), &val, "parity_pendingTransactions", limit, filter)
	return
}

/// Returns all transactions from transaction queue.
///
/// Some of them might not be ready to be included in a block yet.
func (c *RpcParityClient) AllTransactions() (val []types.TransactionDetail, err error) {
	err = c.CallContext(context.Background(), &val, "parity_allTransactions")
	return
}

/// Same as parity_allTransactions, but return only transactions hashes.
func (c *RpcParityClient) AllTransactionHashes() (val []common.Hash, err error) {
	err = c.CallContext(context.Background(), &val, "parity_allTransactionHashes")
	return
}

/// Returns all future transactions from transaction queue (deprecated)
func (c *RpcParityClient) FutureTransactions() (val []types.TransactionDetail, err error) {
	err = c.CallContext(context.Background(), &val, "parity_futureTransactions")
	return
}

/// Returns propagation statistics on transactions pending in the queue.
func (c *RpcParityClient) PendingTransactionsStats() (val map[common.Hash](types.TransactionStats), err error) {
	err = c.CallContext(context.Background(), &val, "parity_pendingTransactionsStats")
	return
}

/// Returns a list of current and past local transactions with status details.
func (c *RpcParityClient) LocalTransactions() (val map[common.Hash]types.LocalTransactionStatus, err error) {
	err = c.CallContext(context.Background(), &val, "parity_localTransactions")
	return
}

/// Returns current WS Server interface and port or an error if ws server is disabled.
func (c *RpcParityClient) WsUrl() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "parity_wsUrl")
	return
}

/// Returns next nonce for particular sender. Should include all transactions in the queue.
func (c *RpcParityClient) NextNonce(address common.Address) (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "parity_nextNonce", address)
	val = (*big.Int)(_val)
	return
}

/// Get the mode. Returns one of: "active", "passive", "dark", "offline".
func (c *RpcParityClient) Mode() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "parity_mode")
	return
}

/// Get the chain name. Returns one of the pre-configured chain names or a filename.
func (c *RpcParityClient) Chain() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "parity_chain")
	return
}

/// Get the enode of this node.
func (c *RpcParityClient) Enode() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "parity_enode")
	return
}

/// Get the current chain status.
func (c *RpcParityClient) ChainStatus() (val types.ChainStatus, err error) {
	err = c.CallContext(context.Background(), &val, "parity_chainStatus")
	return
}

/// Get node kind info.
func (c *RpcParityClient) NodeKind() (val types.NodeKind, err error) {
	err = c.CallContext(context.Background(), &val, "parity_nodeKind")
	return
}

/// Get block header.
/// Same as `eth_getBlockByNumber` but without uncles and transactions.
func (c *RpcParityClient) BlockHeader(blockNum *types.BlockNumberOrHash) (val types.RichHeader, err error) {
	err = c.CallContext(context.Background(), &val, "parity_getBlockHeaderByNumber", getRealBlockNumberOrHash(blockNum))
	return
}

/// Get block receipts.
/// Allows you to fetch receipts from the entire block at once.
/// If no parameter is provided defaults to `latest`.
func (c *RpcParityClient) BlockReceipts(blockNum *types.BlockNumberOrHash) (val []types.Receipt, err error) {
	err = c.CallContext(context.Background(), &val, "parity_getBlockReceipts", getRealBlockNumberOrHash(blockNum))
	return
}

/// Call contract, returning the output data.
func (c *RpcParityClient) Call(requests []types.CallRequest, blockNum *types.BlockNumberOrHash) (val [][]byte, err error) {
	var _val []hexutil.Bytes
	err = c.CallContext(context.Background(), &_val, "parity_call", requests, getRealBlockNumberOrHash(blockNum))

	for _, _valItem := range _val {
		val = append(val, []byte(_valItem))
	}
	return
}

/// Used for submitting a proof-of-work solution (similar to `eth_submitWork`,
/// but returns block hash on success, and returns an explicit error message on failure).
// FIXME: nonce shoule be H64 hash
func (c *RpcParityClient) SubmitWorkDetail(nonce string, powHash common.Hash, mixHash common.Hash) (val common.Hash, err error) {
	err = c.CallContext(context.Background(), &val, "parity_submitWorkDetail", nonce, powHash, mixHash)
	return
}

/// Returns the status of the node. Used as the health endpoint.
///
/// The RPC returns successful response if:
/// - The node have a peer (unless running a dev chain)
/// - The node is not syncing.
///
/// Otherwise the RPC returns error.
func (c *RpcParityClient) Status() (err error) {
	err = c.CallContext(context.Background(), &[]interface{}{}, "parity_nodeStatus")
	return
}

/// Extracts Address and public key from signature using the r, s and v params. Equivalent to Solidity erecover
/// as well as checks the signature for chain replay protection
func (c *RpcParityClient) VerifySignature(isPrefixed bool, message []byte, r common.Hash, s common.Hash, v uint64) (val types.RecoveredAccount, err error) {
	_message := (hexutil.Bytes)(message)
	_v := (hexutil.Uint64)(v)
	err = c.CallContext(context.Background(), &val, "parity_verifySignature", isPrefixed, _message, r, s, _v)
	return
}
