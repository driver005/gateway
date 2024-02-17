package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type ClaimTypeValue string

// @oas:schema:AdminPostOrdersOrderClaimsReq
// type: object
// description: "The details of the claim to be created."
// required:
//   - type
//   - claim_items
//
// properties:
//
//	type:
//	  description: >-
//	    The type of the Claim. This will determine how the Claim is treated: `replace` Claims will result in a Fulfillment with new items being created, while a `refund` Claim will refund the amount paid for the claimed items.
//	  type: string
//	  enum:
//	    - replace
//	    - refund
//	claim_items:
//	  description: The Claim Items that the Claim will consist of.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - item_id
//	      - quantity
//	    properties:
//	      item_id:
//	        description: The ID of the Line Item that will be claimed.
//	        type: string
//	      quantity:
//	        description: The number of items that will be returned
//	        type: integer
//	      note:
//	        description: Short text describing the Claim Item in further detail.
//	        type: string
//	      reason:
//	        description: The reason for the Claim
//	        type: string
//	        enum:
//	          - missing_item
//	          - wrong_item
//	          - production_failure
//	          - other
//	      tags:
//	        description: A list of tags to add to the Claim Item
//	        type: array
//	        items:
//	          type: string
//	      images:
//	        description: A list of image URL's that will be associated with the Claim
//	        items:
//	          type: string
//	return_shipping:
//	   description: Optional details for the Return Shipping Method, if the items are to be sent back. Providing this field will result in a return being created and associated with the claim.
//	   type: object
//	   properties:
//	     option_id:
//	       type: string
//	       description: The ID of the Shipping Option to create the Shipping Method from.
//	     price:
//	       type: integer
//	       description: The price to charge for the Shipping Method.
//	additional_items:
//	   description: The new items to send to the Customer. This is only used if the claim's type is `replace`.
//	   type: array
//	   items:
//	     type: object
//	     required:
//	       - variant_id
//	       - quantity
//	     properties:
//	       variant_id:
//	         description: The ID of the Product Variant.
//	         type: string
//	       quantity:
//	         description: The quantity of the Product Variant.
//	         type: integer
//	shipping_methods:
//	   description: The Shipping Methods to send the additional Line Items with. This is only used if the claim's type is `replace`.
//	   type: array
//	   items:
//	      type: object
//	      properties:
//	        id:
//	          description: The ID of an existing Shipping Method
//	          type: string
//	        option_id:
//	          description: The ID of the Shipping Option to create a Shipping Method from
//	          type: string
//	        price:
//	          description: The price to charge for the Shipping Method
//	          type: integer
//	        data:
//	          description: An optional set of key-value pairs to hold additional information.
//	          type: object
//	shipping_address:
//	   description: "An optional shipping address to send the claimed items to. If not provided, the parent order's shipping address will be used."
//	   $ref: "#/components/schemas/AddressPayload"
//	refund_amount:
//	   description: The amount to refund the customer. This is used when the claim's type is `refund`.
//	   type: integer
//	no_notification:
//	   description: If set to true no notification will be send related to this Claim.
//	   type: boolean
//	return_location_id:
//	   description: The ID of the location used for the associated return.
//	   type: string
//	metadata:
//	   description: An optional set of key-value pairs to hold additional information.
//	   type: object
//	   externalDocs:
//	     description: "Learn about the metadata attribute, and how to delete and update it."
//	     url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
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

// @oas:schema:AdminPostOrdersOrderClaimsClaimReq
// type: object
// properties:
//
//	claim_items:
//	  description: The Claim Items that the Claim will consist of.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - id
//	      - images
//	      - tags
//	    properties:
//	      id:
//	        description: The ID of the Claim Item.
//	        type: string
//	      item_id:
//	        description: The ID of the Line Item that will be claimed.
//	        type: string
//	      quantity:
//	        description: The number of items that will be returned
//	        type: integer
//	      note:
//	        description: Short text describing the Claim Item in further detail.
//	        type: string
//	      reason:
//	        description: The reason for the Claim
//	        type: string
//	        enum:
//	          - missing_item
//	          - wrong_item
//	          - production_failure
//	          - other
//	      tags:
//	        description: A list o tags to add to the Claim Item
//	        type: array
//	        items:
//	          type: object
//	          properties:
//	            id:
//	              type: string
//	              description: Tag ID
//	            value:
//	              type: string
//	              description: Tag value
//	      images:
//	        description: A list of image URL's that will be associated with the Claim
//	        type: array
//	        items:
//	          type: object
//	          properties:
//	            id:
//	              type: string
//	              description: Image ID
//	            url:
//	              type: string
//	              description: Image URL
//	      metadata:
//	        description: An optional set of key-value pairs to hold additional information.
//	        type: object
//	        externalDocs:
//	          description: "Learn about the metadata attribute, and how to delete and update it."
//	          url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	shipping_methods:
//	  description: The Shipping Methods to send the additional Line Items with.
//	  type: array
//	  items:
//	     type: object
//	     properties:
//	       id:
//	         description: The ID of an existing Shipping Method
//	         type: string
//	       option_id:
//	         description: The ID of the Shipping Option to create a Shipping Method from
//	         type: string
//	       price:
//	         description: The price to charge for the Shipping Method
//	         type: integer
//	       data:
//	         description: An optional set of key-value pairs to hold additional information.
//	         type: object
//	no_notification:
//	  description: If set to true no notification will be send related to this Swap.
//	  type: boolean
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
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
