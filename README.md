# web3go
The web3go's goal is to build an go sdk for supporting all ethereum rpc

## struct field type rule
1. For convinenet for developer, the web3go use standard type instead of hex types in rpc response, for example `*hexutil.Big` will be `*big.Int` in client interface. As well as the struct fields, but `marshal/unmarshal` result still be hex format.
2. The types of struct fields according to geth and parity and use the minimal type, such geth is `hexutil.Uint64` and parity is `*hexutil.Big`, then the filed type will be `uint64`
3. The slice item always be pointer if item is struct to avoid value copy when iteration
