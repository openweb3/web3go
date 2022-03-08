package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type CallType string
type RewardType string
type CreateType string

const (
	CALL_NONE         CallType = "none"
	CALL_CALL         CallType = "call"
	CALL_CALLCODE     CallType = "callCode"
	CALL_DELEGATECALL CallType = "delegateCall"
	CALL_STATICCALL   CallType = "staticCall"
)

const (
	CREATE_NONE    CreateType = "none"
	CREATE_CREATE  CreateType = "create"
	CREATE_CREATE2 CreateType = "create2"
)

const (
	REWARD_BLOCK     RewardType = "block"
	REWARD_UNCLE     RewardType = "uncle"
	REWARD_EMPTYSTEP RewardType = "emptyStep"
	REWARD_EXTERNAL  RewardType = "external"
)

//go:generate gencodec -type Call -field-override callMarshaling -out gen_call_json.go
type Call struct {
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Value    *big.Int       `json:"value"`
	Gas      *big.Int       `json:"gas"`
	Input    []byte         `json:"input"`
	CallType CallType       `json:"callType"`
}

//go:generate gencodec -type Create -field-override createMarshaling -out gen_create_json.go
type Create struct {
	From       common.Address `json:"from"`
	Value      *big.Int       `json:"value"`
	Gas        *big.Int       `json:"gas"`
	Init       []byte         `json:"init"`
	CreateType *CreateType    `json:"createType,omitempty"` // omit for openethereum, valid for conflux-espace
}

//go:generate gencodec -type Suicide -field-override suicideMarshaling -out gen_suicide_json.go
type Suicide struct {
	Address       common.Address `json:"address"`
	RefundAddress common.Address `json:"refundAddress"`
	Balance       *big.Int       `json:"balance"`
}

//go:generate gencodec -type Reward -field-override rewardMarshaling -out gen_reward_json.go
type Reward struct {
	Author     common.Address `json:"author"`
	Value      *big.Int       `json:"value"`
	RewardType RewardType     `json:"rewardType"`
}

type callMarshaling struct {
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Value    *hexutil.Big   `json:"value"`
	Gas      *hexutil.Big   `json:"gas"`
	Input    hexutil.Bytes  `json:"input"`
	CallType CallType       `json:"callType"`
}

type createMarshaling struct {
	From  common.Address `json:"from"`
	Value *hexutil.Big   `json:"value"`
	Gas   *hexutil.Big   `json:"gas"`
	Init  hexutil.Bytes  `json:"init"`
}

type suicideMarshaling struct {
	Address       common.Address `json:"address"`
	RefundAddress common.Address `json:"refundAddress"`
	Balance       *hexutil.Big   `json:"balance"`
}

type rewardMarshaling struct {
	Author     common.Address `json:"author"`
	Value      *hexutil.Big   `json:"value"`
	RewardType RewardType     `json:"rewardType"`
}
