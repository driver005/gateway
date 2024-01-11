package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Money Amounts represents an amount that a given Product Variant can be purcased for. Each Money Amount either has a Currency or Region associated with it to indicate the pricing in a given Currency or, for fully region-based pricing, the given price in a specific Region. If region-based pricing is used the amount will be in the currency defined for the Reigon.
type MoneyAmount struct {
	core.Model

	// The amount in the smallest currecny unit (e.g. cents 100 cents to charge $1) that the Product Variant will cost.
	Amount float64 `json:"amount"`

	// Currency
	Currency *Currency `json:"currency" gorm:"foreignKey:code;references:currency_code"`

	// The 3 character currency code that the Money Amount is given in.
	CurrencyCode string `json:"currency_code"`

	// The maximum quantity that the Money Amount applies to. If this value is not set, the Money Amount applies to all quantities.
	MaxQuantity int `json:"max_quantity" gorm:"default:null"`

	// The minimum quantity that the Money Amount applies to. If this value is not set, the Money Amount applies to all quantities.
	MinQuantity int `json:"min_quantity" gorm:"default:null"`

	// Price Lists represents a set of prices that overrides the default price for one or more product variants.
	PriceList *PriceList `json:"price_list" gorm:"foreignKey:id;references:price_list_id"`

	// The ID of the price list associated with the money amount
	PriceListId uuid.NullUUID `json:"price_list_id" gorm:"default:null"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" gorm:"foreignKey:id;references:region_id"`

	// The region's ID
	RegionId uuid.NullUUID `json:"region_id" gorm:"default:null"`

	// The Product Variant contained in the Line Item. Available if the relation `variant` is expanded.
	Variant *ProductVariant `json:"variant" gorm:"foreignKey:id;references:variant_id"`

	// The id of the Product Variant contained in the Line Item.
	VariantId uuid.NullUUID `json:"variant_id" gorm:"default:null"`
}
