package client

import (
	"context"

	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
)

type BaseClient struct {
	*providers.MiddlewarableProvider
	context context.Context
}

// WithContext creates a new Client with specified context
func (client *BaseClient) SetContext(ctx context.Context) {
	client.context = ctx
}

func (client *BaseClient) getContext() context.Context {
	if client.context == nil {
		return context.Background()
	}
	return client.context
}
