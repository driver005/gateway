package types

import "github.com/driver005/gateway/core"

type CreateProductCollection struct {
	Title    string     `json:"title"`
	Handle   string     `json:"handle,omitempty" validate:"omitempty"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type UpdateProductCollection struct {
	Title    string     `json:"title,omitempty" validate:"omitempty"`
	Handle   string     `json:"handle,omitempty" validate:"omitempty"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}
