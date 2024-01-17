package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Product Variants represent a Product with a specific set of Product Option configurations. The maximum number of Product Variants that a Product can have is given by the number of available Product Option combinations.
type ProductVariant struct {
	core.Model

	// Whether the Product Variant should be purchasable when `inventory_quantity` is 0.
	AllowBackorder bool `json:"allow_backorder" gorm:"default:null"`

	// Whether the Product Variant should be purchasable.
	Purchasable bool `json:"purchasable" gorm:"default:null"`

	// A generic field for a GTIN number that can be used to identify the Product Variant.
	Barcode string `json:"barcode" gorm:"default:null"`

	// An EAN barcode number that can be used to identify the Product Variant.
	Ean string `json:"ean" gorm:"default:null"`

	// The height of the Product Variant. May be used in shipping rate calculations.
	Height float64 `json:"height" gorm:"default:null"`

	// The Harmonized System code of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	HsCode string `json:"hs_code" gorm:"default:null"`

	// The current quantity of the item that is stocked.
	InventoryQuantity int `json:"inventory_quantity"`

	// The length of the Product Variant. May be used in shipping rate calculations.
	Length float64 `json:"length" gorm:"default:null"`

	// Whether Medusa should manage inventory for the Product Variant.
	ManageInventory bool `json:"manage_inventory" gorm:"default:null"`

	// The material and composition that the Product Variant is made of, May be used by Fulfillment Providers to pass customs information to shipping carriers.
	Material string `json:"material" gorm:"default:null"`

	// The Manufacturers Identification code that identifies the manufacturer of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	MIdCode uuid.UUID `json:"mid_code" gorm:"default:null"`

	// The Product Option Values specified for the Product Variant. Available if the relation `options` is expanded.
	Options []ProductOptionValue `json:"options" gorm:"foreignKey:id"`

	// The country in which the Product Variant was produced. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	OriginCountry string `json:"origin_country" gorm:"default:null"`

	// The Money Amounts defined for the Product Variant. Each Money Amount represents a price in a given currency or a price in a specific Region. Available if the relation `prices` is expanded.
	Prices []MoneyAmount `json:"prices" gorm:"foreignKey:id"`

	// A product object. Available if the relation `product` is expanded.
	Product *Product `json:"product" gorm:"foreignKey:id;references:product_id"`

	// The ID of the Product that the Product Variant belongs to.
	ProductId uuid.NullUUID `json:"product_id"`

	// The unique stock keeping unit used to identify the Product Variant. This will usually be a unqiue identifer for the item that is to be shipped, and can be referenced across multiple systems.
	Sku string `json:"sku" gorm:"default:null"`

	// A title that can be displayed for easy identification of the Product Variant.
	Title string `json:"title"`

	// A UPC barcode number that can be used to identify the Product Variant.
	Upc string `json:"upc" gorm:"default:null"`

	// The ranking of this variant
	VariantRank int `json:"variant_rank" gorm:"default:null"`

	// The weight of the Product Variant. May be used in shipping rate calculations.
	Weight float64 `json:"weight" gorm:"default:null"`

	// The width of the Product Variant. May be used in shipping rate calculations.
	Width float64 `json:"width" gorm:"default:null"`
}
