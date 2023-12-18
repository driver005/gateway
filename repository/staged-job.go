package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type StagedJobRepo struct {
	Repository[models.StagedJob]
}

func StagedJobRepository(db *gorm.DB) StagedJobRepo {
	return StagedJobRepo{*NewRepository[models.StagedJob](db)}
}
