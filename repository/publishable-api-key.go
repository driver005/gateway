package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type PublishableApiKeyRepo struct {
	sql.Repository[models.PublishableApiKey]
}

func PublishableApiKeyRepository(db *gorm.DB) *PublishableApiKeyRepo {
	return &PublishableApiKeyRepo{*sql.NewRepository[models.PublishableApiKey](db)}
}
