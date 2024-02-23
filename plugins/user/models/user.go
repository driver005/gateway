package models

import (
	"github.com/driver005/gateway/core"
	"gorm.io/gorm"
)

type User struct {
	core.BaseModel

	FirstName string         `gorm:"column:first_name;type:text" json:"first_name"`
	LastName  string         `gorm:"column:last_name;type:text" json:"last_name"`
	Email     string         `gorm:"column:email;type:text;not null;index:IDX_user_email,priority:1" json:"email"`
	AvatarURL string         `gorm:"column:avatar_url;type:text" json:"avatar_url"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;index:IDX_user_deleted_at,priority:1" json:"deleted_at"`
}

func (*User) TableName() string {
	return "user"
}
