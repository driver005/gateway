package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

type FilterableInvite struct {
	core.FilterModel
}

type ListInvite struct {
	*models.Invite
	Token string `json:"token"`
}

type CreateInviteInput struct {
	Email string          `json:"email"`
	Role  models.UserRole `json:"role"`
}

type AcceptInvite struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}
