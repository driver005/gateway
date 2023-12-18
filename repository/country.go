package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type CountryRepo struct {
	Repository[models.Country]
}

func CountryRepository(db *gorm.DB) CountryRepo {
	return CountryRepo{*NewRepository[models.Country](db)}
}
