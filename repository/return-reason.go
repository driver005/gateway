package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ReturnReasonRepo struct {
	Repository[models.ReturnReason]
}

func ReturnReasonRepository(db *gorm.DB) ReturnReasonRepo {
	return ReturnReasonRepo{*NewRepository[models.ReturnReason](db)}
}
