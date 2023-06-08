package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonLocalizedTrace(t *testing.T) {
	expect := `{"action":{"callType":"call","from":"0x1207bd45c1002dc88bf592ced9b35ec914bceb4e","gas":"0x43908","input":"0xa9059cbb000000000000000000000000a435a8e30e4c37a9d4480a42dda3955c22a76ed5000000000000000000000000000000000000000000000002b5e3af16b1880000","to":"0xfbef97434ffd0587e5a1c88efd5f7bdc405ba6fa","value":"0x0"},"blockHash":"0xc1836951b6504f293c97cd862a8208211a2b04ecf08220b9990355f00d6c8db5","blockNumber":"0x7805bc7","result":{"gasUsed":"0x37970","output":"0x0000000000000000000000000000000000000000000000000000000000000001"},"subtraces":"0x0","traceAddress":[],"transactionHash":"0xdd6aa06a4112d384d41665494381c98807387a3fa3aa3cde3a9992955a303b14","transactionPosition":"0x0","type":"call","valid":true}`
	var erp LocalizedTrace
	err := json.Unmarshal([]byte(expect), &erp)
	assert.NoError(t, err)

	actual, err := json.Marshal(erp)
	assert.NoError(t, err)

	fExpect, fActual := FormatJson(expect), FormatJson(string(actual))
	assert.Equal(t, fExpect, fActual)
}

// Format Json into string with orderd by field names
func FormatJson(input string) string {
	var oInput interface{}
	json.Unmarshal([]byte(input), &oInput)

	formated, _ := json.Marshal(oInput)
	return string(formated)
}
