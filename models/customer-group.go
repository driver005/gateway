package models

import "github.com/driver005/gateway/core"

// CustomerGroup - Represents a customer group
type CustomerGroup struct {
	core.Model

	// The name of the customer group
	Name string `json:"name"`

	// The customers that belong to the customer group. Available if the relation `customers` is expanded.
	Customers []Customer `json:"customers" gorm:"many2many:customer_group"`

	// The price lists that are associated with the customer group. Available if the relation `price_lists` is expanded.
	PriceLists []PriceList `json:"price_lists" gorm:"foreignKey:id"`
}
