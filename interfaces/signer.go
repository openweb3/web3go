package interfaces

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
)

type Signer interface {
	Address() common.Address
	SignTransaction(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
	SignMessage(text []byte) ([]byte, error)
}
