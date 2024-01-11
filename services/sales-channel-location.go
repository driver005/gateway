package services

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type SalesChannelLocationService struct {
	ctx context.Context
	r   Registry
}

func NewSalesChannelLocationService(
	r Registry,
) *SalesChannelLocationService {
	return &SalesChannelLocationService{
		context.Background(),
		r,
	}
}

func (s *SalesChannelLocationService) SetContext(context context.Context) *SalesChannelLocationService {
	s.ctx = context
	return s
}

func (s *SalesChannelLocationService) RemoveLocation(locationId uuid.UUID, salesChannelId uuid.UUID) *utils.ApplictaionError {
	var res []models.SalesChannelLocation
	query := sql.BuildQuery(
		models.SalesChannelLocation{
			SalesChannelId: uuid.NullUUID{UUID: salesChannelId},
			LocationId:     uuid.NullUUID{UUID: locationId},
		},
		sql.Options{},
	)
	if err := s.r.SalesChannelLocationRepository().Find(s.ctx, res, query); err != nil {
		return err
	}

	if len(res) > 0 {
		if err := s.r.SalesChannelLocationRepository().RemoveSlice(s.ctx, res); err != nil {
			return err
		}
	}
	return nil
}

func (s *SalesChannelLocationService) AssociateLocation(salesChannelId uuid.UUID, locationId uuid.UUID) *utils.ApplictaionError {
	salesChannel, err := s.r.SalesChannelService().SetContext(s.ctx).RetrieveById(salesChannelId, sql.Options{})
	if err != nil {
		return err
	}
	if s.r.StockLocationService() != nil {
		_, err = s.r.StockLocationService().Retrieve(s.ctx, locationId, sql.Options{})
		if err != nil {
			return err
		}
	}

	salesChannelLocation := &models.SalesChannelLocation{
		SalesChannelId: uuid.NullUUID{UUID: salesChannel.Id},
		LocationId:     uuid.NullUUID{UUID: locationId},
	}
	if err := s.r.SalesChannelLocationRepository().Save(s.ctx, salesChannelLocation); err != nil {
		return err
	}
	return nil
}

func (s *SalesChannelLocationService) ListLocationIds(salesChannelId uuid.UUID) (uuid.UUIDs, *utils.ApplictaionError) {
	salesChannels, err := s.r.SalesChannelService().SetContext(s.ctx).RetrieveById(salesChannelId, sql.Options{
		Selects: []string{"id"},
	})
	if err != nil {
		return nil, err
	}

	var locations []models.SalesChannelLocation
	query := sql.BuildQuery(
		models.SalesChannelLocation{
			SalesChannelId: uuid.NullUUID{UUID: salesChannels.Id},
		},
		sql.Options{
			Selects: []string{"location_id"},
		},
	)
	if err := s.r.SalesChannelLocationRepository().Find(s.ctx, locations, query); err != nil {
		return nil, err
	}

	var res uuid.UUIDs
	for _, l := range locations {
		res = append(res, l.LocationId.UUID)
	}
	return res, nil
}

func (s *SalesChannelLocationService) ListSalesChannelIds(locationId uuid.UUID) (uuid.UUIDs, *utils.ApplictaionError) {
	location, err := s.r.StockLocationService().Retrieve(s.ctx, locationId, sql.Options{})
	if err != nil {
		return nil, err
	}
	var locations []models.SalesChannelLocation
	query := sql.BuildQuery(
		models.SalesChannelLocation{
			LocationId: uuid.NullUUID{UUID: location.Id},
		},
		sql.Options{
			Selects: []string{"sales_channel_id"},
		},
	)
	if err := s.r.SalesChannelLocationRepository().Find(s.ctx, locations, query); err != nil {
		return nil, err
	}

	var res uuid.UUIDs
	for _, l := range locations {
		res = append(res, l.SalesChannelId.UUID)
	}
	return res, nil
}
