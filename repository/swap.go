package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type SwapRepo struct {
	sql.Repository[models.Swap]
}

func SwapRepository(db *gorm.DB) *SwapRepo {
	return &SwapRepo{*sql.NewRepository[models.Swap](db)}
}
