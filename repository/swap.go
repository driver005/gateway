package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type SwapRepo struct {
	Repository[models.Swap]
}

func SwapRepository(db *gorm.DB) SwapRepo {
	return SwapRepo{*NewRepository[models.Swap](db)}
}
