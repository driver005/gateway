package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type UserRepo struct {
	sql.Repository[models.User]
}

func UserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{*sql.NewRepository[models.User](db)}
}
