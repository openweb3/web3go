package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"

	"github.com/openweb3/go-rpc-provider"
	"github.com/stretchr/testify/assert"
)

func TestUnMarshalCallRequest(t *testing.T) {
	goods := []CallRequest{
		{
			Data: nil,
		},
		{
			Data: []byte{0x1, 0x2, 0x3},
		},
		{
			Data: nil,
		},
	}

	for _, item := range goods {
		b, e := json.Marshal(item)
		if e != nil {
			t.Fatal(e)
		}
		fmt.Printf("marshaled %s\n", b)

		item = CallRequest{}
		e = json.Unmarshal(b, &item)
		if e != nil {
			t.Fatal(e)
		}
		fmt.Printf("unmarshaled %+v\n", item)
	}
}

func TestBlockNumberOrHashMarshal(t *testing.T) {
	latest := rpc.LatestBlockNumber

	table := []struct {
		in  BlockNumberOrHash
		out string
		err bool
	}{
		{
			in:  BlockNumberOrHash{BlockNumber: &latest},
			out: `"latest"`,
		},
		{
			in:  BlockNumberOrHash(rpc.BlockNumberOrHashWithNumber(10)),
			out: `"0xa"`,
		},
		{
			in:  BlockNumberOrHash(rpc.BlockNumberOrHashWithHash(common.HexToHash("0xae7fbe443ce1e7b7ad867e246f79f4ea316fbcc545f1e6238bbfa697d623b6b9"), true)),
			out: `{"blockHash":"0xae7fbe443ce1e7b7ad867e246f79f4ea316fbcc545f1e6238bbfa697d623b6b9","requireCanonical":true}`,
		},
	}

	for _, v := range table {
		j, e := json.Marshal(v.in)
		if v.err {
			if e == nil {
				t.Fatal("expect error, got nil")
			}
			continue
		}
		if e != nil {
			t.Fatal(e)
		}
		assert.EqualValues(t, string(j), v.out)
	}
}

func TestReceiptMarshal(t *testing.T) {
	fail := uint64(0)
	r := Receipt{
		Status: &fail,
	}

	j, _ := json.Marshal(r)
	fmt.Printf("%s\n", j)
}

