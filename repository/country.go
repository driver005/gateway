package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type CountryRepo struct {
	sql.Repository[models.Country]
}

func CountryRepository(db *gorm.DB) *CountryRepo {
	return &CountryRepo{*sql.NewRepository[models.Country](db)}
}
