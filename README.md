# web3go
The web3go's goal is to build an go sdk for supporting all ethereum rpc

## Struct Fields Type Rule
1. For convinenet for developer, the web3go use standard type instead of hex types in rpc response, for example `*hexutil.Big` will be `*big.Int` in client interface. As well as the struct fields, but `marshal/unmarshal` result still be hex format.
2. The types of struct fields according to geth and parity and use the minimal type, such geth is `hexutil.Uint64` and parity is `*hexutil.Big`, then the filed type will be `uint64`
3. The slice item always be pointer if item is struct to avoid value copy when iteration

## Usage
create simple client
```golang
    NewClient("http://localhost:8545")
```
create client with retry and timeout options, it will retry when call/batchcall failed and timeout in specific times
```golang
    NewClientWithOption("http://localhost:8545", nil)
```
create client with provider
```golang
    p, e := providers.NewBaseProvider(context.Background(), "http://localhost:8545")
	if e != nil {
		return e
	}
	NewClientWithProvider(p)
```
for custom pre/post call/batchcall behaviors, you can use MiddlewareProvider to hook on call/batchcall, such as log requests and so on
```golang
	p, e := providers.NewBaseProvider(context.Background(), "http://localhost:8545")
	if e != nil {
		return e
	}
	mp := providers.NewMiddlewarableProvider(p)
	mp.HookCall(callLogMiddleware)
	NewClientWithProvider(p)
```
the callLogMiddleware is like
```golang
func callLogMiddleware(f providers.CallFunc) providers.CallFunc {
	return func(resultPtr interface{}, method string, args ...interface{}) error {
		fmt.Printf("request %v %v\n", method, args)
		err := f(resultPtr, method, args...)
		j, _ := json.Marshal(resultPtr)
		fmt.Printf("response %s\n", j)
		return err
	}
}
```