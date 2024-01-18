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

type CreateProductCategoryInput struct {
	ProductCategoryInput
	Name string `json:"name"`
}

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
