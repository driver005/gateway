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

func (s *GiftCardService) ListAndCount(selector *models.GiftCard, config sql.Options) ([]models.GiftCard, *int64, *utils.ApplictaionError) {
	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	var res []models.GiftCard

	query := sql.BuildQuery(selector, config)

	count, err := s.r.GiftCardRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *GiftCardService) List(selector *models.GiftCard, config sql.Options) ([]models.GiftCard, *utils.ApplictaionError) {
	result, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *GiftCardService) CreateTransaction(data *models.GiftCardTransaction) (uuid.UUID, *utils.ApplictaionError) {
	if err := s.r.GiftCardTransactionRepository().Save(s.ctx, data); err != nil {
		return uuid.Nil, err
	}

	return data.Id, nil
}

func (s *GiftCardService) Create(data *models.GiftCard) (*models.GiftCard, *utils.ApplictaionError) {
	region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(data.RegionId.UUID, sql.Options{})
	if err != nil {
		return nil, err
	}

	data.RegionId = uuid.NullUUID{UUID: region.Id}

	data.Code = s.GenerateCode()
	data.TaxRate = s.ResolveTaxRate(data.TaxRate, region)

	// s.eventBus_.WithTransaction(manager).Emit(GiftCardService.Events.CREATED, {
	// 	id: result.id,
	// })

	if err := s.r.GiftCardRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}

	return data, nil
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

func (s *GiftCardService) Retrieve(selector *models.GiftCard, config sql.Options) (*models.GiftCard, *utils.ApplictaionError) {
	var res *models.GiftCard
	query := sql.BuildQuery(selector, config)
	if err := s.r.GiftCardRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GiftCardService) RetrieveById(id uuid.UUID, config sql.Options) (*models.GiftCard, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"id" must be defined`,
			"500",
			nil,
		)
	}

	return s.Retrieve(&models.GiftCard{Model: core.Model{Id: id}}, config)
}

func (s *GiftCardService) RetrieveByCode(code string, config sql.Options) (*models.GiftCard, *utils.ApplictaionError) {
	if code == "" {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"code" must be defined`,
			"500",
			nil,
		)
	}

	return s.Retrieve(&models.GiftCard{Code: code}, config)
}

func (s *GiftCardService) Update(id uuid.UUID, Update *models.GiftCard) (*models.GiftCard, *utils.ApplictaionError) {
	giftCard, err := s.RetrieveById(id, sql.Options{})
	if err != nil {
		return nil, err
	}

	if Update.RegionId.UUID != uuid.Nil && Update.RegionId != giftCard.RegionId {
		region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(Update.RegionId.UUID, sql.Options{})
		if err != nil {
			return nil, err
		}
		Update.RegionId = uuid.NullUUID{UUID: region.Id}
	}
	if Update.Balance != 0.0 {
		if Update.Balance < 0 || giftCard.Value < Update.Balance {
			return nil, utils.NewApplictaionError(
				utils.INVALID_ARGUMENT,
				"new balance is invalid",
				"500",
				nil,
			)
		}
		Update.Balance = giftCard.Balance
	}

	Update.Id = giftCard.Id

	if err := s.r.GiftCardRepository().Save(s.ctx, Update); err != nil {
		return nil, err
	}

	return Update, nil
}

func (s *GiftCardService) Delete(id uuid.UUID) *utils.ApplictaionError {
	res, err := s.RetrieveById(id, sql.Options{})
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
