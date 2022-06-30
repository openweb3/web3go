module github.com/openweb3/web3go

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.15
	github.com/google/uuid v1.1.5
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/mcuadros/go-defaults v1.2.0
	github.com/miguelmota/go-ethereum-hdwallet v0.1.1
	github.com/openweb3/go-rpc-provider v0.2.2
	github.com/openweb3/go-sdk-common v0.0.0-20220630071703-d06847a2fb80
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/tyler-smith/go-bip39 v1.0.1-0.20181017060643-dbb3b84ba2ef
	gotest.tools v2.2.0+incompatible

)

// replace github.com/openweb3/go-sdk-common => ../go-sdk-common
// replace github.com/openweb3/go-rpc-provider v0.1.2 => ../go-rpc-provider
