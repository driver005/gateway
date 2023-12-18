package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// A value given to a Product Variant's option set. Product Variant have a Product Option Value for each of the Product Options defined on the Product.
type ProductOptionValue struct {
	core.Model

	// Product Options define properties that may vary between different variants of a Product. Common Product Options are "Size" and "Color", but Medusa doesn't limit what Product Options that can be defined.
	Option *ProductOption `json:"option" gorm:"foreignKey:id;references:option_id"`

	// The ID of the Product Option that the Product Option Value is defined for.
	OptionId uuid.NullUUID `json:"option_id"`

	// The value that the Product Variant has defined for the specific Product Option (e.g. if the Product Option is "Size" this value could be "Small", "Medium" or "Large").
	Value string `json:"value"`

	// Product Variants represent a Product with a specific set of Product Option configurations. The maximum number of Product Variants that a Product can have is given by the number of available Product Option combinations.
	Variant *ProductVariant `json:"variant" gorm:"foreignKey:id;references:variant_id"`

	// The ID of the Product Variant that the Product Option Value is defined for.
	VariantId uuid.NullUUID `json:"variant_id"`
}
