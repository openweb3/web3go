package web3go

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client, err := NewClient("https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa")
	if err != nil {
		t.Fatal(err)
	}

	p := client.Provider()
	mp := pproviders.NewMiddlewarableProvider(p)
	mp.HookCallContext(callcontextLogMiddleware)
	client.SetProvider(mp)

	_, err = client.Eth.ClientVersion()
	if err != nil {
		t.Fatal(err)
	}
}

func callcontextLogMiddleware(f pproviders.CallContextFunc) pproviders.CallContextFunc {
	return func(ctx context.Context, resultPtr interface{}, method string, args ...interface{}) error {
		fmt.Printf("request %v %v\n", method, args)
		err := f(ctx, resultPtr, method, args...)
		j, _ := json.Marshal(resultPtr)
		fmt.Printf("response %s\n", j)
		return err
	}
}

func _TestSendTxByArgsUseClientWithOption(t *testing.T) {
	mnemonic := "crisp shove million stem shiver side hospital split play lottery join vintage"
	sm := signers.MustNewSignerManagerByMnemonic(mnemonic, 10, nil)
	c, err := NewClientWithOption("https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa", *(new(ClientOption).WithLooger(os.Stdout).WithSignerManager(sm)))
	assert.NoError(t, err)

	from := sm.List()[0].Address()
	to := sm.List()[1].Address()
	hash, err := c.Eth.SendTransactionByArgs(types.TransactionArgs{
		From: &from,
		To:   &to,
	})
	assert.NoError(t, err)
	fmt.Printf("hash: %s\n", hash)
}

func TestTraceSetAuth(t *testing.T) {
	client, err := NewClientWithOption("https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa", *(new(ClientOption).WithLooger(os.Stdout)))
	if err != nil {
		t.Fatal(err)
	}

	to := common.HexToAddress("0x00000000863B56a3C1f0F1be8BC4F8b7BD78F57a")
	block := types.BlockNumberOrHashWithNumber(types.LatestBlockNumber)
	traces, err := client.Eth.EstimateGas(types.CallRequest{
		To:   &to,
		Data: hexutil.MustDecode("0x4cb0a33fc6cce7302152bfb28296f583e60b753fd0ca8bcf21b01c712c94cc66"),
	}, &block, &types.StateOverride{
		common.HexToAddress("0x00000000863B56a3C1f0F1be8BC4F8b7BD78F57a"): {
			Balance: (*hexutil.Big)(big.NewInt(1000000000000000000)),
		},
	}, &types.BlockOverrides{})

	assert.NoError(t, err)
	j, _ := json.Marshal(traces)
	fmt.Printf("traces: %s\n", j)
}
