package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go/types/enums"
)

// #[derive(Default, Serialize)]
// #[serde(rename_all = "camelCase")]
//
//	pub struct AccountPendingTransactions {
//	    pub pending_transactions: Vec<Transaction>,
//	    pub first_tx_status: Option<TransactionStatus>,
//	    pub pending_count: U64,
//	}

// #[derive(Serialize)]
// #[serde(rename_all = "camelCase")]
// pub enum TransactionStatus {
//     Packed,
//     Ready,
//     Pending(PendingReason),
// }

// #[derive(Serialize)]
// #[serde(rename_all = "camelCase")]
// pub enum PendingReason {
//     FutureNonce,
//     NotEnoughCash,
//     OldEpochHeight,
//     // The tx status in the pool is inaccurate due to chain switch or sponsor
//     // balance change. This tx will not be packed even if it should have
//     // been ready, and the user needs to send a new transaction to trigger
//     // the status change.
//     OutdatedStatus,
// }

type TransactionStatus struct {
	Status        enums.TransactionStatus
	PendingReason enums.PendingReason
}

func (t TransactionStatus) MarshalJSON() ([]byte, error) {
	if t.Status == enums.TransactionStatusPending {
		return json.Marshal(struct {
			Pending string `json:"pending"`
		}{
			Pending: string(t.PendingReason),
		})
	}
	return json.Marshal(t.Status)
}

func (t *TransactionStatus) UnmarshalJSON(data []byte) error {
	var rawStatus string
	if err := json.Unmarshal(data, &rawStatus); err == nil {
		t.Status = enums.TransactionStatus(rawStatus)
		return nil
	}

	var objStatus struct {
		Pending string `json:"pending"`
	}
	if err := json.Unmarshal(data, &objStatus); err != nil {
		return err
	}

	t.Status = enums.TransactionStatusPending
	t.PendingReason = enums.PendingReason(objStatus.Pending)
	return nil
}

//go:generate gencodec -type AccountPendingTransactions -field-override accountPendingTransactionsMarshaling -out gen_account_pending_transactions_json.go
type AccountPendingTransactions struct {
	PendingTransactions []Transaction      `json:"pendingTransactions"`
	FirstTxStatus       *TransactionStatus `json:"firstTxStatus,omitempty"`
	PendingCount        uint64             `json:"pendingCount"`
}

type accountPendingTransactionsMarshaling struct {
	PendingTransactions []Transaction      `json:"pendingTransactions"`
	FirstTxStatus       *TransactionStatus `json:"firstTxStatus,omitempty"`
	PendingCount        hexutil.Uint64     `json:"pendingCount"`
}
