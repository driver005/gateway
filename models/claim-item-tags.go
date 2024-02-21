package models

import "github.com/google/uuid"

type ClaimItemTag struct {
	ItemId uuid.NullUUID `gorm:"column:item_id;type:character varying;primaryKey;index:IDX_c2c0f3edf39515bd15432afe6e,priority:1" json:"item_id"`
	TagId  uuid.NullUUID `gorm:"column:tag_id;type:character varying;primaryKey;index:IDX_dc9bbf9fcb9ba458d25d512811,priority:1" json:"tag_id"`
}
