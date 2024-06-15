package services

import (
	"context"

	"github.com/driver005/gateway/plugins/customer/models"
	"github.com/driver005/gateway/utils"
	"gorm.io/gorm"
)

type Service struct {
	Context context.Context
	DB      *gorm.DB
}

func (s *Service) Create(data *models.Customer) error {
	if err := data.Create(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) Update(data *models.Customer) error {
	if err := data.Update(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) Upsert(data *models.Customer) error {
	if err := data.Upsert(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(data *models.Customer) error {
	if err := data.Delete(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) Find(data *models.Customer, query *utils.Query) (*models.Customer, error) {
	if err := data.FindOne(s.DB, s.Context, query); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) List(data *models.Customer, query *utils.Query) ([]models.Customer, error) {
	value, err := data.Find(s.DB, s.Context, query)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (s *Service) ListAndCount(data *models.Customer, query *utils.Query) ([]models.Customer, *int64, error) {
	value, count, err := data.FindAndCount(s.DB, s.Context, query)
	if err != nil {
		return nil, nil, err
	}

	return value, count, nil
}

func (s *Service) CreateCustomerGroup(data *models.CustomerGroup) error {
	if err := data.Create(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateCustomerGroup(data *models.CustomerGroup) error {
	if err := data.Update(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) AddCustomerToGroup(data *models.CustomerGroup, customer *models.Customer) error {
	data.Customer = append(data.Customer, *customer)
	if err := data.Update(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveCustomerToGroup(data *models.CustomerGroup, customer *models.Customer) error {
	for i, v := range data.Customer {
		if v.Id == customer.Id {
			data.Customer = append(data.Customer[:i], data.Customer[i+1:]...)
		}
	}

	if err := data.Update(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateAddress(data *models.CustomerAddress) error {
	if err := data.Update(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) AddAddress(data *models.Customer, address *models.CustomerAddress) error {
	data.CustomerAddress = append(data.CustomerAddress, *address)
	if err := data.Update(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveAddress(data *models.Customer, address *models.CustomerAddress) error {
	for i, v := range data.CustomerAddress {
		if v.Id == address.Id {
			data.CustomerAddress = append(data.CustomerAddress[:i], data.CustomerAddress[i+1:]...)
		}
	}

	if err := data.Update(s.DB, s.Context); err != nil {
		return err
	}

	return nil
}
