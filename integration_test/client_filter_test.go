package integrationtest

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
)

func _TestFilters(t *testing.T) {
	client, err := web3go.NewClientWithOption("https://evmtestnet-internal.confluxrpc.com", web3go.ClientOption{
		Option: providers.Option{
			Logger: os.Stdout,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	logId, err := client.Filter.NewLogFilter(&types.FilterQuery{
		Topics: [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
	})
	assert.NoError(t, err)
	fmt.Printf("new log filter: %v\n", *logId)

	blockId, err := client.Filter.NewBlockFilter()
	assert.NoError(t, err)
	fmt.Printf("new block filter: %v\n", *blockId)

	pendingTxId, err := client.Filter.NewPendingTransactionFilter()
	assert.NoError(t, err)
	fmt.Printf("new pending tx filter: %v\n", *pendingTxId)

	for i := 0; i < 10; i++ {
		logChanges, err := client.Filter.GetFilterLogs(*logId)
		assert.NoError(t, err)
		fmt.Printf("filterd logs: %v\n", logChanges)

		blockChanges, err := client.Filter.GetFilterChanges(*blockId)
		assert.NoError(t, err)
		fmt.Printf("block changes: %v\n", blockChanges)

		pendingTxChanges, err := client.Filter.GetFilterChanges(*pendingTxId)
		assert.NoError(t, err)
		fmt.Printf("pending tx changes: %v\n", pendingTxChanges)
		time.Sleep(time.Second * 2)
	}

	ok, err := client.Filter.UninstallFilter(*logId)
	assert.NoError(t, err)
	fmt.Printf("uninstall filter result: %v\n", ok)
}
