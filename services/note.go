package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type NoteService struct {
	ctx  context.Context
	repo *repository.NoteRepo
}

func NewNoteService(
	ctx context.Context,
	repo *repository.NoteRepo,
) *NoteService {
	return &NoteService{
		ctx,
		repo,
	}
}
