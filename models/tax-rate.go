package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// TaxRate - A Tax Rate can be used to associate a certain rate to charge on products within a given Region
type TaxRate struct {
	core.Model

	// The numeric rate to charge
	Rate float64 `json:"rate" gorm:"default:null"`

	// A code to identify the tax type by
	Code string `json:"code" gorm:"default:null"`

	// A human friendly name for the tax
	Name string `json:"name"`

	// The id of the Region that the rate belongs to
	RegionId uuid.NullUUID `json:"region_id"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" gorm:"foreignKey:id;references:region_id"`

	// The products that belong to this tax rate. Available if the relation `products` is expanded.
	Products []Product `json:"products" gorm:"foreignKey:id"`

	// The product types that belong to this tax rate. Available if the relation `product_types` is expanded.
	ProductTypes []ProductType `json:"product_types" gorm:"foreignKey:id"`

	// The shipping options that belong to this tax rate. Available if the relation `shipping_options` is expanded.
	ShippingOptions []ShippingOption `json:"shipping_options" gorm:"foreignKey:id"`

	// The count of products
	ProductCount int32 `json:"product_count" gorm:"default:null"`

	// The count of product types
	ProductTypeCount int32 `json:"product_type_count" gorm:"default:null"`

	// The count of shipping options
	ShippingOptionCount int32 `json:"shipping_option_count" gorm:"default:null"`
}
