package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ImageRepo struct {
	Repository[models.Image]
}

func ImageRepository(db *gorm.DB) ImageRepo {
	return ImageRepo{*NewRepository[models.Image](db)}
}
