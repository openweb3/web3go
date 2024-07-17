package types

import "encoding/json"

// GethDebugTracingOptions represents the tracing options for Geth debug tracing
type GethDebugTracingOptions struct {
	GethDefaultTracingOptions
	Tracer       string                 `json:"tracer,omitempty"`
	TracerConfig *GethDebugTracerConfig `json:"tracerConfig,omitempty"`
	Timeout      *string                `json:"timeout,omitempty"`
}

// GethDefaultTracingOptions represents the default tracing options for the struct logger
type GethDefaultTracingOptions struct {
	EnableMemory      *bool   `json:"enableMemory,omitempty"`
	DisableMemory     *bool   `json:"disableMemory,omitempty"`
	DisableStack      *bool   `json:"disableStack,omitempty"`
	DisableStorage    *bool   `json:"disableStorage,omitempty"`
	EnableReturnData  *bool   `json:"enableReturnData,omitempty"`
	DisableReturnData *bool   `json:"disableReturnData,omitempty"`
	Debug             *bool   `json:"debug,omitempty"`
	Limit             *uint64 `json:"limit,omitempty"`
}

// GethDebugTracerConfig is a wrapper around json.RawMessage for tracer configuration
type GethDebugTracerConfig struct {
	CallConfig     *CallConfig
	PreStateConfig *PreStateConfig
}

func (g GethDebugTracerConfig) MarshalJSON() ([]byte, error) {
	if g.CallConfig != nil {
		return json.Marshal(g.CallConfig)
	}
	return json.Marshal(g.PreStateConfig)
}

func (g *GethDebugTracerConfig) UnmarshalJSON(data []byte) error {
	temp := struct {
		*CallConfig
		*PreStateConfig
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*g = GethDebugTracerConfig{
		CallConfig:     temp.CallConfig,
		PreStateConfig: temp.PreStateConfig,
	}
	return nil
}

// CallConfig represents the configuration for the call tracer.
type CallConfig struct {
	// OnlyTopCall, when set to true, will only trace the primary (top-level) call and not any sub-calls.
	// It eliminates the additional processing for each call frame.
	OnlyTopCall *bool `json:"onlyTopCall,omitempty"`

	// WithLog, when set to true, will include the logs emitted by the call.
	WithLog *bool `json:"withLog,omitempty"`
}

type PreStateConfig struct {
	/// If `diffMode` is set to true, the response frame includes all the account and storage diffs
	/// for the transaction. If it's missing or set to false it only returns the accounts and
	/// storage necessary to execute the transaction.
	DiffMode *bool `json:"diffMode,omitempty"`
}
