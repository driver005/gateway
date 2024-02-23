package models

import "github.com/google/uuid"

type PriceListCustomerGroup struct {
	PriceListId     uuid.NullUUID `gorm:"column:price_list_id;type:character varying;primaryKey;index:IDX_52875734e9dd69064f0041f4d9,priority:1" json:"price_list_id"`
	CustomerGroupId uuid.NullUUID `gorm:"column:customer_group_id;type:character varying;primaryKey;index:IDX_c5516f550433c9b1c2630d787a,priority:1" json:"customer_group_id"`
}
