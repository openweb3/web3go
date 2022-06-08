package interfaces

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
)

type Signer interface {
	Address() common.Address
	SignTransaction(*types.Transaction) (*types.Transaction, error)
	SignMessage(text []byte) ([]byte, error)
}
