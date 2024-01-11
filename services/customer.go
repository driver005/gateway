package services

import (
	"context"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService struct {
	ctx context.Context
	r   Registry
}

func NewCustomerService(
	r Registry,
) *CustomerService {
	return &CustomerService{
		context.Background(),
		r,
	}
}

func (s *CustomerService) SetContext(context context.Context) *CustomerService {
	s.ctx = context
	return s
}

func (s *CustomerService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *CustomerService) GenerateResetPasswordToken(customerId uuid.UUID) (*string, *utils.ApplictaionError) {
	customer, err := s.RetrieveById(customerId, sql.Options{
		Selects: []string{"id", "has_account", "password_hash", "email", "first_name", "last_name"},
	})
	if err != nil {
		return nil, err
	}

	if !customer.HasAccount {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"you must have an account to reset the password. Create an account first",
			"500",
			nil,
		)
	}

	expiry := math.Floor(float64(time.Now().Unix())/1000.0) + 60*15
	tocken, er := s.r.TockenService().SetContext(s.ctx).SignTokenWithSecret(
		map[string]interface{}{
			"customer_id": customer.Id,
			"exp":         expiry,
		},
		[]byte(customer.PasswordHash),
	)
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			er.Error(),
			"500",
			nil,
		)
	}

	return tocken, nil
}

func (s *CustomerService) List(selector models.Customer, config sql.Options, q *string, groups []string) ([]models.Customer, *utils.ApplictaionError) {
	res, _, err := s.ListAndCount(selector, config, q, groups)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerService) ListAndCount(selector models.Customer, config sql.Options, q *string, groups []string) ([]models.Customer, *int64, *utils.ApplictaionError) {
	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	return s.r.CustomerRepository().ListAndCount(s.ctx, selector, config, q, groups)
}

func (s *CustomerService) Count() (*int64, *utils.ApplictaionError) {
	count, err := s.r.CustomerRepository().Count(s.ctx, sql.Query{})
	if err != nil {
		return nil, err
	}
	return count, nil
}

func (s *CustomerService) Retrieve(selector models.Customer, config sql.Options) (*models.Customer, *utils.ApplictaionError) {
	var res *models.Customer
	query := sql.BuildQuery[models.Customer](selector, config)

	if err := s.r.CustomerRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerService) RetrieveByEmail(email string, config sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if email == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"email" must be defined`,
			"500",
			nil,
		)
	}
	var res *models.Customer

	query := sql.BuildQuery[models.Customer](models.Customer{Email: strings.ToLower(email)}, config)

	if err := s.r.CustomerRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerService) RetrieveUnregisteredByEmail(email string, config sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if email == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"email" must be defined`,
			"500",
			nil,
		)
	}
	return s.Retrieve(models.Customer{Email: strings.ToLower(email), HasAccount: false}, config)
}

func (s *CustomerService) RetrieveRegisteredByEmail(email string, config sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if email == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"email" must be defined`,
			"500",
			nil,
		)
	}
	return s.Retrieve(models.Customer{Email: strings.ToLower(email), HasAccount: true}, config)
}

func (s *CustomerService) ListByEmail(email string, config sql.Options) ([]models.Customer, *utils.ApplictaionError) {
	if email == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"email" must be defined`,
			"500",
			nil,
		)
	}

	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(2)
	}

	return s.List(models.Customer{Email: strings.ToLower(email)}, config, nil, nil)
}

