package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type InviteRepo struct {
	sql.Repository[models.Invite]
}

func InviteRepository(db *gorm.DB) *InviteRepo {
	return &InviteRepo{*sql.NewRepository[models.Invite](db)}
}
