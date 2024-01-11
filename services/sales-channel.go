package services

import (
	"context"
	"reflect"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type SalesChannelService struct {
	ctx context.Context
	r   Registry
}

func NewSalesChannelService(
	r Registry,
) *SalesChannelService {
	return &SalesChannelService{
		context.Background(),
		r,
	}
}

func (s *SalesChannelService) SetContext(context context.Context) *SalesChannelService {
	s.ctx = context
	return s
}

func (s *SalesChannelService) Retrieve(selector *models.SalesChannel, config sql.Options) (*models.SalesChannel, *utils.ApplictaionError) {
	var res *models.SalesChannel
	query := sql.BuildQuery(selector, config)
	if err := s.r.SalesChannelRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SalesChannelService) RetrieveById(salesChannelId uuid.UUID, config sql.Options) (*models.SalesChannel, *utils.ApplictaionError) {
	if salesChannelId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"optionId" must be defined`,
			"500",
			nil,
		)
	}

	return s.Retrieve(&models.SalesChannel{Model: core.Model{Id: salesChannelId}}, config)
}

func (s *SalesChannelService) RetrieveByName(name string, config sql.Options) (*models.SalesChannel, *utils.ApplictaionError) {
	if name == "" {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"name" must be defined`,
			"500",
			nil,
		)
	}
	return s.Retrieve(&models.SalesChannel{Name: name}, config)
}

func (s *SalesChannelService) List(selector models.SalesChannel, config sql.Options) ([]models.SalesChannel, *utils.ApplictaionError) {
	res, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SalesChannelService) ListAndCount(selector models.SalesChannel, config sql.Options) ([]models.SalesChannel, *int64, *utils.ApplictaionError) {
	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	var res []models.SalesChannel

	query := sql.BuildQuery(selector, config)

	count, err := s.r.SalesChannelRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *SalesChannelService) Create(data *models.SalesChannel) (*models.SalesChannel, *utils.ApplictaionError) {
	// err := s.EventBusService.withTransaction(manager).emit(SalesChannelService.Events.CREATED, map[string]interface{}{
	// 	"id": salesChannel.ID,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	if err := s.r.SalesChannelRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}
	return data, nil

}

func (s *SalesChannelService) Update(salesChannelId uuid.UUID, Update *models.SalesChannel) (*models.SalesChannel, *utils.ApplictaionError) {
	salesChannel, err := s.RetrieveById(salesChannelId, sql.Options{})
	if err != nil {
		return nil, err
	}

	Update.Id = salesChannel.Id

	if err := s.r.SalesChannelRepository().Save(s.ctx, Update); err != nil {
		return nil, err
	}
	// err = s.EventBusService.emit(SalesChannelService.Events.UPDATED, map[string]interface{}{
	// 	"id": result.ID,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return Update, nil

}

func (s *SalesChannelService) Delete(salesChannelId uuid.UUID) *utils.ApplictaionError {
	salesChannel, err := s.RetrieveById(salesChannelId, sql.Options{
		Relations: []string{"locations"},
	})
	if err != nil {
		return err
	}
	store, err := s.r.StoreService().SetContext(s.ctx).Retrieve(sql.Options{
		Selects: []string{"default_sales_channel_id"},
	})
	if err != nil {
		return err
	}
	if salesChannel.Id == store.DefaultSalesChannelId.UUID {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"You cannot Delete the default sales channel",
			"500",
			nil,
		)
	}
	if err := s.r.SalesChannelRepository().SoftRemove(s.ctx, salesChannel); err != nil {
		return err
	}
	// err = s.EventBusService.emit(SalesChannelService.Events.DELETED, map[string]interface{}{
	// 	"id": salesChannelId,
	// })
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s *SalesChannelService) CreateDefault() (*models.SalesChannel, *utils.ApplictaionError) {
	store, err := s.r.StoreService().SetContext(s.ctx).Retrieve(sql.Options{
		Relations: []string{"default_sales_channel"},
	})
	if err != nil {
		return nil, err
	}
	if store.DefaultSalesChannelId.UUID != uuid.Nil {
		return store.DefaultSalesChannel, nil
	}
	defaultSalesChannel, err := s.Create(&models.SalesChannel{
		Description: "Created by Medusa",
		Name:        "Default Sales Channel",
		IsDisabled:  false,
	})
	if err != nil {
		return nil, err
	}
	_, err = s.r.StoreService().SetContext(s.ctx).Update(&models.Store{
		DefaultSalesChannelId: uuid.NullUUID{UUID: defaultSalesChannel.Id},
	})
	if err != nil {
		return nil, err
	}
	return defaultSalesChannel, nil

}

func (s *SalesChannelService) RetrieveDefault() (*models.SalesChannel, *utils.ApplictaionError) {
	store, err := s.r.StoreService().SetContext(s.ctx).Retrieve(sql.Options{
		Relations: []string{"default_sales_channel"},
	})
	if err != nil {
		return nil, err
	}
	if store.DefaultSalesChannelId.UUID == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			"Default Sales channel was not found",
			"500",
			nil,
		)
	}
	return store.DefaultSalesChannel, nil
}

func (s *SalesChannelService) ListProductIdsBySalesChannelIds(salesChannelIds uuid.UUIDs) (map[string][]string, *utils.ApplictaionError) {
	return s.r.SalesChannelRepository().ListProductIdsBySalesChannelIds(salesChannelIds)
}

func (s *SalesChannelService) AddProducts(salesChannelId uuid.UUID, productIds uuid.UUIDs) (*models.SalesChannel, *utils.ApplictaionError) {
	isMedusaV2Enabled := true
	err := s.r.SalesChannelRepository().AddProducts(salesChannelId, productIds, isMedusaV2Enabled)
	if err != nil {
		return nil, err
	}
	return s.RetrieveById(salesChannelId, sql.Options{})
}

func (s *SalesChannelService) RemoveProducts(salesChannelId uuid.UUID, productIds uuid.UUIDs) (*models.SalesChannel, *utils.ApplictaionError) {
	err := s.r.SalesChannelRepository().RemoveProducts(salesChannelId, productIds)
	if err != nil {
		return nil, err
	}
	return s.RetrieveById(salesChannelId, sql.Options{})
}
