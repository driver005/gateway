package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type CustomerRepo struct {
	Repository[models.Customer]
}

func CustomerRepository(db *gorm.DB) CustomerRepo {
	return CustomerRepo{*NewRepository[models.Customer](db)}
}

func (r *CustomerRepo) ListAndCount(ctx context.Context, selector models.Customer, config Options, q *string, groups []string) ([]models.Customer, *int64, error) {
	var res []models.Customer

	if q != nil {
		v := ILike(*q)
		selector.Email = v
		selector.FirstName = v
		selector.LastName = v
	}
	for _, g := range groups {
		var group models.CustomerGroup
		if err := group.ParseUUID(g); err != nil {
			return nil, nil, err
		}

		selector.Groups = append(selector.Groups, group)
	}

	query := BuildQuery[models.Customer](selector, config)

	count, err := r.FindAndCount(ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}
