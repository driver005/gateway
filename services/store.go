package services

import (
	"context"
	"reflect"
	"strings"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
)

type StoreService struct {
	ctx context.Context
	r   Registry
}

func NewStoreService(
	r Registry,
) *StoreService {
	return &StoreService{
		context.Background(),
		r,
	}
}

func (s *StoreService) SetContext(context context.Context) *StoreService {
	s.ctx = context
	return s
}

func (s *StoreService) Create() (*models.Store, *utils.ApplictaionError) {
	store, err := s.Retrieve(&sql.Options{})
	if err == nil {
		return store, nil
	}

	var model *models.Store = &models.Store{}
	var usd *models.Currency = &models.Currency{}

	query := sql.BuildQuery(models.Currency{Code: "usd"}, &sql.Options{})
	if err := s.r.CurrencyRepository().FindOne(s.ctx, usd, query); err != nil {
		return nil, err
	}

	model.Currencies = []models.Currency{*usd}

	if err := s.r.StoreRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *StoreService) Retrieve(config *sql.Options) (*models.Store, *utils.ApplictaionError) {
	var stores []models.Store
	if err := s.r.StoreRepository().Find(s.ctx, &stores, sql.Query{Where: "NOT id = null"}); err != nil {
		return nil, err
	}
	if len(stores) == 0 {
		return nil, utils.NewApplictaionError(
			utils.DB_ERROR,
			"Store does not exist",
			nil,
		)
	}
	return &stores[0], nil
}

func (s *StoreService) GetDefaultCurrency(code string) *models.Currency {
	currencyObject := utils.Currencies[strings.ToUpper(code)]
	return &models.Currency{
		Code:         strings.ToLower(currencyObject.Code),
		Symbol:       currencyObject.Symbol,
		SymbolNative: currencyObject.SymbolNative,
		Name:         currencyObject.Name,
	}
}

func (s *StoreService) Update(data *types.UpdateStoreInput) (*models.Store, *utils.ApplictaionError) {
	store, err := s.Retrieve(&sql.Options{
		Relations: []string{"currencies"},
	})
	if err == nil {
		return store, nil
	}

	if data.Currencies != nil {
		defaultCurr := store.DefaultCurrencyCode
		if data.DefaultCurrencyCode == "" {
			defaultCurr = store.DefaultCurrencyCode
		}
		hasDefCurrency := false
		for _, c := range data.Currencies {
			if strings.EqualFold(c, defaultCurr) {
				hasDefCurrency = true
				break
			}
		}
		if !hasDefCurrency {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"You are not allowed to remove default currency from store currencies without replacing it as well",
				"500",
				nil,
			)
		}
		var currencies []models.Currency
		for _, curr := range data.Currencies {
			currency, err := s.r.CurrencyService().RetrieveByCode(curr)
			if err != nil {
				return nil, err
			}
			if currency == nil {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					"Currency with code "+curr+" does not exist",
					"500",
					nil,
				)
			}
			currencies = append(currencies, *currency)
		}
		store.Currencies = currencies
	}
	if !reflect.ValueOf(data.DefaultCurrencyCode).IsZero() {
		hasDefCurrency := false
		for _, c := range store.Currencies {
			if strings.EqualFold(c.Code, data.DefaultCurrencyCode) {
				hasDefCurrency = true
				break
			}
		}
		if !hasDefCurrency {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Store does not have currency: "+data.DefaultCurrencyCode,
				"500",
				nil,
			)
		}

		currency, err := s.r.CurrencyService().RetrieveByCode(data.DefaultCurrencyCode)
		if err != nil {
			return nil, err
		}

		store.DefaultCurrency = currency
		store.DefaultCurrencyCode = currency.Code
	}

	if !reflect.ValueOf(data.Name).IsZero() {
		store.Name = data.Name
	}
	if !reflect.ValueOf(data.SwapLinkTemplate).IsZero() {
		store.SwapLinkTemplate = data.SwapLinkTemplate
	}
	if !reflect.ValueOf(data.PaymentLinkTemplate).IsZero() {
		store.PaymentLinkTemplate = data.PaymentLinkTemplate
	}
	if !reflect.ValueOf(data.InviteLinkTemplate).IsZero() {
		store.InviteLinkTemplate = data.InviteLinkTemplate
	}
	if data.Metadata != nil {
		store.Metadata = utils.MergeMaps(store.Metadata, data.Metadata)
	}

	if err := s.r.StoreRepository().Update(s.ctx, store); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *StoreService) AddCurrency(code string) (*models.Store, *utils.ApplictaionError) {
	var currency *models.Currency = &models.Currency{}
	query := sql.BuildQuery(models.Currency{Code: strings.ToLower(code)}, &sql.Options{})
	if err := s.r.CurrencyRepository().FindOne(s.ctx, currency, query); err != nil {
		return nil, err
	}
	if reflect.DeepEqual(currency, &models.Currency{}) {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Currency "+code+" not found",
			nil,
		)
	}
	store, err := s.Retrieve(&sql.Options{
		Relations: []string{"currencies"},
	})
	if err != nil {
		return nil, err
	}
	doesStoreInclCurrency := false
	for _, c := range store.Currencies {
		if strings.EqualFold(c.Code, currency.Code) {
			doesStoreInclCurrency = true
			break
		}
	}
	if doesStoreInclCurrency {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Currency already added",
			nil,
		)
	}
	store.Currencies = append(store.Currencies, *currency)
	if err := s.r.StoreRepository().Upsert(s.ctx, store); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *StoreService) RemoveCurrency(code string) (*models.Store, *utils.ApplictaionError) {
	store, err := s.Retrieve(&sql.Options{
		Relations: []string{"currencies"},
	})
	if err != nil {
		return nil, err
	}
	doesCurrencyExists := false
	for _, c := range store.Currencies {
		if c.Code == strings.ToLower(code) {
			doesCurrencyExists = true
			break
		}
	}
	if !doesCurrencyExists {
		return store, nil
	}

	var currencies []models.Currency
	for _, c := range store.Currencies {
		if c.Code != code {
			currencies = append(currencies, c)
		}
	}
	store.Currencies = currencies
	if err := s.r.StoreRepository().Upsert(s.ctx, store); err != nil {
		return nil, err
	}
	return store, nil
}
