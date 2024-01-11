package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type StagedJobRepo struct {
	sql.Repository[models.StagedJob]
}

func StagedJobRepository(db *gorm.DB) *StagedJobRepo {
	return &StagedJobRepo{*sql.NewRepository[models.StagedJob](db)}
}
