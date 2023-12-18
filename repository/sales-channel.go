package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type SalesChannelRepo struct {
	Repository[models.SalesChannel]
}

func SalesChannelRepository(db *gorm.DB) SalesChannelRepo {
	return SalesChannelRepo{*NewRepository[models.SalesChannel](db)}
}
