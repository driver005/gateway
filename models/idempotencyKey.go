package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

// @oas:schema:IdempotencyKey
// title: "Idempotency Key"
// description: "Idempotency Key is used to continue a process in case of any failure that might occur."
// type: object
// required:
//   - created_at
//   - id
//   - idempotency_key
//   - locked_at
//   - recovery_point
//   - response_code
//   - response_body
//   - request_method
//   - request_params
//   - request_path
//
// properties:
//
//	id:
//	  description: The idempotency key's ID
//	  type: string
//	  example: ikey_01G8X9A7ESKAJXG2H0E6F1MW7A
//	idempotency_key:
//	  description: The unique randomly generated key used to determine the state of a process.
//	  type: string
//	  externalDocs:
//	    url: https://docs.medusajs.com/development/idempotency-key/overview.md
//	    description: Learn more how to use the idempotency key.
//	created_at:
//	  description: Date which the idempotency key was locked.
//	  type: string
//	  format: date-time
//	locked_at:
//	  description: Date which the idempotency key was locked.
//	  nullable: true
//	  type: string
//	  format: date-time
//	request_method:
//	  description: The method of the request
//	  nullable: true
//	  type: string
//	  example: POST
//	request_params:
//	  description: The parameters passed to the request
//	  nullable: true
//	  type: object
//	  example:
//	    id: cart_01G8ZH853Y6TFXWPG5EYE81X63
//	request_path:
//	  description: The request's path
//	  nullable: true
//	  type: string
//	  example: /store/carts/cart_01G8ZH853Y6TFXWPG5EYE81X63/complete
//	response_code:
//	  description: The response's code.
//	  nullable: true
//	  type: string
//	  example: 200
//	response_body:
//	  description: The response's body
//	  nullable: true
//	  type: object
//	  example:
//	    id: cart_01G8ZH853Y6TFXWPG5EYE81X63
//	recovery_point:
//	  description: Where to continue from.
//	  type: string
//	  default: started
type IdempotencyKey struct {
	core.Model

	IdempotencyKey string     `json:"idempotency_key" gorm:"column:idempotency_key"`
	LockedAt       *time.Time `json:"locked_at" gorm:"column:locked_at"`
	RequestMethod  string     `json:"request_method" gorm:"column:request_method"`
	RequestParams  core.JSONB `json:"request_params" gorm:"column:request_params"`
	RequestPath    string     `json:"request_path" gorm:"column:request_path"`
	ResponseCode   int        `json:"response_code" gorm:"column:response_code"`
	ResponseBody   core.JSONB `json:"response_body" gorm:"column:response_body"`
	RecoveryPoint  string     `json:"recovery_point" gorm:"column:recovery_point;default:'started'"`
}
