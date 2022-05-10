package web3go

import (
	"time"
)

// ClientOption for set keystore path and flags for retry
type ClientOption struct {
	// KeystorePath string
	// retry
	RetryCount    int           `default:"3"`
	RetryInterval time.Duration `default:"1s"`
	// timeout of request
	RequestTimeout time.Duration `default:"3s"`
	// Maximum number of connections may be established. The default value is 512.
	MaxConnectionNum int
}

func (o *ClientOption) WithRetry(retryCount int, retryInterval time.Duration) *ClientOption {
	o.RetryCount = retryCount
	o.RetryInterval = retryInterval
	return o
}

func (o *ClientOption) WithTimout(requestTimeout time.Duration) *ClientOption {
	o.RequestTimeout = requestTimeout
	return o
}

// func (o *ClientOption) WrapProvider(provider interfaces.Provider) *providers.MiddlewarableProvider {
// 	retryProvider := providers.NewRetriableProvider(provider, o.RetryCount, o.RetryInterval)
// }
