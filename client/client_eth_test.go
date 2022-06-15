package client

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"

	ethrpctypes "github.com/ethereum/go-ethereum/core/types"
	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go/providers"
	"github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/assert"
)

func TestSendTransaction(t *testing.T) {
	ast := assert.New(t)
	provider := pproviders.MustNewBaseProvider(context.Background(), "https://ropsten.infura.io/v3/cb2c1b76cb894b699f20a602f35731f1")
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)
	sm, err := signers.NewSignerManagerByPrivateKeyStrings([]string{"9ec393923a14eeb557600010ea05d635c667a6995418f8a8f4bdecc63dfe0bb9"})
	ast.NoError(err)

	provider = providers.NewSignableProvider(provider, sm)
	provider = pproviders.NewLoggerProvider(provider, os.Stdout)

	c := NewRpcEthClient(provider)

	from := common.HexToAddress("0xe6D148D8398c4cb456196C776D2d9093Dd62C9B0")
	pendingBlock := types.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber)
	nonce, err := c.TransactionCount(from, &pendingBlock)
	ast.NoError(err)

	tx := ethrpctypes.NewTransaction(nonce.Uint64(), from, big.NewInt(1000000), 1000000, big.NewInt(1), nil)
	txhash, err := c.SendTransaction(from, *tx)
	ast.NoError(err)

	fmt.Printf("txhash: %s\n", txhash)

	nonce, err = c.TransactionCount(from, &pendingBlock)
	ast.NoError(err)
	dtx := &ethrpctypes.DynamicFeeTx{
		Nonce: nonce.Uint64(),
		// GasTipCap: big.NewInt(100000),
		// GasFeeCap: big.NewInt(1000000),
		Gas:   21000,
		To:    &from,
		Value: big.NewInt(1),
		Data:  nil,
	}
	txhash, err = c.SendTransaction(from, *ethrpctypes.NewTx(dtx))
	ast.NoError(err)
	fmt.Printf("txhash: %s\n", txhash)
}
