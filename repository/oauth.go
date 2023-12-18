package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type OAuthRepo struct {
	Repository[models.OAuth]
}

func OAuthRepository(db *gorm.DB) OAuthRepo {
	return OAuthRepo{*NewRepository[models.OAuth](db)}
}
