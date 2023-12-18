package types

import "github.com/driver005/gateway/core"

type StringComparisonOperator string

type DateComparisonOperator string

type FilterableLineItemAdjustmentProps struct {
	core.FilterModel

	ItemID      string `json:"item_id,,omitempty"`
	Description string `json:"description,,omitempty"`
	ResourceID  string `json:"resource_id,,omitempty"`
}
