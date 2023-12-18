package services

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/icza/gox/gox"
)

type CurrencyService struct {
	ctx  context.Context
	repo *repository.CurrencyRepo
}

func NewCurrencyService(
	ctx context.Context,
	repo *repository.CurrencyRepo,
) *CurrencyService {
	return &CurrencyService{
		ctx,
		repo,
	}
}

func (s *CurrencyService) RetrieveByCode(code string) (*models.Currency, error) {
	if code == "" {
		return nil, errors.New(`"email" must be defined`)
	}
	var res *models.Currency

	query := repository.BuildQuery[models.Currency](models.Currency{Code: strings.ToLower(code)}, repository.Options{})

	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CurrencyService) ListAndCount(selector models.Currency, config repository.Options) ([]models.Currency, *int64, error) {
	var res []models.Currency

	if reflect.DeepEqual(config, repository.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(20)
	}

	query := repository.BuildQuery[models.Currency](selector, config)
	count, err := s.repo.FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *CurrencyService) Update(code string, IncludesTax bool) (*models.Currency, error) {
	if code == "" {
		return nil, errors.New(`"code" must be defined`)
	}

	currency, err := s.RetrieveByCode(code)
	if err != nil {
		return nil, err
	}

	currency.IncludesTax = IncludesTax

	if err := s.repo.Upsert(s.ctx, currency); err != nil {
		return nil, err
	}

	return currency, nil
}
