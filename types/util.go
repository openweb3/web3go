package types

import "math/big"

func BigIntToBlockNumber(input *big.Int) *BlockNumber {
	if input == nil {
		return nil
	}
	v := BlockNumber(input.Int64())
	return &v
}
