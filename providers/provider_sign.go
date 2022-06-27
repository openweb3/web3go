package providers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/go-rpc-provider"
	pinterfaces "github.com/openweb3/go-rpc-provider/interfaces"
	"github.com/openweb3/web3go/interfaces"
	signers "github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"

	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
)

type SignableMiddleware struct {
	manager  signers.SignerManager
	provider pinterfaces.Provider
}

var (
	ErrNoSigner      error = errors.New("signer not found")
	ErrChainNotReady error = errors.New("chain is not ready")
	ErrNoTxArgs      error = errors.New("no transaction args")
)

const (
	METHOD_SEND_TRANSACTION     = "eth_sendTransaction"
	METHOD_SEND_RAW_TRANSACTION = "eth_sendRawTransaction"
	METHOD_CHAIN_ID             = "eth_chainId"
)

func NewSignableProvider(p pinterfaces.Provider, signManager *signers.SignerManager) *pproviders.MiddlewarableProvider {
	mp := pproviders.NewMiddlewarableProvider(p)

	mid := &SignableMiddleware{
		manager:  *signManager,
		provider: p,
	}
	mp.HookCallContext(mid.CallContextMiddleware)
	mp.HookBatchCallContext(mid.BatchCallContextMiddleware)

	return mp
}

func (s *SignableMiddleware) CallContextMiddleware(call pproviders.CallContextFunc) pproviders.CallContextFunc {
	return func(ctx context.Context, resultPtr interface{}, method string, args ...interface{}) error {
		if method == METHOD_SEND_TRANSACTION {
			rawTx, err := s.signTxAndEncode(args[0])
			if err != nil && err != ErrNoSigner {
				return err
			}
			args[0] = rawTx
			method = METHOD_SEND_RAW_TRANSACTION
		}
		return call(ctx, resultPtr, method, args...)
	}
}

func (s *SignableMiddleware) BatchCallContextMiddleware(batchCall pproviders.BatchCallContextFunc) pproviders.BatchCallContextFunc {
	return func(ctx context.Context, b []rpc.BatchElem) error {
		for i := range b {
			if b[i].Method == METHOD_SEND_TRANSACTION {

				if len(b[i].Args) == 0 {
					return ErrNoTxArgs
				}

				rawTx, err := s.signTxAndEncode(b[i].Args[0])
				if err != nil && err != ErrNoSigner {
					return err
				}
				b[i].Args[0] = rawTx
			}
		}
		return batchCall(ctx, b)
	}
}

func (s *SignableMiddleware) signTxAndEncode(tx interface{}) (hexutil.Bytes, error) {

	var txArgs types.TransactionArgs

	switch tx.(type) {
	case map[string]interface{}:
		j, err := json.Marshal(tx)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(j, &txArgs); err != nil {
			return nil, err
		}
	case types.TransactionArgs:
		txArgs = tx.(types.TransactionArgs)
	case *types.TransactionArgs:
		txArgs = *tx.(*types.TransactionArgs)
	}

	var signer interfaces.Signer
	if txArgs.From == nil {
		signers := s.manager.List()
		if len(signers) == 0 {
			return nil, ErrNoSigner
		}
		signer = signers[0]
	} else {
		var err error
		signer, err = s.manager.Get(*txArgs.From)
		if err != nil {
			return nil, err
		}
	}

	if signer == nil {
		return nil, ErrNoSigner
	}

	// get chainId from chain
	var chainId *hexutil.Big
	if err := s.provider.CallContext(context.Background(), &chainId, METHOD_CHAIN_ID); err != nil {
		return nil, err
	}

	if chainId == nil {
		return nil, ErrChainNotReady
	}

	tx2, err := txArgs.ToTransaction()
	if err != nil {
		return nil, err
	}

	tx2, err = signer.SignTransaction(tx2, chainId.ToInt())
	if err != nil {
		return nil, err
	}

	rawTx, err := tx2.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return rawTx, nil
}
