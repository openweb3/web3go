package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethrpctypes "github.com/ethereum/go-ethereum/rpc"
	"gotest.tools/assert"
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
	latest := ethrpctypes.LatestBlockNumber

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
			in:  BlockNumberOrHash(ethrpctypes.BlockNumberOrHashWithNumber(10)),
			out: `"0xa"`,
		},
		{
			in:  BlockNumberOrHash(ethrpctypes.BlockNumberOrHashWithHash(common.HexToHash("0xae7fbe443ce1e7b7ad867e246f79f4ea316fbcc545f1e6238bbfa697d623b6b9"), true)),
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
		assert.DeepEqual(t, string(j), v.out)
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
