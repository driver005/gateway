package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type PaymentProviderService struct {
	ctx               context.Context
	repo              *repository.PaymentProviderRepo
	paymentRepository *repository.PaymentRepo
	refundRepository  *repository.RefundRepo
}

func NewPaymentProviderService(
	ctx context.Context,
	repo *repository.PaymentProviderRepo,
	paymentRepository *repository.PaymentRepo,
	refundRepository *repository.RefundRepo,
) *PaymentProviderService {
	return &PaymentProviderService{
		ctx,
		repo,
		paymentRepository,
		refundRepository,
	}
}
