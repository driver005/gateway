package services

import (
	"context"
	"reflect"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type PublishableApiKeyService struct {
	ctx context.Context
	r   Registry
}

func NewPublishableApiKeyService(
	r Registry,
) *PublishableApiKeyService {
	return &PublishableApiKeyService{
		context.Background(),
		r,
	}
}

func (s *PublishableApiKeyService) SetContext(context context.Context) *PublishableApiKeyService {
	s.ctx = context
	return s
}

func (s *PublishableApiKeyService) Create(data *types.CreatePublishableApiKeyInput, loggedInUserId uuid.UUID) (*models.PublishableApiKey, *utils.ApplictaionError) {
	model := &models.PublishableApiKey{
		CreatedBy: uuid.NullUUID{UUID: loggedInUserId},
		Title:     data.Title,
	}

	if err := s.r.PublishableApiKeyRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	// err := s.EventBusService.WithTransaction(manager).Emit("publishable_api_key.created", map[string]string{"id": publishableApiKey.ID})
	// if err != nil {
	// 	return nil, err
	// }

	return model, nil
}

func (s *PublishableApiKeyService) RetrieveById(id uuid.UUID, config *sql.Options) (*models.PublishableApiKey, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			nil,
		)
	}
	return s.Retrieve(&models.PublishableApiKey{Model: core.Model{Id: id}}, config)
}

func (s *PublishableApiKeyService) Retrieve(selector *models.PublishableApiKey, config *sql.Options) (*models.PublishableApiKey, *utils.ApplictaionError) {
	var res *models.PublishableApiKey
	query := sql.BuildQuery(selector, config)

	if err := s.r.PublishableApiKeyRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PublishableApiKeyService) ListAndCount(selector *models.PublishableApiKey, config *sql.Options, q *string) ([]models.PublishableApiKey, *int64, *utils.ApplictaionError) {
	var res []models.PublishableApiKey

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(20)
	}

	if q != nil {
		v := sql.ILike(*q)
		selector.Title = v
	}

	query := sql.BuildQuery(selector, config)

	count, err := s.r.PublishableApiKeyRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *PublishableApiKeyService) Update(id uuid.UUID, data *types.UpdatePublishableApiKeyInput) (*models.PublishableApiKey, *utils.ApplictaionError) {
	pubKey, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	pubKey.Title = data.Title

	if err := s.r.PublishableApiKeyRepository().Save(s.ctx, pubKey); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return pubKey, nil
}

func (s *PublishableApiKeyService) Delete(id uuid.UUID) *utils.ApplictaionError {
	data, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	if data == nil {
		return nil
	}

	if err := s.r.PublishableApiKeyRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *PublishableApiKeyService) Revoke(id uuid.UUID, loggedInUserId uuid.UUID) *utils.ApplictaionError {
	pubKey, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}
	if pubKey.RevokedAt != nil {
		return utils.NewApplictaionError(
			utils.CONFLICT,
			"PublishableApiKey has already been revoked.",
			nil,
		)
	}
	now := time.Now()
	pubKey.RevokedAt = &now
	pubKey.RevokedBy = uuid.NullUUID{UUID: loggedInUserId}

	if err := s.r.PublishableApiKeyRepository().Save(s.ctx, pubKey); err != nil {
		return err
	}
	// err = s.EventBusService.WithTransaction(manager).Emit("publishable_api_key.revoked", map[string]string{"id": pubKey.ID})
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s *PublishableApiKeyService) IsValid(id uuid.UUID) (bool, *utils.ApplictaionError) {
	pubKey, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return false, err
	}
	return pubKey.RevokedBy.UUID == uuid.Nil, nil
}

func (s *PublishableApiKeyService) AddSalesChannels(id uuid.UUID, salesChannelIds uuid.UUIDs) *utils.ApplictaionError {
	if err := s.r.PublishableApiKeySalesChannelRepository().AddSalesChannels(id, salesChannelIds); err != nil {
		return err
	}
	return nil
}

func (s *PublishableApiKeyService) RemoveSalesChannels(id uuid.UUID, salesChannelIds uuid.UUIDs) *utils.ApplictaionError {
	if err := s.r.PublishableApiKeySalesChannelRepository().RemoveSalesChannels(id, salesChannelIds); err != nil {
		return err
	}
	return nil
}

func (s *PublishableApiKeyService) ListSalesChannels(id uuid.UUID, q *string) ([]models.SalesChannel, *utils.ApplictaionError) {
	return s.r.PublishableApiKeySalesChannelRepository().FindSalesChannels(id, q)
}

func (s *PublishableApiKeyService) GetResourceScopes(id uuid.UUID) (map[string]uuid.UUIDs, *utils.ApplictaionError) {
	var res []models.PublishableApiKeySalesChannel

	query := sql.BuildQuery(models.PublishableApiKeySalesChannel{PublishableKeyId: uuid.NullUUID{UUID: id}}, &sql.Options{})

	if err := s.r.PublishableApiKeySalesChannelRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}

	var salesChannelIds uuid.UUIDs
	for _, salesChannel := range res {
		salesChannelIds = append(salesChannelIds, salesChannel.SalesChannelId.UUID)
	}
	return map[string]uuid.UUIDs{"sales_channel_ids": salesChannelIds}, nil
}
