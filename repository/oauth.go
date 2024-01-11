package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type OAuthRepo struct {
	sql.Repository[models.OAuth]
}

func OAuthRepository(db *gorm.DB) *OAuthRepo {
	return &OAuthRepo{*sql.NewRepository[models.OAuth](db)}
}
