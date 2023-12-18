package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

type FilterableProduct struct {
	core.FilterModel
	Q                       string                 `json:"q,omitempty"`
	Status                  []models.ProductStatus `json:"status,omitempty"`
	PriceListID             []string               `json:"price_list_id,omitempty"`
	CollectionID            []string               `json:"collection_id,omitempty"`
	Tags                    []string               `json:"tags,omitempty"`
	Title                   string                 `json:"title,omitempty"`
	Description             string                 `json:"description,omitempty"`
	Handle                  string                 `json:"handle,omitempty"`
	IsGiftcard              *bool                  `json:"is_giftcard,omitempty"`
	TypeID                  []string               `json:"type_id,omitempty"`
	SalesChannelID          []string               `json:"sales_channel_id,omitempty"`
	DiscountConditionID     string                 `json:"discount_condition_id,omitempty"`
	CategoryID              []string               `json:"category_id,omitempty"`
	IncludeCategoryChildren *bool                  `json:"include_category_children,omitempty"`
}
