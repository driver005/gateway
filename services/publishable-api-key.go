package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type PublishableApiKeyService struct {
	ctx                                     context.Context
	repo                                    *repository.PublishableApiKeyRepo
	publishableApiKeySalesChannelRepository *repository.PublishableApiKeySalesChannelRepo
}

func NewPublishableApiKeyService(
	ctx context.Context,
	repo *repository.PublishableApiKeyRepo,
	publishableApiKeySalesChannelRepository *repository.PublishableApiKeySalesChannelRepo,
) *PublishableApiKeyService {
	return &PublishableApiKeyService{
		ctx,
		repo,
		publishableApiKeySalesChannelRepository,
	}
}
