package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type StagedJobService struct {
	ctx  context.Context
	repo *repository.StagedJobRepo
}

func NewStagedJobService(ctx context.Context, repo *repository.StagedJobRepo) *StagedJobService {
	return &StagedJobService{
		ctx,
		repo,
	}
}
