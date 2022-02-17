package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// TODO: missing types:
// LocalTransactionStatus
// RichHeader
// RichBlock
type LocalTransactionStatus interface{}
type RichHeader interface{}
type RichBlock interface{}

// TODO: struct is better
type SemverVersion string

// FIXME: need to confirm if the fields be camleCase when json marshal/unmarshal
type ClientVersion struct {
	ParityClient        ParityClientData `json:"ParityClient,omitempty"`
	ParityUnknownFormat string           `json:"ParityUnknownFormat,omitempty"`
	Other               string           `json:"Other,omitempty"`
}

type ParityClientData struct {
	Name                   string        `json:"name"`
	Identity               *string       `json:"identity"`
	Semver                 SemverVersion `json:"semver"`
	Os                     string        `json:"os"`
	Compiler               string        `json:"compiler"`
	CanHandleLargeRequests bool          `json:"canHandleLargeRequests"`
}

type Peers struct {
	Active    uint       `json:"active"`
	Connected uint       `json:"connected"`
	Max       uint32     `json:"max"`
	Peers     []PeerInfo `json:"peers"`
}
type PeerInfo struct {
	Id        *string           `json:"id"`
	Name      ClientVersion     `json:"name"`
	Caps      []string          `json:"caps"`
	Network   PeerNetworkInfo   `json:"network"`
	Protocols PeerProtocolsInfo `json:"protocols"`
}
type PeerNetworkInfo struct {
	RemoteAddress string `json:"remoteAddress"`
	LocalAddress  string `json:"localAddress"`
}
type PeerProtocolsInfo struct {
	Eth *EthProtocolInfo `json:"eth"`
}

//go:generate gencodec -type EthProtocolInfo -field-override ethProtocolInfoMarshaling -out gen_eth_protocol_info_json.go
type EthProtocolInfo struct {
	Version    uint32   `json:"version"`
	Difficulty *big.Int `json:"difficulty"`
	Head       string   `json:"head"`
}

type ethProtocolInfoMarshaling struct {
	Version    uint32       `json:"version"`
	Difficulty *hexutil.Big `json:"difficulty"`
	Head       string       `json:"head"`
}

type RpcSettings struct {
	Enabled   bool   `json:"enabled"`
	Interface string `json:"interface"`
	Port      uint64 `json:"port"`
}

//go:generate gencodec -type Histogram -field-override histogramMarshaling -out gen_histogram_json.go
type Histogram struct {
	BucketBounds []*big.Int `json:"bucketBounds"`
	Counts       []uint     `json:"counts"`
}

type histogramMarshaling struct {
	BucketBounds []*hexutil.Big `json:"bucketBounds"`
	Counts       []uint         `json:"counts"`
}

type TransactionStats struct {
	FirstSeen    uint64          `json:"firstSeen"`
	PropagatedTo map[string]uint `json:"propagatedTo"`
}

type ChainStatus struct {
	BlockGap BlockGap `json:"blockGap"`
}

type BlockGap struct {
	First *big.Int
	Last  *big.Int
}

func (b BlockGap) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		(*hexutil.Big)(b.First),
		(*hexutil.Big)(b.Last),
	})
}

func (b *BlockGap) UnmarshalJSON(data []byte) error {
	var fields [2]*hexutil.Big
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	b.First = (*big.Int)(fields[0])
	b.Last = (*big.Int)(fields[1])
	return nil
}

type NodeKind struct {
	Capability   string `json:"capability"`
	Availability string `json:"availability"`
}

type RecoveredAccount struct {
	Address                common.Address `json:"address"`
	PublicKey              string         `json:"publicKey"`
	IsValidForCurrentChain bool           `json:"isValidForCurrentChain"`
}
