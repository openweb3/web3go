package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethrpctypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

type mockPopulateReader struct{}

func (m *mockPopulateReader) ChainId() (val *uint64, err error) {
	chainId := uint64(0x12) //18
	return &chainId, nil
}

func (m *mockPopulateReader) GasPrice() (val *big.Int, err error) {
	return big.NewInt(0x1c), nil //28
}

func (m *mockPopulateReader) MaxPriorityFeePerGas() (val *big.Int, err error) {
	return big.NewInt(0x1e), nil //30
}

func (m *mockPopulateReader) EstimateGas(callRequest CallRequest, blockNum *BlockNumberOrHash, overrides *StateOverride, blockOverrides *BlockOverrides) (val *big.Int, err error) {
	return big.NewInt(0x26), nil //38
}

func (m *mockPopulateReader) TransactionCount(addr common.Address, blockNum *BlockNumberOrHash) (val *big.Int, err error) {
	return big.NewInt(0x30), nil //48
}

func (m *mockPopulateReader) BlockByNumber(num BlockNumber, isFull bool) (val *Block, err error) {
	return &Block{BaseFeePerGas: big.NewInt(0x3a)}, nil //58
}

type mockPopulateReaderNo1559 struct {
	mockPopulateReader
}

func (m *mockPopulateReaderNo1559) BlockByNumber(num BlockNumber, isFull bool) (val *Block, err error) {
	return &Block{BaseFeePerGas: nil}, nil
}

func TestPopulateNon1559SetCodeAndDynamic(t *testing.T) {
	reader := &mockPopulateReaderNo1559{}

	t.Run("setcode inferred without explicit fee returns error", func(t *testing.T) {
		args := &TransactionArgs{
			From:              &common.Address{},
			To:                &common.Address{},
			AuthorizationList: []ethrpctypes.SetCodeAuthorization{{ChainID: *uint256.NewInt(1), Address: common.Address{}}},
		}

		err := args.Populate(reader)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "setCode transaction requires explicit maxFeePerGas and maxPriorityFeePerGas on non-1559 chain")
	})

	t.Run("setcode inferred with explicit fee succeeds", func(t *testing.T) {
		args := &TransactionArgs{
			From:                 &common.Address{},
			To:                   &common.Address{},
			MaxFeePerGas:         (*hexutil.Big)(big.NewInt(44)),
			MaxPriorityFeePerGas: (*hexutil.Big)(big.NewInt(22)),
			AuthorizationList:    []ethrpctypes.SetCodeAuthorization{{ChainID: *uint256.NewInt(1), Address: common.Address{}}},
		}

		err := args.Populate(reader)
		assert.NoError(t, err)
		assert.NotNil(t, args.TxType)
		assert.Equal(t, uint8(ethrpctypes.SetCodeTxType), *args.TxType)
	})

	t.Run("dynamic fee without explicit fee returns error", func(t *testing.T) {
		args := &TransactionArgs{
			From:   &common.Address{},
			To:     &common.Address{},
			TxType: TxTypePtr(ethrpctypes.DynamicFeeTxType),
		}

		err := args.Populate(reader)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "dynamic fee transaction requires explicit maxFeePerGas and maxPriorityFeePerGas on non-1559 chain")
	})
}

func TestConvertDynamicFeeTxToArgs(t *testing.T) {
	dtx := &ethrpctypes.DynamicFeeTx{}

	actual := ConvertTransactionToArgs(common.Address{}, ethrpctypes.NewTx(dtx))
	expect := `{"from":"0x0000000000000000000000000000000000000000","to":null,"gas":null,"value":"0x0","nonce":null,"data":"0x","accessList":[],"type":2}`

	argsJ, _ := json.Marshal(actual)
	assert.Equal(t, expect, string(argsJ))
}

func TestConvertLegacyTxToArgs(t *testing.T) {
	dtx := &ethrpctypes.LegacyTx{}

	actual := ConvertTransactionToArgs(common.Address{}, ethrpctypes.NewTx(dtx))
	expect := `{"from":"0x0000000000000000000000000000000000000000","to":null,"gas":null,"value":"0x0","nonce":null,"data":"0x","type":0}`

	argsJ, _ := json.Marshal(actual)
	assert.Equal(t, expect, string(argsJ))
}

