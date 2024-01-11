package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ReturnRepo struct {
	sql.Repository[models.Return]
}

func ReturnRepository(db *gorm.DB) *ReturnRepo {
	return &ReturnRepo{*sql.NewRepository[models.Return](db)}
}
