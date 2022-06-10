package providers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/go-rpc-provider"
	pinterfaces "github.com/openweb3/go-rpc-provider/interfaces"
	signers "github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"

	pproviders "github.com/openweb3/go-rpc-provider/provider_wrapper"
)

type SignableMiddleware struct {
	manager signers.SignerManager
}

var (
	ErrNoSigner error = errors.New("signer not found")
)

const (
	METHOD_SEND_TRANSACTION = "eth_sendTransaction"
)

func NewSignableProvider(p pinterfaces.Provider, signManager *signers.SignerManager) *pproviders.MiddlewarableProvider {
	mp := pproviders.NewMiddlewarableProvider(p)

	mid := &SignableMiddleware{
		manager: *signManager,
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
			method = "eth_sendRawTransaction"
		}
		return call(ctx, resultPtr, method, args...)
	}
}

func (s *SignableMiddleware) BatchCallContextMiddleware(batchCall pproviders.BatchCallContextFunc) pproviders.BatchCallContextFunc {
	return func(ctx context.Context, b []rpc.BatchElem) error {
		for i := range b {
			if b[i].Method == METHOD_SEND_TRANSACTION {

				if len(b[i].Args) == 0 {
					return errors.New("no args")
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
	m := map[string]interface{}{}

	// tx maybe a struct or a map, so we need to convert it to map[string]interface{}
	j, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(j, &m); err != nil {
		return nil, err
	}

	fmt.Printf("tx: %v\n", m)

	from := common.HexToAddress(m["from"].(string))
	chainId := m["chainId"].(string)

	signer, err := s.manager.Get(from)
	if err != nil {
		return nil, err
	}

	fmt.Printf("find signer of %v: %v\n", from, signer)

	if signer != nil {
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		tx2 := &types.Transaction{}
		json.Unmarshal(j, tx2)

		chainIdInBig, ok := new(big.Int).SetString(chainId, 0)
		if !ok {
			return nil, errors.New("invalid chainId")
		}

		tx2, err = signer.SignTransaction(tx2, chainIdInBig)
		if err != nil {
			return nil, err
		}
		rawTx, err := tx2.MarshalBinary()
		if err != nil {
			return nil, err
		}

		// tx2J, _ := json.Marshal(tx2)
		// fmt.Printf("signed tx: %s\n", tx2J)
		// fmt.Printf("signed raw tx: %x\n", rawTx)

		return rawTx, nil
	}
	return nil, ErrNoSigner
}
