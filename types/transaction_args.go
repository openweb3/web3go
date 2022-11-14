package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/openweb3/go-rpc-provider"

	"github.com/pkg/errors"
)

type ReaderForPopulate interface {
	ChainId() (val *uint64, err error)
	GasPrice() (val *big.Int, err error)
	EstimateGas(callRequest CallRequest, blockNum *BlockNumberOrHash) (val *big.Int, err error)
	TransactionCount(addr common.Address, blockNum *BlockNumberOrHash) (val *big.Int, err error)
	BlockByNumber(blockNumber BlockNumber, isFull bool) (val *Block, err error)
}

type TransactionArgs struct {
	From                 *common.Address `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice,omitempty"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas,omitempty"`
	Value                *hexutil.Big    `json:"value"`
	Nonce                *hexutil.Uint64 `json:"nonce"`

	// use pointer to keep same with go-ethereum behavior, beacuse nil *hexutil.Bytes will be marshaled to nil,
	// but nil hexutil.Bytes will be marshaled to '0x'
	Data *hexutil.Bytes `json:"data"`

	// Introduced by AccessListTxType transaction.
	AccessList *types.AccessList `json:"accessList,omitempty"`
	ChainID    *hexutil.Big      `json:"chainId,omitempty"`

	TxType *uint8 `json:"type"`
}

// data retrieves the transaction calldata. Input field is preferred.
func (args *TransactionArgs) data() []byte {
	if args.Data != nil {
		return *args.Data
	}
	return nil
}

// ToTransaction converts the arguments to a transaction.
// This assumes that setDefaults has been called.
func (args *TransactionArgs) ToTransaction() (*types.Transaction, error) {

	if args.Nonce == nil || args.Gas == nil {
		return nil, errors.New("nonce and gas are required")
	}

	al := types.AccessList{}
	if args.AccessList != nil {
		al = *args.AccessList
	}

	genDynamicFeeTx := func() types.TxData {
		return &types.DynamicFeeTx{
			To:         args.To,
			ChainID:    (*big.Int)(args.ChainID),
			Nonce:      uint64(*args.Nonce),
			Gas:        uint64(*args.Gas),
			GasFeeCap:  (*big.Int)(args.MaxFeePerGas),
			GasTipCap:  (*big.Int)(args.MaxPriorityFeePerGas),
			Value:      (*big.Int)(args.Value),
			Data:       args.data(),
			AccessList: al,
		}
	}

	genAccessListTx := func() types.TxData {
		return &types.AccessListTx{
			To:         args.To,
			ChainID:    (*big.Int)(args.ChainID),
			Nonce:      uint64(*args.Nonce),
			Gas:        uint64(*args.Gas),
			GasPrice:   (*big.Int)(args.GasPrice),
			Value:      (*big.Int)(args.Value),
			Data:       args.data(),
			AccessList: al,
		}
	}

	genLegacyTx := func() types.TxData {
		return &types.LegacyTx{
			To:       args.To,
			Nonce:    uint64(*args.Nonce),
			Gas:      uint64(*args.Gas),
			GasPrice: (*big.Int)(args.GasPrice),
			Value:    (*big.Int)(args.Value),
			Data:     args.data(),
		}
	}

	switch *args.TxType {
	case types.LegacyTxType:
		return types.NewTx(genLegacyTx()), nil
	case types.AccessListTxType:
		return types.NewTx(genAccessListTx()), nil
	case types.DynamicFeeTxType:
		return types.NewTx(genDynamicFeeTx()), nil
	}

	return nil, errors.New("unknown transaction type")
}

func (args *TransactionArgs) Populate(reader ReaderForPopulate) error {

	if args.From == nil {
		return errors.New("from is required")
	}

	if err := args.populateGasPrice(reader); err != nil {
		return errors.Wrap(err, "failed to populate gas price")
	}

	if args.Value == nil {
		args.Value = new(hexutil.Big)
	}

	// try get pending nonce first, if failed, try get nonce of latest block
	if args.Nonce == nil {
		pending := BlockNumberOrHashWithNumber(PendingBlockNumber)
		nonce, err := reader.TransactionCount(*args.From, &pending)
		if err != nil {
			nonce, err = reader.TransactionCount(*args.From, nil)
			if err != nil {
				return errors.Wrap(err, "failed to get nonce of both pending and latest block")
			}
		}
		temp := nonce.Uint64()
		args.Nonce = (*hexutil.Uint64)(&temp)
	}

	if args.To == nil && len(args.data()) == 0 {
		return errors.New(`contract creation without any data provided`)
	}
	// Estimate the gas usage if necessary.
	if args.Gas == nil {
		// These fields are immutable during the estimation, safe to
		// pass the pointer directly.
		data := args.data()
		callArgs := CallRequest{
			From:                 args.From,
			To:                   args.To,
			GasPrice:             (*big.Int)(args.GasPrice),
			MaxFeePerGas:         (*big.Int)(args.MaxFeePerGas),
			MaxPriorityFeePerGas: (*big.Int)(args.MaxPriorityFeePerGas),
			Value:                (*big.Int)(args.Value),
			Data:                 data,
			AccessList:           args.AccessList,
		}

		latest := BlockNumberOrHashWithNumber(rpc.LatestBlockNumber)
		estimated, err := reader.EstimateGas(callArgs, &latest)
		if err != nil {
			return errors.Wrap(err, "failed to estimate")
		}

		temp := estimated.Uint64()
		args.Gas = (*hexutil.Uint64)(&temp)
	}
	if args.ChainID == nil {
		id, err := reader.ChainId()
		if err != nil {
			return errors.Wrap(err, "failed to get chaind")
		}
		temp := big.NewInt(0).SetUint64(*id)
		args.ChainID = (*hexutil.Big)(temp)
	}
	return nil
}