func TestConvertAccesslistTxToArgs(t *testing.T) {
	dtx := &ethrpctypes.AccessListTx{}

	actual := ConvertTransactionToArgs(common.Address{}, ethrpctypes.NewTx(dtx))
	expect := `{"from":"0x0000000000000000000000000000000000000000","to":null,"gas":null,"value":"0x0","nonce":null,"data":"0x","accessList":[],"type":1}`

	argsJ, _ := json.Marshal(actual)
	assert.Equal(t, expect, string(argsJ))
}

func TestConvertSetCodeTxToArgs(t *testing.T) {
	t.Run("empty fields", func(t *testing.T) {
		dtx := &ethrpctypes.SetCodeTx{AccessList: ethrpctypes.AccessList{}, AuthList: []ethrpctypes.SetCodeAuthorization{}}

		actual := ConvertTransactionToArgs(common.Address{}, ethrpctypes.NewTx(dtx))
		expect := `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":null,"value":"0x0","nonce":null,"data":"0x","accessList":[],"type":4}`

		argsJ, err := json.Marshal(actual)
		assert.NoError(t, err)
		assert.Equal(t, expect, string(argsJ))
	})

	t.Run("with authorization and gas fields", func(t *testing.T) {
		to := common.HexToAddress("0x1111111111111111111111111111111111111111")
		dtx := &ethrpctypes.SetCodeTx{
			ChainID:   uint256.NewInt(1),
			Nonce:     5,
			GasTipCap: uint256.NewInt(1000000000),
			GasFeeCap: uint256.NewInt(2000000000),
			Gas:       21000,
			To:        to,
			Value:     uint256.NewInt(0),
			AccessList: ethrpctypes.AccessList{
				{Address: to, StorageKeys: []common.Hash{common.HexToHash("0x01")}},
			},
			AuthList: []ethrpctypes.SetCodeAuthorization{
				{
					ChainID: *uint256.NewInt(1),
					Address: common.HexToAddress("0x2222222222222222222222222222222222222222"),
					Nonce:   10,
					V:       1,
					R:       *uint256.NewInt(0x1234),
					S:       *uint256.NewInt(0x5678),
				},
			},
		}

		from := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		actual := ConvertTransactionToArgs(from, ethrpctypes.NewTx(dtx))

		assert.Equal(t, &from, actual.From)
		assert.Equal(t, &to, actual.To)
		assert.NotNil(t, actual.Nonce)
		assert.Equal(t, uint64(5), uint64(*actual.Nonce))
		assert.NotNil(t, actual.Gas)
		assert.Equal(t, uint64(21000), uint64(*actual.Gas))
		assert.Equal(t, big.NewInt(1000000000), (*big.Int)(actual.MaxPriorityFeePerGas))
		assert.Equal(t, big.NewInt(2000000000), (*big.Int)(actual.MaxFeePerGas))
		assert.Equal(t, big.NewInt(1), (*big.Int)(actual.ChainID))
		assert.NotNil(t, actual.AccessList)
		assert.Equal(t, 1, len(*actual.AccessList))
		assert.Equal(t, 1, len(actual.AuthorizationList))
		assert.Equal(t, *uint256.NewInt(1), actual.AuthorizationList[0].ChainID)
		assert.Equal(t, common.HexToAddress("0x2222222222222222222222222222222222222222"), actual.AuthorizationList[0].Address)
		assert.Equal(t, uint64(10), actual.AuthorizationList[0].Nonce)
		assert.Equal(t, uint8(1), actual.AuthorizationList[0].V)

		argsJ, err := json.Marshal(actual)
		assert.NoError(t, err)
		var raw map[string]json.RawMessage
		err = json.Unmarshal(argsJ, &raw)
		assert.NoError(t, err)
		assert.Contains(t, raw, "authorizationList")
		assert.Contains(t, raw, "maxFeePerGas")
		assert.Contains(t, raw, "maxPriorityFeePerGas")
		assert.Contains(t, raw, "accessList")
	})
}

