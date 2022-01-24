package providers

import (
	"context"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/openweb3/web3go/interfaces"
)

func NewBaseProvider(ctx context.Context, url string) (interfaces.Provider, error) {
	return rpc.DialContext(ctx, url)
}
