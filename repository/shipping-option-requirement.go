package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ShippingOptionRequirementRepo struct {
	Repository[models.ShippingOptionRequirement]
}

func ShippingOptionRequirementRepository(db *gorm.DB) ShippingOptionRequirementRepo {
	return ShippingOptionRequirementRepo{*NewRepository[models.ShippingOptionRequirement](db)}
}
