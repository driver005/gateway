package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type DraftOrderCreate struct {
	Status              string
	Email               string
	BillingAddressId    uuid.UUID
	BillingAddress      *models.Address
	ShippingAddressId   uuid.UUID
	ShippingAddress     *models.Address
	Items               []models.LineItem
	RegionId            uuid.UUID
	Discounts           []models.Discount
	CustomerId          uuid.UUID
	NoNotificationOrder bool
	ShippingMethods     []models.ShippingMethod
	Metadata            core.JSONB
	IdempotencyKey      string
}
