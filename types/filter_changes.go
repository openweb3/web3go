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
	if err := json.Unmarshal(input, &f.Logs); err == nil {
		return nil
	}
	if err := json.Unmarshal(input, &f.Hashes); err == nil {
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
