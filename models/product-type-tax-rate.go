package models

import "github.com/google/uuid"

// ProductTypeTaxRate - Associates a tax rate with a product type to indicate that the product type is taxed in a certain way
type ProductTypeTaxRate struct {

	// The ID of the Product type
	ProductTypeId uuid.NullUUID `json:"product_type_id"`

	ProductType *ProductType `json:"product_type" gorm:"foreignKey:id;references:product_type_id"`

	// The id of the Tax Rate
	RateId uuid.NullUUID `json:"rate_id"`

	TaxRate *TaxRate `json:"tax_rate" gorm:"foreignKey:id;references:rate_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" gorm:"default:null"`
}
