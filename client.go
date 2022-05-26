package web3go

import (
	"github.com/openweb3/go-rpc-provider/interfaces"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	client "github.com/openweb3/web3go/client"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	*providers.MiddlewarableProvider
	option *providers.Option
	Eth    *client.RpcEthClient
	Trace  *client.RpcTraceClient
	Parity *client.RpcParityClient
}

func NewClient(rawurl string) (*Client, error) {
	return NewClientWithOption(rawurl, providers.Option{})
}

func NewClientWithOption(rawurl string, option providers.Option) (*Client, error) {
	p, err := providers.NewProviderWithOption(rawurl, option)
	if err != nil {
		return nil, err
	}

	ec := NewClientWithProvider(p)
	ec.option = &option

	return ec, nil
}

func NewClientWithProvider(p interfaces.Provider) *Client {
	c := &Client{}
	c.SetProvider(p)
	return c
}

func (c *Client) SetProvider(p interfaces.Provider) {
	c.MiddlewarableProvider = providers.NewMiddlewarableProvider(p)
	c.Eth = client.NewRpcEthClient(p)
	c.Trace = client.NewRpcTraceClient(p)
	c.Parity = client.NewRpcParityClient(p)
}

func (c *Client) Provider() *providers.MiddlewarableProvider {
	return c.MiddlewarableProvider
}
