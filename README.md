# web3go
web3go's goal is to build an go sdk for supporting all ethereum rpc

# struct field type rule
1. set types fields according to geth and parity and use the minimal type, such geth is hexutil.Uint64 and parity is *hexutil.Big, then the filed type will be hexutil.Uint64
2. slice item always be pointer if item is struct