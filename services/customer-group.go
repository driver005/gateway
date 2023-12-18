package services

import (
	"context"
	"errors"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerGroupService struct {
	ctx             context.Context
	repo            *repository.CustomerGroupRepo
	customerService *CustomerService
}

func NewCustomerGroupService(
	ctx context.Context,
	repo *repository.CustomerGroupRepo,
	customerService *CustomerService,
) *CustomerGroupService {
	return &CustomerGroupService{
		ctx,
		repo,
		customerService,
	}
}

func (s *CustomerGroupService) Retrieve(customerGroupId uuid.UUID, config repository.Options) (*models.CustomerGroup, error) {
	if customerGroupId == uuid.Nil {
		return nil, errors.New(`"customerGroupId" must be defined`)
	}

	var res *models.CustomerGroup
	query := repository.BuildQuery(map[string]interface{}{"id": customerGroupId}, config)
	if err := s.repo.FindOne(s.ctx, res, query); err == nil {
		return nil, errors.New(`CustomerGroup with id ` + customerGroupId.String() + ` was not found`)
	}
	return res, nil
}

func (s *CustomerGroupService) Create(model *models.CustomerGroup) (*models.CustomerGroup, error) {
	if err := s.repo.Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *CustomerGroupService) AddCustomers(id uuid.UUID, customerIds uuid.UUIDs) (*models.CustomerGroup, error) {
	res, err := s.repo.AddCustomers(s.ctx, id, customerIds)
	if err != nil {
		return nil, s.handleCreationFail(id, customerIds, err)
	}
	return res, nil
}

func (s *CustomerGroupService) Update(customerGroupId uuid.UUID, update *models.CustomerGroup) (*models.CustomerGroup, error) {
	update.Id = customerGroupId

	if err := s.repo.FindOne(s.ctx, update, repository.Query{}); err != nil {
		return nil, err
	}

	if err := s.repo.Upsert(s.ctx, update); err != nil {
		return nil, err
	}
	return update, nil

}

func (s *CustomerGroupService) Delete(groupId uuid.UUID) error {
	data, err := s.Retrieve(groupId, repository.Options{})
	if err != nil {
		return err
	}

	if err := s.repo.SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *CustomerGroupService) List(selector models.CustomerGroup, config repository.Options, q *string) ([]models.CustomerGroup, error) {
	customerGroups, _, err := s.ListAndCount(selector, config, q)
	if err != nil {
		return nil, err
	}
	return customerGroups, nil
}

func (s *CustomerGroupService) ListAndCount(selector models.CustomerGroup, config repository.Options, q *string) ([]models.CustomerGroup, *int64, error) {
	var res []models.CustomerGroup

	if q != nil {
		v := repository.ILike(*q)
		selector.Name = v
	}

	config.Relations = []string{}

	query := repository.BuildQuery(selector, config)

	count, err := s.repo.FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *CustomerGroupService) RemoveCustomer(id uuid.UUID, customerIds uuid.UUIDs) (*models.CustomerGroup, error) {
	res, err := s.repo.RemoveCustomers(s.ctx, id, customerIds)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerGroupService) handleCreationFail(id uuid.UUID, ids uuid.UUIDs, err error) error {
	if err == gorm.ErrForeignKeyViolated {
		s.Retrieve(id, repository.Options{})
		var nonExistingCustomers uuid.UUIDs
		_, err := s.customerService.List(models.Customer{Model: core.Model{Id: id}}, repository.Options{}, nil, []string{})
		if err != nil {
			nonExistingCustomers = append(nonExistingCustomers, id)
		}
		return errors.New(`The following customer ids do not exist: ` + strings.Join(nonExistingCustomers.Strings(), ", "))
	}
	return err
}
