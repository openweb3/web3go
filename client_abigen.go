package web3go

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/web3go/types"
)

// ClientForContract is used for abi-binding struct generated by abigen.
// abigen is a source code generator to convert Ethereum contract definitions into easy to use, compile-time type-safe Go packages.
// Please see https://geth.ethereum.org/docs/dapp/native-bindings page for details
type ClientForContract struct {
	raw *Client
}

func NewClientForContract(raw *Client) *ClientForContract {
	return &ClientForContract{
		raw: raw,
	}
}

func (c *ClientForContract) TransactionReceipt(ctx context.Context, txHash common.Hash) (*ethtypes.Receipt, error) {
	r, err := c.raw.Eth.TransactionReceipt(txHash)
	if err != nil {
		return nil, err
	}
	ethReceipt := toEthReceipt(*r)
	return &ethReceipt, nil
}

func (c *ClientForContract) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	bnInt64 := types.BlockNumber(blockNumber.Int64())
	bnOrHash := types.BlockNumberOrHashWithNumber(bnInt64)
	return c.raw.Eth.CodeAt(account, &bnOrHash)
}

func (c *ClientForContract) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	cr := convertCallMsg2CallRequest(call)
	bn := types.BlockNumberOrHashWithNumber(types.NewBlockNumber(blockNumber.Int64()))
	return c.raw.Eth.Call(cr, &bn)
}

// PendingCallContract executes an Ethereum contract call against the pending state.
func (c *ClientForContract) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	cr := convertCallMsg2CallRequest(call)
	pending := types.BlockNumberOrHashWithNumber(types.PendingBlockNumber)
	return c.raw.Eth.Call(cr, &pending)
}

// HeaderByNumber returns a block header from the current canonical chain. If
// number is nil, the latest known header is returned.
func (c *ClientForContract) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	b, err := c.raw.Eth.BlockByNumber(types.NewBlockNumber(number.Int64()), false)
	if err != nil {
		return nil, err
	}
	h, err := b.Header()
	if err != nil {
		return nil, err
	}
	return h, nil
}

// PendingCodeAt returns the code of the given account in the pending state.
func (c *ClientForContract) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	pending := types.BlockNumberOrHashWithNumber(types.PendingBlockNumber)
	return c.raw.Eth.CodeAt(account, &pending)
}

// PendingNonceAt retrieves the current pending nonce associated with an account.
func (c *ClientForContract) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	pending := types.BlockNumberOrHashWithNumber(types.PendingBlockNumber)
	nonce, err := c.raw.Eth.TransactionCount(account, &pending)
	if err != nil {
		return 0, err
	}
	return nonce.Uint64(), nil
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (c *ClientForContract) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.raw.Eth.GasPrice()
}

// SuggestGasTipCap retrieves the currently suggested 1559 priority fee to allow
// a timely execution of a transaction.
func (c *ClientForContract) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return c.raw.Eth.MaxPriorityFeePerGas()
}

// EstimateGas tries to estimate the gas needed to execute a specific
// transaction based on the current pending state of the backend blockchain.
// There is no guarantee that this is the true gas limit requirement as other
// transactions may be added or removed by miners, but it should provide a basis
// for setting a reasonable default.
func (c *ClientForContract) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	cr := convertCallMsg2CallRequest(call)

	pending := types.BlockNumberOrHashWithNumber(types.PendingBlockNumber)
	val, err := c.raw.Eth.EstimateGas(cr, &pending)
	if err != nil {
		return 0, err
	}
	return val.Uint64(), nil
}

// SendTransaction injects the transaction into the pending pool for execution.
func (c *ClientForContract) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	// if not signed, sign it by first account
	if v, _, _ := tx.RawSignatureValues(); v == nil {
		sm, err := c.raw.GetSignerManager()
		if err != nil {
			return err
		}

		if len(sm.List()) == 0 {
			return errors.New("no signer available")
		}

		account := sm.List()[0].Address()
		_, err = c.raw.Eth.SendTransaction(account, *tx)
		return err
	}

	// otherwise, send raw transaction
	rawTx, err := tx.MarshalBinary()
	if err != nil {
		return err
	}

	_, err = c.raw.Eth.SendRawTransaction(rawTx)
	return err
}

