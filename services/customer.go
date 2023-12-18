package services

import (
	"context"
	"errors"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService struct {
	ctx               context.Context
	repo              *repository.CustomerRepo
	addressRepository *repository.AddressRepo
	tockenService     *TockenService
}

func NewCustomerService(
	ctx context.Context,
	repo *repository.CustomerRepo,
	addressRepository *repository.AddressRepo,
	tockenService *TockenService,
) *CustomerService {
	return &CustomerService{
		ctx,
		repo,
		addressRepository,
		tockenService,
	}
}

func (s *CustomerService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *CustomerService) GenerateResetPasswordToken(customerId uuid.UUID) (*string, error) {
	customer, err := s.RetrieveById(customerId, repository.Options{
		Selects: []string{"id", "has_account", "password_hash", "email", "first_name", "last_name"},
	})
	if err != nil {
		return nil, err
	}

	if !customer.HasAccount {
		return nil, errors.New("you must have an account to reset the password. create an account first")
	}

	expiry := math.Floor(float64(time.Now().Unix())/1000.0) + 60*15
	tocken, err := s.tockenService.SignTokenWithSecret(
		map[string]interface{}{
			"customer_id": customer.Id,
			"exp":         expiry,
		},
		[]byte(customer.PasswordHash),
	)
	if err != nil {
		return nil, err
	}

	return tocken, nil
}

func (s *CustomerService) List(selector models.Customer, config repository.Options, q *string, groups []string) ([]models.Customer, error) {
	if reflect.DeepEqual(config, repository.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	res, _, err := s.repo.ListAndCount(s.ctx, selector, config, q, groups)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerService) ListAndCount(selector models.Customer, config repository.Options, q *string, groups []string) ([]models.Customer, *int64, error) {
	if reflect.DeepEqual(config, repository.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	return s.repo.ListAndCount(s.ctx, selector, config, q, groups)
}

func (s *CustomerService) Count() (*int64, error) {
	count, err := s.repo.Count(s.ctx, repository.Query{})
	if err != nil {
		return nil, err
	}
	return count, nil
}

func (s *CustomerService) Retrieve(selector models.Customer, config repository.Options) (*models.Customer, error) {
	var res *models.Customer
	query := repository.BuildQuery[models.Customer](selector, config)

	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerService) RetrieveByEmail(email string, config repository.Options) (*models.Customer, error) {
	if email == "" {
		return nil, errors.New(`"email" must be defined`)
	}
	var res *models.Customer

	query := repository.BuildQuery[models.Customer](models.Customer{Email: strings.ToLower(email)}, config)

	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerService) RetrieveUnregisteredByEmail(email string, config repository.Options) (*models.Customer, error) {
	if email == "" {
		return nil, errors.New(`"email" must be defined`)
	}
	return s.Retrieve(models.Customer{Email: strings.ToLower(email), HasAccount: false}, config)
}

func (s *CustomerService) RetrieveRegisteredByEmail(email string, config repository.Options) (*models.Customer, error) {
	if email == "" {
		return nil, errors.New(`"email" must be defined`)
	}
	return s.Retrieve(models.Customer{Email: strings.ToLower(email), HasAccount: true}, config)
}

func (s *CustomerService) ListByEmail(email string, config repository.Options) ([]models.Customer, error) {
	if email == "" {
		return nil, errors.New(`"email" must be defined`)
	}

	if reflect.DeepEqual(config, repository.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(2)
	}

	return s.List(models.Customer{Email: strings.ToLower(email)}, config, nil, nil)
}

func (s *CustomerService) RetrieveByPhone(phone string, config repository.Options) (*models.Customer, error) {
	if phone == "" {
		return nil, errors.New(`"phone" must be defined`)
	}
	return s.Retrieve(models.Customer{Phone: phone}, config)
}

func (s *CustomerService) RetrieveById(id uuid.UUID, config repository.Options) (*models.Customer, error) {
	if id == uuid.Nil {
		return nil, errors.New(`"id" must be defined`)
	}

	return s.Retrieve(models.Customer{Model: core.Model{Id: id}}, config)
}

func (s *CustomerService) Create(model *models.Customer) (*models.Customer, error) {
	if err := validator.New().Var(model.Email, "required,email"); err != nil {
		return nil, err
	}

	model.Email = strings.ToLower(model.Email)

	existing, _ := s.ListByEmail(model.Email, repository.Options{})

	if len(existing) != 0 {
		for _, exit := range existing {
			if exit.HasAccount && model.Password != "" {
				return nil, errors.New(`a customer with the given email already has an account. Log in instead`)
			} else if !exit.HasAccount && model.Password == "" {
				return nil, errors.New(`guest customer with email already exists`)
			}
		}
	}

	if model.Password != "" {
		hashedPassword, err := s.HashPassword(model.Password)
		if err != nil {
			return nil, err
		}

		model.PasswordHash = hashedPassword
	}

	if err := s.repo.Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *CustomerService) Update(userId uuid.UUID, update *models.Customer) (*models.Customer, error) {
	if update.Email != "" {
		return nil, errors.New(`"You are not allowed to update email"`)
	}

	if update.PasswordHash != "" {
		return nil, errors.New("use dedicated field, `password` for password operations")
	}

	if userId == uuid.Nil {
		return nil, errors.New(`"userId" must be defined`)
	}

	update.Id = userId

	if err := s.repo.FindOne(s.ctx, update, repository.Query{}); err != nil {
		return nil, err
	}

	if update.Password != "" {
		hashedPassword, err := s.HashPassword(update.Password)
		if err != nil {
			return nil, err
		}

		update.PasswordHash = hashedPassword
	}

	if err := s.repo.Upsert(s.ctx, update); err != nil {
		return nil, err
	}

	return update, nil
}

func (s *CustomerService) UpdateBillingAddress(model *models.Customer, address models.Address, id uuid.UUID) error {
	if reflect.DeepEqual(address, models.Address{}) && id == uuid.Nil {
		model.BillingAddressId = uuid.NullUUID{}
	}

	var addr *models.Address

	if id != uuid.Nil {
		addr.Id = id

		if err := s.addressRepository.FindOne(s.ctx, addr, repository.Query{}); err != nil {
			return errors.New("address with id ${id} was not found")
		}
	} else {
		addr = &address
	}

	addr.CountryCode = strings.ToLower(addr.CountryCode)

	if addr.Id != uuid.Nil {
		model.BillingAddressId = uuid.NullUUID{
			UUID:  addr.Id,
			Valid: true,
		}
	} else {
		if model.BillingAddressId.UUID != uuid.Nil {
			addr.Id = model.BillingAddressId.UUID

			if err := s.addressRepository.Save(s.ctx, addr); err != nil {
				return err
			}
		} else {
			if err := s.addressRepository.Save(s.ctx, addr); err != nil {
				return err
			}

			model.BillingAddress = addr
		}
	}

	return nil
}

func (s *CustomerService) UpdateAddress(customerId uuid.UUID, addressId uuid.UUID, model *models.Address) (*models.Address, error) {
	if addressId == uuid.Nil {
		return nil, errors.New(`"userId" must be defined`)
	}

	if customerId == uuid.Nil {
		return nil, errors.New(`"customerId" must be defined`)
	}

	model.Id = addressId
	model.CustomerId.UUID = customerId

	model.CountryCode = strings.ToLower(model.CountryCode)

	if err := s.addressRepository.FindOne(s.ctx, model, repository.Query{}); err != nil {
		return nil, errors.New("could not find address for customer")
	}

	if err := s.addressRepository.Upsert(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *CustomerService) RemoveAddress(customerId uuid.UUID, addressId uuid.UUID) (err error) {
	if addressId == uuid.Nil {
		return errors.New(`"userId" must be defined`)
	}

	if customerId == uuid.Nil {
		return errors.New(`"customerId" must be defined`)
	}
	var model *models.Address
	model.Id = addressId
	model.CustomerId.UUID = customerId

	if err := s.addressRepository.FindOne(s.ctx, model, repository.Query{}); err != nil {
		return errors.New("could not find address for customer")
	}

	if err := s.addressRepository.SoftRemove(s.ctx, model); err != nil {
		return err
	}

	return nil
}

func (s *CustomerService) AddAddress(customerId uuid.UUID, address *models.Address) (*models.Customer, *models.Address, error) {
	if customerId == uuid.Nil {
		return nil, nil, errors.New(`"customerId" must be defined`)
	}

	address.CountryCode = strings.ToLower(address.CountryCode)

	customer, err := s.RetrieveById(customerId, repository.Options{
		Relations: []string{"shipping_addresses"},
	})
	if err != nil {
		return nil, nil, errors.New("could not find address for customer")
	}

	if reflect.DeepEqual(customer.ShippingAddresses, address) {
		address.CustomerId.UUID = customerId

		if err := s.addressRepository.Save(s.ctx, address); err != nil {
			return nil, address, err
		}
	}

	return customer, nil, nil
}

func (s *CustomerService) Delete(customerId uuid.UUID) error {
	data, err := s.RetrieveById(customerId, repository.Options{})
	if err != nil {
		return err
	}

	if err := s.repo.SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}
