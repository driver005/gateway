package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"gorm.io/gorm"
)

type CustomerRepo struct {
	sql.Repository[models.Customer]
}

func CustomerRepository(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{*sql.NewRepository[models.Customer](db)}
}

func (r *CustomerRepo) ListAndCount(ctx context.Context, selector *types.FilterableCustomer, config *sql.Options) ([]models.Customer, *int64, *utils.ApplictaionError) {
	var res []models.Customer

	if config.Q != nil {
		v := sql.ILike(*config.Q)
		selector.Email = v
		selector.FirstName = v
		selector.LastName = v
	}
	// for _, g := range selector.Groups {
	// 	var group models.CustomerGroup

	// 	selector.Groups = append(selector.Groups, group)
	// }

	query := sql.BuildQuery(selector, config)

	count, err := r.FindAndCount(ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}