func TestCallRequestMarshalUnmarshal(t *testing.T) {
	from := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	to := common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd")
	gas := uint64(21000)
	nonce := uint64(5)
	txType := uint64(2)

	t.Run("unmarshal from raw JSON with all fields", func(t *testing.T) {
		raw := `{
			"from": "0x1234567890abcdef1234567890abcdef12345678",
			"to": "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
			"gas": "0x5208",
			"gasPrice": "0x3b9aca00",
			"maxFeePerGas": "0x77359400",
			"maxPriorityFeePerGas": "0x3b9aca00",
			"value": "0xde0b6b3a7640000",
			"nonce": "0x5",
			"data": "0xa9059cbb",
			"input": "0xdeadbeef",
			"chainId": "0x1",
			"type": "0x2",
			"accessList": [{"address": "0x1234567890abcdef1234567890abcdef12345678", "storageKeys": ["0x0000000000000000000000000000000000000000000000000000000000000001"]}],
			"authorizationList": [{
				"chainId": "0x1",
				"address": "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
				"nonce": "0xa",
				"yParity": "0x1",
				"r": "0x1234",
				"s": "0x5678"
			}]
		}`

		var cr CallRequest
		err := json.Unmarshal([]byte(raw), &cr)
		assert.NoError(t, err)
		assert.Equal(t, &from, cr.From)
		assert.Equal(t, &to, cr.To)
		assert.Equal(t, &gas, cr.Gas)
		assert.Equal(t, big.NewInt(1000000000), cr.GasPrice)
		assert.Equal(t, big.NewInt(2000000000), cr.MaxFeePerGas)
		assert.Equal(t, big.NewInt(1000000000), cr.MaxPriorityFeePerGas)
		assert.Equal(t, new(big.Int).SetUint64(1000000000000000000), cr.Value)
		assert.Equal(t, &nonce, cr.Nonce)
		assert.Equal(t, []byte{0xa9, 0x05, 0x9c, 0xbb}, cr.Data)
		assert.Equal(t, []byte{0xde, 0xad, 0xbe, 0xef}, cr.Input)
		assert.Equal(t, big.NewInt(1), cr.ChainID)
		assert.Equal(t, &txType, cr.Type)
		assert.NotNil(t, cr.AccessList)
		assert.Equal(t, 1, len(*cr.AccessList))
		assert.Equal(t, 1, len(cr.AuthorizationList))
		assert.Equal(t, *uint256.NewInt(1), cr.AuthorizationList[0].ChainID)
		assert.Equal(t, to, cr.AuthorizationList[0].Address)
		assert.Equal(t, uint64(10), cr.AuthorizationList[0].Nonce)
		assert.Equal(t, uint8(1), cr.AuthorizationList[0].V)
		assert.Equal(t, *uint256.NewInt(0x1234), cr.AuthorizationList[0].R)
		assert.Equal(t, *uint256.NewInt(0x5678), cr.AuthorizationList[0].S)
	})

	t.Run("unmarshal minimal JSON", func(t *testing.T) {
		raw := `{"to": "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"}`

		var cr CallRequest
		err := json.Unmarshal([]byte(raw), &cr)
		assert.NoError(t, err)
		assert.Nil(t, cr.From)
		assert.Equal(t, &to, cr.To)
		assert.Nil(t, cr.Gas)
		assert.Nil(t, cr.GasPrice)
		assert.Nil(t, cr.Value)
		assert.Nil(t, cr.Data)
		assert.Nil(t, cr.Type)
	})

	t.Run("unmarshal empty object", func(t *testing.T) {
		var cr CallRequest
		err := json.Unmarshal([]byte(`{}`), &cr)
		assert.NoError(t, err)
		assert.Nil(t, cr.From)
		assert.Nil(t, cr.To)
		assert.Nil(t, cr.Gas)
	})

	t.Run("marshal then unmarshal round-trip", func(t *testing.T) {
		storageKey := common.HexToHash("0x01")
		accessList := ethtypes.AccessList{
			{Address: from, StorageKeys: []common.Hash{storageKey}},
		}
		authList := []ethtypes.SetCodeAuthorization{
			{
				ChainID: *uint256.NewInt(1),
				Address: to,
				Nonce:   10,
				V:       1,
				R:       *uint256.NewInt(0x1234),
				S:       *uint256.NewInt(0x5678),
			},
		}
		original := CallRequest{
			From:                 &from,
			To:                   &to,
			Gas:                  &gas,
			GasPrice:             big.NewInt(1000000000),
			MaxFeePerGas:         big.NewInt(2000000000),
			MaxPriorityFeePerGas: big.NewInt(1000000000),
			Value:                big.NewInt(1e18),
			Nonce:                &nonce,
			Data:                 []byte{0xa9, 0x05, 0x9c, 0xbb},
			Input:                []byte{0xde, 0xad, 0xbe, 0xef},
			AccessList:           &accessList,
			AuthorizationList:    authList,
			ChainID:              big.NewInt(1),
			Type:                 &txType,
		}

		j, err := json.Marshal(original)
		assert.NoError(t, err)

		var decoded CallRequest
		err = json.Unmarshal(j, &decoded)
		assert.NoError(t, err)

		assert.Equal(t, original.From, decoded.From)
		assert.Equal(t, original.To, decoded.To)
		assert.Equal(t, original.Gas, decoded.Gas)
		assert.Equal(t, original.GasPrice.Cmp(decoded.GasPrice), 0)
		assert.Equal(t, original.MaxFeePerGas.Cmp(decoded.MaxFeePerGas), 0)
		assert.Equal(t, original.MaxPriorityFeePerGas.Cmp(decoded.MaxPriorityFeePerGas), 0)
		assert.Equal(t, original.Value.Cmp(decoded.Value), 0)
		assert.Equal(t, original.Nonce, decoded.Nonce)
		assert.Equal(t, original.Data, decoded.Data)
		assert.Equal(t, original.Input, decoded.Input)
		assert.Equal(t, original.ChainID.Cmp(decoded.ChainID), 0)
		assert.Equal(t, original.Type, decoded.Type)
		assert.Equal(t, *original.AccessList, *decoded.AccessList)
		assert.Equal(t, original.AuthorizationList, decoded.AuthorizationList)
	})

	t.Run("marshal omits nil/empty fields", func(t *testing.T) {
		cr := CallRequest{To: &to}
		j, err := json.Marshal(cr)
		assert.NoError(t, err)

		var raw map[string]json.RawMessage
		err = json.Unmarshal(j, &raw)
		assert.NoError(t, err)
		assert.Contains(t, raw, "to")
		assert.NotContains(t, raw, "from")
		assert.NotContains(t, raw, "gas")
		assert.NotContains(t, raw, "gasPrice")
		assert.NotContains(t, raw, "value")
		assert.NotContains(t, raw, "data")
		assert.NotContains(t, raw, "nonce")
		assert.NotContains(t, raw, "type")
		assert.NotContains(t, raw, "chainId")
		assert.NotContains(t, raw, "accessList")
		assert.NotContains(t, raw, "authorizationList")
	})

	t.Run("unmarshal EIP-7702 authorizationList with multiple entries", func(t *testing.T) {
		delegateTo := common.HexToAddress("0x1111111111111111111111111111111111111111")
		txType7702 := uint64(4)
		raw := `{
			"from": "0x1234567890abcdef1234567890abcdef12345678",
			"to": "0x1111111111111111111111111111111111111111",
			"type": "0x4",
			"authorizationList": [
				{
					"chainId": "0x1",
					"address": "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
					"nonce": "0x0",
					"yParity": "0x0",
					"r": "0xaaaa",
					"s": "0xbbbb"
				},
				{
					"chainId": "0xaa36a7",
					"address": "0x1111111111111111111111111111111111111111",
					"nonce": "0x1",
					"yParity": "0x1",
					"r": "0xcccc",
					"s": "0xdddd"
				}
			]
		}`

		var cr CallRequest
		err := json.Unmarshal([]byte(raw), &cr)
		assert.NoError(t, err)
		assert.Equal(t, &from, cr.From)
		assert.Equal(t, &delegateTo, cr.To)
		assert.Equal(t, &txType7702, cr.Type)
		assert.Nil(t, cr.AccessList)

		assert.Equal(t, 2, len(cr.AuthorizationList))

		auth0 := cr.AuthorizationList[0]
		assert.Equal(t, *uint256.NewInt(1), auth0.ChainID)
		assert.Equal(t, to, auth0.Address)
		assert.Equal(t, uint64(0), auth0.Nonce)
		assert.Equal(t, uint8(0), auth0.V)
		assert.Equal(t, *uint256.NewInt(0xaaaa), auth0.R)
		assert.Equal(t, *uint256.NewInt(0xbbbb), auth0.S)

		auth1 := cr.AuthorizationList[1]
		assert.Equal(t, *uint256.NewInt(0xaa36a7), auth1.ChainID)
		assert.Equal(t, delegateTo, auth1.Address)
		assert.Equal(t, uint64(1), auth1.Nonce)
		assert.Equal(t, uint8(1), auth1.V)
		assert.Equal(t, *uint256.NewInt(0xcccc), auth1.R)
		assert.Equal(t, *uint256.NewInt(0xdddd), auth1.S)
	})
}

