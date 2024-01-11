package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type PaymentRepo struct {
	sql.Repository[models.Payment]
}

func PaymentRepository(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{*sql.NewRepository[models.Payment](db)}
}
