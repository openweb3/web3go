package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"

	"github.com/openweb3/go-rpc-provider/interfaces"
	"github.com/openweb3/web3go/types"
)

type RpcEthClient struct {
	*providers.MiddlewarableProvider
}

func NewRpcEthClient(provider interfaces.Provider) *RpcEthClient {
	return &RpcEthClient{
		MiddlewarableProvider: providers.NewMiddlewarableProvider(provider),
	}
}

func (c *RpcEthClient) ClientVersion() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "web3_clientVersion")
	return
}

func (c *RpcEthClient) NetVersion() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "net_version")
	return
}

/// Returns protocol version encoded as a string (quotes are necessary).
func (c *RpcEthClient) ProtocolVersion() (val string, err error) {
	err = c.CallContext(context.Background(), &val, "eth_protocolVersion")
	return
}

/// Returns an object with data about the sync status or false. (wtf?)
func (c *RpcEthClient) Syncing() (val types.SyncStatus, err error) {
	err = c.CallContext(context.Background(), &val, "eth_syncing")
	return
}

/// Returns the number of hashes per second that the node is mining with.
func (c *RpcEthClient) Hashrate() (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_hashrate")
	val = (*big.Int)(_val)
	return
}

/// Returns block author.
func (c *RpcEthClient) Author() (val common.Address, err error) {
	err = c.CallContext(context.Background(), &val, "eth_coinbase")
	return
}

/// Returns true if client is actively mining new blocks.
func (c *RpcEthClient) IsMining() (val bool, err error) {
	err = c.CallContext(context.Background(), &val, "eth_mining")
	return
}

/// Returns the chain ID used for transaction signing at the
/// current best block. None is returned if not
/// available.
func (c *RpcEthClient) ChainId() (val *uint64, err error) {
	var _val *hexutil.Uint64
	err = c.CallContext(context.Background(), &_val, "eth_chainId")
	val = (*uint64)(_val)
	return
}

/// Returns current gas_price.
func (c *RpcEthClient) GasPrice() (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_gasPrice")
	val = (*big.Int)(_val)
	return
}

/// Returns current max_priority_fee
func (c *RpcEthClient) MaxPriorityFeePerGas() (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_maxPriorityFeePerGas")
	val = (*big.Int)(_val)
	return
}

/// Returns accounts list.
func (c *RpcEthClient) Accounts() (val []common.Address, err error) {
	err = c.CallContext(context.Background(), &val, "eth_accounts")
	return
}

/// Returns highest block number.
func (c *RpcEthClient) BlockNumber() (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_blockNumber")
	val = (*big.Int)(_val)
	return
}

/// Returns balance of the given account.
func (c *RpcEthClient) Balance(addr common.Address, block *types.BlockNumberOrHash) (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_getBalance", addr, getRealBlockNumberOrHash(block))
	val = (*big.Int)(_val)
	return
}

/// Returns content of the storage at given address.
func (c *RpcEthClient) StorageAt(addr common.Address, location *big.Int, block *types.BlockNumberOrHash) (val common.Hash, err error) {
	_location := (*hexutil.Big)(location)
	err = c.CallContext(context.Background(), &val, "eth_getStorageAt", addr, _location, getRealBlockNumberOrHash(block))
	return
}

/// Returns block with given hash.
func (c *RpcEthClient) BlockByHash(blockHash common.Hash, isFull bool) (val *types.Block, err error) {
	block := &types.Block{}
	block.Transactions = *types.NewTxOrHashList(isFull)
	err = c.CallContext(context.Background(), &block, "eth_getBlockByHash", blockHash, isFull)
	return block, err
}

/// Returns block with given number.
func (c *RpcEthClient) BlockByNumber(blockNumber types.BlockNumber, isFull bool) (val *types.Block, err error) {
	block := &types.Block{}
	block.Transactions = *types.NewTxOrHashList(isFull)
	err = c.CallContext(context.Background(), &block, "eth_getBlockByNumber", blockNumber, isFull)
	return block, err
}

/// Returns the number of transactions sent from given address at given time
/// (block number).
// TODO: nil *types.BlockNumberOrHash will be marshaled to null, which is not allowed in geth, but could work in ganache and not treat as latest, what behavior in conflux-rust should be investigate
func (c *RpcEthClient) TransactionCount(addr common.Address, blockNum *types.BlockNumberOrHash) (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_getTransactionCount", addr, getRealBlockNumberOrHash(blockNum))
	val = (*big.Int)(_val)
	return
}

/// Returns the number of transactions in a block with given hash.
func (c *RpcEthClient) BlockTransactionCountByHash(blockHash common.Hash) (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_getBlockTransactionCountByHash", blockHash)
	val = (*big.Int)(_val)
	return
}

/// Returns the number of transactions in a block with given block number.
func (c *RpcEthClient) BlockTransactionCountByNumber(blockNum types.BlockNumber) (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_getBlockTransactionCountByNumber", blockNum)
	val = (*big.Int)(_val)
	return
}

/// Returns the number of uncles in a block with given hash.
func (c *RpcEthClient) BlockUnclesCountByHash(blockHash common.Hash) (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_getUncleCountByBlockHash", blockHash)
	val = (*big.Int)(_val)
	return
}

/// Returns the number of uncles in a block with given block number.
func (c *RpcEthClient) BlockUnclesCountByNumber(blockNum types.BlockNumber) (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_getUncleCountByBlockNumber", blockNum)
	val = (*big.Int)(_val)
	return
}