func TestUnmarshalFilterChanges(t *testing.T) {
	var f FilterChanges

	err := json.Unmarshal([]byte(`[
		"0x3236cdebd39cf8470de82dc7b35bab55e29622dd03a941fbaa9de291b8fa6787",
		"0xcc9a4e25d1164841ae8ed2424cb2a007b6cb67feb3723366a0459c046c8e132f",
		"0xe9f3bd4466b8b8625f76e4adde77aef7cc369086cb325147dff42a5cd8b04a60",
		"0x026a7262eaeb9cf63ee4fc51b40b95e49948f791a416245507069607d6882147"
	]`), &f)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(f.Logs))
	assert.Equal(t, 4, len(f.Hashes))

	f = FilterChanges{}
	logs := []Log{
		{
			Address: common.HexToAddress("0x09b5928d6ab3381c7d090b6fbe528db136e0bea3"),
		},
	}
	j, _ := json.Marshal(logs)
	err = json.Unmarshal(j, &f)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(f.Logs))
	assert.Equal(t, 0, len(f.Hashes))
}

func TestUnmarshalRpcID(t *testing.T) {
	j, _ := json.Marshal("0x39")
	var val *rpc.ID
	err := json.Unmarshal(j, &val)
	assert.NoError(t, err)
	assert.Equal(t, *val, rpc.ID("0x39"))
}

func TestUnmarshalFeeHistory(t *testing.T) {
	expect := `{"oldestBlock":"0x1302340","reward":[["0x989680","0x77359400"],["0x9402a0","0x51a875fa"],["0x55d4a80","0xa9eaec6d"]],"baseFeePerGas":["0x1d31b3ab6","0x1dadaf3bc","0x1e975439b","0x1b0249037"],"gasUsedRatio":[0.5663567666666667,0.6230082666666666,0.03160246666666667]}`

	var feeHistory FeeHistory
	err := json.Unmarshal([]byte(expect), &feeHistory)
	assert.NoError(t, err)

	j, _ := json.Marshal(feeHistory)
	assert.Equal(t, expect, string(j))
}
