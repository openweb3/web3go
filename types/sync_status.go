package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type SyncStatus struct {
	IsSyncing bool
	SyncInfo  *SyncProgress
}

type SyncProgress struct {
	CurrentBlock  hexutil.Uint64 `json:"currentBlock"`  // Current block number where sync is at
	HighestBlock  hexutil.Uint64 `json:"highestBlock"`  // Highest alleged block number in the chain
	StartingBlock hexutil.Uint64 `json:"startingBlock"` // Block number where sync began
}

func (s *SyncStatus) UnmarshalJSON(data []byte) error {
	var isSyncing bool
	var e error
	if e = json.Unmarshal(data, &isSyncing); e == nil {
		s.IsSyncing = isSyncing
		return nil
	}

	var syncInfo SyncProgress
	if e = json.Unmarshal(data, &syncInfo); e == nil {
		s.IsSyncing = true
		s.SyncInfo = &syncInfo
	}
	return e
}

func (s SyncStatus) MarshalJSON() ([]byte, error) {
	if s.IsSyncing {
		return json.Marshal(s.SyncInfo)
	}
	return json.Marshal(s.IsSyncing)
}
