package repository

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO: ADD
type CustomerGroupRepo struct {
	sql.Repository[models.CustomerGroup]
}

func CustomerGroupRepository(db *gorm.DB) *CustomerGroupRepo {
	return &CustomerGroupRepo{*sql.NewRepository[models.CustomerGroup](db)}
}

func (r *CustomerGroupRepo) AddCustomers(ctx context.Context, groupId uuid.UUID, customerIds []uuid.UUID) (*models.CustomerGroup, *utils.ApplictaionError) {
	var customerGroup *models.CustomerGroup
	if err := r.FindOne(ctx, customerGroup, sql.Query{}); err != nil {
		return nil, err
	}

	for _, id := range customerIds {
		customerGroup.Customers = append(customerGroup.Customers, models.Customer{SoftDeletableModel: core.SoftDeletableModel{Id: id}})
	}

	if err := r.Update(ctx, customerGroup); err != nil {
		return nil, err
	}

	return customerGroup, nil
}

func (r *CustomerGroupRepo) RemoveCustomers(ctx context.Context, groupId uuid.UUID, customerIds []uuid.UUID) (*models.CustomerGroup, *utils.ApplictaionError) {
	var customerGroup *models.CustomerGroup
	if err := r.FindOne(ctx, customerGroup, sql.Query{}); err != nil {
		return nil, err
	}

	for index, customer := range customerGroup.Customers {
		for _, id := range customerIds {
			if customer.Id == id {
				customerGroup.Customers = utils.Remove[models.Customer](customerGroup.Customers, index)
			}
		}
	}

	if err := r.Update(ctx, customerGroup); err != nil {
		return nil, err
	}

	return customerGroup, nil
}
