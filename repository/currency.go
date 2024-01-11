package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type CurrencyRepo struct {
	sql.Repository[models.Currency]
}

func CurrencyRepository(db *gorm.DB) *CurrencyRepo {
	return &CurrencyRepo{*sql.NewRepository[models.Currency](db)}
}
