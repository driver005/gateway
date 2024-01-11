package types

import (
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

const TempReorderRank = 99999

type ProductCategoryInput struct {
	Handle           string                  `json:"handle,omitempty"`
	IsInternal       bool                    `json:"is_internal,omitempty"`
	IsActive         bool                    `json:"is_active,omitempty"`
	ParentCategoryId uuid.UUID               `json:"parent_category_id,omitempty"`
	ParentCategory   *models.ProductCategory `json:"parent_category,omitempty"`
	Rank             int                     `json:"rank,omitempty"`
	Metadata         map[string]interface{}  `json:"metadata,omitempty"`
}

type CreateProductCategoryInput struct {
	ProductCategoryInput
	Name string `json:"name"`
}

type UpdateProductCategoryInput struct {
	ProductCategoryInput
	Name *string `json:"name,omitempty"`
}

type AdminProductCategoriesReqBase struct {
	Description      string    `json:"description,omitempty"`
	Handle           string    `json:"handle,omitempty"`
	IsInternal       bool      `json:"is_internal,omitempty"`
	IsActive         bool      `json:"is_active,omitempty"`
	ParentCategoryId uuid.UUID `json:"parent_category_id,omitempty"`
}

type ProductBatchProductCategory struct {
	Id string `json:"id"`
}

type ReorderConditions struct {
	TargetCategoryId    uuid.UUID `json:"targetCategoryId"`
	OriginalParentId    uuid.UUID `json:"originalParentId,omitempty"`
	TargetParentId      uuid.UUID `json:"targetParentId,omitempty"`
	OriginalRank        int64     `json:"originalRank"`
	TargetRank          int64     `json:"targetRank,omitempty"`
	ShouldChangeParent  bool      `json:"shouldChangeParent"`
	ShouldChangeRank    bool      `json:"shouldChangeRank"`
	ShouldIncrementRank bool      `json:"shouldIncrementRank"`
	ShouldDeleteElement bool      `json:"shouldDeleteElement"`
}
