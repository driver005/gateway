package models

import "github.com/driver005/gateway/core"

// ProductCollection - Product Collections represents a group of Products that are related.
type ProductCollection struct {
	core.Model

	// The title that the Product Collection is identified by.
	Title string `json:"title"`

	// A unique string that identifies the Product Collection - can for example be used in slug structures.
	Handle string `json:"handle" gorm:"default:null"`

	// The Products contained in the Product Collection. Available if the relation `products` is expanded.
	Products []Product `json:"products" gorm:"foreignKey:id"`
}
