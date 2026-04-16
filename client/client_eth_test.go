package client

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	rpc "github.com/openweb3/go-rpc-provider"

	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
)

func TestBatchCall(t *testing.T) {
	provider := pproviders.MustNewBaseProvider(context.Background(), "https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa")
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	var batchElems []rpc.BatchElem
	batchElems = append(batchElems, rpc.BatchElem{
		Method: "eth_blockNumber",
		Args:   nil,
		Result: new(hexutil.Big),
	}, rpc.BatchElem{
		Method: "eth_getBlockByNumber",
		Args:   []interface{}{types.LatestBlockNumber, false},
		Result: new(types.Block),
	},
	)
	err := provider.BatchCallContext(context.Background(), batchElems)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, *batchElems[0].Result.(*hexutil.Big))
	assert.NotEqual(t, types.Block{}, *batchElems[1].Result.(*types.Block))
}

func TestAccountPendingTransactions(t *testing.T) {
	// node := "https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa"
	node := "https://evmtestnet-internal.confluxrpc.com"
	provider := pproviders.MustNewBaseProvider(context.Background(), node)
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	c := NewRpcEthClient(provider)

	addr := common.HexToAddress("0xBeD38c825459994002257DFBB88371E243204B6c")
	pendingTxs, err := c.AccountPendingTransactions(addr, nil, nil)
	assert.NoError(t, err)
	fmt.Printf("pending: %+v\n", pendingTxs)
}