func (args *TransactionArgs) populateGasPrice(reader ReaderForPopulate) error {
	if args.GasPrice != nil && (args.MaxFeePerGas != nil || args.MaxPriorityFeePerGas != nil) {
		return errors.New("both gasPrice and (maxFeePerGas or maxPriorityFeePerGas) specified")
	}

	has1559 := args.MaxFeePerGas != nil || args.MaxPriorityFeePerGas != nil

	nullType := uint8(255)
	argsTxType := nullType
	if args.TxType != nil {
		argsTxType = uint8(*args.TxType)
	}

	gasFeeData, err := getFeeData(reader)
	if err != nil {
		return errors.Wrap(err, "failed to get fee data")
	}

	// set the txtype according to feeData
	// - if support1559, then set txtype to 2
	// - if not support1559
	// - - if has maxFeePerGas or maxPriorityFeePerGas, then return error
	// - - if contains accesslist, set txtype to 1
	// - - else set txtype to 0
	if argsTxType == nullType {
		if gasFeeData.isSupport1559() {
			argsTxType = types.DynamicFeeTxType
		} else {
			if has1559 {
				return errors.New("not support 1559 but (maxFeePerGas or maxPriorityFeePerGas) specified")
			}

			if args.AccessList == nil {
				argsTxType = types.LegacyTxType
			} else {
				argsTxType = types.AccessListTxType
			}
		}
	}
	args.TxType = &argsTxType

	// if txtype is DynamicFeeTxType that means support 1559, so if gasPrice is not nil, set max... to gasPrice
	if *args.TxType == types.DynamicFeeTxType {
		if args.GasPrice != nil {
			args.MaxFeePerGas = args.GasPrice
			args.MaxPriorityFeePerGas = args.GasPrice
			args.GasPrice = nil
			return nil
		}

		if args.MaxPriorityFeePerGas == nil {
			args.MaxPriorityFeePerGas = (*hexutil.Big)(gasFeeData.maxPriorityFeePerGas)
		}
		if args.MaxFeePerGas == nil {
			args.MaxFeePerGas = (*hexutil.Big)(gasFeeData.maxFeePerGas)
		}
		if args.MaxFeePerGas.ToInt().Cmp(args.MaxPriorityFeePerGas.ToInt()) < 0 {
			return fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", args.MaxFeePerGas, args.MaxPriorityFeePerGas)
		}
		return nil
	}

	if args.GasPrice != nil {
		return nil
	}

	args.GasPrice = (*hexutil.Big)(gasFeeData.gasPrice)
	return nil
}

// ToTransaction converts the arguments to a transaction.
// This assumes that Populate has been called.
func (args *TransactionArgs) PopulateAndToTransaction(reader ReaderForPopulate) (*types.Transaction, error) {
	if err := args.Populate(reader); err != nil {
		return nil, err
	}
	return args.ToTransaction()
}

func ConvertTransactionToArgs(from common.Address, tx Transaction) *TransactionArgs {
	args := &TransactionArgs{}

	txType := tx.Type()
	args.TxType = &txType

	args.From = &from
	args.To = tx.To()

	nonce := tx.Nonce()
	if nonce != 0 {
		args.Nonce = (*hexutil.Uint64)(&nonce)
	}

	gas := tx.Gas()
	if gas != 0 {
		args.Gas = (*hexutil.Uint64)(&gas)
	}

	args.Value = (*hexutil.Big)(tx.Value())

	data := tx.Data()
	args.Data = (*hexutil.Bytes)(&data)

	switch tx.Type() {
	case types.LegacyTxType:
		if tx.GasPrice().Cmp(big.NewInt(0)) != 0 {
			args.GasPrice = (*hexutil.Big)(tx.GasPrice())
		}
	case types.DynamicFeeTxType:
		if tx.GasTipCap().Cmp(big.NewInt(0)) != 0 {
			args.MaxPriorityFeePerGas = (*hexutil.Big)(tx.GasTipCap())
		}
		if tx.GasFeeCap().Cmp(big.NewInt(0)) != 0 {
			args.MaxFeePerGas = (*hexutil.Big)(tx.GasFeeCap())
		}
		fallthrough
	case types.AccessListTxType:
		al := tx.AccessList()
		args.AccessList = &al
		if tx.ChainId().Cmp(big.NewInt(0)) != 0 {
			args.ChainID = (*hexutil.Big)(tx.ChainId())
		}
	}
	return args
}
