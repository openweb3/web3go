package web3go

import (
	"context"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	client "github.com/openweb3/web3go/client"
	"github.com/openweb3/web3go/interfaces"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	c   interfaces.RpcProvider
	eth *client.RpcEthClient
}

// clientOption for set keystore path and flags for retry
//
// The simplest way to set logger is to use the types.DefaultCallRpcLog and types.DefaultBatchCallRPCLog
type clientOption struct {
	KeystorePath string
	// retry
	RetryCount    int
	RetryInterval time.Duration
	// timeout of request
	RequestTimeout time.Duration
}

type Option func(c *clientOption)

func NewClient(rawurl string, options ...Option) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}

	// Override defaults with any provided options
	o := clientOption{}
	for _, opt := range options {
		opt(&o)
	}

	ec := &Client{c, nil}

	return ec, nil
}

func WithRetry(retryCount int, retryInterval time.Duration) Option {
	return func(c *clientOption) {
		c.RetryCount = retryCount
		c.RetryInterval = retryInterval
	}
}

func WithTimout(requestTimeout time.Duration) Option {
	return func(c *clientOption) {
		c.RequestTimeout = requestTimeout
	}
}
