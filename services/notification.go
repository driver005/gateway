package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type NotificationService struct {
	ctx                            context.Context
	repo                           *repository.NotificationRepo
	notificationProviderRepository *repository.NotificationProviderRepo
}

func NewNotificationService(
	ctx context.Context,
	repo *repository.NotificationRepo,
	notificationProviderRepository *repository.NotificationProviderRepo,
) *NotificationService {
	return &NotificationService{
		ctx,
		repo,
		notificationProviderRepository,
	}
}
