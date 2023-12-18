package models

import "github.com/driver005/gateway/core"

// Images holds a reference to a URL at which the image file can be found.
type Image struct {
	core.Model

	// The URL at which the image file can be found.
	Url string `json:"url"`
}
