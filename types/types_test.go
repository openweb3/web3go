package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gotest.tools/assert"
)

func TestUnMarshalCallRequest(t *testing.T) {
	goods := []CallRequest{
		{
			Data: nil,
		},
		{
			Input: nil,
		},
		{
			Data:  &hexutil.Bytes{0x1, 0x2, 0x3},
			Input: nil,
		},
		{
			Data:  nil,
			Input: &hexutil.Bytes{0x1, 0x2, 0x3},
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

		assert.Equal(t, item.Input, (*hexutil.Bytes)(nil))
	}

	bads := []CallRequest{
		{
			Data:  &hexutil.Bytes{0x1, 0x2},
			Input: &hexutil.Bytes{0x1, 0x2, 0x3},
		},
	}

	for _, item := range bads {
		_, e := json.Marshal(item)
		if e == nil {
			t.Fatalf("expected error, got nil")
		}
	}

}

func TestUnmarshalHexbytesInStruct(t *testing.T) {
	type sWithHexbytes struct {
		H1 hexutil.Bytes  `json:"h1"`
		H2 *hexutil.Bytes `json:"h2"`
	}

	s := sWithHexbytes{}
	b, e := json.Marshal(s)
	if e != nil {
		t.Fatal(e)
	}

	fmt.Printf("marshaled %s\n", b)

	e = json.Unmarshal(b, &s)
	if e != nil {
		t.Fatal(e)
	}
}
