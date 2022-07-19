package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/pkg/errors"
)

// FilterQuery contains options for contract log filtering.
type FilterQuery struct {
	BlockHash *common.Hash     `json:"blockHash,omitempty"` // used by eth_getLogs, return logs only from block with this hash
	FromBlock *BlockNumber     `json:"fromBlock,omitempty"` // beginning of the queried range, nil means latest block
	ToBlock   *BlockNumber     `json:"toBlock,omitempty"`   // end of the range, nil means latest block
	Addresses []common.Address `json:"address,omitempty"`   // restricts matches to events created by specific contracts

	// The Topic list restricts matches to particular event topics. Each event has a list
	// of topics. Topics matches a prefix of that list. An empty element slice matches any
	// topic. Non-empty elements represent an alternative that matches any of the
	// contained topics.
	//
	// Examples:
	// {} or nil          matches any topic list
	// {{A}}              matches topic A in first position
	// {{}, {B}}          matches any topic in first position AND B in second position
	// {{A}, {B}}         matches topic A in first position AND B in second position
	// {{A, B}, {C, D}}   matches topic (A OR B) in first position AND (C OR D) in second position
	Topics [][]common.Hash `json:"topics,omitempty"`
}

func (args *FilterQuery) UnmarshalJSON(data []byte) error {
	var fc filters.FilterCriteria
	if err := json.Unmarshal(data, &fc); err != nil {
		return errors.Wrapf(err, "failed to unmarshal filter criteria")
	}

	args.BlockHash = fc.BlockHash
	args.FromBlock = BigIntToBlockNumber(fc.FromBlock)
	args.ToBlock = BigIntToBlockNumber(fc.ToBlock)
	args.Addresses = fc.Addresses
	args.Topics = fc.Topics
	return nil
}
