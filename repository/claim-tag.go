package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ClaimTagRepo struct {
	sql.Repository[models.ClaimTag]
}

func ClaimTagRepository(db *gorm.DB) *ClaimTagRepo {
	return &ClaimTagRepo{*sql.NewRepository[models.ClaimTag](db)}
}
