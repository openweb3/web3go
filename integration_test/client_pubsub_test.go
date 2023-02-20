package integrationtest

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
)

func _TestSubHeader(t *testing.T) {
	client, err := web3go.NewClientWithOption("wss://evmtestnet-internal.confluxrpc.com/ws ", web3go.ClientOption{
		Option: providers.Option{
			Logger: os.Stdout,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	head := make(chan *types.Header, 100)
	sub, err := client.Eth.SubscribeNewHead(head)
	assert.NoError(t, err)
	defer sub.Unsubscribe()

	for h := range head {
		j, _ := json.Marshal(h)
		fmt.Printf("%s\n", j)
	}
}
