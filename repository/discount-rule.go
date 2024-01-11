package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type DiscountRuleRepo struct {
	sql.Repository[models.DiscountRule]
}

func DiscountRuleRepository(db *gorm.DB) *DiscountRuleRepo {
	return &DiscountRuleRepo{*sql.NewRepository[models.DiscountRule](db)}
}
