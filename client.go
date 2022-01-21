package web3go

import (
	"context"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/mcuadros/go-defaults"
	client "github.com/openweb3/web3go/client"
	"github.com/openweb3/web3go/interfaces"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	c   interfaces.Provider
	Eth *client.RpcEthClient
}

func NewClient(rawurl string) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}
	eth := client.NewRpcEthClient(c)
	ec := &Client{c, eth}

	return ec, nil
}

func NewClientWithOption(rawurl string, option *ClientOption) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}

	if option == nil {
		defaults.SetDefaults(&option)
	}

	eth := client.NewRpcEthClient(c)
	ec := &Client{c, eth}

	return ec, nil
}
