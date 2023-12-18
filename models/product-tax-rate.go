package models

import "github.com/google/uuid"

// ProductTaxRate - Associates a tax rate with a product to indicate that the product is taxed in a certain way
type ProductTaxRate struct {

	// The ID of the Product
	ProductId uuid.NullUUID `json:"product_id"`

	Product *Product `json:"product" gorm:"foreignKey:id;references:product_id"`

	// The ID of the Tax Rate
	RateId uuid.NullUUID `json:"rate_id"`

	TaxRate *TaxRate `json:"tax_rate" gorm:"foreignKey:id;references:rate_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" gorm:"default:null"`
}
