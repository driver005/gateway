package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

// IdempotencyKey - Idempotency Key is used to continue a process in case of any failure that might occur.
type IdempotencyKey struct {
	core.Model

	// The unique randomly generated key used to determine the state of a process.
	IdempotencyKey string `json:"idempotency_key"`

	// Date which the idempotency key was locked.
	LockedAt *time.Time `json:"locked_at" gorm:"default:null"`

	// The method of the request
	RequestMethod string `json:"request_method" gorm:"default:null"`

	// The parameters passed to the request
	RequestParams JSONB `json:"request_params" gorm:"default:null"`

	// The request's path
	RequestPath string `json:"request_path" gorm:"default:null"`

	// The response's code.
	ResponseCode string `json:"response_code" gorm:"default:null"`

	// The response's body
	ResponseBody JSONB `json:"response_body" gorm:"default:null"`

	// Where to continue from.
	RecoveryPoint string `json:"recovery_point" gorm:"default:null"`
}
