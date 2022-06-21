package web3go

import (
	"io"
	"time"

	"github.com/mcuadros/go-defaults"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/signers"
)

type ClientOption struct {
	providers.Option
	SignerManager *signers.SignerManager
}

func (c *ClientOption) setDefault() *ClientOption {
	defaults.SetDefaults(&c.Option)
	return c
}

func (c *ClientOption) WithRetry(retryCount int, retryInterval time.Duration) *ClientOption {
	c.Option.WithRetry(retryCount, retryInterval)
	return c
}

func (c *ClientOption) WithTimout(requestTimeout time.Duration) *ClientOption {
	c.Option.WithTimout(requestTimeout)
	return c
}

func (c *ClientOption) WithMaxConnectionPerHost(maxConnectionPerHost int) *ClientOption {
	c.Option.WithMaxConnectionPerHost(maxConnectionPerHost)
	return c
}

func (c *ClientOption) WithLooger(w io.Writer) *ClientOption {
	c.Option.WithLooger(w)
	return c
}

func (c *ClientOption) WithSignerManager(signerManager *signers.SignerManager) *ClientOption {
	c.SignerManager = signerManager
	return c
}
