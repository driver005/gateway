package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ShippingOptionRequirementRepo struct {
	sql.Repository[models.ShippingOptionRequirement]
}

func ShippingOptionRequirementRepository(db *gorm.DB) *ShippingOptionRequirementRepo {
	return &ShippingOptionRequirementRepo{*sql.NewRepository[models.ShippingOptionRequirement](db)}
}
