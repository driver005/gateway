package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ReturnService struct {
	ctx                  context.Context
	repo                 *repository.ReturnRepo
	returnItemRepository *repository.ReturnItemRepo
}

func NewReturnService(
	ctx context.Context,
	repo *repository.ReturnRepo,
	returnItemRepository *repository.ReturnItemRepo,
) *ReturnService {
	return &ReturnService{
		ctx,
		repo,
		returnItemRepository,
	}
}
