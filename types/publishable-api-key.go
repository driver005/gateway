package types

type CreatePublishableApiKeyInput struct {
	Title string `json:"title"`
}

type UpdatePublishableApiKeyInput struct {
	Title string `json:"title,omitempty" validate:"omitempty"`
}
