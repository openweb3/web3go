package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJsonGethTraceResult(t *testing.T) {
	input := `{"result":{"failed":false,"gas":7397755,"returnValue":"","structLogs":[{"depth":1,"gas":8856328,"gasCost":3,"op":"PUSH4","pc":447,"stack":["0x16873099","0x16873099"]}]}}`

	var g GethTraceResult
	err := json.Unmarshal([]byte(input), &g)
	assert.NoError(t, err)

	j, _ := json.Marshal(g)
	assert.Equal(t, input, string(j))
}
