package models

import "github.com/google/uuid"

type AnalyticsConfig struct {
	UserId    uuid.UUID `json:"user_id" gorm:"uniqueIndex"`
	OptOut    *bool     `json:"opt_out,omitempty" gorm:"default:false"`
	Anonymize *bool     `json:"anonymize,omitempty" gorm:"default:false"`
}
