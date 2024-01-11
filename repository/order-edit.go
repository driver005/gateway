package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type OrderEditRepo struct {
	sql.Repository[models.OrderEdit]
}

func OrderEditRepository(db *gorm.DB) *OrderEditRepo {
	return &OrderEditRepo{*sql.NewRepository[models.OrderEdit](db)}
}
