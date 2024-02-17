package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FilterableShippingProfile struct {
	core.FilterModel
}

// @oas:schema:AdminPostShippingProfilesReq
// type: object
// description: "The details of the shipping profile to create."
// required:
//   - name
//   - type
//
// properties:
//
//	name:
//	  description: The name of the Shipping Profile
//	  type: string
//	type:
//	  description: The type of the Shipping Profile
//	  type: string
//	  enum: [default, gift_card, custom]
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type CreateShippingProfile struct {
	Name     string                     `json:"name"`
	Type     models.ShippingProfileType `json:"type"`
	Metadata core.JSONB                 `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostShippingProfilesProfileReq
// type: object
// description: "The detail to update of the shipping profile."
// properties:
//
//	name:
//	  description: The name of the Shipping Profile
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	type:
//	  description: The type of the Shipping Profile
//	  type: string
//	  enum: [default, gift_card, custom]
//	products:
//	  description: product IDs to associate with the Shipping Profile
//	  type: array
//	shipping_options:
//	  description: Shipping option IDs to associate with the Shipping Profile
//	  type: array
type UpdateShippingProfile struct {
	Name            string                     `json:"name,omitempty" validate:"omitempty"`
	Metadata        core.JSONB                 `json:"metadata,omitempty" validate:"omitempty"`
	Type            models.ShippingProfileType `json:"type,omitempty" validate:"omitempty"`
	Products        uuid.UUIDs                 `json:"products,omitempty" validate:"omitempty"`
	ShippingOptions uuid.UUIDs                 `json:"shipping_options,omitempty" validate:"omitempty"`
}
