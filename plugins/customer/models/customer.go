package models

import "github.com/driver005/gateway/core"

type Customer struct {
	core.SoftDeletableModel

	CompanyName string          `gorm:"column:company_name;type:text" json:"company_name"`
	FirstName   string          `gorm:"column:first_name;type:text" json:"first_name"`
	LastName    string          `gorm:"column:last_name;type:text" json:"last_name"`
	Email       string          `gorm:"column:email;type:text" json:"email"`
	Phone       string          `gorm:"column:phone;type:text" json:"phone"`
	HasAccount  bool            `gorm:"column:has_account;type:boolean;not null" json:"has_account"`
	Groups      []CustomerGroup `gorm:"many2many:customer_group_customers" json:"groups"`
	Addresses   []Address       `gorm:"foreignKey:CustomerId;references:Id;constraint:OnDelete:CASCADE" json:"addresses"`
	CreatedBy   string          `gorm:"column:created_by;type:text" json:"created_by"`
}

func (*Customer) TableName() string {
	return "customer"
}
