# web3go
The web3go's goal is to build a golang SDK for supporting all Ethereum RPC

## Struct Fields Type Rule
For developer convenience, the web3go use standard type instead of hex types in RPC response, for example, `*hexutil.Big` will be `*big.Int` in client interface. As well as the struct fields, but `marshal/unmarshal` result still be hex format.
1. The types of struct fields according to geth and parity and use the minimal type, such geth is `hexutil.Uint64` and parity is `*hexutil.Big`, then the filed type will be `uint64`
2. The slice item always be a pointer if the item is struct to avoid value copy when iteration

## Client

### NewClient

The `NewClient` creates a client with default timeout options.

- the default timeout is 30 seconds 

```golang
	NewClient("http://localhost:8545")
```
### NewClientWithOption
Use `NewClientWithOption` to specify retry, timeout and signer manager options, the signer manager option is used for signing transactions automatically when call `SendTransaction` or `SendTransactionByArgs`

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

## Sign

### Signer
There is an interface `signer` for signing transactions and messages.

- `SignTransaction` to sign a transaction
- `SignMessage` to sign a message

We provide kinds of functions to create a private key signer, which means it will convert anything input to the private key.

*create by private key*
- `NewPrivateKeySigner`

*create by private key string*
- `NewPrivateKeySignerByString`
- `MustNewPrivateKeySignerByString`

*create a random private key*
- `NewRandomPrivateKeySigner`
- `MustNewRandomPrivateKeySigner`

*create by mnemonic*
- `NewPrivateKeySignerByMnemonic`
- `MustNewPrivateKeySignerByMnemonic`

*create by Keystore*
- `NewPrivateKeySignerByKeystoreFile`
- `MustNewPrivateKeySignerByKeystoreFile`
- `NewPrivateKeySignerByKeystore`
- `MustNewPrivateKeySignerByKeystore`

### Signer Manager

Signer Manager is for manager signers conveniently, support get/add/remove/list signer.

And convenient functions to create signer managers, such as create by private key strings and mnemonic

- `NewSignerManager`
- `NewSignerManagerByPrivateKeyStrings`
- `MustNewSignerManagerByPrivateKeyStrings`
- `NewSignerManagerByMnemonic`
- `MustNewSignerManagerByMnemonic`

### Auto Sign

There are two ways to create a client that can be automatically signed when sending transactions.

- Firstly and the simple way is to create a client by `NewClientWithOption` and set the field `SignerManager` of `Option`

```go
	mnemonic := "crisp shove million stem shiver side hospital split play lottery join vintage"
	sm := signers.MustNewSignerManagerByMnemonic(mnemonic, 10, nil)
	option := new(ClientOption).WithLooger(os.Stdout).WithSignerManager(sm)
	c, err := NewClientWithOption("https://evm.confluxrpc.com", *option)

	from := sm.List()[0].Address()
	to := sm.List()[1].Address()
	hash, err := c.Eth.SendTransactionByArgs(types.TransactionArgs{
		From: &from,
		To:   &to,
	})
```

- Another way is to create a `MiddlewareableProvider` by [`NewSignableProvider`](https://github.com/openweb3/web3go/blob/main/providers/provider_sign.go)

```go
	mnemonic := "crisp shove million stem shiver side hospital split play lottery join vintage"
	sm := signers.MustNewSignerManagerByMnemonic(mnemonic, 10, nil)
	p := pproviders.MustNewBaseProvider(context.Background(), "https://evm.confluxrpc.com")
	p = providers.NewSignableProvider(p, sm)
	c := NewClientWithProvider(p)
```

#### Send Transaction

There are two ways to send transactions and auto-sign, `SendTransaction` and `SendTransactionByArgs`

- `SendTransaction` send by [transaction type of `go-ethereum`](https://github.com/openweb3/web3go/blob/08c2cb1790acbc92277f28b3d98e8b7347450cc5/types/types.go#L246) and needs to specify the sender
```golang
	mnemonic := "crisp shove million stem shiver side hospital split play lottery join vintage"
	sm := signers.MustNewSignerManagerByMnemonic(mnemonic, 10, nil)
	c := MustNewClientWithOption("https://ropsten.infura.io/v3/cb2c1b76cb894b699f20a602f35731f1", *(new(ClientOption).WithLoger(os.Stdout).WithSignerManager(sm)))

	// send legacy tx
	tx := ethrpctypes.NewTransaction(nonce.Uint64(), from, big.NewInt(1000000), 1000000, big.NewInt(1), nil)
	txhash, err := c.SendTransaction(from, *tx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("send legacy tx, hash: %s\n", txhash)

	// send dynamic fee tx
	dtx := &ethrpctypes.DynamicFeeTx{
		To:    &from,
		Value: big.NewInt(1),
	}
	txhash, err = c.SendTransaction(from, *ethrpctypes.NewTx(dtx))
	if err != nil {
		panic(err)
	}
	fmt.Printf("send dynamic fee tx, hash: %s\n", txhash)

```
- `SendTransactionByArgs` send by transaction type of `TransactionArgs` which contains all fields a transaction needs
```golang
	mnemonic := "crisp shove million stem shiver side hospital split play lottery join vintage"
	sm := signers.MustNewSignerManagerByMnemonic(mnemonic, 10, nil)
	c := MustNewClientWithOption("https://ropsten.infura.io/v3/cb2c1b76cb894b699f20a602f35731f1", *(new(ClientOption).WithLoger(os.Stdout).WithSignerManager(sm)))

	from := sm.List()[0].Address()
	to := sm.List()[1].Address()
	hash, err := c.Eth.SendTransactionByArgs(types.TransactionArgs{
		From: &from,
		To:   &to,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("hash: %s\n", hash)
```
If the provider of client contains the signer of the transaction's `From`, both of them will populate transaction fields, sign the transaction and call `eth_sendRawTransaction` to send RLP-Encoded transaction. Otherwise will call `eth_sendTransaction`.

## Contract

Invoke with contract please use [abigen](https://geth.ethereum.org/docs/dapp/native-bindings), we provide the methods `ToClientForContract` for generating `bind.ContractBackend` and `bind.SignerFn` for conveniently use in abi-binding struct which is generated by abigen

Please see the example from [example_abigen](https://github.com/openweb3/web3go-example/blob/master/example_abigen)