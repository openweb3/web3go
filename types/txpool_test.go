package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTxPoolInspectJson(t *testing.T) {
	jstr := `{"pending":{"0x00000000863b56a3c1f0f1be8bc4f8b7bd78f57a":{"40":"contract creation: 0 wei + 612412 gas × 6000000000 wei"},"0x0512261a7486b1e29704ac49a5eb355b6fd86872":{"124930":"0x000000000000000000000000000000000000007E: 0 wei + 100187 gas × 20000000000 wei"},"0x201354729f8d0f8b64e9a0c353c672c6a66b3857":{"252350":"0xd10e3Be2bc8f959Bc8C41CF65F60dE721cF89ADF: 0 wei + 65792 gas × 2000000000 wei","252351":"0xd10e3Be2bc8f959Bc8C41CF65F60dE721cF89ADF: 0 wei + 65792 gas × 2000000000 wei","252352":"0xd10e3Be2bc8f959Bc8C41CF65F60dE721cF89ADF: 0 wei + 65780 gas × 2000000000 wei","252353":"0xd10e3Be2bc8f959Bc8C41CF65F60dE721cF89ADF: 0 wei + 65780 gas × 2000000000 wei"}},"queued":{"0x0f87ffcd71859233eb259f42b236c8e9873444e3":{"7":"0x3479BE69e07E838D9738a301Bb0c89e8EA2Bef4a: 1000000000000000 wei + 21000 gas × 10000000000 wei","8":"0x73Aaf691bc33fe38f86260338EF88f9897eCaa4F: 1000000000000000 wei + 21000 gas × 10000000000 wei"},"0x307e8f249bcccfa5b245449256c5d7e6e079943e":{"3":"0x73Aaf691bc33fe38f86260338EF88f9897eCaa4F: 10000000000000000 wei + 21000 gas × 10000000000 wei"}}}`
	var val TxpoolInspect
	err := json.Unmarshal([]byte(jstr), &val)
	if err != nil {
		t.Fatal(err)
	}

	j, _ := json.Marshal(val)
	assert.Equal(t, jstr, string(j))
}
