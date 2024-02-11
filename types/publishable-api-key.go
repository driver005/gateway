package types

type CreatePublishableApiKeyInput struct {
	Title string `json:"title"`
}

type UpdatePublishableApiKeyInput struct {
	Title string `json:"title,omitempty" validate:"omitempty"`
}

type PublishableApiKeySalesChannelsBatch struct {
	SalesChannelIds []ProductBatchSalesChannel `json:"sales_channel_ids"`
}
