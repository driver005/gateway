package types

import "github.com/driver005/gateway/core"

type FilterableNote struct {
	core.FilterModel

	ResourceType string `json:"resource_type,omitempty"`
	ResourceId   string `json:"resource_id,omitempty"`
	Value        string `json:"value,omitempty"`
	AuthorId     string `json:"author_id,omitempty"`
}
