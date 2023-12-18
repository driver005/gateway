package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type OrderEditRepo struct {
	Repository[models.OrderEdit]
}

func OrderEditRepository(db *gorm.DB) OrderEditRepo {
	return OrderEditRepo{*NewRepository[models.OrderEdit](db)}
}
