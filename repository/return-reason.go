package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ReturnReasonRepo struct {
	sql.Repository[models.ReturnReason]
}

func ReturnReasonRepository(db *gorm.DB) *ReturnReasonRepo {
	return &ReturnReasonRepo{*sql.NewRepository[models.ReturnReason](db)}
}
