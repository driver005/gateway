package types

import "github.com/driver005/gateway/core"

type FilterableSwap struct {
	core.FilterModel

	IdempotencyKey string `json:"idempotency_key,omitempty" validate:"omitempty"`
}
