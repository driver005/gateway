package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type PublishableApiKeyRepo struct {
	Repository[models.PublishableApiKey]
}

func PublishableApiKeyRepository(db *gorm.DB) PublishableApiKeyRepo {
	return PublishableApiKeyRepo{*NewRepository[models.PublishableApiKey](db)}
}
