package enums

type TransactionStatus string

const (
	TransactionStatusPacked  TransactionStatus = "packed"
	TransactionStatusReady   TransactionStatus = "ready"
	TransactionStatusPending TransactionStatus = "pending"
)
