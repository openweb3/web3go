package client

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"testing"

	rpc "github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type createAccessListProvider struct {
	method string
	args   []interface{}
}

func (p *createAccessListProvider) CallContext(_ context.Context, result interface{}, method string, args ...interface{}) error {
	p.method = method
	p.args = args
	return json.Unmarshal([]byte(`{"accessList":[],"gasUsed":"0x5208"}`), result)
}

func (*createAccessListProvider) BatchCallContext(context.Context, []rpc.BatchElem) error {
	return errors.New("not implemented")
}

func (*createAccessListProvider) Subscribe(context.Context, string, interface{}, ...interface{}) (*rpc.ClientSubscription, error) {
	return nil, errors.New("not implemented")
}

func (*createAccessListProvider) SubscribeWithReconn(context.Context, string, interface{}, ...interface{}) *rpc.ReconnClientSubscription {
	return nil
}

func (*createAccessListProvider) Close() {}

func TestCreateAccessList(t *testing.T) {
	provider := new(createAccessListProvider)
	client := NewRpcEthClient(provider)
	block := types.BlockNumberOrHashWithNumber(types.LatestBlockNumber)
	overrides := types.StateOverride{}

	result, err := client.CreateAccessList(types.CallRequest{}, &block, &overrides)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, big.NewInt(21000), result.GasUsed)
	assert.Equal(t, "eth_createAccessList", provider.method)
	require.Len(t, provider.args, 3)
	assert.IsType(t, types.CallRequest{}, provider.args[0])
	assert.Equal(t, &block, provider.args[1])
	assert.Same(t, &overrides, provider.args[2])
}
