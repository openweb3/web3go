package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethrpctypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

type mockPopulateReader struct{}

func (m *mockPopulateReader) ChainId() (val *uint64, err error) {
	chainId := uint64(18)
	return &chainId, nil
}

func (m *mockPopulateReader) GasPrice() (val *big.Int, err error) {
	return big.NewInt(28), nil
}

func (m *mockPopulateReader) EstimateGas(callRequest CallRequest, blockNum *BlockNumberOrHash) (val *big.Int, err error) {
	return big.NewInt(38), nil
}

func (m *mockPopulateReader) TransactionCount(addr common.Address, blockNum *BlockNumberOrHash) (val *big.Int, err error) {
	return big.NewInt(48), nil
}

func (m *mockPopulateReader) BlockByNumber(num BlockNumber, isFull bool) (val *Block, err error) {
	return &Block{BaseFeePerGas: big.NewInt(58)}, nil
}

func TestConvertDynamicFeeTxToArgs(t *testing.T) {
	dtx := &ethrpctypes.DynamicFeeTx{}

	actual := ConvertTransactionToArgs(common.Address{}, *ethrpctypes.NewTx(dtx))
	expect := `{"from":"0x0000000000000000000000000000000000000000","to":null,"gas":null,"value":"0x0","nonce":null,"data":"0x","accessList":[],"type":2}`

	argsJ, _ := json.Marshal(actual)
	assert.Equal(t, expect, string(argsJ))
}

func TestConvertLegacyTxToArgs(t *testing.T) {
	dtx := &ethrpctypes.LegacyTx{}

	actual := ConvertTransactionToArgs(common.Address{}, *ethrpctypes.NewTx(dtx))
	expect := `{"from":"0x0000000000000000000000000000000000000000","to":null,"gas":null,"value":"0x0","nonce":null,"data":"0x","type":0}`

	argsJ, _ := json.Marshal(actual)
	assert.Equal(t, expect, string(argsJ))
}

func TestConvertAccesslistTxToArgs(t *testing.T) {
	dtx := &ethrpctypes.AccessListTx{}

	actual := ConvertTransactionToArgs(common.Address{}, *ethrpctypes.NewTx(dtx))
	expect := `{"from":"0x0000000000000000000000000000000000000000","to":null,"gas":null,"value":"0x0","nonce":null,"data":"0x","accessList":[],"type":1}`

	argsJ, _ := json.Marshal(actual)
	assert.Equal(t, expect, string(argsJ))
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
			expectOut: `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":"0x26","maxFeePerGas":"0x59682f74","maxPriorityFeePerGas":"0x59682f00","value":"0x0","nonce":"0x30","data":null,"chainId":"0x12","type":2}`,
		},
		{
			input:     &TransactionArgs{From: &common.Address{}, To: &common.Address{}, GasPrice: (*hexutil.Big)(big.NewInt(33))},
			expectOut: `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":"0x26","maxFeePerGas":"0x21","maxPriorityFeePerGas":"0x21","value":"0x0","nonce":"0x30","data":null,"chainId":"0x12","type":2}`,
		},
		{
			input:     &TransactionArgs{From: &common.Address{}, To: &common.Address{}, MaxFeePerGas: (*hexutil.Big)(big.NewInt(44)), MaxPriorityFeePerGas: (*hexutil.Big)(big.NewInt(22))},
			expectOut: `{"from":"0x0000000000000000000000000000000000000000","to":"0x0000000000000000000000000000000000000000","gas":"0x26","maxFeePerGas":"0x2c","maxPriorityFeePerGas":"0x16","value":"0x0","nonce":"0x30","data":null,"chainId":"0x12","type":2}`,
		},
		{
			input:       &TransactionArgs{From: &common.Address{}, To: &common.Address{}, MaxFeePerGas: (*hexutil.Big)(big.NewInt(44))},
			expectError: true,
		},
		{
			input:       &TransactionArgs{From: &common.Address{}, To: &common.Address{}, GasPrice: (*hexutil.Big)(big.NewInt(33)), MaxFeePerGas: (*hexutil.Big)(big.NewInt(44))},
			expectError: true,
		},
	}

	for _, item := range table {
		err := item.input.Populate(&mockPopulateReader{})
		if item.expectError {
			ast.Error(err)
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
