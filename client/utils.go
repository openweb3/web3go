package client

import (
	rpc "github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/web3go/types"
)

func getRealBlockNumberOrHash(input *types.BlockNumberOrHash) *types.BlockNumberOrHash {
	if input == nil {
		tmp := types.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber)
		return &tmp
	}
	return input
}
