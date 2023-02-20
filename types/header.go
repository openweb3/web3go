package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

//go:generate gencodec -type Header -field-override headerMarshaling -out gen_header_json.go
type Header struct {
	ethtypes.Header
	HeaderExtra
}

type HeaderExtra struct {
	Author *common.Address `json:"author,omitempty"`
	Hash   *common.Hash    `json:"hash,omitempty"`
}

func (h *Header) UnmarshalJSON(data []byte) error {
	var he HeaderExtra
	if err := json.Unmarshal(data, &he); err != nil {
		return err
	}

	var _h ethtypes.Header
	if err := json.Unmarshal(data, &_h); err != nil {
		return err
	}

	if he.Hash != nil {
		h.Header = _h
		h.HeaderExtra = he
	} else {
		h.Header = _h
		h.HeaderExtra.Hash = he.Hash
	}

	return nil
}

func (h Header) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash  common.Hash         `json:"parentHash"       gencodec:"required"`
		UncleHash   common.Hash         `json:"sha3Uncles"       gencodec:"required"`
		Coinbase    common.Address      `json:"miner"            gencodec:"required"`
		Root        common.Hash         `json:"stateRoot"        gencodec:"required"`
		TxHash      common.Hash         `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash common.Hash         `json:"receiptsRoot"     gencodec:"required"`
		Bloom       ethtypes.Bloom      `json:"logsBloom"        gencodec:"required"`
		Difficulty  *hexutil.Big        `json:"difficulty"       gencodec:"required"`
		Number      *hexutil.Big        `json:"number"           gencodec:"required"`
		GasLimit    hexutil.Uint64      `json:"gasLimit"         gencodec:"required"`
		GasUsed     hexutil.Uint64      `json:"gasUsed"          gencodec:"required"`
		Time        hexutil.Uint64      `json:"timestamp"        gencodec:"required"`
		Extra       hexutil.Bytes       `json:"extraData"        gencodec:"required"`
		MixDigest   common.Hash         `json:"mixHash"`
		Nonce       ethtypes.BlockNonce `json:"nonce"`
		BaseFee     *hexutil.Big        `json:"baseFeePerGas" rlp:"optional"`
		Hash        *common.Hash        `json:"hash,omitempty"`
		Author      *common.Address     `json:"author,omitempty"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.UncleHash = h.UncleHash
	enc.Coinbase = h.Coinbase
	enc.Root = h.Root
	enc.TxHash = h.TxHash
	enc.ReceiptHash = h.ReceiptHash
	enc.Bloom = h.Bloom
	enc.Difficulty = (*hexutil.Big)(h.Difficulty)
	enc.Number = (*hexutil.Big)(h.Number)
	enc.GasLimit = hexutil.Uint64(h.GasLimit)
	enc.GasUsed = hexutil.Uint64(h.GasUsed)
	enc.Time = hexutil.Uint64(h.Time)
	enc.Extra = h.Extra
	enc.MixDigest = h.MixDigest
	enc.Nonce = h.Nonce
	enc.BaseFee = (*hexutil.Big)(h.BaseFee)
	enc.Author = h.Author

	if h.HeaderExtra.Hash != nil {
		enc.Hash = h.HeaderExtra.Hash
	} else {
		_hash := h.Header.Hash()
		enc.Hash = &_hash
	}

	return json.Marshal(&enc)
}
