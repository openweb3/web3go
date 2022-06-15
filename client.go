package web3go

import (
	"github.com/openweb3/go-rpc-provider/interfaces"
	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
	client "github.com/openweb3/web3go/client"
	"github.com/openweb3/web3go/providers"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	*pproviders.MiddlewarableProvider
	option *ClinetOption
	Eth    *client.RpcEthClient
	Trace  *client.RpcTraceClient
	Parity *client.RpcParityClient
}

func NewClient(rawurl string) (*Client, error) {
	return NewClientWithOption(rawurl, &ClinetOption{
		&pproviders.Option{}, nil,
	})
}

func NewClientWithOption(rawurl string, option *ClinetOption) (*Client, error) {
	p, err := pproviders.NewProviderWithOption(rawurl, *option.Option)
	if err != nil {
		return nil, err
	}

	if option.SignerManager != nil {
		p = providers.NewSignableProvider(p, option.SignerManager)
	}

	ec := NewClientWithProvider(p)
	ec.option = option

	return ec, nil
}

func NewClientWithProvider(p interfaces.Provider) *Client {
	c := &Client{}
	c.SetProvider(p)
	return c
}

func (c *Client) SetProvider(p interfaces.Provider) {
	if _, ok := p.(*pproviders.MiddlewarableProvider); !ok {
		p = pproviders.NewMiddlewarableProvider(p)
	}

	c.MiddlewarableProvider = p.(*pproviders.MiddlewarableProvider)
	c.Eth = client.NewRpcEthClient(p)
	c.Trace = client.NewRpcTraceClient(p)
	c.Parity = client.NewRpcParityClient(p)
}

func (c *Client) Provider() *pproviders.MiddlewarableProvider {
	return c.MiddlewarableProvider
}
