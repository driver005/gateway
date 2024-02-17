package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableSalesChannel struct {
	core.FilterModel
	Name        string `json:"name,omitempty" validate:"omitempty"`
	Description string `json:"description,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostSalesChannelsReq
// type: object
// description: "The details of the sales channel to create."
// required:
//   - name
//
// properties:
//
//	name:
//	  description: The name of the Sales Channel
//	  type: string
//	description:
//	  description: The description of the Sales Channel
//	  type: string
//	is_disabled:
//	  description: Whether the Sales Channel is disabled.
//	  type: boolean
type CreateSalesChannelInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty" validate:"omitempty"`
	IsDisabled  bool   `json:"is_disabled,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostSalesChannelsSalesChannelReq
// type: object
// description: "The details to update of the sales channel."
// properties:
//
//	name:
//	  type: string
//	  description: The name of the sales channel
//	description:
//	  type: string
//	  description:  The description of the sales channel.
//	is_disabled:
//	  type: boolean
//	  description: Whether the Sales Channel is disabled.
type UpdateSalesChannelInput CreateSalesChannelInput

type ProductBatchSalesChannel struct {
	Id uuid.UUID `json:"id"`
}

// @oas:schema:AdminPostSalesChannelsChannelStockLocationsReq
// type: object
// required:
//   - location_id
//
// properties:
//
//	location_id:
//	  description: The ID of the stock location
//	  type: string
type SalesChannelStockLocations struct {
	LocationId uuid.UUID `json:"location_id"`
}
