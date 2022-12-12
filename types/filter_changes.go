package types

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type FilterChanges struct {
	Logs   []Log
	Hashes []common.Hash
}

func (f *FilterChanges) UnmarshalJSON(input []byte) error {
	logs := []Log{}
	if err := json.Unmarshal(input, &logs); err == nil {
		f.Logs = logs
		return nil
	}

	hashes := []common.Hash{}
	if err := json.Unmarshal(input, &hashes); err == nil {
		f.Hashes = hashes
		return nil
	}

	return fmt.Errorf("failed to unmarshal filter changes by %x", input)
}

func (f *FilterChanges) MarshalJSON() ([]byte, error) {
	if f.Logs != nil {
		return json.Marshal(f.Logs)
	}

	if f.Hashes != nil {
		return json.Marshal(f.Hashes)
	}

	return json.Marshal(nil)
}
