package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ReturnRepo struct {
	Repository[models.Return]
}

func ReturnRepository(db *gorm.DB) ReturnRepo {
	return ReturnRepo{*NewRepository[models.Return](db)}
}
