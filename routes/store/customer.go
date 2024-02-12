package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Customer struct {
	r Registry
}

func NewCustomer(r Registry) *Customer {
	m := Customer{r: r}
	return &m
}

func (m *Customer) SetRoutes(router fiber.Router) {
	route := router.Group("/customers")
	route.Post("", m.Create)

	route.Post("/password-tocken", m.ResetPasswordTocken)
	route.Post("/reste-password", m.ResetPassword)

	route.Use(utils.ConvertMiddleware(m.r.Middleware().AuthenticateCustomer())...)

	route.Get("/me", m.Get)
	route.Post("/me", m.Update)

	route.Post("/me/orders", m.ListOrders)
	route.Post("/me/addresses", m.CreateAddress)
	route.Post("/me/addresses/:address_id", m.UpdateAddress)
	route.Delete("/me/addresses/:address_id", m.DeleteAdress)
	route.Get("/me/payment-methods", m.GetPaymnetMethods)

}

func (m *Customer) Get(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Customer) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateCustomerInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Customer) Update(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	model, err := api.Bind[types.UpdateCustomerInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Customer) GetPaymnetMethods(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	paymentProviders, err := m.r.PaymentProviderService().SetContext(context.Context()).List()
	if err != nil {
		return err
	}

	var methods []types.PaymentMethod
	for _, paymentProvider := range paymentProviders {
		provider, err := m.r.PaymentProviderService().SetContext(context.Context()).RetrieveProvider(paymentProvider.Id)
		if err != nil {
			return err
		}

		pMethods := provider.RetrieveSavedMethods(customer)
		for _, pMethod := range pMethods {
			methods = append(methods, types.PaymentMethod{
				ProviderId: paymentProvider.Id,
				Data:       structs.Map(pMethod),
			})
		}
	}

	return context.Status(fiber.StatusOK).JSON(methods)
}

func (m *Customer) ListOrders(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	model, config, err := api.BindList[types.FilterableOrder](context)
	if err != nil {
		return err
	}

	model.CustomerId = id

	result, count, err := m.r.OrderService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

func (m *Customer) CreateAddress(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	model, err := api.BindCreate[types.CustomerAddAddress](context, m.r.Validator())
	if err != nil {
		return err
	}

	if _, _, err := m.r.CustomerService().SetContext(context.Context()).AddAddress(id, utils.CreateToAddress(model.Address)); err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Customer) UpdateAddress(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	model, addressId, err := api.BindUpdate[types.AddressPayload](context, "address_id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.CustomerService().SetContext(context.Context()).UpdateAddress(id, addressId, utils.ToAddress(model)); err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Customer) DeleteAdress(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	addressId, err := api.BindDelete(context, "address_id")
	if err != nil {
		return err
	}

	if err := m.r.CustomerService().SetContext(context.Context()).RemoveAddress(id, addressId); err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Customer) ResetPassword(context fiber.Ctx) error {
	model, err := api.BindCreate[types.UserResetPasswordRequest](context, m.r.Validator())
	if err != nil {
		return err
	}

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveRegisteredByEmail(model.Email, &sql.Options{Selects: []string{"id", "password_hash"}})
	if err != nil {
		return err
	}

	tocken, claims, er := m.r.TockenService().VerifyTokenWithSecret(model.Token, []byte(customer.PasswordHash))
	if er != nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			er.Error(),
		)
	}

	if tocken == nil || claims["customer_id"] != customer.Id {
		return utils.NewApplictaionError(
			utils.UNAUTHORIZED,
			"Invalid or expired password reset token",
		)
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Update(customer.Id, &types.UpdateCustomerInput{
		Password: model.Password,
	})
	if err != nil {
		return err
	}

	//TODO: Check If working
	result.PasswordHash = ""
	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Customer) ResetPasswordTocken(context fiber.Ctx) error {
	model, err := api.BindCreate[types.UserResetPasswordToken](context, m.r.Validator())
	if err != nil {
		return err
	}

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveRegisteredByEmail(model.Email, &sql.Options{})
	if err != nil {
		return err
	}

	if customer != nil {
		if _, err := m.r.CustomerService().SetContext(context.Context()).GenerateResetPasswordToken(customer.Id); err != nil {
			return err
		}
	}

	return context.SendStatus(204)
}
