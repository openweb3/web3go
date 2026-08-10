package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAccessListResultJSON(t *testing.T) {
	input := `{"accessList":[{"address":"0x1111111111111111111111111111111111111111","storageKeys":["0x0000000000000000000000000000000000000000000000000000000000000001"]}],"gasUsed":"0x5208","error":"execution reverted"}`

	var result CreateAccessListResult
	require.NoError(t, json.Unmarshal([]byte(input), &result))
	require.Len(t, result.AccessList, 1)
	assert.Equal(t, common.HexToAddress("0x1111111111111111111111111111111111111111"), result.AccessList[0].Address)
	assert.Equal(t, big.NewInt(21000), result.GasUsed)
	require.NotNil(t, result.Error)
	assert.Equal(t, "execution reverted", *result.Error)

	encoded, err := json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, input, string(encoded))
}

func TestCreateAccessListResultJSONWithoutError(t *testing.T) {
	input := `{"accessList":[],"gasUsed":"0x0"}`

	var result CreateAccessListResult
	require.NoError(t, json.Unmarshal([]byte(input), &result))
	assert.Empty(t, result.AccessList)
	require.NotNil(t, result.GasUsed)
	assert.Zero(t, result.GasUsed.Sign())
	assert.Nil(t, result.Error)

	encoded, err := json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, input, string(encoded))
}