/// Returns the code at given address at given time (block number).
func (c *RpcEthClient) CodeAt(addr common.Address, blockNum *types.BlockNumberOrHash) (val []byte, err error) {
	var _val hexutil.Bytes
	err = c.CallContext(context.Background(), &_val, "eth_getCode", addr, getRealBlockNumberOrHash(blockNum))
	val = ([]byte)(_val)
	return
}

// SendTransaction sends a transaction by TransactionArgs which contains all fields a transaction needs,
// it will populate empty fields automatically before sending.
func (c *RpcEthClient) SendTransactionByArgs(args types.TransactionArgs) (txHash common.Hash, err error) {
	if err = args.Populate(c); err != nil {
		return
	}
	err = c.CallContext(context.Background(), &txHash, "eth_sendTransaction", args)
	return
}

func (c *RpcEthClient) SendTransaction(from common.Address, tx types.Transaction) (txHash common.Hash, err error) {
	txArgs := types.ConvertTransactionToArgs(from, tx)
	return c.SendTransactionByArgs(*txArgs)
}

/// Sends signed transaction, returning its hash.
func (c *RpcEthClient) SendRawTransaction(rawTx []byte) (val common.Hash, err error) {
	_rawTx := (hexutil.Bytes)(rawTx)
	err = c.CallContext(context.Background(), &val, "eth_sendRawTransaction", _rawTx)
	return
}

/// @alias of `eth_sendRawTransaction`.
func (c *RpcEthClient) SubmitTransaction(rawTx []byte) (val common.Hash, err error) {
	_rawTx := (hexutil.Bytes)(rawTx)
	err = c.CallContext(context.Background(), &val, "eth_submitTransaction", _rawTx)
	return
}

/// Call contract, returning the output data.
func (c *RpcEthClient) Call(callRequest types.CallRequest, blockNum *types.BlockNumberOrHash) (val []byte, err error) {
	var _val hexutil.Bytes
	err = c.CallContext(context.Background(), &_val, "eth_call", callRequest, getRealBlockNumberOrHash(blockNum))
	val = ([]byte)(_val)
	return
}

/// Estimate gas needed for execution of given contract.
func (c *RpcEthClient) EstimateGas(callRequest types.CallRequest, blockNum *types.BlockNumberOrHash) (val *big.Int, err error) {
	var _val *hexutil.Big
	err = c.CallContext(context.Background(), &_val, "eth_estimateGas", callRequest, getRealBlockNumberOrHash(blockNum))
	val = (*big.Int)(_val)
	return
}

/// Get transaction by its hash.
func (c *RpcEthClient) TransactionByHash(txHash common.Hash) (val *types.TransactionDetail, err error) {
	err = c.CallContext(context.Background(), &val, "eth_getTransactionByHash", txHash)
	return
}

/// Returns transaction at given block hash and index.
func (c *RpcEthClient) TransactionByBlockHashAndIndex(blockHash common.Hash, index uint) (val *types.TransactionDetail, err error) {
	err = c.CallContext(context.Background(), &val, "eth_getTransactionByBlockHashAndIndex", blockHash, index)
	return
}

/// Returns transaction by given block number and index.
func (c *RpcEthClient) TransactionByBlockNumberAndIndex(blockNum types.BlockNumber, index uint) (val *types.TransactionDetail, err error) {
	err = c.CallContext(context.Background(), &val, "eth_getTransactionByBlockNumberAndIndex", blockNum, index)
	return
}

/// Returns transaction receipt by transaction hash.
func (c *RpcEthClient) TransactionReceipt(txHash common.Hash) (val *types.Receipt, err error) {
	err = c.CallContext(context.Background(), &val, "eth_getTransactionReceipt", txHash)
	return
}

/// Returns an uncles at given block and index.
func (c *RpcEthClient) UncleByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) (val *types.Block, err error) {
	err = c.CallContext(context.Background(), &val, "eth_getUncleByBlockHashAndIndex", blockHash, index)
	return
}

/// Returns an uncles at given block and index.
func (c *RpcEthClient) UncleByBlockNumberAndIndex(blockNum types.BlockNumber, index uint) (val *types.Block, err error) {
	err = c.CallContext(context.Background(), &val, "eth_getUncleByBlockNumberAndIndex", blockNum, index)
	return
}

/// Returns logs matching given filter object.
func (c *RpcEthClient) Logs(logFilter types.FilterQuery) (val []types.Log, err error) {
	err = c.CallContext(context.Background(), &val, "eth_getLogs", logFilter)
	return
}

/// Used for submitting mining hashrate.
func (c *RpcEthClient) SubmitHashrate(rate *big.Int, id common.Hash) (val bool, err error) {
	_rate := (*hexutil.Big)(rate)
	err = c.CallContext(context.Background(), &val, "eth_submitHashrate", _rate, id)
	return
}

// SubscribeNewHead subscribes to notifications about the current blockchain head
// on the given channel.
func (c *RpcEthClient) SubscribeNewHead(ch chan<- *types.Header) (types.Subscription, error) {
	return c.Subscribe(context.Background(), "eth", ch, "newHeads")
}

// SubscribeNewHead subscribes to notifications about the current blockchain head
// on the given channel.
func (c *RpcEthClient) SubscribeFilterLogs(q types.FilterQuery, ch chan<- types.Log) (types.Subscription, error) {
	return c.Subscribe(context.Background(), "eth", ch, "logs", q)
}
