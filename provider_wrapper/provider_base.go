package providers

import (
	"context"

	rpc "github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/web3go/interfaces"
)

func NewBaseProvider(ctx context.Context, url string) (interfaces.Provider, error) {
	return rpc.DialContext(ctx, url)
}
