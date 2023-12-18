package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type StoreService struct {
	ctx                context.Context
	repo               *repository.StoreRepo
	currencyRepository *repository.CurrencyRepo
}

func NewStoreService(
	ctx context.Context,
	repo *repository.StoreRepo,
	currencyRepository *repository.CurrencyRepo,
) *StoreService {
	return &StoreService{
		ctx,
		repo,
		currencyRepository,
	}
}
