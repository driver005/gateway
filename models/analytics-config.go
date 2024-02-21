package models

import "github.com/google/uuid"

type AnalyticsConfig struct {
	UserId    uuid.NullUUID `json:"user_id" gorm:"column:user_id;uniqueIndex"`
	OptOut    bool          `json:"opt_out" gorm:"column:opt_out;default:false"`
	Anonymize bool          `json:"anonymize" gorm:"column:anonymize;default:false"`
}
