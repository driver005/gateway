package types

import "github.com/google/uuid"

// @oas:schema:AdminPostPublishableApiKeysReq
// type: object
// description: "The details of the publishable API key to create."
// required:
//   - title
//
// properties:
//
//	title:
//	  description: The title of the publishable API key
//	  type: string
type CreatePublishableApiKeyInput struct {
	Title string `json:"title"`
}

// @oas:schema:AdminPostPublishableApiKeysPublishableApiKeyReq
// type: object
// description: "The details to update of the publishable API key."
// properties:
//
//	title:
//	  description: The title of the Publishable API Key.
//	  type: string
type UpdatePublishableApiKeyInput struct {
	Title string `json:"title,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostPublishableApiKeySalesChannelsBatchReq
// type: object
// description: "The details of the sales channels to add to the publishable API key."
// required:
//   - sales_channel_ids
//
// properties:
//
//	sales_channel_ids:
//	  description: The IDs of the sales channels to add to the publishable API key
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - id
//	    properties:
//	      id:
//	        type: string
//	        description: The ID of the sales channel
type PublishableApiKeySalesChannelsBatch struct {
	SalesChannelIds []ProductBatchSalesChannel `json:"sales_channel_ids"`
}

type PublishableApiKeyScopes struct {
	SalesChannelIds uuid.UUIDs `json:"sales_channel_ids"`
}
