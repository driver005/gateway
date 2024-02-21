package models

import "github.com/google/uuid"

type CartDiscount struct {
	CartId     uuid.NullUUID `gorm:"column:cart_id;type:character varying;primaryKey;index:IDX_6680319ebe1f46d18f106191d5,priority:1" json:"cart_id"`
	DiscountId uuid.NullUUID `gorm:"column:discount_id;type:character varying;primaryKey;index:IDX_8df75ef4f35f217768dc113545,priority:1" json:"discount_id"`
}
