module github.com/openweb3/web3go

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.15
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/mcuadros/go-defaults v1.2.0
	github.com/openweb3/go-rpc-provider v0.3.1
	github.com/openweb3/go-sdk-common v0.0.0-20220720074746-a7134e1d372c
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
)

// replace github.com/openweb3/go-sdk-common => ../go-sdk-common
// replace github.com/openweb3/go-rpc-provider v0.2.2 => ../go-rpc-provider
// replace github.com/openweb3/go-sdk-common v0.0.0 => ../go-sdk-common
