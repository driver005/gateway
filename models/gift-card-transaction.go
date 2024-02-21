package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:GiftCardTransaction
// title: "Gift Card Transaction"
// description: "Gift Card Transactions are created once a Customer uses a Gift Card to pay for their Order."
// type: object
// required:
//   - amount
//   - created_at
//   - gift_card_id
//   - id
//   - is_taxable
//   - order_id
//   - tax_rate
//
// properties:
//
//	id:
//	  description: The gift card transaction's ID
//	  type: string
//	  example: gct_01G8X9A7ESKAJXG2H0E6F1MW7A
//	gift_card_id:
//	  description: The ID of the Gift Card that was used in the transaction.
//	  type: string
//	  example: gift_01G8XKBPBQY2R7RBET4J7E0XQZ
//	gift_card:
//	  description: The details of the gift card associated used in this transaction.
//	  x-expandable: "gift_card"
//	  nullable: true
//	  $ref: "#/components/schemas/GiftCard"
//	order_id:
//	  description: The ID of the order that the gift card was used for payment.
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order that the gift card was used for payment.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	amount:
//	  description: The amount that was used from the Gift Card.
//	  type: integer
//	  example: 10
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	is_taxable:
//	  description: Whether the transaction is taxable or not.
//	  nullable: true
//	  type: boolean
//	  example: false
//	tax_rate:
//	  description: The tax rate of the transaction
//	  nullable: true
//	  type: number
//	  example: 0
type GiftCardTransaction struct {
	core.SoftDeletableModel

	GiftCardId uuid.NullUUID `json:"gift_card_id" gorm:"column:gift_card_id"`
	GiftCard   *GiftCard     `json:"gift_card" gorm:"foreignKey:GiftCardId"`
	OrderId    uuid.NullUUID `json:"order_id" gorm:"column:order_id"`
	Order      *Order        `json:"order" gorm:"foreignKey:OrderId"`
	Amount     float64       `json:"amount" gorm:"column:amount"`
	IsTaxable  bool          `json:"is_taxable" gorm:"column:is_taxable"`
	TaxRate    float64       `json:"tax_rate" gorm:"column:tax_rate"`
}
