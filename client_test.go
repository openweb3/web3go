package web3go

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
)

func TestClient(t *testing.T) {
	client, err := NewClient("https://evm.confluxrpc.com")
	if err != nil {
		t.Fatal(err)
	}

	p := client.Provider()
	mp := providers.NewMiddlewarableProvider(p)
	mp.HookCallContext(callcontextLogMiddleware)
	client.SetProvider(mp)

	_, err = client.Eth.ClientVersion()
	if err != nil {
		t.Fatal(err)
	}
}

func callcontextLogMiddleware(f providers.CallContextFunc) providers.CallContextFunc {
	return func(ctx context.Context, resultPtr interface{}, method string, args ...interface{}) error {
		fmt.Printf("request %v %v\n", method, args)
		err := f(ctx, resultPtr, method, args...)
		j, _ := json.Marshal(resultPtr)
		fmt.Printf("response %s\n", j)
		return err
	}
}
