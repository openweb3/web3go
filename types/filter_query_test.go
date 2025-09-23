package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterQueryMarshalUnmarshal(t *testing.T) {
	// Prepare test JSON data
	jsonStr := `{"fromBlock":"earliest","toBlock":"latest"}`

	// Unmarshal step
	var filterQuery FilterQuery
	err := json.Unmarshal([]byte(jsonStr), &filterQuery)
	assert.NoError(t, err)

	fmt.Printf("unmarshaled filterQuery: %+v\n", filterQuery)

	// Marshal step
	j, err := json.Marshal(filterQuery)
	assert.NoError(t, err)
	fmt.Printf("marshaled filterQuery: %s\n", j)

	// Verify marshal result meets expectations
	assert.JSONEq(t, jsonStr, string(j))
}
