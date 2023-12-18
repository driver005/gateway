package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type InviteRepo struct {
	Repository[models.Invite]
}

func InviteRepository(db *gorm.DB) InviteRepo {
	return InviteRepo{*NewRepository[models.Invite](db)}
}
