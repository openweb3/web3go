package types

import (
	"math/big"

	"github.com/openweb3/go-rpc-provider"
)

type gasFeeData struct {
	gasPrice             *big.Int
	maxFeePerGas         *big.Int
	maxPriorityFeePerGas *big.Int
}

func getFeeData(r ReaderForPopulate) (*gasFeeData, error) {
	data := &gasFeeData{}

	gasPrice, err := r.GasPrice()
	if err != nil {
		return nil, err
	}
	data.gasPrice = gasPrice

	block, err := r.BlockByNumber(rpc.LatestBlockNumber, false)
	if err != nil {
		return nil, err
	}
	basefee := block.BaseFeePerGas

	if basefee == nil {
		return data, nil
	}

	data.maxPriorityFeePerGas = big.NewInt(1.5e9)
	data.maxFeePerGas = new(big.Int).Mul(basefee, big.NewInt(2))
	data.maxFeePerGas = new(big.Int).Add(data.maxFeePerGas, data.maxPriorityFeePerGas)
	return data, nil
}

func (g gasFeeData) isSupport1559() bool {
	return g.maxPriorityFeePerGas != nil && g.maxFeePerGas != nil
}
