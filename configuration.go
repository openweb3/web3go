package web3go

import "time"

// ClientOption for set keystore path and flags for retry
//
// The simplest way to set logger is to use the types.DefaultCallRpcLog and types.DefaultBatchCallRPCLog
type ClientOption struct {
	// KeystorePath string
	// retry
	RetryCount    int           `default:"1"`
	RetryInterval time.Duration `default:"1000"`
	// timeout of request
	RequestTimeout time.Duration `default:"1000"`
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
