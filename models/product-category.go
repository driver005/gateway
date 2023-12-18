package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// A product category can be used to categorize products into a hierarchy of categories.
type ProductCategory struct {
	core.Model

	// The product category's name
	Name string `json:"name"`

	// The product category's description.
	Description string `json:"description,omitempty"`

	// A unique string that identifies the Product Category - can for example be used in slug structures.
	Handle string `json:"handle"`

	// A string for Materialized Paths - used for finding ancestors and descendents
	Mpath string `json:"mpath"`

	// A flag to make product category an internal category for admins
	IsInternal bool `json:"is_internal"`

	// A flag to make product category visible/hidden in the store front
	IsActive bool `json:"is_active"`

	// An integer that depicts the rank of category in a tree node
	Rank int32 `json:"rank,omitempty"`

	// The details of the category's children.
	CategoryChildren []ProductCategory `json:"category_children"`

	// The ID of the parent category.
	ParentCategoryId uuid.NullUUID `json:"parent_category_id"`

	// The details of the parent of this category.
	ParentCategory *ProductCategory `json:"parent_category,omitempty" gorm:"foreignKey:id;references:parent_category_id"`

	// The details of the products that belong to this category.
	Products []Product `json:"products,omitempty"`
}
