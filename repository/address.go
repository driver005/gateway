package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type AddressRepo struct {
	Repository[models.Address]
}

func AddressRepository(db *gorm.DB) AddressRepo {
	return AddressRepo{*NewRepository[models.Address](db)}
}
