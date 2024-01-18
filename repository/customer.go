package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"gorm.io/gorm"
)

type CustomerRepo struct {
	sql.Repository[models.Customer]
}

func CustomerRepository(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{*sql.NewRepository[models.Customer](db)}
}

func (r *CustomerRepo) ListAndCount(ctx context.Context, selector models.Customer, config *sql.Options, groups []string) ([]models.Customer, *int64, *utils.ApplictaionError) {
	var res []models.Customer

	if config.Q != nil {
		v := sql.ILike(*config.Q)
		selector.Email = v
		selector.FirstName = v
		selector.LastName = v
	}
	for _, g := range groups {
		var group models.CustomerGroup
		if err := group.ParseUUID(g); err != nil {
			return nil, nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				err.Error(),
				"500",
				nil,
			)
		}

		selector.Groups = append(selector.Groups, group)
	}

	query := sql.BuildQuery[models.Customer](selector, config)

	count, err := r.FindAndCount(ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}
