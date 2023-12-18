package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// ShippingMethod - Shipping Methods represent a way in which an Order or Return can be shipped. Shipping Methods are built from a Shipping Option, but may contain additional details, that can be necessary for the Fulfillment Provider to handle the shipment.
type ShippingMethod struct {
	core.Model

	// The id of the Shipping Option that the Shipping Method is built from.
	ShippingOptionId uuid.NullUUID `json:"shipping_option_id"`

	ShippingOption *ShippingOption `json:"shipping_option" gorm:"foreignKey:id;references:shipping_option_id"`

	// The id of the Order that the Shipping Method is used on.
	OrderId uuid.NullUUID `json:"order_id" gorm:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// The id of the Return that the Shipping Method is used on.
	ReturnId uuid.NullUUID `json:"return_id" gorm:"default:null"`

	// A return object. Available if the relation `return_order` is expanded.
	ReturnOrder *Return `json:"return_order" gorm:"foreignKey:id;references:return_id"`

	// The id of the Swap that the Shipping Method is used on.
	SwapId uuid.NullUUID `json:"swap_id" gorm:"default:null"`

	// A swap object. Available if the relation `swap` is expanded.
	Swap *Swap `json:"swap" gorm:"foreignKey:id;references:swap_id"`

	// The id of the Cart that the Shipping Method is used on.
	CartId uuid.NullUUID `json:"cart_id" gorm:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart *Cart `json:"cart" gorm:"foreignKey:id;references:cart_id"`

	// The id of the Claim that the Shipping Method is used on.
	ClaimOrderId uuid.NullUUID `json:"claim_order_id" gorm:"default:null"`

	// A claim order object. Available if the relation `claim_order` is expanded.
	ClaimOrder *ClaimOrder `json:"claim_order" gorm:"foreignKey:id;references:claim_order_id"`

	// Available if the relation `tax_lines` is expanded.
	TaxLines []ShippingMethodTaxLine `json:"tax_lines" gorm:"foreignKey:id"`

	// The amount to charge for the Shipping Method. The currency of the price is defined by the Region that the Order that the Shipping Method belongs to is a part of.
	Price float64 `json:"price"`

	// Additional data that the Fulfillment Provider needs to fulfill the shipment. This is used in combination with the Shipping Options data, and may contain information such as a drop point id.
	Data JSONB `json:"data" gorm:"default:null"`

	// [EXPERIMENTAL] Indicates if the shipping method price include tax
	IncludesTax bool `json:"includes_tax" gorm:"default:null"`
}
