package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type AddressRepo struct {
	sql.Repository[models.Address]
}

func AddressRepository(db *gorm.DB) *AddressRepo {
	return &AddressRepo{*sql.NewRepository[models.Address](db)}
}
