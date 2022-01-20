package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestUnMarshalCallRequest(t *testing.T) {
	cr := CallRequest{
		Data: nil,
	}
	b, e := json.Marshal(cr)
	if e != nil {
		t.Fatal(e)
	}

	e = json.Unmarshal(b, &cr)
	if e != nil {
		t.Fatal(e)
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
