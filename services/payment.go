package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type PaymentService struct {
	ctx  context.Context
	repo *repository.PaymentRepo
}

func NewPaymentService(
	ctx context.Context,
	repo *repository.PaymentRepo,
) *PaymentService {
	return &PaymentService{
		ctx,
		repo,
	}
}
