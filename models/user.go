package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
)

// User - Represents a User who can manage store settings.
type User struct {
	core.Model

	Role UserRole `json:"role" gorm:"default:member"`

	// The email of the User
	Email string `json:"email"`

	// The first name of the User
	FirstName string `json:"first_name" gorm:"default:null"`

	// The last name of the User
	LastName string `json:"last_name" gorm:"default:null"`

	// Password of the user
	Password string `json:"password" gorm:"-" sql:"-"`

	// Password hash of the user
	PasswordHash string `json:"password_hash" gorm:"default:null"`

	// An API token associated with the user.
	ApiToken string `json:"api_token" gorm:"default:null"`
}

type UserRole string

// Defines values for UserRole.
const (
	UserRoleAdmin     UserRole = "admin"
	UserRoleMember    UserRole = "member"
	UserRoleDeveloper UserRole = "developer"
)

func (sp *UserRole) Scan(value interface{}) error {
	*sp = UserRole(value.([]byte))
	return nil
}

func (sp UserRole) Value() (driver.Value, error) {
	return string(sp), nil
}
