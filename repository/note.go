package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type NoteRepo struct {
	Repository[models.Note]
}

func NoteRepository(db *gorm.DB) NoteRepo {
	return NoteRepo{*NewRepository[models.Note](db)}
}
