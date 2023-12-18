package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type MoneyAmountRepo struct {
	Repository[models.MoneyAmount]
}

func MoneyAmountRepository(db *gorm.DB) MoneyAmountRepo {
	return MoneyAmountRepo{*NewRepository[models.MoneyAmount](db)}
}
