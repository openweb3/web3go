package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/client"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
)

func TestTraceTransaction(t *testing.T) {
	provider := pproviders.MustNewBaseProvider(context.Background(), "https://evmtestnet-internal.confluxrpc.com")
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	c := client.NewRpcDebugClient(provider)
	val, err := c.TraceTransaction(common.HexToHash("0x6b9a69106704eec878731c00b251c3a67e1fbb561bda279e92ee3f6071da6500"))
	assert.NoError(t, err)

	fmt.Printf("val %v\n", val)

	for _, tracer := range []string{
		"4byteTracer",
		"callTracer",
		"prestateTracer",
		"noopTracer",
		"muxTracer",
	} {

		onlyTopCall := true
		val, err := c.TraceTransaction(common.HexToHash("0x6b9a69106704eec878731c00b251c3a67e1fbb561bda279e92ee3f6071da6500"), &types.GethDebugTracingOptions{
			Tracer:       tracer,
			TracerConfig: &types.GethDebugTracerConfig{CallConfig: &types.CallConfig{OnlyTopCall: &onlyTopCall}},
		})
		assert.NoError(t, err)

		j, _ := json.Marshal(val)
		fmt.Printf("val %s\n", j)
	}
}

func TestTraceBlockByHash(t *testing.T) {
	provider := pproviders.MustNewBaseProvider(context.Background(), "https://evmtestnet-internal.confluxrpc.com")
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	c := client.NewRpcDebugClient(provider)

	val, err := c.TraceBlockByHash(common.HexToHash("0x92346af4d942871946186ef86c7136ffa45abf1424a605df2723a9ed5694d1f8"))
	assert.NoError(t, err)

	j, err := json.Marshal(val)
	assert.NoError(t, err)
	fmt.Printf("val %s\n", j)
}

func TestTraceBlockByNumber(t *testing.T) {
	provider := pproviders.MustNewBaseProvider(context.Background(), "https://evmtestnet-internal.confluxrpc.com")
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	c := client.NewRpcDebugClient(provider)

	val, err := c.TraceBlockByNumber(types.NewBlockNumber(179904465))
	assert.NoError(t, err)

	j, err := json.Marshal(val)
	assert.NoError(t, err)
	fmt.Printf("val %s\n", j)
}

func TestTraceCall(t *testing.T) {
	provider := pproviders.MustNewBaseProvider(context.Background(), "https://evmtestnet-internal.confluxrpc.com")
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	c := client.NewRpcDebugClient(provider)

	blkNum := types.NewBlockNumber(179904465)
	to := common.HexToAddress("0x807da62384be660ded0319d613d8b37cf3892d20")
	val, err := c.TraceCall(types.CallRequest{
		To: &to,
	}, &blkNum)
	assert.NoError(t, err)

	j, err := json.Marshal(val)
	assert.NoError(t, err)
	fmt.Printf("val %s\n", j)
}
