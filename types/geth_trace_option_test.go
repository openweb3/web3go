package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJSONGethDebugTracerConfig(t *testing.T) {
	var g GethDebugTracerConfig
	err := json.Unmarshal([]byte("null"), &g)
	assert.NoError(t, err)

	err = json.Unmarshal([]byte(`{"diffMode":true}`), &g)
	assert.NoError(t, err)

	j, _ := json.Marshal(g)
	fmt.Printf("g %s\n", j)

	err = json.Unmarshal([]byte(`{"onlyTopCall":true}`), &g)
	assert.NoError(t, err)

	j, _ = json.Marshal(g)
	fmt.Printf("g %s\n", j)
}
