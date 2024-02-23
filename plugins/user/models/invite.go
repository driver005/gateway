package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type Invite struct {
	core.BaseModel

	Email     string         `gorm:"column:email;type:text;not null;uniqueIndex:IDX_invite_email,priority:1" json:"email"`
	Accepted  bool           `gorm:"column:accepted;type:boolean;not null" json:"accepted"`
	Token     string         `gorm:"column:token;type:text;not null;index:IDX_invite_token,priority:1" json:"token"`
	ExpiresAt time.Time      `gorm:"column:expires_at;type:timestamp with time zone;not null" json:"expires_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_invite_deleted_at,priority:1" json:"deleted_at"`
}

func (*Invite) TableName() string {
	return "invite"
}
