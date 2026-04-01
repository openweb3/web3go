package interfaces

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	web3types "github.com/openweb3/web3go/types"
)

type Signer interface {
	Address() common.Address
	SignTransaction(tx *web3types.Transaction, chainID *big.Int) (*web3types.Transaction, error)
	SignMessage(text []byte) ([]byte, error)
	SignHash(hash []byte) ([]byte, error)
	SignSetCodeAuthorization(auth ethtypes.SetCodeAuthorization) (ethtypes.SetCodeAuthorization, error)
}
