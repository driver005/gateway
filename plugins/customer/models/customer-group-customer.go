package models

import "github.com/driver005/gateway/core"

type CustomerGroupCustomer struct {
	core.BaseModel

	CustomerId      string         `gorm:"column:customer_id;type:text;not null;index:IDX_customer_group_customer_customer_id,priority:1" json:"customer_id"`
	CustomerGroupId string         `gorm:"column:customer_group_id;type:text;not null;index:IDX_customer_group_customer_group_id,priority:1" json:"customer_group_id"`
	Customer        *Customer      `gorm:"foreignKey:CustomerId;references:Id;constraint:OnDelete:CASCADE"`
	CustomerGroup   *CustomerGroup `gorm:"foreignKey:CustomerGroupId;references:Id;constraint:OnDelete:CASCADE"`
	CreatedBy       string         `gorm:"column:created_by;type:text" json:"created_by"`
}

func (*CustomerGroupCustomer) TableName() string {
	return "customer_group_customer"
}
