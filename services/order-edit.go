package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type OrderEditService struct {
	ctx  context.Context
	repo *repository.OrderEditRepo
}

func NewOrderEditService(
	ctx context.Context,
	repo *repository.OrderEditRepo,
) *OrderEditService {
	return &OrderEditService{
		ctx,
		repo,
	}
}
