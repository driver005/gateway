package models

import "github.com/google/uuid"

type CustomerGroupCustomer struct {
	CustomerGroupId uuid.NullUUID `gorm:"column:customer_group_id;type:character varying;primaryKey;index:IDX_620330964db8d2999e67b0dbe3,priority:1" json:"customer_group_id"`
	CustomerId      uuid.NullUUID `gorm:"column:customer_id;type:character varying;primaryKey;index:IDX_3c6412d076292f439269abe1a2,priority:1" json:"customer_id"`
}
