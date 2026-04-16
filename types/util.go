package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
)

func BigIntToBlockNumber(input *big.Int) *BlockNumber {
	if input == nil {
		return nil
	}
	v := BlockNumber(input.Int64())
	return &v
}

func BigIntToUint256(input *big.Int) *uint256.Int {
	if input == nil {
		return nil
	}
	return uint256.NewInt(0).SetBytes(input.Bytes())
}

func HexBigToUint256(v *hexutil.Big) *uint256.Int {
	if v == nil {
		return uint256.NewInt(0)
	}
	return BigIntToUint256(v.ToInt())
}

func Pointer[T any](val T) *T {
	return &val
}
