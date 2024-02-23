package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/plugins/order/types"
)

type Order struct {
	core.SoftDeletableModel

	RegionId          string              `gorm:"column:region_id;type:text;index:IDX_order_region_id,priority:1" json:"region_id"`
	CustomerId        string              `gorm:"column:customer_id;type:text;index:IDX_order_customer_id,priority:1" json:"customer_id"`
	OriginalOrderId   string              `gorm:"column:original_order_id;type:text;index:IDX_order_original_order_id,priority:1" json:"original_order_id"`
	Version           int32               `gorm:"column:version;type:integer;not null;default:1" json:"version"`
	SalesChannelId    string              `gorm:"column:sales_channel_id;type:text" json:"sales_channel_id"`
	Status            string              `gorm:"column:status;type:text;not null;default:pending" json:"status"`
	Email             string              `gorm:"column:email;type:text" json:"email"`
	CurrencyCode      string              `gorm:"column:currency_code;type:text;not null;index:IDX_order_currency_code,priority:1" json:"currency_code"`
	ShippingAddressId string              `gorm:"column:shipping_address_id;type:text;index:IDX_order_shipping_address_id,priority:1" json:"shipping_address_id"`
	ShippingAddress   *Address            `gorm:"foreignkey:ShippingAddressId" json:"shipping_address"`
	BillingAddressId  string              `gorm:"column:billing_address_id;type:text;index:IDX_order_billing_address_id,priority:1" json:"billing_address_id"`
	BillingAddress    *Address            `gorm:"foreignkey:BillingAddressId" json:"billing_address"`
	NoNotification    bool                `gorm:"column:no_notification;type:boolean" json:"no_notification"`
	Summary           *types.OrderSummary `gorm:"column:summary;type:jsonb;not null" json:"summary"`
	Items             []OrderDetail       `gorm:"foreignkey:OrderId" json:"items"`
	ShippingMethods   []ShippingMethod    `gorm:"foreignkey:OrderId" json:"shipping_methods"`
	Transactions      []Transaction       `gorm:"foreignkey:OrderId" json:"transactions"`
	CanceledAt        *time.Time          `gorm:"column:canceled_at;type:timestamp with time zone" json:"canceled_at"`
}

func (*Order) TableName() string {
	return "order"
}
