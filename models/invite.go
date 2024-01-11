package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

// Invite - Represents an invite
type Invite struct {
	core.Model

	// The email of the user being invited.
	UserEmail string `json:"user_email"`

	// The user's role.
	Role UserRole `json:"role" gorm:"default:null"`

	// Whether the invite was accepted or not.
	Accepted bool `json:"accepted" gorm:"default:null"`

	// The token used to accept the invite.
	Token string `json:"token" gorm:"default:null"`

	// The date the invite expires at.
	ExpiresAt *time.Time `json:"expores_at" gorm:"default:null"`
}
