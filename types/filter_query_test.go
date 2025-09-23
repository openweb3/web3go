package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterQueryMarshalUnmarshal(t *testing.T) {
	// 准备测试 JSON 数据
	jsonStr := `{"fromBlock":"earliest","toBlock":"latest"}`

	// Unmarshal 步骤
	var filterQuery FilterQuery
	err := json.Unmarshal([]byte(jsonStr), &filterQuery)
	assert.NoError(t, err)

	fmt.Printf("unmarshaled filterQuery: %+v\n", filterQuery)

	// Marshal 步骤
	j, err := json.Marshal(filterQuery)
	assert.NoError(t, err)
	fmt.Printf("marshaled filterQuery: %s\n", j)

	// 验证 marshal 结果是否符合预期
	assert.JSONEq(t, jsonStr, string(j))
}