func TestPopulate(t *testing.T) {
	ast := assert.New(t)

	table := []struct {
		input       *TransactionArgs
		expectOut   string
		expectError bool
	}{
		{
			input:       &TransactionArgs{},
			expectError: true,
		},
		{
			input:     &TransactionArgs{From: &common.Address{}, To: &common.Address{}},
			expectOut: `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":"0x26","maxFeePerGas":"0x92","maxPriorityFeePerGas":"0x1e","value":"0x0","nonce":"0x30","data":null,"chainId":"0x12","type":2}`,
		},
		{
			input:     &TransactionArgs{From: &common.Address{}, To: &common.Address{}, GasPrice: (*hexutil.Big)(big.NewInt(33))},
			expectOut: `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":"0x26","gasPrice":"0x21","value":"0x0","nonce":"0x30","data":null,"chainId":"0x12","type":0}`,
		},
		{
			input:     &TransactionArgs{From: &common.Address{}, To: &common.Address{}, MaxFeePerGas: (*hexutil.Big)(big.NewInt(44)), MaxPriorityFeePerGas: (*hexutil.Big)(big.NewInt(22))},
			expectOut: `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":"0x26","maxFeePerGas":"0x2c","maxPriorityFeePerGas":"0x16","value":"0x0","nonce":"0x30","data":null,"chainId":"0x12","type":2}`,
		},
		{
			input:     &TransactionArgs{From: &common.Address{}, To: &common.Address{}, MaxFeePerGas: (*hexutil.Big)(big.NewInt(44))},
			expectOut: `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":"0x26","maxFeePerGas":"0x2c","maxPriorityFeePerGas":"0x1e","value":"0x0","nonce":"0x30","data":null,"chainId":"0x12","type":2}`,
		},
		{
			input:       &TransactionArgs{From: &common.Address{}, To: &common.Address{}, GasPrice: (*hexutil.Big)(big.NewInt(33)), MaxFeePerGas: (*hexutil.Big)(big.NewInt(44))},
			expectError: true,
		},
		// SetCodeTxType: inferred from AuthorizationList, gas populated via 1559 logic
		{
			input: &TransactionArgs{
				From:              &common.Address{},
				To:                &common.Address{},
				AuthorizationList: []ethrpctypes.SetCodeAuthorization{{ChainID: *uint256.NewInt(1), Address: common.Address{}}},
			},
			expectOut: `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":"0x26","maxFeePerGas":"0x92","maxPriorityFeePerGas":"0x1e","value":"0x0","nonce":"0x30","data":null,"authorizationList":[{"chainId":"0x1","address":"0x0000000000000000000000000000000000000000","nonce":"0x0","yParity":"0x0","r":"0x0","s":"0x0"}],"chainId":"0x12","type":4}`,
		},
		// SetCodeTxType: explicit type with GasPrice converts to maxFeePerGas/maxPriorityFeePerGas
		{
			input: &TransactionArgs{
				From:              &common.Address{},
				To:                &common.Address{},
				TxType:            TxTypePtr(ethrpctypes.SetCodeTxType),
				GasPrice:          (*hexutil.Big)(big.NewInt(33)),
				AuthorizationList: []ethrpctypes.SetCodeAuthorization{{ChainID: *uint256.NewInt(1), Address: common.Address{}}},
			},
			expectOut: `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":"0x26","maxFeePerGas":"0x21","maxPriorityFeePerGas":"0x21","value":"0x0","nonce":"0x30","data":null,"authorizationList":[{"chainId":"0x1","address":"0x0000000000000000000000000000000000000000","nonce":"0x0","yParity":"0x0","r":"0x0","s":"0x0"}],"chainId":"0x12","type":4}`,
		},
	}

	for i, item := range table {
		err := item.input.Populate(&mockPopulateReader{})
		if item.expectError {
			ast.Error(err, i)
			continue
		}

		ast.NoError(err)
		actual, _ := json.Marshal(item.input)
		ast.Equal(item.expectOut, string(actual))
	}
}

func TestJsonMarshalHexBytes(t *testing.T) {
	j, _ := json.Marshal((*hexutil.Bytes)(nil))
	fmt.Printf("%s\n", string(j))
}
