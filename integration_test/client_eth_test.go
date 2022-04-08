package integrationtest

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/go-sdk-common/rpctest"
	"github.com/openweb3/web3go/client"
	providers "github.com/openweb3/web3go/provider_wrapper"
	"github.com/openweb3/web3go/types"
)

func int2Hexbig(result interface{}) (handlerdResult interface{}) {
	return (*hexutil.Big)(result.(*big.Int))
}

func u64ToHexU64(result interface{}) (handlerdResult interface{}) {
	return (*hexutil.Uint64)(result.(*uint64))
}

func bytes2HexBytes(result interface{}) (handlerdResult interface{}) {
	return (hexutil.Bytes)(result.([]byte))
}

func getEthTestConfig() rpctest.RpcTestConfig {

	var rpc2Func map[string]string = map[string]string{
		"web3_clientVersion":                      "ClientVersion",
		"net_version":                             "NetVersion",
		"eth_protocolVersion":                     "ProtocolVersion",
		"eth_syncing":                             "Syncing",
		"eth_hashrate":                            "Hashrate",
		"eth_coinbase":                            "Author",
		"eth_mining":                              "IsMining",
		"eth_chainId":                             "ChainId",
		"eth_gasPrice":                            "GasPrice",
		"eth_maxPriorityFeePerGas":                "MaxPriorityFeePerGas",
		"eth_accounts":                            "Accounts",
		"eth_blockNumber":                         "BlockNumber",
		"eth_getBalance":                          "Balance",
		"eth_getStorageAt":                        "StorageAt",
		"eth_getBlockByHash":                      "BlockByHash",
		"eth_getBlockByNumber":                    "BlockByNumber",
		"eth_getTransactionCount":                 "TransactionCount",
		"eth_getBlockTransactionCountByHash":      "BlockTransactionCountByHash",
		"eth_getBlockTransactionCountByNumber":    "BlockTransactionCountByNumber",
		"eth_getUncleCountByBlockHash":            "BlockUnclesCountByHash",
		"eth_getUncleCountByBlockNumber":          "BlockUnclesCountByNumber",
		"eth_getCode":                             "CodeAt",
		"eth_sendRawTransaction":                  "SendRawTransaction",
		"eth_submitTransaction":                   "SubmitTransaction",
		"eth_call":                                "Call",
		"eth_estimateGas":                         "EstimateGas",
		"eth_getTransactionByHash":                "TransactionByHash",
		"eth_getTransactionByBlockHashAndIndex":   "TransactionByBlockHashAndIndex",
		"eth_getTransactionByBlockNumberAndIndex": "TransactionByBlockNumberAndIndex",
		"eth_getTransactionReceipt":               "TransactionReceipt",
		"eth_getUncleByBlockHashAndIndex":         "UncleByBlockHashAndIndex",
		"eth_getUncleByBlockNumberAndIndex":       "UncleByBlockNumberAndIndex",
		"eth_getLogs":                             "Logs",
		"eth_submitHashrate":                      "SubmitHashrate",
	}

	rpc2FuncSelector := map[string]func(params []interface{}) (string, []interface{}){}
	rpc2FuncResultHandler := map[string]func(result interface{}) (handlerdResult interface{}){
		"eth_hashrate":                         int2Hexbig,
		"eth_chainId":                          u64ToHexU64,
		"eth_gasPrice":                         int2Hexbig,
		"eth_maxPriorityFeePerGas":             int2Hexbig,
		"eth_blockNumber":                      int2Hexbig,
		"eth_getBalance":                       int2Hexbig,
		"eth_getTransactionCount":              int2Hexbig,
		"eth_getBlockTransactionCountByHash":   int2Hexbig,
		"eth_getBlockTransactionCountByNumber": int2Hexbig,
		"eth_getUncleCountByBlockHash":         int2Hexbig,
		"eth_getUncleCountByBlockNumber":       int2Hexbig,
		"eth_getCode":                          bytes2HexBytes,
		"eth_call":                             bytes2HexBytes,
		"eth_estimateGas":                      int2Hexbig,
	}

	var ignoreRpc map[string]bool = map[string]bool{}
	var onlyTestRpc map[string]bool = map[string]bool{
		// "eth_getTransactionByHash": true,
	}

	provider, _ := providers.NewBaseProvider(context.Background(), "http://47.93.101.243/eth/")
	return rpctest.RpcTestConfig{
		ExamplesUrl: "https://raw.githubusercontent.com/Conflux-Chain/jsonrpc-spec/main/src/eth/examples.json",
		Client:      client.NewRpcEthClient(provider),

		Rpc2Func:              rpc2Func,
		Rpc2FuncSelector:      rpc2FuncSelector,
		Rpc2FuncResultHandler: rpc2FuncResultHandler,
		IgnoreRpcs:            ignoreRpc,
		OnlyTestRpcs:          onlyTestRpc,
	}

}

// TODO: Open after rpc mock server ready
func TestClienEth(t *testing.T) {
	config := getEthTestConfig()
	rpctest.DoClientTest(t, config)
}

func TestClienEthA(t *testing.T) {
	v := types.Transaction{}
	fmt.Printf("%v\n", reflect.ValueOf(v).Kind())

	fmt.Printf("%v\n", reflect.ValueOf(v).Elem().Kind())
}
