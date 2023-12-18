package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type PaymentCollectionService struct {
	ctx  context.Context
	repo *repository.PaymentCollectionRepo
}

func NewPaymentCollectionService(
	ctx context.Context,
	repo *repository.PaymentCollectionRepo,
) *PaymentCollectionService {
	return &PaymentCollectionService{
		ctx,
		repo,
	}
}
