package internal

import (
	"context"
	"time"

	rpc "github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/web3go/interfaces"
)

// TimeoutableProvider overwrite Call by CallContext with timeout context, make it to internal package to prevent external package to use it.
type TimeoutableProvider struct {
	Inner   interfaces.Provider
	timeout time.Duration
}

func NewTimeoutableProvider(inner interfaces.Provider, timeout time.Duration) *TimeoutableProvider {
	return &TimeoutableProvider{inner, timeout}
}

func (p *TimeoutableProvider) Call(resultPtr interface{}, method string, args ...interface{}) error {
	ctx, f := context.WithTimeout(context.Background(), p.timeout)
	defer f()
	return p.CallContext(ctx, resultPtr, method, args...)
}

func (p *TimeoutableProvider) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return p.Inner.CallContext(ctx, result, method, args...)
}

func (p *TimeoutableProvider) BatchCall(b []rpc.BatchElem) error {
	ctx, f := context.WithTimeout(context.Background(), p.timeout)
	defer f()
	return p.BatchCallContext(ctx, b)
}

func (p *TimeoutableProvider) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return p.Inner.BatchCallContext(ctx, b)
}

func (p *TimeoutableProvider) Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error) {
	return p.Inner.Subscribe(ctx, namespace, channel, args...)
}

func (p *TimeoutableProvider) Close() {
	p.Inner.Close()
}
