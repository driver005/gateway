package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type ClaimTypeValue string

type CreateClaimInput struct {
	Type              models.ClaimStatus                   `json:"type"`
	ClaimItems        []CreateClaimItemInput               `json:"claim_items"`
	ReturnShipping    *CreateClaimReturnShippingInput      `json:"return_shipping,omitempty" validate:"omitempty"`
	AdditionalItems   []CreateClaimItemAdditionalItemInput `json:"additional_items,omitempty" validate:"omitempty"`
	ShippingMethods   []CreateClaimShippingMethodInput     `json:"shipping_methods,omitempty" validate:"omitempty"`
	RefundAmount      float64                              `json:"refund_amount,omitempty" validate:"omitempty"`
	ShippingAddress   *AddressPayload                      `json:"shipping_address,omitempty" validate:"omitempty"`
	NoNotification    bool                                 `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata          core.JSONB                           `json:"metadata,omitempty" validate:"omitempty"`
	Order             *models.Order                        `json:"order"`
	ClaimOrderId      uuid.UUID                            `json:"claim_order_id,omitempty" validate:"omitempty"`
	ShippingAddressId uuid.UUID                            `json:"shipping_address_id,omitempty" validate:"omitempty"`
	ReturnLocationId  uuid.UUID                            `json:"return_location_id,omitempty" validate:"omitempty"`
}

type CreateClaimReturnShippingInput struct {
	OptionId uuid.UUID `json:"option_id,omitempty" validate:"omitempty"`
	Price    float64   `json:"price,omitempty" validate:"omitempty"`
}

type CreateClaimShippingMethodInput struct {
	Id       uuid.UUID  `json:"id,omitempty" validate:"omitempty"`
	OptionId uuid.UUID  `json:"option_id,omitempty" validate:"omitempty"`
	Price    float64    `json:"price,omitempty" validate:"omitempty"`
	Data     core.JSONB `json:"data,omitempty" validate:"omitempty"`
}

type CreateClaimItemInput struct {
	ItemId       uuid.UUID              `json:"item_id"`
	Quantity     int                    `json:"quantity"`
	ClaimOrderId uuid.UUID              `json:"claim_order_id"`
	Reason       models.ClaimReasonType `json:"reason"`
	Note         string                 `json:"note,omitempty" validate:"omitempty"`
	Tags         []string               `json:"tags,omitempty" validate:"omitempty"`
	Images       []string               `json:"images,omitempty" validate:"omitempty"`
}

type CreateClaimItemAdditionalItemInput struct {
	VariantId uuid.UUID `json:"variant_id"`
	Quantity  int       `json:"quantity"`
}

type UpdateClaimInput struct {
	ClaimItems      []UpdateClaimItemInput           `json:"claim_items,omitempty" validate:"omitempty"`
	ShippingMethods []UpdateClaimShippingMethodInput `json:"shipping_methods,omitempty" validate:"omitempty"`
	NoNotification  bool                             `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata        core.JSONB                       `json:"metadata,omitempty" validate:"omitempty"`
}

type UpdateClaimShippingMethodInput struct {
	Id       uuid.UUID  `json:"id,omitempty" validate:"omitempty"`
	OptionId uuid.UUID  `json:"option_id,omitempty" validate:"omitempty"`
	Price    float64    `json:"price,omitempty" validate:"omitempty"`
	Data     core.JSONB `json:"data,omitempty" validate:"omitempty"`
}

type UpdateClaimItemInput struct {
	Id       uuid.UUID                   `json:"id,omitempty" validate:"omitempty"`
	Note     string                      `json:"note,omitempty" validate:"omitempty"`
	Reason   models.ClaimReasonType      `json:"reason,omitempty" validate:"omitempty"`
	Images   []UpdateClaimItemImageInput `json:"images,omitempty" validate:"omitempty"`
	Tags     []UpdateClaimItemTagInput   `json:"tags,omitempty" validate:"omitempty"`
	Metadata core.JSONB                  `json:"metadata,omitempty" validate:"omitempty"`
}

type UpdateClaimItemImageInput struct {
	Id  uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	URL string    `json:"url,omitempty" validate:"omitempty"`
}

type UpdateClaimItemTagInput struct {
	Id    uuid.UUID `json:"id,omitempty" validate:"omitempty"`
	Value string    `json:"value,omitempty" validate:"omitempty"`
}
