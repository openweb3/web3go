package web3go

import (
	"context"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
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
}

func NewClient(rawurl string) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}
	eth := client.NewRpcEthClient(c)
	ec := &Client{c, nil, eth}

	return ec, nil
}

func NewClientWithOption(rawurl string, option *ClientOption) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}

	defaults.SetDefaults(option)
	p := wrapProvider(c, option)

	eth := client.NewRpcEthClient(p)
	ec := &Client{c, nil, eth}

	return ec, nil
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
