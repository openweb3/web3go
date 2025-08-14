package web3go

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

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

func TestSendTxByArgsUseClientWithOption(t *testing.T) {
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

func TestGetBlockByHash(t *testing.T) {
	client, err := NewClient("https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa")
	if err != nil {
		t.Fatal(err)
	}

	block, err := client.Eth.BlockByNumber(types.LatestBlockNumber, false)
	assert.NoError(t, err)
	j, _ := json.Marshal(block)
	fmt.Printf("block: %+v\n", string(j))
}

func TestGetBlockByHash2(t *testing.T) {
	client, err := NewClient("https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa")
	if err != nil {
		t.Fatal(err)
	}

	var result interface{}
	err = client.Eth.CallContext(context.Background(), &result, "eth_getBlockByNumber", "latest", false)
	assert.NoError(t, err)

	j, _ := json.Marshal(result)
	fmt.Printf("block: %+v\n", string(j))
}
