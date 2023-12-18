package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

// TODO: ADD
type PublishableApiKeySalesChannelRepo struct {
	Repository[models.PublishableApiKeySalesChannel]
}

func PublishableApiKeySalesChannelRepository(db *gorm.DB) PublishableApiKeySalesChannelRepo {
	return PublishableApiKeySalesChannelRepo{*NewRepository[models.PublishableApiKeySalesChannel](db)}
}
