package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	Repository[models.User]
}

func UserRepository(db *gorm.DB) UserRepo {
	return UserRepo{*NewRepository[models.User](db)}
}
