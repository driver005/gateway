package models

import "github.com/google/uuid"

type RegionFulfillmentProvider struct {
	RegionId   uuid.NullUUID `gorm:"column:region_id;type:character varying;primaryKey;index:IDX_c556e14eff4d6f03db593df955,priority:1" json:"region_id"`
	ProviderId uuid.NullUUID `gorm:"column:provider_id;type:character varying;primaryKey;index:IDX_37f361c38a18d12a3fa3158d0c,priority:1" json:"provider_id"`
}
