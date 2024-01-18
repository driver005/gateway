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
	"github.com/driver005/gateway/types"
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
	customer, err := s.RetrieveById(customerId, &sql.Options{
		Selects: []string{"id", "has_account", "password_hash", "email", "first_name", "last_name"},
	})
	if err != nil {
		return nil, err
	}

	if !customer.HasAccount {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"you must have an account to reset the password. Create an account first",
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
			nil,
		)
	}

	return tocken, nil
}

func (s *CustomerService) List(selector models.Customer, config *sql.Options, groups []string) ([]models.Customer, *utils.ApplictaionError) {
	res, _, err := s.ListAndCount(selector, config, groups)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerService) ListAndCount(selector models.Customer, config *sql.Options, groups []string) ([]models.Customer, *int64, *utils.ApplictaionError) {
	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	return s.r.CustomerRepository().ListAndCount(s.ctx, selector, config, groups)
}

func (s *CustomerService) Count() (*int64, *utils.ApplictaionError) {
	count, err := s.r.CustomerRepository().Count(s.ctx, sql.Query{})
	if err != nil {
		return nil, err
	}
	return count, nil
}

func (s *CustomerService) Retrieve(selector models.Customer, config *sql.Options) (*models.Customer, *utils.ApplictaionError) {
	var res *models.Customer
	query := sql.BuildQuery[models.Customer](selector, config)

	if err := s.r.CustomerRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CustomerService) RetrieveByEmail(email string, config *sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if email == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"email" must be defined`,
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

func (s *CustomerService) RetrieveUnregisteredByEmail(email string, config *sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if email == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"email" must be defined`,
			nil,
		)
	}
	return s.Retrieve(models.Customer{Email: strings.ToLower(email), HasAccount: false}, config)
}

func (s *CustomerService) RetrieveRegisteredByEmail(email string, config *sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if email == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"email" must be defined`,
			nil,
		)
	}
	return s.Retrieve(models.Customer{Email: strings.ToLower(email), HasAccount: true}, config)
}

func (s *CustomerService) ListByEmail(email string, config *sql.Options) ([]models.Customer, *utils.ApplictaionError) {
	if email == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"email" must be defined`,
			nil,
		)
	}

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(2)
	}

	return s.List(models.Customer{Email: strings.ToLower(email)}, config, nil)
}

func (s *CustomerService) RetrieveByPhone(phone string, config *sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if phone == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"phone" must be defined`,
			nil,
		)
	}
	return s.Retrieve(models.Customer{Phone: phone}, config)
}

func (s *CustomerService) RetrieveById(id uuid.UUID, config *sql.Options) (*models.Customer, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			nil,
		)
	}

	return s.Retrieve(models.Customer{Model: core.Model{Id: id}}, config)
}

func (s *CustomerService) Create(data *types.CreateCustomerInput) (*models.Customer, *utils.ApplictaionError) {
	if err := validator.New().Var(data.Email, "required,email"); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	var model *models.Customer

	model.Email = strings.ToLower(data.Email)

	existing, _ := s.ListByEmail(model.Email, &sql.Options{})

	if len(existing) != 0 {
		for _, exit := range existing {
			if exit.HasAccount && reflect.ValueOf(data.Password).IsZero() {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					`a customer with the given email already has an account. Log in instead`,
					"500",
					nil,
				)
			} else if !exit.HasAccount && reflect.ValueOf(data.Password).IsZero() {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					`guest customer with email already exists`,
					"500",
					nil,
				)
			}
		}
	}

	if !reflect.ValueOf(data.Password).IsZero() {
		hashedPassword, err := s.HashPassword(data.Password)
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

	if data.Metadata != nil {
		model.Metadata = utils.MergeMaps(model.Metadata, data.Metadata)
	}

	if !reflect.ValueOf(data.HasAccount).IsZero() {
		model.HasAccount = data.HasAccount
	}

	if !reflect.ValueOf(data.FirstName).IsZero() {
		model.FirstName = data.FirstName
	}
	if !reflect.ValueOf(data.LastName).IsZero() {
		model.LastName = data.LastName
	}
	if !reflect.ValueOf(data.Phone).IsZero() {
		model.Phone = data.Phone
	}

	if err := s.r.CustomerRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *CustomerService) Update(userId uuid.UUID, data *types.UpdateCustomerInput) (*models.Customer, *utils.ApplictaionError) {
	model, err := s.RetrieveById(userId, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if !reflect.ValueOf(data.Email).IsZero() {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"You are not allowed to Update email"`,
			nil,
		)
	}

	if userId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"userId" must be defined`,
			nil,
		)
	}

	if !reflect.ValueOf(data.Password).IsZero() {
		hashedPassword, err := s.HashPassword(data.Password)
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

	if data.Metadata != nil {
		model.Metadata = utils.MergeMaps(model.Metadata, data.Metadata)
	}

	if data.BillingAddress != nil || data.BillingAddressId != uuid.Nil {
		addr := utils.ToAddress(data.BillingAddress)
		if data.BillingAddressId != uuid.Nil {
			addr.Id = data.BillingAddressId
		}
		if err := s.UpdateBillingAddress(model, uuid.Nil, addr); err != nil {
			return nil, err
		}
	}
	if data.Groups != nil {
		for _, g := range data.Groups {
			model.Groups = append(model.Groups, models.CustomerGroup{
				Model: core.Model{Id: g.Id},
			})
		}
	}

	if !reflect.ValueOf(data.FirstName).IsZero() {
		model.FirstName = data.FirstName
	}
	if !reflect.ValueOf(data.LastName).IsZero() {
		model.LastName = data.LastName
	}
	if !reflect.ValueOf(data.Phone).IsZero() {
		model.Phone = data.Phone
	}

	if err := s.r.CustomerRepository().Upsert(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *CustomerService) UpdateBillingAddress(model *models.Customer, id uuid.UUID, address *models.Address) *utils.ApplictaionError {
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
		addr = address
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
			nil,
		)
	}

	if customerId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"customerId" must be defined`,
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
			nil,
		)
	}

	if customerId == uuid.Nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"customerId" must be defined`,
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
			nil,
		)
	}

	address.CountryCode = strings.ToLower(address.CountryCode)

	customer, err := s.RetrieveById(customerId, &sql.Options{
		Relations: []string{"shipping_addresses"},
	})
	if err != nil {
		return nil, nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"could not find address for customer",
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
	data, err := s.RetrieveById(customerId, &sql.Options{})
	if err != nil {
		return err
	}

	if err := s.r.CustomerRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}
