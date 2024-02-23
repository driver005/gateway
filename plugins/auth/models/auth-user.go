package models

import "github.com/driver005/gateway/core"

type AuthUser struct {
	Id               string     `gorm:"column:id;type:text;primaryKey" json:"id"`
	EntityId         string     `gorm:"column:entity_id;type:text;not null;uniqueIndex:IDX_auth_user_provider_scope_entity_id,priority:1" json:"entity_id"`
	Provider         string     `gorm:"column:provider;type:text;not null;uniqueIndex:IDX_auth_user_provider_scope_entity_id,priority:2" json:"provider"`
	Scope            string     `gorm:"column:scope;type:text;not null;uniqueIndex:IDX_auth_user_provider_scope_entity_id,priority:3" json:"scope"`
	UserMetadata     core.JSONB `gorm:"column:user_metadata;type:jsonb" json:"user_metadata"`
	AppMetadata      core.JSONB `gorm:"column:app_metadata;type:jsonb;not null" json:"app_metadata"`
	ProviderMetadata core.JSONB `gorm:"column:provider_metadata;type:jsonb" json:"provider_metadata"`
}
