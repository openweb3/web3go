package interfaces

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	web3types "github.com/openweb3/web3go/types"
)

type Signer interface {
	// Address returns the signer account address.
	Address() common.Address

	// SignTransaction signs the transaction with the specified chain ID.
	SignTransaction(tx *web3types.Transaction, chainID *big.Int) (*web3types.Transaction, error)

	// SignMessage signs a message using the Ethereum signed message prefix (EIP-191).
	SignMessage(text []byte) ([]byte, error)

	// SignHash signs a precomputed 32-byte hash directly without message prefixing.
	SignHash(hash common.Hash) ([]byte, error)

	// SignSetCodeAuthorization signs an EIP-7702 SetCode authorization tuple.
	SignSetCodeAuthorization(auth ethtypes.SetCodeAuthorization) (ethtypes.SetCodeAuthorization, error)
}
