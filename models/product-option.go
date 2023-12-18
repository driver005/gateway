package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Product Options define properties that may vary between different variants of a Product. Common Product Options are "Size" and "Color", but Medusa doesn't limit what Product Options that can be defined.
type ProductOption struct {
	core.Model

	// A product object. Available if the relation `product` is expanded.
	Product *Product `json:"product" gorm:"foreignKey:id;references:product_id"`

	// The ID of the Product that the Product Option is defined for.
	ProductId uuid.NullUUID `json:"product_id"`

	// The title that the Product Option is defined by (e.g. "Size").
	Title string `json:"title"`

	// The Product Option Values that are defined for the Product Option. Available if the relation `values` is expanded.
	Values []ProductOptionValue `json:"values" gorm:"foreignKey:id"`
}
