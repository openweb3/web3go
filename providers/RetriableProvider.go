package providers

import (
	"time"

	"github.com/openweb3/web3go/interfaces"
)

type RetriableProvider struct {
}

func NewRetriableProvider(p interfaces.RpcProvider, maxRetry int, interval time.Duration) *RetriableProvider {
	return &RetriableProvider{}
}
