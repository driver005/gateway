package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type IdempotencyKeyRepo struct {
	Repository[models.IdempotencyKey]
}

func IdempotencyKeyRepository(db *gorm.DB) IdempotencyKeyRepo {
	return IdempotencyKeyRepo{*NewRepository[models.IdempotencyKey](db)}
}
