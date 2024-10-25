package enums

type PendingReason string

const (
	PENDING_REASON_FUTURE_NONCE     PendingReason = "futureNonce"
	PENDING_REASON_NOT_ENOUGH_CASH  PendingReason = "notEnoughCash"
	PENDING_REASON_OLD_EPOCH_HEIGHT PendingReason = "oldEpochHeight"
	PENDING_REASON_OUTDATED_STATUS  PendingReason = "outdatedStatus"
)
