package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type BatchJobRepo struct {
	sql.Repository[models.BatchJob]
}

func BatchJobRepository(db *gorm.DB) *BatchJobRepo {
	return &BatchJobRepo{*sql.NewRepository[models.BatchJob](db)}
}