// FilterLogs executes a log filter operation, blocking during execution and
// returning all the results in one batch.
//
// TODO(karalabe): Deprecate when the subscription one can return past data too.
func (c *ClientForContract) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]ethtypes.Log, error) {
	q := convertFilterQuery(query)
	logs, err := c.raw.Eth.Logs(q)
	if err != nil {
		return nil, err
	}
	return toEthLogs(logs), nil
}

// SubscribeFilterLogs creates a background log filtering operation, returning
// a subscription immediately, which can be used to stream the found events.
func (c *ClientForContract) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- ethtypes.Log) (ethereum.Subscription, error) {
	q := convertFilterQuery(query)
	return c.raw.Subscribe(context.Background(), "eth", ch, "logs", q)
}

func convertFilterQuery(query ethereum.FilterQuery) types.FilterQuery {
	fromBlock := types.NewBlockNumber(query.FromBlock.Int64())
	toBlock := types.NewBlockNumber(query.ToBlock.Int64())
	return types.FilterQuery{
		BlockHash: query.BlockHash,
		FromBlock: &fromBlock,
		ToBlock:   &toBlock,
		Addresses: query.Addresses,
		Topics:    query.Topics,
	}
}

func convertCallMsg2CallRequest(call ethereum.CallMsg) types.CallRequest {
	return types.CallRequest{
		From:                 &call.From,
		To:                   call.To,
		Gas:                  &call.Gas,
		GasPrice:             call.GasPrice,
		MaxFeePerGas:         call.GasFeeCap,
		MaxPriorityFeePerGas: call.GasTipCap,
		Value:                call.Value,
		Data:                 call.Data,
		Input:                call.Data,
		AccessList:           &call.AccessList,
	}
}

func toEthReceipt(r types.Receipt) ethtypes.Receipt {
	eReceipt := ethtypes.Receipt{
		PostState:         r.Root,
		CumulativeGasUsed: r.CumulativeGasUsed,
		Bloom:             r.LogsBloom,
		Logs:              toEthLogPtrs(r.Logs),
		TxHash:            r.TransactionHash,
		GasUsed:           r.GasUsed,
		BlockHash:         r.BlockHash,
		BlockNumber:       new(big.Int).SetUint64(r.BlockNumber),
		TransactionIndex:  uint(r.TransactionIndex),
	}
	if r.Type != nil {
		eReceipt.Type = uint8(*r.Type)
	}

	if r.Status != nil {
		eReceipt.Status = *r.Status
	}

	if r.ContractAddress != nil {
		eReceipt.ContractAddress = *r.ContractAddress
	}
	return eReceipt
}

func toEthLogs(logs []types.Log) []ethtypes.Log {
	eLogs := make([]ethtypes.Log, len(logs))
	for i, l := range logs {
		eLogs[i] = toEthLog(l)
	}
	return eLogs
}

func toEthLogPtrs(logs []*types.Log) []*ethtypes.Log {
	eLogs := make([]*ethtypes.Log, len(logs))
	for i, l := range logs {
		if l == nil {
			eLogs[i] = nil
			continue
		}
		el := toEthLog(*l)
		eLogs[i] = &el
	}
	return eLogs
}

func toEthLog(log types.Log) ethtypes.Log {
	return ethtypes.Log{
		Address:     log.Address,
		Topics:      log.Topics,
		Data:        log.Data,
		BlockNumber: log.BlockNumber,
		TxHash:      log.TxHash,
		TxIndex:     log.TxIndex,
		BlockHash:   log.BlockHash,
		Index:       log.Index,
		Removed:     log.Removed,
	}
}
