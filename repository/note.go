package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type NoteRepo struct {
	sql.Repository[models.Note]
}

func NoteRepository(db *gorm.DB) *NoteRepo {
	return &NoteRepo{*sql.NewRepository[models.Note](db)}
}
