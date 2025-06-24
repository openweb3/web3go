package types

import (
	"math/big"

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
