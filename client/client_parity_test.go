package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	ethrpctypes "github.com/ethereum/go-ethereum/rpc"
	providers "github.com/openweb3/web3go/provider_wrapper"
)

func TestParityClient(t *testing.T) {
	p, err := providers.NewBaseProvider(context.Background(), "http://net8889eth.confluxrpc.com")
	if err != nil {
		t.Fatal(err)
	}
	traceclient := NewRpcParityClient(p)

	blockNum := ethrpctypes.BlockNumberOrHashWithNumber(2883646)
	val, _ := traceclient.BlockReceipts(&blockNum)
	j, _ := json.MarshalIndent(val, "", "  ")
	fmt.Printf("block receipts: %v\n", string(j))
}
