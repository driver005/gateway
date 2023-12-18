package models

import "github.com/driver005/gateway/core"

// Product Tags can be added to Products for easy filtering and grouping.
type ProductTag struct {
	core.Model

	// The value that the Product Tag represents
	Value string `json:"value"`
}
