package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// GiftCardTransaction - Gift Card Transactions are created once a Customer uses a Gift Card to pay for their Order
type GiftCardTransaction struct {
	core.Model

	// The ID of the Gift Card that was used in the transaction.
	GiftCardId uuid.NullUUID `json:"gift_card_id"`

	// A gift card object. Available if the relation `gift_card` is expanded.
	GiftCard *GiftCard `json:"gift_card" gorm:"foreignKey:id;references:gift_card_id"`

	// The ID of the Order that the Gift Card was used to pay for.
	OrderId uuid.NullUUID `json:"order_id" gorm:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// The amount that was used from the Gift Card.
	Amount float64 `json:"amount"`

	// Whether the transaction is taxable or not.
	IsTaxable bool `json:"is_taxable" gorm:"default:null"`

	// The tax rate of the transaction
	TaxRate float64 `json:"tax_rate" gorm:"default:null"`
}
