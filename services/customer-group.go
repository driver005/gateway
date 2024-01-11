package services

import (
	"context"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type CustomerGroupService struct {
	ctx context.Context
	r   Registry
}

func NewCustomerGroupService(
	r Registry,
) *CustomerGroupService {
	return &CustomerGroupService{
		context.Background(),
		r,
	}
}

func (s *CustomerGroupService) SetContext(context context.Context) *CustomerGroupService {
	s.ctx = context
	return s
}

func (s *CustomerGroupService) Retrieve(customerGroupId uuid.UUID, config sql.Options) (*models.CustomerGroup, *utils.ApplictaionError) {
	if customerGroupId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"customerGroupId" must be defined`,
			"500",
			nil,
		)
	}

	var res *models.CustomerGroup
	query := sql.BuildQuery(models.CustomerGroup{Model: core.Model{Id: customerGroupId}}, config)
	if err := s.r.CustomerGroupRepository().FindOne(s.ctx, res, query); err == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`CustomerGroup with id `+customerGroupId.String()+` was not found`,
			"500",
			nil,
		)
	}
	return res, nil
}

func (s *CustomerGroupService) Create(model *models.CustomerGroup) (*models.CustomerGroup, *utils.ApplictaionError) {
	if err := s.r.CustomerGroupRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *CustomerGroupService) AddCustomers(id uuid.UUID, customerIds uuid.UUIDs) (*models.CustomerGroup, *utils.ApplictaionError) {
	res, err := s.r.CustomerGroupRepository().AddCustomers(s.ctx, id, customerIds)
	if err != nil {
		return nil, s.handleCreationFail(id, customerIds, err)
	}
	return res, nil
}

func (s *CustomerGroupService) Update(customerGroupId uuid.UUID, Update *models.CustomerGroup) (*models.CustomerGroup, *utils.ApplictaionError) {
	Update.Id = customerGroupId

	if err := s.r.CustomerGroupRepository().FindOne(s.ctx, Update, sql.Query{}); err != nil {
		return nil, err
	}

	if err := s.r.CustomerGroupRepository().Upsert(s.ctx, Update); err != nil {
		return nil, err
	}
	return Update, nil

}

func (s *CustomerGroupService) Delete(groupId uuid.UUID) *utils.ApplictaionError {
	data, err := s.Retrieve(groupId, sql.Options{})
	if err != nil {
		return err
	}

	if err := s.r.CustomerGroupRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *CustomerGroupService) List(selector models.CustomerGroup, config sql.Options, q *string) ([]models.CustomerGroup, *utils.ApplictaionError) {
	customerGroups, _, err := s.ListAndCount(selector, config, q)
	if err != nil {
		return nil, err
	}
	return customerGroups, nil
}

func (s *CustomerGroupService) ListAndCount(selector models.CustomerGroup, config sql.Options, q *string) ([]models.CustomerGroup, *int64, *utils.ApplictaionError) {
	var res []models.CustomerGroup

	if q != nil {
		v := sql.ILike(*q)
		selector.Name = v
	}

	config.Relations = []string{}

	query := sql.BuildQuery(selector, config)

	count, err := s.r.CustomerGroupRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *CustomerGroupService) RemoveCustomer(id uuid.UUID, customerIds uuid.UUIDs) (*models.CustomerGroup, *utils.ApplictaionError) {
	res, err := s.r.CustomerGroupRepository().RemoveCustomers(s.ctx, id, customerIds)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerGroupService) handleCreationFail(id uuid.UUID, ids uuid.UUIDs, err *utils.ApplictaionError) *utils.ApplictaionError {
	if err.Type == utils.DB_ERROR {
		s.Retrieve(id, sql.Options{})
		var nonExistingCustomers uuid.UUIDs
		_, err := s.r.CustomerService().SetContext(s.ctx).List(models.Customer{Model: core.Model{Id: id}}, sql.Options{}, nil, []string{})
		if err != nil {
			nonExistingCustomers = append(nonExistingCustomers, id)
		}
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			`The following customer ids do not exist: `+strings.Join(nonExistingCustomers.Strings(), ", "),
			"500",
			nil,
		)
	}
	return err
}