func (s *CustomerService) RetrieveByPhone(phone string, config sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if phone == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"phone" must be defined`,
			"500",
			nil,
		)
	}
	return s.Retrieve(models.Customer{Phone: phone}, config)
}

func (s *CustomerService) RetrieveById(id uuid.UUID, config sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}

	return s.Retrieve(models.Customer{Model: core.Model{Id: id}}, config)
}

func (s *CustomerService) Create(model *models.Customer) (*models.Customer, *utils.ApplictaionError) {
	if err := validator.New().Var(model.Email, "required,email"); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			"500",
			nil,
		)
	}

	model.Email = strings.ToLower(model.Email)

	existing, _ := s.ListByEmail(model.Email, sql.Options{})

	if len(existing) != 0 {
		for _, exit := range existing {
			if exit.HasAccount && model.Password != "" {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					`a customer with the given email already has an account. Log in instead`,
					"500",
					nil,
				)
			} else if !exit.HasAccount && model.Password == "" {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					`guest customer with email already exists`,
					"500",
					nil,
				)
			}
		}
	}

	if model.Password != "" {
		hashedPassword, err := s.HashPassword(model.Password)
		if err != nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				err.Error(),
				"500",
				nil,
			)
		}

		model.PasswordHash = hashedPassword
	}

	if err := s.r.CustomerRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *CustomerService) Update(userId uuid.UUID, Update *models.Customer) (*models.Customer, *utils.ApplictaionError) {
	if Update.Email != "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"You are not allowed to Update email"`,
			"500",
			nil,
		)
	}

	if Update.PasswordHash != "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"use dedicated field, `password` for password operations",
			"500",
			nil,
		)
	}

	if userId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"userId" must be defined`,
			"500",
			nil,
		)
	}

	Update.Id = userId

	if err := s.r.CustomerRepository().FindOne(s.ctx, Update, sql.Query{}); err != nil {
		return nil, err
	}

	if Update.Password != "" {
		hashedPassword, err := s.HashPassword(Update.Password)
		if err != nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				err.Error(),
				"500",
				nil,
			)
		}

		Update.PasswordHash = hashedPassword
	}

	if err := s.r.CustomerRepository().Upsert(s.ctx, Update); err != nil {
		return nil, err
	}

	return Update, nil
}

func (s *CustomerService) UpdateBillingAddress(model *models.Customer, address models.Address, id uuid.UUID) *utils.ApplictaionError {
	if reflect.DeepEqual(address, models.Address{}) && id == uuid.Nil {
		model.BillingAddressId = uuid.NullUUID{}
	}

	var addr *models.Address

	if id != uuid.Nil {
		addr.Id = id

		if err := s.r.AddressRepository().FindOne(s.ctx, addr, sql.Query{}); err != nil {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				"address with id ${id} was not found",
				"500",
				nil,
			)
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

			if err := s.r.AddressRepository().Save(s.ctx, addr); err != nil {
				return err
			}
		} else {
			if err := s.r.AddressRepository().Save(s.ctx, addr); err != nil {
				return err
			}

			model.BillingAddress = addr
		}
	}

	return nil
}

func (s *CustomerService) UpdateAddress(customerId uuid.UUID, addressId uuid.UUID, model *models.Address) (*models.Address, *utils.ApplictaionError) {
	if addressId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"userId" must be defined`,
			"500",
			nil,
		)
	}

	if customerId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"customerId" must be defined`,
			"500",
			nil,
		)
	}

	model.Id = addressId
	model.CustomerId.UUID = customerId

	model.CountryCode = strings.ToLower(model.CountryCode)

	if err := s.r.AddressRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"could not find address for customer",
			"500",
			nil,
		)
	}

	if err := s.r.AddressRepository().Upsert(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *CustomerService) RemoveAddress(customerId uuid.UUID, addressId uuid.UUID) (err *utils.ApplictaionError) {
	if addressId == uuid.Nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"userId" must be defined`,
			"500",
			nil,
		)
	}

	if customerId == uuid.Nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"customerId" must be defined`,
			"500",
			nil,
		)
	}
	var model *models.Address
	model.Id = addressId
	model.CustomerId.UUID = customerId

	if err := s.r.AddressRepository().FindOne(s.ctx, model, sql.Query{}); err != nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"could not find address for customer",
			"500",
			nil,
		)
	}

	if err := s.r.AddressRepository().SoftRemove(s.ctx, model); err != nil {
		return err
	}

	return nil
}

func (s *CustomerService) AddAddress(customerId uuid.UUID, address *models.Address) (*models.Customer, *models.Address, *utils.ApplictaionError) {
	if customerId == uuid.Nil {
		return nil, nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"customerId" must be defined`,
			"500",
			nil,
		)
	}

	address.CountryCode = strings.ToLower(address.CountryCode)

	customer, err := s.RetrieveById(customerId, sql.Options{
		Relations: []string{"shipping_addresses"},
	})
	if err != nil {
		return nil, nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"could not find address for customer",
			"500",
			nil,
		)
	}

	if reflect.DeepEqual(customer.ShippingAddresses, address) {
		address.CustomerId.UUID = customerId

		if err := s.r.AddressRepository().Save(s.ctx, address); err != nil {
			return nil, address, err
		}
	}

	return customer, nil, nil
}

func (s *CustomerService) Delete(customerId uuid.UUID) *utils.ApplictaionError {
	data, err := s.RetrieveById(customerId, sql.Options{})
	if err != nil {
		return err
	}

	if err := s.r.CustomerRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}
