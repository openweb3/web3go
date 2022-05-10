package providers

import (
	"context"
	"net/url"

	rpc "github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/web3go/interfaces"
	"github.com/valyala/fasthttp"
)

// NewBaseProvider returns a new BaseProvider.
// maxConnsPerHost is the maximum number of concurrent connections, only works for http(s) protocal
func NewBaseProvider(ctx context.Context, nodeUrl string, maxConnectionNum ...int) (interfaces.Provider, error) {
	if len(maxConnectionNum) > 0 && maxConnectionNum[0] > 0 {
		if u, err := url.Parse(nodeUrl); err == nil {
			if u.Scheme == "http" || u.Scheme == "https" {
				fasthttpClient := new(fasthttp.Client)
				fasthttpClient.MaxConnsPerHost = maxConnectionNum[0]
				return rpc.DialHTTPWithClient(nodeUrl, fasthttpClient)
			}
		}
	}
	return rpc.DialContext(ctx, nodeUrl)
}
