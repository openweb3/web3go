# web3go
The web3go's goal is to build a golang SDK for supporting all Ethereum RPC

## Struct Fields Type Rule
For developer convenience, the web3go use standard type instead of hex types in RPC response, for example, `*hexutil`.Big` will be `*big.Int` in client interface. As well as the struct fields, but `marshal/unmarshal` result still be hex format.
1. The types of struct fields according to geth and parity and use the minimal type, such geth is `hexutil.Uint64` and parity is `*hexutil.Big`, then the filed type will be `uint64`
2. The slice item always be a pointer if the item is struct to avoid value copy when iteration

## Usage

### NewClient

The `NewClient` creates a client with default timeout options.

- the default timeout is 30 seconds 

```golang
    NewClient("http://localhost:8545")
```
### NewClientWithOption
Use `NewClientWithOption` to specify retry and timeout options

```golang
    NewClientWithOption("http://localhost:8545", providers.Option{...})
```

The provider of both clients created by `NewClient` and `NewClientWithOption` are [middlewarable providers](https://github.com/openweb3/go-rpc-provider).

Middlewarable providers could hook CallContext/BatchCallContext/Subscribe, such as log RPC request and response or cache environment variable in the context.

For custom pre/post Call/Batchcall behaviors, you can use HookCallContext of Middlewarable, such as log requests and so on, see more from [go-rpc-provider](https://github.com/openweb3/go-rpc-provider)
```golang
	p, e := providers.NewBaseProvider(context.Background(), "http://localhost:8545")
	if e != nil {
		return e
	}
	p.HookCallContext(callContextLogMiddleware)
	NewClientWithProvider(p)
```

the callLogMiddleware is like
```golang
	func callContextLogMiddleware(f providers.CallFunc) providers.CallFunc {
		return func(ctx context.Context, resultPtr interface{}, method string, args ...interface{}) error {
			fmt.Printf("request %v %v\n", method, args)
			err := f(ctx, resultPtr, method, args...)
			j, _ := json.Marshal(resultPtr)
			fmt.Printf("response %s\n", j)
			return err
		}
	}
```
### NewClientWithProvider

You also could set your customer provider by `NewClientWithProvider`


