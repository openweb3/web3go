package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
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
			Data:  []byte{0x1, 0x2, 0x3},
			Input: nil,
		},
		{
			Data:  nil,
			Input: []byte{0x1, 0x2, 0x3},
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

		// assert.Equal(t, item.Input, ([]byte)(nil))
		if !bytes.Equal(item.Input, ([]byte)(nil)) {
			t.Fatal("item.Input not nil")
		}
	}

	bads := []CallRequest{
		{
			Data:  []byte{0x1, 0x2},
			Input: []byte{0x1, 0x2, 0x3},
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

func TestMarshalSlice(t *testing.T) {
	var a []string
	fmt.Printf("%v\n", reflect.TypeOf(a).Kind())
	j, e := json.Marshal(a)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Print(string(j))

	type sWithSlice struct {
		S []string `json:"s"`
		B []byte   `json:"b"`
	}

	j, e = json.Marshal(sWithSlice{})
	if e != nil {
		t.Fatal(e)
	}
	fmt.Print(string(j))
}
