package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	providers "github.com/openweb3/web3go/provider_wrapper"
	"github.com/openweb3/web3go/types"
)

func _TestParity(t *testing.T) {
	url := "http://localhost:8545"
	// url := "http://net8889eth.confluxrpc.com"
	p, err := providers.NewBaseProvider(context.Background(), url)
	if err != nil {
		panic(err)
	}
	client := NewRpcParityClient(p)

	blockNum := types.BlockNumberOrHashWithNumber(1)
	phrase := "stylus outing overhand dime radial seducing harmless uselessly evasive tastiness eradicate imperfect"
	account := common.HexToAddress("0x00a329c0648769a73afac7f9381e08fb43dbea72")

	t.Run("TransactionsLimit", func(t *testing.T) {
		transactionsLimitVal, err1 := client.TransactionsLimit()
		j, err2 := json.MarshalIndent(transactionsLimitVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("TransactionsLimit result: %v\n", string(j))
	})

	t.Run("ExtraData", func(t *testing.T) {
		extraDataVal, err1 := client.ExtraData()
		j, err2 := json.MarshalIndent(extraDataVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("ExtraData result: %v\n", string(j))
	})

	t.Run("GasFloorTarget", func(t *testing.T) {
		gasFloorTargetVal, err1 := client.GasFloorTarget()
		j, err2 := json.MarshalIndent(gasFloorTargetVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("GasFloorTarget result: %v\n", string(j))
	})

	t.Run("GasCeilTarget", func(t *testing.T) {
		gasCeilTargetVal, err1 := client.GasCeilTarget()
		j, err2 := json.MarshalIndent(gasCeilTargetVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("GasCeilTarget result: %v\n", string(j))
	})

	t.Run("MinGasPrice", func(t *testing.T) {
		minGasPriceVal, err1 := client.MinGasPrice()
		j, err2 := json.MarshalIndent(minGasPriceVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("MinGasPrice result: %v\n", string(j))
	})

	t.Run("DevLogs", func(t *testing.T) {
		devLogsVal, err1 := client.DevLogs()
		j, err2 := json.MarshalIndent(devLogsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("DevLogs result: %v\n", string(j))
	})

	t.Run("DevLogsLevels", func(t *testing.T) {
		devLogsLevelsVal, err1 := client.DevLogsLevels()
		j, err2 := json.MarshalIndent(devLogsLevelsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("DevLogsLevels result: %v\n", string(j))
	})

	t.Run("NetChain", func(t *testing.T) {
		netChainVal, err1 := client.NetChain()
		j, err2 := json.MarshalIndent(netChainVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("NetChain result: %v\n", string(j))
	})

	t.Run("NetPeers", func(t *testing.T) {
		netPeersVal, err1 := client.NetPeers()
		j, err2 := json.MarshalIndent(netPeersVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("NetPeers result: %v\n", string(j))
	})

	t.Run("NetPort", func(t *testing.T) {
		netPortVal, err1 := client.NetPort()
		j, err2 := json.MarshalIndent(netPortVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("NetPort result: %v\n", string(j))
	})

	t.Run("RpcSettings", func(t *testing.T) {
		rpcSettingsVal, err1 := client.RpcSettings()
		j, err2 := json.MarshalIndent(rpcSettingsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("RpcSettings result: %v\n", string(j))
	})

	t.Run("NodeName", func(t *testing.T) {
		nodeNameVal, err1 := client.NodeName()
		j, err2 := json.MarshalIndent(nodeNameVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("NodeName result: %v\n", string(j))
	})

	t.Run("DefaultExtraData", func(t *testing.T) {
		defaultExtraDataVal, err1 := client.DefaultExtraData()
		j, err2 := json.MarshalIndent(defaultExtraDataVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("DefaultExtraData result: %v\n", string(j))
	})

	t.Run("GasPriceHistogram", func(t *testing.T) {
		gasPriceHistogramVal, err1 := client.GasPriceHistogram()
		j, err2 := json.MarshalIndent(gasPriceHistogramVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("GasPriceHistogram result: %v\n", string(j))
	})

	t.Run("UnsignedTransactionsCount", func(t *testing.T) {
		unsignedTransactionsCountVal, err1 := client.UnsignedTransactionsCount()
		j, err2 := json.MarshalIndent(unsignedTransactionsCountVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("UnsignedTransactionsCount result: %v\n", string(j))
	})

	t.Run("GenerateSecretPhrase", func(t *testing.T) {
		generateSecretPhraseVal, err1 := client.GenerateSecretPhrase()
		j, err2 := json.MarshalIndent(generateSecretPhraseVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("GenerateSecretPhrase result: %v\n", string(j))
	})

	t.Run("PhraseToAddress", func(t *testing.T) {
		phraseToAddressVal, err1 := client.PhraseToAddress(phrase)
		j, err2 := json.MarshalIndent(phraseToAddressVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("PhraseToAddress result: %v\n", string(j))
	})

	t.Run("RegistryAddress", func(t *testing.T) {
		registryAddressVal, err1 := client.RegistryAddress()
		j, err2 := json.MarshalIndent(registryAddressVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("RegistryAddress result: %v\n", string(j))
	})

	t.Run("ListAccounts", func(t *testing.T) {
		listAccountsVal, err1 := client.ListAccounts(10, nil, nil)
		j, err2 := json.MarshalIndent(listAccountsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("ListAccounts result: %v\n", string(j))
	})

	t.Run("ListStorageKeys", func(t *testing.T) {
		listStorageKeysVal, err1 := client.ListStorageKeys(account, 5, nil, nil)
		j, err2 := json.MarshalIndent(listStorageKeysVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("ListStorageKeys result: %v\n", string(j))
	})

	t.Run("EncryptMessage", func(t *testing.T) {
		publickey := "0xD219959D466D666060284733A80DDF025529FEAA8337169540B3267B8763652A13D878C40830DD0952639A65986DBEC611CF2171A03CFDC37F5A40537068AA4F"
		message := []byte("hello world") // "hello world"
		encryptMessageVal, err1 := client.EncryptMessage(publickey, message)
		j, err2 := json.MarshalIndent(hexutil.Bytes(encryptMessageVal), "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("EncryptMessage result: %v\n", string(j))
	})

	t.Run("PendingTransactions", func(t *testing.T) {
		filter := &types.TransactionFilter{
			To: &types.ActionArgument{
				Eq: common.HexToAddress("0xe8b2d01ffa0a15736b2370b6e5064f9702c891b6"),
			},
			Gas: &types.ValueFilterArgument{
				Gt: big.NewInt(0x493e0),
			},
		}
		pendingTransactionsVal, err1 := client.PendingTransactions(nil, filter)
		j, err2 := json.MarshalIndent(pendingTransactionsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("PendingTransactions result: %v\n", string(j))
	})

	t.Run("AllTransactions", func(t *testing.T) {
		allTransactionsVal, err1 := client.AllTransactions()
		j, err2 := json.MarshalIndent(allTransactionsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("AllTransactions result: %v\n", string(j))
	})

	t.Run("AllTransactionHashes", func(t *testing.T) {
		allTransactionHashesVal, err1 := client.AllTransactionHashes()
		j, err2 := json.MarshalIndent(allTransactionHashesVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("AllTransactionHashes result: %v\n", string(j))
	})

	t.Run("FutureTransactions", func(t *testing.T) {
		futureTransactionsVal, err1 := client.FutureTransactions()
		j, err2 := json.MarshalIndent(futureTransactionsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("FutureTransactions result: %v\n", string(j))
	})

	t.Run("PendingTransactionsStats", func(t *testing.T) {
		pendingTransactionsStatsVal, err1 := client.PendingTransactionsStats()
		j, err2 := json.MarshalIndent(pendingTransactionsStatsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("PendingTransactionsStats result: %v\n", string(j))
	})

	t.Run("LocalTransactions", func(t *testing.T) {
		localTransactionsVal, err1 := client.LocalTransactions()
		j, err2 := json.MarshalIndent(localTransactionsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("LocalTransactions result: %v\n", string(j))
	})

	t.Run("WsUrl", func(t *testing.T) {
		wsUrlVal, err1 := client.WsUrl()
		j, err2 := json.MarshalIndent(wsUrlVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("WsUrl result: %v\n", string(j))
	})

	t.Run("NextNonce", func(t *testing.T) {
		nextNonceVal, err1 := client.NextNonce(account)
		j, err2 := json.MarshalIndent(nextNonceVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("NextNonce result: %v\n", string(j))
	})

	t.Run("Mode", func(t *testing.T) {
		modeVal, err1 := client.Mode()
		j, err2 := json.MarshalIndent(modeVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("Mode result: %v\n", string(j))
	})

	t.Run("Chain", func(t *testing.T) {
		chainVal, err1 := client.Chain()
		j, err2 := json.MarshalIndent(chainVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("Chain result: %v\n", string(j))
	})

	t.Run("Enode", func(t *testing.T) {
		enodeVal, err1 := client.Enode()
		j, err2 := json.MarshalIndent(enodeVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("Enode result: %v\n", string(j))
	})

	t.Run("ChainStatus", func(t *testing.T) {
		chainStatusVal, err1 := client.ChainStatus()
		j, err2 := json.MarshalIndent(chainStatusVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("ChainStatus result: %v\n", string(j))
	})

	t.Run("NodeKind", func(t *testing.T) {
		nodeKindVal, err1 := client.NodeKind()
		j, err2 := json.MarshalIndent(nodeKindVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("NodeKind result: %v\n", string(j))
	})

	t.Run("BlockHeader", func(t *testing.T) {
		blockHeaderVal, err1 := client.BlockHeader(&blockNum)
		j, err2 := json.MarshalIndent(blockHeaderVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("BlockHeader result: %v\n", string(j))
	})

	t.Run("BlockReceipts", func(t *testing.T) {
		blockReceiptsVal, err1 := client.BlockReceipts(&blockNum)
		j, err2 := json.MarshalIndent(blockReceiptsVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("BlockReceipts result: %v\n", string(j))
	})

	t.Run("Call", func(t *testing.T) {
		callVal, err1 := client.Call([]types.CallRequest{{
			From: &account,
			To:   nil,
			Data: common.Hex2Bytes("0x3859818153F3"),
		}}, &blockNum)
		j, err2 := json.MarshalIndent(callVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("Call result: %v\n", string(j))
	})

	t.Run("SubmitWorkDetail", func(t *testing.T) {
		submitWorkDetailVal, err1 := client.SubmitWorkDetail("0x0000000000000001",
			common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
			common.HexToHash("0xD1FE5700000000000000000000000000D1FE5700000000000000000000000000"))
		j, err2 := json.MarshalIndent(submitWorkDetailVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("SubmitWorkDetail result: %v\n", string(j))
	})

	t.Run("Status", func(t *testing.T) {
		err1 := client.Status()
		if err1 != nil {
			t.Fatal(err1)
		}
	})

	t.Run("VerifySignature", func(t *testing.T) {
		verifySignatureVal, err1 := client.VerifySignature(
			true,
			common.Hex2Bytes("0xbc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98a"),
			common.HexToHash("0x4355c47d63924e8a72e509b65029052eb6c299d53a04e167c5775fd466751c9d"),
			common.HexToHash("0x07299936d304c153f6443dfa05f40ff007d72911b6f72307f996231605b91562"),
			0x45,
		)
		j, err2 := json.MarshalIndent(verifySignatureVal, "", "  ")
		if err1 != nil || err2 != nil {
			t.Fatal(err1, err2)
		}
		fmt.Printf("VerifySignature result: %v\n", string(j))
	})
}
