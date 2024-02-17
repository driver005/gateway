package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

const TempReorderRank = 99999

type FilterableProductCategory struct {
	core.FilterModel
	Handle                 string    `json:"handle,omitempty" validate:"omitempty,string"`
	IncludeDescendantsTree bool      `json:"include_descendants_tree,omitempty" validate:"omitempty,bool"`
	IsInternal             bool      `json:"is_internal,omitempty" validate:"omitempty,bool"`
	IsActive               bool      `json:"is_active,omitempty" validate:"omitempty,bool"`
	ParentCategoryId       uuid.UUID `json:"parent_category_id,omitempty" validate:"omitempty,string"`
}

type ProductCategoryInput struct {
	Handle           string                  `json:"handle,omitempty" validate:"omitempty"`
	IsInternal       bool                    `json:"is_internal,omitempty" validate:"omitempty"`
	IsActive         bool                    `json:"is_active,omitempty" validate:"omitempty"`
	ParentCategoryId uuid.UUID               `json:"parent_category_id,omitempty" validate:"omitempty"`
	ParentCategory   *models.ProductCategory `json:"parent_category,omitempty" validate:"omitempty"`
	Rank             int64                   `json:"rank,omitempty" validate:"omitempty"`
	Metadata         core.JSONB              `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostProductCategoriesReq
// type: object
// description: "The details of the product category to create."
// required:
//   - name
//
// properties:
//
//	name:
//	  type: string
//	  description: The name of the product category
//	description:
//	  type: string
//	  description: The description of the product category.
//	handle:
//	  type: string
//	  description: The handle of the product category. If none is provided, the kebab-case version of the name will be used. This field can be used as a slug in URLs.
//	is_internal:
//	  type: boolean
//	  description: >-
//	    If set to `true`, the product category will only be available to admins.
//	is_active:
//	  type: boolean
//	  description: >-
//	    If set to `false`, the product category will not be available in the storefront.
//	parent_category_id:
//	  type: string
//	  description: The ID of the parent product category
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type CreateProductCategoryInput struct {
	ProductCategoryInput
	Name string `json:"name"`
}

// @oas:schema:AdminPostProductCategoriesCategoryReq
// type: object
// description: "The details to update of the product category."
// properties:
//
//	name:
//	  type: string
//	  description:  The name to identify the Product Category by.
//	description:
//	  type: string
//	  description: An optional text field to describe the Product Category by.
//	handle:
//	  type: string
//	  description:  A handle to be used in slugs.
//	is_internal:
//	  type: boolean
//	  description: A flag to make product category an internal category for admins
//	is_active:
//	  type: boolean
//	  description: A flag to make product category visible/hidden in the store front
//	parent_category_id:
//	  type: string
//	  description: The ID of the parent product category
//	rank:
//	  type: number
//	  description: The rank of the category in the tree node (starting from 0)
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdateProductCategoryInput struct {
	ProductCategoryInput
	Name string `json:"name,omitempty" validate:"omitempty"`
}

type AdminProductCategoriesReqBase struct {
	Description      string    `json:"description,omitempty" validate:"omitempty"`
	Handle           string    `json:"handle,omitempty" validate:"omitempty"`
	IsInternal       bool      `json:"is_internal,omitempty" validate:"omitempty"`
	IsActive         bool      `json:"is_active,omitempty" validate:"omitempty"`
	ParentCategoryId uuid.UUID `json:"parent_category_id,omitempty" validate:"omitempty"`
}

type ProductBatchProductCategory struct {
	Id uuid.UUID `json:"id"`
}

type ReorderConditions struct {
	TargetCategoryId    uuid.UUID `json:"targetCategoryId uuid.UUID"`
	OriginalParentId    uuid.UUID `json:"originalParentId uuid.UUID,omitempty" validate:"omitempty"`
	TargetParentId      uuid.UUID `json:"targetParentId uuid.UUID,omitempty" validate:"omitempty"`
	OriginalRank        int64     `json:"originalRank"`
	TargetRank          int64     `json:"targetRank,omitempty" validate:"omitempty"`
	ShouldChangeParent  bool      `json:"shouldChangeParent"`
	ShouldChangeRank    bool      `json:"shouldChangeRank"`
	ShouldIncrementRank bool      `json:"shouldIncrementRank"`
	ShouldDeleteElement bool      `json:"shouldDeleteElement"`
}
