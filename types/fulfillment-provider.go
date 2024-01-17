package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FulfillmentOptions struct {
	ProviderId uuid.UUID  `json:"provider_id"`
	Options    core.JSONB `json:"options"`
}
