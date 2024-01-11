package types

import "github.com/google/uuid"

type FulfillmentOptions struct {
	ProviderId uuid.UUID
	Options    map[string]interface{}
}
