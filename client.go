package web3go

import (
	"context"

	"github.com/mcuadros/go-defaults"
	client "github.com/openweb3/web3go/client"
	"github.com/openweb3/web3go/interfaces"
	"github.com/openweb3/web3go/internal"
	providers "github.com/openweb3/web3go/provider_wrapper"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	provider interfaces.Provider
	option   *ClientOption
	Eth      *client.RpcEthClient
	Trace    *client.RpcTraceClient
	Parity   *client.RpcParityClient
}

func NewClient(rawurl string) (*Client, error) {
	p, err := providers.NewBaseProvider(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}

	ec := NewClientWithProvider(p)
	return ec, nil
}

func NewClientWithOption(rawurl string, option *ClientOption) (*Client, error) {
	p, err := providers.NewBaseProvider(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}

	if option == nil {
		option = &ClientOption{}
	}

	defaults.SetDefaults(option)
	p = wrapProvider(p, option)

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
	c.provider = p
	c.Eth = client.NewRpcEthClient(p)
	c.Trace = client.NewRpcTraceClient(p)
	c.Parity = client.NewRpcParityClient(p)
}

func (c *Client) Provider() interfaces.Provider {
	return c.provider
}

// wrapProvider wrap provider accroding to option
func wrapProvider(p interfaces.Provider, option *ClientOption) interfaces.Provider {
	if option == nil {
		return p
	}

	p = internal.NewTimeoutableProvider(p, option.RequestTimeout)
	p = providers.NewRetriableProvider(p, option.RetryCount, option.RetryInterval)
	return p
}
