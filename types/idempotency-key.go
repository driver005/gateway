package types

import (
	"github.com/driver005/gateway/core"
)

type CreateIdempotencyKeyInput struct {
	RequestMethod  string     `json:"request_method,omitempty"`
	RequestParams  core.JSONB `json:"request_params,omitempty"`
	RequestPath    string     `json:"request_path,omitempty"`
	IdempotencyKey string     `json:"idempotency_key,omitempty"`
}

type IdempotencyCallbackResult struct {
	RecoveryPoint string     `json:"recovery_point,omitempty"`
	ResponseCode  int        `json:"response_code,omitempty"`
	ResponseBody  core.JSONB `json:"response_body,omitempty"`
}
