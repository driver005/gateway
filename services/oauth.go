package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type OAuthService struct {
	ctx  context.Context
	repo *repository.OAuthRepo
}

func NewOAuthService(
	ctx context.Context,
	repo *repository.OAuthRepo,
) *OAuthService {
	return &OAuthService{
		ctx,
		repo,
	}
}
