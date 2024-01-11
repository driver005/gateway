package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublishableApiKeySalesChannelRepo struct {
	sql.Repository[models.PublishableApiKeySalesChannel]
}

func PublishableApiKeySalesChannelRepository(db *gorm.DB) *PublishableApiKeySalesChannelRepo {
	return &PublishableApiKeySalesChannelRepo{*sql.NewRepository[models.PublishableApiKeySalesChannel](db)}
}
func (r *PublishableApiKeySalesChannelRepo) FindSalesChannels(id uuid.UUID, q *string) ([]models.SalesChannel, *utils.ApplictaionError) {
	var records []models.SalesChannel
	err := r.Db().Model(&models.PublishableApiKeySalesChannel{}).
		Joins("SalesChannel").
		Where("publishable_api_key_id = ?", id).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if q != nil {
				return db.Where("description LIKE ? OR name LIKE ?", "%"+*q+"%", "%"+*q+"%")
			}
			return db
		}).
		Find(&records).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}

	return records, nil
}

func (r *PublishableApiKeySalesChannelRepo) AddSalesChannels(id uuid.UUID, salesChannelIds uuid.UUIDs) *utils.ApplictaionError {
	for _, salesChannelId := range salesChannelIds {
		err := r.Db().Create(&models.PublishableApiKeySalesChannel{
			SalesChannelId:   uuid.NullUUID{UUID: salesChannelId},
			PublishableKeyId: uuid.NullUUID{UUID: id},
		}).Error
		if err != nil {
			return r.HandleDBError(err)
		}
	}

	return nil
}

func (r *PublishableApiKeySalesChannelRepo) RemoveSalesChannels(id uuid.UUID, salesChannelIds uuid.UUIDs) *utils.ApplictaionError {
	err := r.Db().Where("publishable_api_key_id = ? AND sales_channel_id IN ?", id, salesChannelIds.Strings()).
		Delete(&models.PublishableApiKeySalesChannel{}).Error

	return r.HandleDBError(err)
}
