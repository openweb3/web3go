module github.com/openweb3/web3go

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.15
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/mcuadros/go-defaults v1.2.0
	github.com/openweb3/go-rpc-provider v0.1.2
	github.com/openweb3/go-sdk-common v0.0.0-20220413032440-b5356d1d9613
	github.com/pkg/errors v0.9.1
	gotest.tools v2.2.0+incompatible
)

// replace github.com/openweb3/go-sdk-common => ../go-sdk-common
