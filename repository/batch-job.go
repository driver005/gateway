package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type BatchJobRepo struct {
	Repository[models.BatchJob]
}

func BatchJobRepository(db *gorm.DB) BatchJobRepo {
	return BatchJobRepo{*NewRepository[models.BatchJob](db)}
}
