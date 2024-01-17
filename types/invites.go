package types

import "github.com/driver005/gateway/models"

type ListInvite struct {
	*models.Invite
	Token string `json:"token"`
}
