package services

import (
	"context"
	"reflect"
	"strings"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/icza/gox/gox"
)

type CurrencyService struct {
	ctx context.Context
	r   Registry
}

func NewCurrencyService(
	r Registry,
) *CurrencyService {
	return &CurrencyService{
		context.Background(),
		r,
	}
}

func (s *CurrencyService) SetContext(context context.Context) *CurrencyService {
	s.ctx = context
	return s
}

func (s *CurrencyService) RetrieveByCode(code string) (*models.Currency, *utils.ApplictaionError) {
	if code == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"code" must be defined`,
			"500",
			nil,
		)
	}
	var res *models.Currency

	query := sql.BuildQuery[models.Currency](models.Currency{Code: strings.ToLower(code)}, sql.Options{})

	if err := s.r.CurrencyRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CurrencyService) ListAndCount(selector models.Currency, config sql.Options) ([]models.Currency, *int64, *utils.ApplictaionError) {
	var res []models.Currency

	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(20)
	}

	query := sql.BuildQuery[models.Currency](selector, config)
	count, err := s.r.CurrencyRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *CurrencyService) Update(code string, IncludesTax bool) (*models.Currency, *utils.ApplictaionError) {
	if code == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"code" must be defined`,
			"500",
			nil,
		)
	}

	currency, err := s.RetrieveByCode(code)
	if err != nil {
		return nil, err
	}

	currency.IncludesTax = IncludesTax

	if err := s.r.CurrencyRepository().Upsert(s.ctx, currency); err != nil {
		return nil, err
	}

	return currency, nil
}
