// Copyright 2019 OpenWeb3. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package web3go

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/go-rpc-provider/interfaces"
	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
	client "github.com/openweb3/web3go/client"
	"github.com/openweb3/web3go/providers"
	"github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	*pproviders.MiddlewarableProvider
	option *ClientOption
	Eth    *client.RpcEthClient
	Trace  *client.RpcTraceClient
	Parity *client.RpcParityClient
	Filter *client.RpcFilterClient
}

var (
	ErrNotFound = errors.New("not found")
)

func NewClient(rawurl string) (*Client, error) {
	return NewClientWithOption(rawurl, ClientOption{})
}

func MustNewClient(rawurl string) *Client {
	c, err := NewClient(rawurl)
	if err != nil {
		panic(err)
	}
	return c
}

func NewClientWithOption(rawurl string, option ClientOption) (*Client, error) {

	option.setDefault()

	p, err := pproviders.NewProviderWithOption(rawurl, option.Option)
	if err != nil {
		return nil, err
	}

	if option.SignerManager != nil {
		p = providers.NewSignableProvider(p, option.SignerManager)
	}

	ec := NewClientWithProvider(p)
	ec.option = &option

	return ec, nil
}

func MustNewClientWithOption(rawurl string, option ClientOption) *Client {
	c, err := NewClientWithOption(rawurl, option)
	if err != nil {
		panic(err)
	}
	return c
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
	c.Filter = client.NewRpcFilterClient(p)
}

func (c *Client) Provider() *pproviders.MiddlewarableProvider {
	return c.MiddlewarableProvider
}

// GetSignerManager returns signer manager if exist in option, otherwise return error
func (c *Client) GetSignerManager() (*signers.SignerManager, error) {
	if c.option.SignerManager != nil {
		return c.option.SignerManager, nil
	}
	return nil, ErrNotFound
}

// ToClientForContract returns ClientForContract and SignerFn for use by abi-binding struct generated by abigen.
// abigen is a source code generator to convert Ethereum contract definitions into easy to use, compile-time type-safe Go packages.
// Please see https://geth.ethereum.org/docs/dapp/native-bindings page for details
func (c *Client) ToClientForContract() (*ClientForContract, bind.SignerFn) {
	sm, err := c.GetSignerManager()
	if err != nil {
		return NewClientForContract(c), nil
	}

	signFunc := func(addr common.Address, t *types.Transaction) (*types.Transaction, error) {
		chainId, err := c.Eth.ChainId()
		if err != nil {
			return nil, err
		}
		s, err := sm.Get(addr)
		if err != nil {
			return nil, err
		}
		return s.SignTransaction(t, new(big.Int).SetUint64(*chainId))
	}

	return NewClientForContract(c), signFunc
}
