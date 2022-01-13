package web3go

import (
	"context"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	client "github.com/openweb3/web3go/client"
	"github.com/openweb3/web3go/interfaces"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	c   interfaces.RpcProvider
	Eth *client.RpcEthClient
}

// type Option func(c *ClientOption)

func NewClient(rawurl string) (*Client, error) {
	return NewClientWithOption(rawurl, ClientOption{})
}

func NewClientWithOption(rawurl string, option ClientOption) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}

	// Override defaults with any provided options
	// o := ClientOption{}
	// for _, opt := range options {
	// 	opt(&o)
	// }

	eth := client.NewRpcEthClient(c)
	ec := &Client{c, eth}

	return ec, nil
}

// func (c *Client) Eth() *client.RpcEthClient {
// 	return c.eth
// }
