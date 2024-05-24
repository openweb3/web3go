package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

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
