package client

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	rpc "github.com/openweb3/go-rpc-provider"

	ethrpctypes "github.com/ethereum/go-ethereum/core/types"
	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/providers"
	"github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
)

func TestSendTransaction(t *testing.T) {
	ast := assert.New(t)
	sm, err := signers.NewSignerManagerByPrivateKeyStrings([]string{"9ec393923a14eeb557600010ea05d635c667a6995418f8a8f4bdecc63dfe0bb9"})
	ast.NoError(err)

	provider := pproviders.MustNewBaseProvider(context.Background(), "https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa")
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)
	provider = providers.NewSignableProvider(provider, sm)
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	c := NewRpcEthClient(provider)

	from := common.HexToAddress("0xe6D148D8398c4cb456196C776D2d9093Dd62C9B0")
	pendingBlock := types.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber)
	nonce, err := c.TransactionCount(from, &pendingBlock)
	ast.NoError(err)

	// legacy tx
	tx := ethrpctypes.NewTransaction(nonce.Uint64(), from, big.NewInt(1000000), 1000000, big.NewInt(1), nil)
	txhash, err := c.SendTransaction(from, *tx)
	ast.NoError(err)
	fmt.Printf("txhash: %s\n", txhash)

	// dynamic fee tx
	dtx := &ethrpctypes.DynamicFeeTx{
		To:    &from,
		Value: big.NewInt(1),
	}
	txhash, err = c.SendTransaction(from, *ethrpctypes.NewTx(dtx))
	ast.NoError(err)
	fmt.Printf("txhash: %s\n", txhash)
}

func TestBatchCall(t *testing.T) {
	provider := pproviders.MustNewBaseProvider(context.Background(), "https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa")
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	var batchElems []rpc.BatchElem
	batchElems = append(batchElems, rpc.BatchElem{
		Method: "eth_blockNumber",
		Args:   nil,
		Result: new(hexutil.Big),
	}, rpc.BatchElem{
		Method: "eth_getBlockByNumber",
		Args:   []interface{}{types.LatestBlockNumber, false},
		Result: new(types.Block),
	},
	)
	err := provider.BatchCallContext(context.Background(), batchElems)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, *batchElems[0].Result.(*hexutil.Big))
	assert.NotEqual(t, types.Block{}, *batchElems[1].Result.(*types.Block))
}

func TestAccountPendingTransactions(t *testing.T) {
	// node := "https://sepolia.infura.io/v3/d91582da330a4812be53d698a34741aa"
	node := "https://evmtestnet-internal.confluxrpc.com"
	provider := pproviders.MustNewBaseProvider(context.Background(), node)
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	c := NewRpcEthClient(provider)

	addr := common.HexToAddress("0xBeD38c825459994002257DFBB88371E243204B6c")
	pendingTxs, err := c.AccountPendingTransactions(addr, nil, nil)
	assert.NoError(t, err)
	fmt.Printf("pending: %+v\n", pendingTxs)
}
