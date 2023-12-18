package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type DiscountRuleRepo struct {
	Repository[models.DiscountRule]
}

func DiscountRuleRepository(db *gorm.DB) DiscountRuleRepo {
	return DiscountRuleRepo{*NewRepository[models.DiscountRule](db)}
}
