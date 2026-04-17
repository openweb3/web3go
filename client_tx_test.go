//go:build tx

package web3go

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethrpctypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"

	rpc "github.com/openweb3/go-rpc-provider"
	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
	client "github.com/openweb3/web3go/client"
	"github.com/openweb3/web3go/providers"
	"github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
)

func TestSendTransactionUseEthClient(t *testing.T) {
	mnemonic := "crisp shove million stem shiver side hospital split play lottery join vintage"
	sm := signers.MustNewSignerManagerByMnemonic(mnemonic, 10, nil)
	from := sm.List()[0].Address()
	to := common.Address{}

	provider := pproviders.MustNewBaseProvider(context.Background(), "https://evmtestnet.confluxrpc.com")
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)
	provider = providers.NewSignableProvider(provider, sm)
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	c := client.NewRpcEthClient(provider)

	pendingBlock := types.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber)
	nonce, err := c.TransactionCount(from, &pendingBlock)
	assert.NoError(t, err)

	// legacy tx
	tx := ethrpctypes.NewTransaction(nonce.Uint64(), to, big.NewInt(1), 1000000, big.NewInt(20000000000), nil)
	txhash, err := c.SendTransaction(from, tx)
	assert.NoError(t, err)
	fmt.Printf("txhash: %s\n", txhash)

	// dynamic fee tx
	dtx := &ethrpctypes.DynamicFeeTx{
		To:    &to,
		Value: big.NewInt(1),
	}
	txhash, err = c.SendTransaction(from, ethrpctypes.NewTx(dtx))
	assert.NoError(t, err)
	fmt.Printf("txhash: %s\n", txhash)
}

func TestSendTxByArgsUseClientWithOption(t *testing.T) {
	mnemonic := "crisp shove million stem shiver side hospital split play lottery join vintage"
	sm := signers.MustNewSignerManagerByMnemonic(mnemonic, 10, nil)
	c, err := NewClientWithOption("https://evmtestnet.confluxrpc.com", *(new(ClientOption).WithLooger(os.Stdout).WithSignerManager(sm)))
	assert.NoError(t, err)

	from := sm.List()[0].Address()
	to := common.Address{}

	t.Run("send simple tx", func(t *testing.T) {
		hash, err := c.Eth.SendTransactionByArgs(types.TransactionArgs{
			From: &from,
			To:   &to,
		})
		assert.NoError(t, err)
		fmt.Printf("hash: %s\n", hash)
	})

	t.Run("send 7702 tx", func(t *testing.T) {
		chainId, err := c.Eth.ChainId()
		assert.NoError(t, err, "Failed to retrieve chain ID")
		fmt.Println("Chain ID:", *chainId)

		pengdingBlockNumber := types.BlockNumberOrHashWithNumber(types.PendingBlockNumber)
		nonce, err := c.Eth.TransactionCount(from, &pengdingBlockNumber)
		assert.NoError(t, err, "Failed to retrieve pending nonce")
		fmt.Println("Pending nonce:", nonce)

		usdt0Addr := common.HexToAddress("0x7d682e65EFC5C13Bf4E394B8f376C48e6baE0355")
		auth := ethrpctypes.SetCodeAuthorization{
			ChainID: *uint256.NewInt(*chainId),
			Address: usdt0Addr,
			Nonce:   nonce.Uint64() + 1,
		}

		nonceHex := hexutil.Uint64(nonce.Uint64())
		tx := types.TransactionArgs{
			From:              &from,
			To:                &to,
			Nonce:             &nonceHex,
			AuthorizationList: []ethrpctypes.SetCodeAuthorization{auth},
		}
		assert.Nil(t, tx.TxType, "TxType should be nil before Populate")

		err = tx.Populate(c.Eth)
		assert.NoError(t, err, "Failed to populate tx args")
		assert.NotNil(t, tx.TxType)
		assert.Equal(t, uint8(ethrpctypes.SetCodeTxType), *tx.TxType, "TxType should be auto-inferred as SetCodeTxType")

		txHash, err := c.Eth.SendTransactionByArgs(tx)
		assert.NoError(t, err, "Failed to send 7702 tx")
		fmt.Println("7702 tx sent:", txHash)
	})
}
