package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJsonTrace(t *testing.T) {
	inputs := []string{
		`{"type":"call","action":{"from":"0x69a44e15f5718853e757866d000a98141d49da0d","to":"0x2c32e41753902cb99ae7107d238b9fde53075655","value":"0x0","gas":"0x2baf8","input":"0x91f87991","callType":"call"},"result":{"gasUsed":"0x1547","output":"0x08c379a0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000145245564552545f4552524f525f4d455353414745000000000000000000000000"},"error":"Reverted","traceAddress":[],"subtraces":0,"transactionPosition":0,"transactionHash":"0x924cf10d6b1ee67ab4d6b930d5d525c9e96eee26d064c8a12dd73f00e0241c32","blockNumber":246730020,"blockHash":"0x024eb5698c878cb765ec9c01084403e0c4ee4f88b5006841f2a889a3be12b939","valid":false}`,
		`{"type":"call","action":{"from":"0x69a44e15f5718853e757866d000a98141d49da0d","to":"0x2c32e41753902cb99ae7107d238b9fde53075655","value":"0x0","gas":"0x2baf8","input":"0x42358443","callType":"call"},"result":{"gasUsed":"0x2baf8","output":"0x4f75744f66476173"},"error":"OutOfGas","traceAddress":[],"subtraces":0,"transactionPosition":0,"transactionHash":"0x61b0b980edfe5df726fab8c056fc2bfc9014e115673f8ac9a836316ec0200e6c","blockNumber":246730050,"blockHash":"0xd9f8b00091612a3d0b6499f60f7fd650f066d0235af4adddbc84eb4d01da8130","valid":false}`,
	}

	for _, input := range inputs {
		var val LocalizedTrace
		err := json.Unmarshal([]byte(input), &val)
		assert.NoError(t, err)

		j, _ := json.Marshal(val)
		assert.Equal(t, input, string(j))
	}
}
