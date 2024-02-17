package services

import (
	"context"
	"reflect"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

type GiftCardService struct {
	ctx context.Context
	r   Registry
}

func NewGiftCardService(
	r Registry,
) *GiftCardService {
	return &GiftCardService{
		context.Background(),
		r,
	}
}

func (s *GiftCardService) SetContext(context context.Context) *GiftCardService {
	s.ctx = context
	return s
}

func (s *GiftCardService) GenerateCode() string {
	code := randstr.Hex(4) + "-" + randstr.Hex(4) + "-" + randstr.Hex(4) + "-" + randstr.Hex(4)
	return code
}

func (s *GiftCardService) ListAndCount(selector *types.FilterableGiftCard, config *sql.Options) ([]models.GiftCard, *int64, *utils.ApplictaionError) {
	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 50
		config.Order = "created_at DESC"
	}

	var res []models.GiftCard

	query := sql.BuildQuery(selector, config)

	count, err := s.r.GiftCardRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *GiftCardService) List(selector *types.FilterableGiftCard, config *sql.Options) ([]models.GiftCard, *utils.ApplictaionError) {
	result, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *GiftCardService) CreateTransaction(data *types.CreateGiftCardTransactionInput) (uuid.UUID, *utils.ApplictaionError) {
	model := &models.GiftCardTransaction{
		Model: core.Model{
			CreatedAt: data.CreatedAt,
		},
		GiftCardId: uuid.NullUUID{UUID: data.GiftCardId},
		OrderId:    uuid.NullUUID{UUID: data.OrderId},
		Amount:     data.Amount,
		IsTaxable:  data.IsTaxable,
		TaxRate:    data.TaxRate,
	}
	if err := s.r.GiftCardTransactionRepository().Save(s.ctx, model); err != nil {
		return uuid.Nil, err
	}

	return model.Id, nil
}

func (s *GiftCardService) Create(data *types.CreateGiftCardInput) (*models.GiftCard, *utils.ApplictaionError) {
	var model *models.GiftCard
	region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(data.RegionId, &sql.Options{})
	if err != nil {
		return nil, err
	}

	model.RegionId = uuid.NullUUID{UUID: region.Id}

	model.Code = s.GenerateCode()
	if !reflect.ValueOf(data.TaxRate).IsZero() {
		model.TaxRate = s.ResolveTaxRate(data.TaxRate, region)
	}

	// s.eventBus_.WithTransaction(manager).Emit(GiftCardService.Events.CREATED, {
	// 	id: result.id,
	// })

	if !reflect.ValueOf(data.OrderId).IsZero() {
		model.OrderId = uuid.NullUUID{UUID: data.OrderId}
	}
	if !reflect.ValueOf(data.Value).IsZero() {
		model.Value = data.Value
	}
	if !reflect.ValueOf(data.Balance).IsZero() {
		model.Balance = data.Balance
	}
	if !reflect.ValueOf(data.EndsAt).IsZero() {
		model.EndsAt = data.EndsAt
	}
	if !reflect.ValueOf(data.IsDisabled).IsZero() {
		model.IsDisabled = data.IsDisabled
	}
	if !reflect.ValueOf(data.RegionId).IsZero() {
		model.RegionId = uuid.NullUUID{UUID: data.RegionId}
	}
	if data.Metadata != nil {
		model.Metadata = utils.MergeMaps(model.Metadata, data.Metadata)
	}

	if err := s.r.GiftCardRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *GiftCardService) ResolveTaxRate(giftCardTaxRate float64, region *models.Region) float64 {
	if !region.GiftCardsTaxable {
		return 0
	}
	if giftCardTaxRate != 0 {
		return giftCardTaxRate
	}
	return region.TaxRate
}

func (s *GiftCardService) Retrieve(selector *models.GiftCard, config *sql.Options) (*models.GiftCard, *utils.ApplictaionError) {
	var res *models.GiftCard
	query := sql.BuildQuery(selector, config)
	if err := s.r.GiftCardRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GiftCardService) RetrieveById(id uuid.UUID, config *sql.Options) (*models.GiftCard, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"id" must be defined`,
			nil,
		)
	}

	return s.Retrieve(&models.GiftCard{Model: core.Model{Id: id}}, config)
}

func (s *GiftCardService) RetrieveByCode(code string, config *sql.Options) (*models.GiftCard, *utils.ApplictaionError) {
	if code == "" {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"code" must be defined`,
			nil,
		)
	}

	return s.Retrieve(&models.GiftCard{Code: code}, config)
}

func (s *GiftCardService) Update(id uuid.UUID, data *types.UpdateGiftCardInput) (*models.GiftCard, *utils.ApplictaionError) {
	giftCard, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if data.RegionId != uuid.Nil && data.RegionId != giftCard.RegionId.UUID {
		region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(data.RegionId, &sql.Options{})
		if err != nil {
			return nil, err
		}
		giftCard.RegionId = uuid.NullUUID{UUID: region.Id}
	}
	if !reflect.ValueOf(data.Balance).IsZero() {
		if data.Balance < 0 || giftCard.Value < data.Balance {
			return nil, utils.NewApplictaionError(
				utils.INVALID_ARGUMENT,
				"new balance is invalid",
				"500",
				nil,
			)
		}
		giftCard.Balance = data.Balance
	}

	if !reflect.ValueOf(data.EndsAt).IsZero() {
		giftCard.EndsAt = data.EndsAt
	}
	if !reflect.ValueOf(data.IsDisabled).IsZero() {
		giftCard.IsDisabled = data.IsDisabled
	}
	if !reflect.ValueOf(data.RegionId).IsZero() {
		giftCard.RegionId = uuid.NullUUID{UUID: data.RegionId}
	}
	if data.Metadata != nil {
		giftCard.Metadata = utils.MergeMaps(giftCard.Metadata, data.Metadata)
	}

	if err := s.r.GiftCardRepository().Update(s.ctx, giftCard); err != nil {
		return nil, err
	}

	return giftCard, nil
}

func (s *GiftCardService) Delete(id uuid.UUID) *utils.ApplictaionError {
	res, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	if res == nil {
		return nil
	}

	if err := s.r.GiftCardRepository().SoftRemove(s.ctx, res); err != nil {
		return err
	}

	return nil
}
