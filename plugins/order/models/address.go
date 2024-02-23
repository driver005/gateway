package models

import (
	"github.com/driver005/gateway/core"
)

type Address struct {
	core.BaseModel

	CustomerId  string `gorm:"column:customer_id;type:text;index:IDX_order_address_customer_id,priority:1" json:"customer_id"`
	Company     string `gorm:"column:company;type:text" json:"company"`
	FirstName   string `gorm:"column:first_name;type:text" json:"first_name"`
	LastName    string `gorm:"column:last_name;type:text" json:"last_name"`
	Address1    string `gorm:"column:address_1;type:text" json:"address_1"`
	Address2    string `gorm:"column:address_2;type:text" json:"address_2"`
	City        string `gorm:"column:city;type:text" json:"city"`
	CountryCode string `gorm:"column:country_code;type:text" json:"country_code"`
	Province    string `gorm:"column:province;type:text" json:"province"`
	PostalCode  string `gorm:"column:postal_code;type:text" json:"postal_code"`
	Phone       string `gorm:"column:phone;type:text" json:"phone"`
}

func (*Address) TableName() string {
	return "order_address"
}
