package models

import "github.com/driver005/gateway/core"

type CustomerGroup struct {
	core.SoftDeletableModel

	Id        string     `gorm:"column:id;type:text;primaryKey" json:"id"`
	Name      string     `gorm:"column:name;type:text;uniqueIndex:IDX_customer_group_name,priority:1" json:"name"`
	Customers []Customer `gorm:"many2many:customer_group_customers" json:"customers"`
	CreatedBy string     `gorm:"column:created_by;type:text" json:"created_by"`
}

func (*CustomerGroup) TableName() string {
	return "customer_group"
}
