package types

import "github.com/google/uuid"

type CreatePublishableApiKeyInput struct {
	Title string `json:"title"`
}

type UpdatePublishableApiKeyInput struct {
	Title string `json:"title,omitempty" validate:"omitempty"`
}

type PublishableApiKeySalesChannelsBatch struct {
	SalesChannelIds []ProductBatchSalesChannel `json:"sales_channel_ids"`
}

type PublishableApiKeyScopes struct {
	SalesChannelIds uuid.UUIDs `json:"sales_channel_ids"`
}
