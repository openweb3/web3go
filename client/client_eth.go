package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go/interfaces"
	"github.com/openweb3/web3go/types"
)

type RpcEthClient struct {
	core interfaces.RpcProvider
}

func (c *RpcEthClient) ClientVersion() (val string, err error) {
	err = c.core.Call(&val, "web3_clientVersion")
	return
}

func (c *RpcEthClient) NetVersion() (val string, err error) {
	err = c.core.Call(&val, "net_version")
	return
}

/// Returns protocol version encoded as a string (quotes are necessary).
func (c *RpcEthClient) ProtocolVersion() (val string, err error) {
	err = c.core.Call(&val, "eth_protocolVersion")
	return
}

/// Returns an object with data about the sync status or false. (wtf?)
func (c *RpcEthClient) Syncing() (val types.SyncProgress, err error) {
	err = c.core.Call(&val, "eth_syncing")
	return
}

/// Returns the number of hashes per second that the node is mining with.
func (c *RpcEthClient) Hashrate() (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_hashrate")
	return
}

/// Returns block author.
func (c *RpcEthClient) Author() (val common.Address, err error) {
	err = c.core.Call(&val, "eth_coinbase")
	return
}

/// Returns true if client is actively mining new blocks.
func (c *RpcEthClient) IsMining() (val bool, err error) {
	err = c.core.Call(&val, "eth_mining")
	return
}

/// Returns the chain ID used for transaction signing at the
/// current best block. None is returned if not
/// available.
func (c *RpcEthClient) ChainId() (val *hexutil.Uint64, err error) {
	err = c.core.Call(&val, "eth_chainId")
	return
}

/// Returns current gas_price.
func (c *RpcEthClient) GasPrice() (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_gasPrice")
	return
}

/// Returns current max_priority_fee
func (c *RpcEthClient) MaxPriorityFeePerGas() (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_maxPriorityFeePerGas")
	return
}

/// Returns accounts list.
func (c *RpcEthClient) Accounts() (val []common.Address, err error) {
	err = c.core.Call(&val, "eth_accounts")
	return
}

/// Returns highest block number.
func (c *RpcEthClient) BlockNumber() (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_blockNumber")
	return
}

/// Returns balance of the given account.
func (c *RpcEthClient) Balance(addr common.Address, block *types.BlockNumber) (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_getBalance", addr, block)
	return
}

/// Returns content of the storage at given address.
func (c *RpcEthClient) StorageAt(addr common.Address, location *hexutil.Big, block *types.BlockNumber) (val common.Hash, err error) {
	err = c.core.Call(&val, "eth_getStorageAt", addr, location, block)
	return
}

// TODO: return summary if isFull is false
/// Returns block with given hash.
func (c *RpcEthClient) BlockByHash(blockHash common.Hash, isFull bool) (val *types.Block, err error) {
	err = c.core.Call(&val, "eth_getBlockByHash", blockHash, isFull)
	return
}

/// Returns block with given number.
func (c *RpcEthClient) BlockByNumber(blockNumber types.BlockNumber, isFull bool) (val *types.Block, err error) {
	err = c.core.Call(&val, "eth_getBlockByNumber", blockNumber, isFull)
	return
}

/// Returns the number of transactions sent from given address at given time
/// (block number).
func (c *RpcEthClient) TransactionCount(addr common.Address, blockNumber *types.BlockNumber) (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_getTransactionCount", addr, blockNumber)
	return
}

/// Returns the number of transactions in a block with given hash.
func (c *RpcEthClient) BlockTransactionCountByHash(blockHash common.Hash) (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_getBlockTransactionCountByHash", blockHash)
	return
}

/// Returns the number of transactions in a block with given block number.
func (c *RpcEthClient) BlockTransactionCountByNumber(blockNumber types.BlockNumber) (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_getBlockTransactionCountByNumber", blockNumber)
	return
}

