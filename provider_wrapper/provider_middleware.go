package providers

import (
	"context"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/openweb3/web3go/interfaces"
)

type MiddlewarableProvider struct {
	Inner                   interfaces.Provider
	callMiddles             []CallMiddleware
	callContextMiddles      []CallContextMiddleware
	batchCallMiddles        []BatchCallMiddleware
	batchCallContextMiddles []BatchCallContextMiddleware
	// subscribeMiddlewares    []SubscribeMiddleware
}

func NewMiddlewarableProvider(p interfaces.Provider) *MiddlewarableProvider {
	return &MiddlewarableProvider{Inner: p}
}

type CallFunc func(resultPtr interface{}, method string, args ...interface{}) error
type CallContextFunc func(ctx context.Context, result interface{}, method string, args ...interface{}) error
type BatchCallFunc func(b []rpc.BatchElem) error
type BatchCallContextFunc func(ctx context.Context, b []rpc.BatchElem) error

// type SubscribeFunc func(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error)

type CallMiddleware func(CallFunc) CallFunc
type CallContextMiddleware func(CallContextFunc) CallContextFunc
type BatchCallMiddleware func(BatchCallFunc) BatchCallFunc
type BatchCallContextMiddleware func(BatchCallContextFunc) BatchCallContextFunc

// type SubscribeMiddleware func(SubscribeFunc) SubscribeFunc

func (mp *MiddlewarableProvider) HookCall(cm CallMiddleware) {
	mp.callMiddles = append(mp.callMiddles, cm)
}

func (mp *MiddlewarableProvider) HookCallContext(cm CallContextMiddleware) {
	mp.callContextMiddles = append(mp.callContextMiddles, cm)
}

func (mp *MiddlewarableProvider) HookBatchCall(cm BatchCallMiddleware) {
	mp.batchCallMiddles = append(mp.batchCallMiddles, cm)
}

func (mp *MiddlewarableProvider) HookBatchCallContext(cm BatchCallContextMiddleware) {
	mp.batchCallContextMiddles = append(mp.batchCallContextMiddles, cm)
}

func (mp *MiddlewarableProvider) Call(resultPtr interface{}, method string, args ...interface{}) error {
	// callMiddles: A -> B -> C, execute order is A -> B -> C
	var nestedWare CallFunc = mp.Inner.Call
	for i := len(mp.callMiddles) - 1; i >= 0; i-- {
		nestedWare = mp.callMiddles[i](nestedWare)
	}

	return nestedWare(resultPtr, method, args...)
}

func (mp *MiddlewarableProvider) CallContext(ctx context.Context, resultPtr interface{}, method string, args ...interface{}) error {
	var nestedWare CallContextFunc = mp.Inner.CallContext
	for i := len(mp.callContextMiddles) - 1; i >= 0; i-- {
		nestedWare = mp.callContextMiddles[i](nestedWare)
	}

	return nestedWare(ctx, resultPtr, method, args...)
}

func (mp *MiddlewarableProvider) BatchCall(b []rpc.BatchElem) error {
	var nestedWare BatchCallFunc = mp.Inner.BatchCall
	for i := len(mp.batchCallMiddles) - 1; i >= 0; i-- {
		nestedWare = mp.batchCallMiddles[i](nestedWare)
	}

	return nestedWare(b)
}

func (mp *MiddlewarableProvider) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	var nestedWare BatchCallContextFunc = mp.Inner.BatchCallContext
	for i := len(mp.batchCallContextMiddles) - 1; i >= 0; i-- {
		nestedWare = mp.batchCallContextMiddles[i](nestedWare)
	}

	return nestedWare(ctx, b)
}

func (mp *MiddlewarableProvider) Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error) {
	return mp.Inner.Subscribe(ctx, namespace, channel, args...)
}

func (mp *MiddlewarableProvider) Close() {
	mp.Inner.Close()
}
