package models

import (
	"github.com/driver005/gateway/core"
)

type Cart struct {
	core.SoftDeletableModel

	RegionId          string           `gorm:"column:region_id;type:text;index:IDX_cart_region_id,priority:1" json:"region_id"`
	CustomerId        string           `gorm:"column:customer_id;type:text;index:IDX_cart_customer_id,priority:1" json:"customer_id"`
	SalesChannelId    string           `gorm:"column:sales_channel_id;type:text;index:IDX_cart_sales_channel_id,priority:1" json:"sales_channel_id"`
	Email             string           `gorm:"column:email;type:text" json:"email"`
	CurrencyCode      string           `gorm:"column:currency_code;type:text;not null;index:IDX_cart_currency_code,priority:1" json:"currency_code"`
	ShippingAddressId string           `gorm:"column:shipping_address_id;type:text;index:IDX_cart_shipping_address_id,priority:1" json:"shipping_address_id"`
	BillingAddressId  string           `gorm:"column:billing_address_id;type:text;index:IDX_cart_billing_address_id,priority:1" json:"billing_address_id"`
	ShippingAddress   *Address         `gorm:"foreignKey:ShippingAddressId" json:"shipping_address"`
	BillingAddress    *Address         `gorm:"foreignKey:BillingAddressId" json:"billing_address"`
	LineItems         []LineItem       `gorm:"foreignKey:CartId" json:"line_items"`
	ShippingMethods   []ShippingMethod `gorm:"foreignKey:CartId" json:"shipping_methods"`
}

func (*Cart) TableName() string {
	return "cart"
}
