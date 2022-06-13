package client

import (
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/openweb3/web3go/types"
)

func getRealBlockNumberOrHash(input *types.BlockNumberOrHash) *types.BlockNumberOrHash {
	if input == nil {
		tmp := types.BlockNumberOrHashWithNumber(ethrpc.LatestBlockNumber)
		return &tmp
	}
	return input
}
