package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type UpdateStoreInput struct {
	Name                  string     `json:"name,omitempty" validate:"omitempty"`
	SwapLinkTemplate      string     `json:"swap_link_template,omitempty" validate:"omitempty"`
	PaymentLinkTemplate   string     `json:"payment_link_template,omitempty" validate:"omitempty"`
	InviteLinkTemplate    string     `json:"invite_link_template,omitempty" validate:"omitempty"`
	DefaultCurrencyCode   string     `json:"default_currency_code,omitempty" validate:"omitempty"`
	Currencies            []string   `json:"currencies,omitempty" validate:"omitempty"`
	Metadata              core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	DefaultSalesChannelId uuid.UUID  `json:"default_sales_channel_id,omitempty" validate:"omitempty"`
}

// type ExtendedStoreDTO struct {
// 	*models.Store
// 	PaymentProviders     []PaymentProvider
// 	FulfillmentProviders []FulfillmentProvider
// 	FeatureFlags         FeatureFlagsResponse
// 	Modules              ModulesResponse
// }
