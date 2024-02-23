package models

import "github.com/google/uuid"

type RegionPaymentProvider struct {
	RegionId   uuid.NullUUID `gorm:"column:region_id;type:character varying;primaryKey;index:IDX_8aaa78ba90d3802edac317df86,priority:1" json:"region_id"`
	ProviderId uuid.NullUUID `gorm:"column:provider_id;type:character varying;primaryKey;index:IDX_3a6947180aeec283cd92c59ebb,priority:1" json:"provider_id"`
}
