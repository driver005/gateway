package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type CurrencyRepo struct {
	Repository[models.Currency]
}

func CurrencyRepository(db *gorm.DB) CurrencyRepo {
	return CurrencyRepo{*NewRepository[models.Currency](db)}
}
