package web3go

import (
	"io"
	"time"

	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/signers"
)

type ClinetOption struct {
	*providers.Option
	SignerManager *signers.SignerManager
}

func (c *ClinetOption) WithRetry(retryCount int, retryInterval time.Duration) *ClinetOption {
	c.Option.WithRetry(retryCount, retryInterval)
	return c
}

func (c *ClinetOption) WithTimout(requestTimeout time.Duration) *ClinetOption {
	c.Option.WithTimout(requestTimeout)
	return c
}

func (c *ClinetOption) WithMaxConnectionPerHost(maxConnectionPerHost int) *ClinetOption {
	c.Option.WithMaxConnectionPerHost(maxConnectionPerHost)
	return c
}

func (c *ClinetOption) WithLooger(w io.Writer) *ClinetOption {
	c.Option.WithLooger(w)
	return c
}

func (c *ClinetOption) WithSignerManager(signerManager *signers.SignerManager) *ClinetOption {
	c.SignerManager = signerManager
	return c
}