/// Returns the number of uncles in a block with given hash.
func (c *RpcEthClient) BlockUnclesCountByHash(blockHash common.Hash) (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_getUncleCountByBlockHash", blockHash)
	return
}

/// Returns the number of uncles in a block with given block number.
func (c *RpcEthClient) BlockUnclesCountByNumber(blockNumber types.BlockNumber) (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_getUncleCountByBlockNumber", blockNumber)
	return
}

/// Returns the code at given address at given time (block number).
func (c *RpcEthClient) CodeAt(addr common.Address, blockNumber *types.BlockNumber) (val hexutil.Bytes, err error) {
	err = c.core.Call(&val, "eth_getCode", addr, blockNumber)
	return
}

/// Sends signed transaction, returning its hash.
func (c *RpcEthClient) SendRawTransaction(rawTx hexutil.Bytes) (val common.Hash, err error) {
	err = c.core.Call(&val, "eth_sendRawTransaction", rawTx)
	return
}

/// @alias of eth_sendRawTransaction.
func (c *RpcEthClient) SubmitTransaction(rawTx hexutil.Bytes) (val common.Hash, err error) {
	err = c.core.Call(&val, "eth_submitTransaction", rawTx)
	return
}

/// Call contract, returning the output data.
func (c *RpcEthClient) Call(request types.TransactionArgs, blockNumber *types.BlockNumber) (val hexutil.Bytes, err error) {
	err = c.core.Call(&val, "eth_call", request, blockNumber)
	return
}

/// Estimate gas needed for execution of given contract.
func (c *RpcEthClient) EstimateGas(request types.TransactionArgs, blockNumber *types.BlockNumber) (val *hexutil.Big, err error) {
	err = c.core.Call(&val, "eth_estimateGas", request, blockNumber)
	return
}

/// Get transaction by its hash.
func (c *RpcEthClient) TransactionByHash(txHash common.Hash) (val *types.Transaction, err error) {
	err = c.core.Call(&val, "eth_getTransactionByHash", txHash)
	return
}

/// Returns transaction at given block hash and index.
func (c *RpcEthClient) TransactionByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) (val *types.Transaction, err error) {
	err = c.core.Call(&val, "eth_getTransactionByBlockHashAndIndex", blockHash, index)
	return
}

/// Returns transaction by given block number and index.
func (c *RpcEthClient) TransactionByBlockNumberAndIndex(blockNumber types.BlockNumber, index hexutil.Uint) (val *types.Transaction, err error) {
	err = c.core.Call(&val, "eth_getTransactionByBlockNumberAndIndex", blockNumber, index)
	return
}

/// Returns transaction receipt by transaction hash.
func (c *RpcEthClient) TransactionReceipt(txHash common.Hash) (val *types.Receipt, err error) {
	err = c.core.Call(&val, "eth_getTransactionReceipt", txHash)
	return
}

/// Returns an uncles at given block and index.
func (c *RpcEthClient) UncleByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) (val *types.Block, err error) {
	err = c.core.Call(&val, "eth_getUncleByBlockHashAndIndex", blockHash, index)
	return
}

/// Returns an uncles at given block and index.
func (c *RpcEthClient) UncleByBlockNumberAndIndex(blockNumber types.BlockNumber, index hexutil.Uint) (val *types.Block, err error) {
	err = c.core.Call(&val, "eth_getUncleByBlockNumberAndIndex", blockNumber, index)
	return
}

/// Returns logs matching given filter object.
func (c *RpcEthClient) Logs(filter types.EthRpcLogFilter) (val []types.Log, err error) {
	err = c.core.Call(&val, "eth_getLogs", filter)
	return
}

/// Used for submitting mining hashrate.
func (c *RpcEthClient) SubmitHashrate(rate *hexutil.Big, id common.Hash) (val bool, err error) {
	err = c.core.Call(&val, "eth_submitHashrate", rate, id)
	return
}
