package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ethrpctypes "github.com/ethereum/go-ethereum/rpc"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/go-sdk-common/rpctest"
	"github.com/openweb3/web3go/client"
)

func int2Hexbig(val interface{}) (converted interface{}) {
	return (*hexutil.Big)(val.(*big.Int))
}

func hexBig2Int(val interface{}) (converted interface{}) {
	return (*big.Int)(val.(*hexutil.Big))
}

func u64ToHexU64(val interface{}) (converted interface{}) {
	return (*hexutil.Uint64)(val.(*uint64))
}

func bytes2HexBytes(val interface{}) (converted interface{}) {
	return (hexutil.Bytes)(val.([]byte))
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

	rpc2FuncSelector := map[string]func(params []interface{}) (realFuncName string, realParams []interface{}){
		"eth_getBalance": func(params []interface{}) (realFuncName string, realParams []interface{}) {
			if len(params) == 1 {
				return "Balance", append(params, ethrpctypes.LatestBlockNumber)
			}
			return "Balance", params
		},
		"eth_getStorageAt": func(params []interface{}) (realFuncName string, realParams []interface{}) {
			params[1] = hexutil.MustDecodeBig(params[1].(string))
			if len(params) == 2 {
				return "StorageAt", append(params, ethrpctypes.LatestBlockNumber)
			}
			return "StorageAt", params
		},
		"eth_getTransactionCount": func(params []interface{}) (realFuncName string, realParams []interface{}) {
			if len(params) == 1 {
				return "TransactionCount", append(params, ethrpctypes.LatestBlockNumber)
			}
			return "TransactionCount", params
		},
		"eth_getCode": func(params []interface{}) (realFuncName string, realParams []interface{}) {
			if len(params) == 1 {
				return "CodeAt", append(params, ethrpctypes.LatestBlockNumber)
			}
			return "CodeAt", params
		},
		"eth_call": func(params []interface{}) (realFuncName string, realParams []interface{}) {
			if len(params) == 1 {
				return "Call", append(params, ethrpctypes.LatestBlockNumber)
			}
			return "Call", params
		},
		"eth_estimateGas": func(params []interface{}) (realFuncName string, realParams []interface{}) {
			if len(params) == 1 {
				return "EstimateGas", append(params, ethrpctypes.LatestBlockNumber)
			}
			return "EstimateGas", params
		},
	}

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

	ignoreRpc := map[string]bool{}
	onlyTestRpc := map[string]bool{}
	ignoreExamples := map[string]bool{
		"eth_getCode-0x1e6309dc46a2a4936abda54b69c91d7a3c75a39e":            true,
		"eth_getStorageAt-0x1e6309dc46a2a4936abda54b69c91d7a3c75a39e,0x100": true,
		"eth_getLogs-1649755528773":                                         true,
	}
	onlyExamples := map[string]bool{}

	provider, _ := providers.NewBaseProvider(context.Background(), "http://47.93.101.243/eth/")
	middled := providers.NewMiddlewarableProvider(provider)
	middled.HookCallContext(callcontextFuncLogMiddle)
	provider = middled

	return rpctest.RpcTestConfig{
		ExamplesUrl: "https://raw.githubusercontent.com/Conflux-Chain/jsonrpc-spec/main/src/eth/examples.json",
		Client:      client.NewRpcEthClient(provider),

		Rpc2Func:              rpc2Func,
		Rpc2FuncSelector:      rpc2FuncSelector,
		Rpc2FuncResultHandler: rpc2FuncResultHandler,
		IgnoreRpcs:            ignoreRpc,
		IgnoreExamples:        ignoreExamples,
		OnlyTestRpcs:          onlyTestRpc,
		OnlyExamples:          onlyExamples,
	}

}

func callcontextFuncLogMiddle(f providers.CallContextFunc) providers.CallContextFunc {
	return func(ctx context.Context, resultPtr interface{}, method string, args ...interface{}) error {
		jArgs, _ := json.Marshal(args)
		fmt.Printf("\n-- rpc call %v %s--\n", method, jArgs)
		err := f(ctx, resultPtr, method, args...)
		j, _ := json.Marshal(resultPtr)
		fmt.Printf("\trpc response %s\n", j)
		fmt.Printf("\trpc error %v\n", err)
		return err
	}
}

func TestClienEth(t *testing.T) {
	config := getEthTestConfig()
	rpctest.DoClientTest(t, config)
}
