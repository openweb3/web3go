package providers

import (
	"context"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/openweb3/web3go/interfaces"
	"github.com/pkg/errors"
)

type RetriableProvider struct {
	MiddlewarableProvider

	maxRetry int
	interval time.Duration
}

func NewRetriableProvider(inner interfaces.Provider, maxRetry int, interval time.Duration) *RetriableProvider {
	m := NewMiddlewarableProvider(inner)

	r := &RetriableProvider{*m, maxRetry, interval}
	r.HookCall(r.callMiddleware)
	r.HookCallContext(r.callContextMiddleware)
	r.HookBatchCall(r.batchCallMiddleware)
	r.HookBatchCallContext(r.batchCallContextMiddleware)
	return r
}

func (r *RetriableProvider) callMiddleware(call CallFunc) CallFunc {
	return func(resultPtr interface{}, method string, args ...interface{}) error {
		handler := func() error {
			return call(resultPtr, method, args...)
		}
		return retry(r.maxRetry, r.interval, handler)
	}
}

func (r *RetriableProvider) callContextMiddleware(call CallContextFunc) CallContextFunc {
	return func(ctx context.Context, resultPtr interface{}, method string, args ...interface{}) error {
		handler := func() error {
			return call(ctx, resultPtr, method, args...)
		}
		return retry(r.maxRetry, r.interval, handler)
	}
}

func (r *RetriableProvider) batchCallMiddleware(call BatchCallFunc) BatchCallFunc {
	return func(b []rpc.BatchElem) error {
		handler := func() error {
			return call(b)
		}
		return retry(r.maxRetry, r.interval, handler)
	}
}

func (r *RetriableProvider) batchCallContextMiddleware(call BatchCallContextFunc) BatchCallContextFunc {
	return func(ctx context.Context, b []rpc.BatchElem) error {
		handler := func() error {
			return call(ctx, b)
		}
		return retry(r.maxRetry, r.interval, handler)
	}
}

func retry(maxRetry int, interval time.Duration, handler func() error) error {
	remain := maxRetry
	for {
		err := handler()
		if err == nil {
			return nil
		}

		if utils.IsRPCJSONError(err) {
			return err
		}

		remain--
		if remain == 0 {
			return errors.Wrapf(err, "failed after %v retries", maxRetry)
		}

		if interval > 0 {
			time.Sleep(interval)
		}
	}
}
