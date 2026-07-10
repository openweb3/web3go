package client

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	rpc "github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type recordingProvider struct {
	method string
	args   []interface{}
	result json.RawMessage
}

func (p *recordingProvider) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	p.method = method
	p.args = append([]interface{}{}, args...)
	return json.Unmarshal(p.result, result)
}

func (p *recordingProvider) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return nil
}

func (p *recordingProvider) Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error) {
	return nil, nil
}

func (p *recordingProvider) SubscribeWithReconn(ctx context.Context, namespace string, channel interface{}, args ...interface{}) *rpc.ReconnClientSubscription {
	return nil
}

func (p *recordingProvider) Close() {}

func TestRpcDebugClientBlockProperties(t *testing.T) {
	block := types.BlockNumberOrHashWithNumber(types.NewBlockNumber(436))
	provider := &recordingProvider{
		result: json.RawMessage(`[
			{
				"txHash": "0x3719bb0b4385a7e0266d1e266166d821b351a38a1a78f2e36df99c73bbbc15ae",
				"innerBlockHash": "0x446012a81945dc9cde4eca03697e43d5f80beed878d78b9829b54cb9a1f9f7a4",
				"coinbase": "0x1d69d968e3673e188b2d2d42b6a385686186258f",
				"difficulty": "0x4",
				"gasLimit": "0x3938700",
				"timestamp": "0x68ee1848",
				"baseFeePerGas": "0x1"
			}
		]`),
	}

	c := NewRpcDebugClient(provider)
	props, err := c.BlockProperties(block)
	require.NoError(t, err)

	assert.Equal(t, "debug_blockProperties", provider.method)
	require.Len(t, provider.args, 1)
	assert.Equal(t, block, provider.args[0])
	require.Len(t, props, 1)
	assert.Equal(t, common.HexToHash("0x3719bb0b4385a7e0266d1e266166d821b351a38a1a78f2e36df99c73bbbc15ae"), *props[0].TxHash)
	assert.Equal(t, common.HexToHash("0x446012a81945dc9cde4eca03697e43d5f80beed878d78b9829b54cb9a1f9f7a4"), props[0].InnerBlockHash)
	assert.Equal(t, common.HexToAddress("0x1d69d968e3673e188b2d2d42b6a385686186258f"), props[0].Coinbase)
	assert.Equal(t, big.NewInt(4), props[0].Difficulty)
	assert.Equal(t, big.NewInt(60000000), props[0].GasLimit)
	assert.Equal(t, uint64(1760434248), props[0].Timestamp)
	assert.Equal(t, big.NewInt(1), props[0].BaseFeePerGas)
}

func TestRpcDebugClientBlockPropertiesNullResult(t *testing.T) {
	block := types.BlockNumberOrHashWithHash(common.HexToHash("0x446012a81945dc9cde4eca03697e43d5f80beed878d78b9829b54cb9a1f9f7a4"), false)
	provider := &recordingProvider{result: json.RawMessage(`null`)}

	c := NewRpcDebugClient(provider)
	props, err := c.BlockProperties(block)
	require.NoError(t, err)

	assert.Equal(t, "debug_blockProperties", provider.method)
	require.Len(t, provider.args, 1)
	assert.Equal(t, block, provider.args[0])
	assert.Nil(t, props)
}
