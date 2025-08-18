package enums

// SetAuthOutcome represents the outcome of the set auth action
type SetAuthOutcome string

const (
	SetAuthOutcomeSuccess              SetAuthOutcome = "success"
	SetAuthOutcomeInvalidChainId       SetAuthOutcome = "invalid_chain_id"
	SetAuthOutcomeNonceOverflow        SetAuthOutcome = "nonce_overflow"
	SetAuthOutcomeAccountCanNotSetAuth SetAuthOutcome = "account_can_not_set_auth"
	SetAuthOutcomeInvalidNonce         SetAuthOutcome = "invalid_nonce"
	SetAuthOutcomeInvalidSignature     SetAuthOutcome = "invalid_signature"
)
