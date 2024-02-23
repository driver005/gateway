package models

import "github.com/google/uuid"

type OrderDiscount struct {
	OrderId    uuid.NullUUID `gorm:"column:order_id;type:character varying;primaryKey;index:IDX_e7b488cebe333f449398769b2c,priority:1" json:"order_id"`
	DiscountId uuid.NullUUID `gorm:"column:discount_id;type:character varying;primaryKey;index:IDX_0fc1ec4e3db9001ad60c19daf1,priority:1" json:"discount_id"`
}
