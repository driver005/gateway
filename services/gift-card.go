package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type GiftCardService struct {
	ctx                           context.Context
	repo                          *repository.GiftCardRepo
	giftCardTransactionRepository *repository.GiftCardTransactionRepo
}

func NewGiftCardService(
	ctx context.Context,
	repo *repository.GiftCardRepo,
	giftCardTransactionRepository *repository.GiftCardTransactionRepo,
) *GiftCardService {
	return &GiftCardService{
		ctx,
		repo,
		giftCardTransactionRepository,
	}
}
