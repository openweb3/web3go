package providers

import (
	"context"
)

type SignableProvider struct {
	MiddlewarableProvider
}

func (p *SignableProvider) Call(resultPtr interface{}, method string, args ...interface{}) error {
	// hook eth_sendTransaction
	panic("not implemented") // TODO: Implement
}

func (p *SignableProvider) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	// hook eth_sendTransaction
	panic("not implemented") // TODO: Implement
}
