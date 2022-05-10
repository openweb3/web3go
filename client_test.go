package web3go

import (
	"encoding/json"
	"fmt"
	"testing"

	providers "github.com/openweb3/web3go/provider_wrapper"
)

func TestClient(t *testing.T) {
	client, err := NewClientWithOption("https://evm.confluxrpc.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	p := client.Provider()
	mp := providers.NewMiddlewarableProvider(p)
	mp.HookCall(callLogMiddleware)
	client.SetProvider(mp)

	_, err = client.Eth.ClientVersion()
	if err != nil {
		t.Fatal(err)
	}
}

func callLogMiddleware(f providers.CallFunc) providers.CallFunc {
	return func(resultPtr interface{}, method string, args ...interface{}) error {
		fmt.Printf("request %v %v\n", method, args)
		err := f(resultPtr, method, args...)
		j, _ := json.Marshal(resultPtr)
		fmt.Printf("response %s\n", j)
		return err
	}
}
