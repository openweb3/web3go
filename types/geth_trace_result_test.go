package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJsonGethTraceResult(t *testing.T) {
	input := `{"result":{"failed":false,"gas":7397755,"returnValue":"","structLogs":[]}}`

	var g GethTraceResult
	err := json.Unmarshal([]byte(input), &g)
	assert.NoError(t, err)

	j, _ := json.Marshal(g)
	fmt.Printf("%s\n", j)
}
