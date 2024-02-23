package models

import "github.com/google/uuid"

type DiscountRegion struct {
	DiscountId uuid.NullUUID `gorm:"column:discount_id;type:character varying;primaryKey;index:IDX_f4194aa81073f3fab8aa86906f,priority:1" json:"discount_id"`
	RegionId   uuid.NullUUID `gorm:"column:region_id;type:character varying;primaryKey;index:IDX_a21a7ffbe420d492eb46c305fe,priority:1" json:"region_id"`
}
