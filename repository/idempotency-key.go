package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type IdempotencyKeyRepo struct {
	sql.Repository[models.IdempotencyKey]
}

func IdempotencyKeyRepository(db *gorm.DB) *IdempotencyKeyRepo {
	return &IdempotencyKeyRepo{*sql.NewRepository[models.IdempotencyKey](db)}
}
