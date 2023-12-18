package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

// A Publishable API key defines scopes that resources are available in. Then, it can be used in request to infer the resources without having to directly pass them. For example, a publishable API key can be associated with one or more sales channels. Then, when the publishable API key is passed in the header of a request, it is inferred what sales channel is being used without having to pass the sales channel as a query or body parameter of the request. Publishable API keys can only be used with sales channels, at the moment.
type PublishableApiKey struct {
	core.Model

	// The unique identifier of the user that created the key.
	CreatedBy string `json:"created_by"`
	
	// The unique identifier of the user that revoked the key.
	RevokedBy string `json:"revoked_by"`

	// The date with timezone at which the key was revoked.
	RevokedAt time.Time `json:"revoked_at"`

	// The key's title.
	Title string `json:"title"`
}
