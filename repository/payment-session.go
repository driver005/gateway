package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type PaymentSessionRepo struct {
	sql.Repository[models.PaymentSession]
}

func PaymentSessionRepository(db *gorm.DB) *PaymentSessionRepo {
	return &PaymentSessionRepo{*sql.NewRepository[models.PaymentSession](db)}
}
