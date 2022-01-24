package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gotest.tools/assert"
)

var executeStack []byte = make([]byte, 0)

func TestHookCall(t *testing.T) {
	p, e := rpc.DialContext(context.Background(), "http://localhost:8545")
	if e != nil {
		t.Fatal(e)
	}
	mp := NewMiddlewarableProvider(p)

	mp.HookCall(callMiddle1)
	mp.HookCall(callMiddle2)
	mp.HookCall(callMiddle3)

	b := new(hexutil.Big)
	mp.Call(b, "eth_blockNumber")

	assert.DeepEqual(t, executeStack, []byte{1, 2, 3, 4, 5, 6})
}

func callMiddle1(f CallFunc) CallFunc {
	return func(resultPtr interface{}, method string, args ...interface{}) error {
		executeStack = append(executeStack, 1)
		fmt.Printf("call %v %v\n", method, args)
		err := f(resultPtr, method, args...)
		j, _ := json.Marshal(resultPtr)
		fmt.Printf("response %s", j)
		executeStack = append(executeStack, 6)
		return err
	}
}

func callMiddle2(f CallFunc) CallFunc {
	return func(resultPtr interface{}, method string, args ...interface{}) error {
		executeStack = append(executeStack, 2)
		fmt.Println("foo1 middle start")
		e := f(resultPtr, method, args...)
		fmt.Println("foo1 middle end")
		executeStack = append(executeStack, 5)
		return e
	}
}

func callMiddle3(f CallFunc) CallFunc {
	return func(resultPtr interface{}, method string, args ...interface{}) error {
		executeStack = append(executeStack, 3)
		fmt.Println("foo2 middle start")
		e := f(resultPtr, method, args...)
		fmt.Println("foo2 middle end")
		executeStack = append(executeStack, 4)
		return e
	}
}
