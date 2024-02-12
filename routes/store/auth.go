package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Auth struct {
	r Registry
}

func NewAuth(r Registry) *Auth {
	m := Auth{r: r}
	return &m
}

func (m *Auth) SetRoutes(router fiber.Router) {
	route := router.Group("/auth")
	route.Get("", m.GetSession, m.r.Middleware().AuthenticateCustomer()...)
	route.Post("", m.CreateSession)
	route.Delete("", m.DeleteSession)
	route.Post("/tocken", m.GetTocken)
	route.Post("/:email", m.Exist)
}

func (m *Auth) CreateSession(context fiber.Ctx) error {
	sess, er := m.r.Session().Get(context)
	if er != nil {
		return er
	}

	req, err := api.BindCreate[types.CreateAuth](context, m.r.Validator())
	if err != nil {
		return err
	}

	data := m.r.AuthService().SetContext(context.Context()).AuthenticateCustomer(req.Email, req.Password)
	if !data.Success {
		return context.SendStatus(fiber.StatusUnauthorized)
	}

	sess.Set("customer_id", data.Customer.Id)
	if err := sess.Save(); err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(data.Customer.Id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(&result)
}

func (m *Auth) DeleteSession(context fiber.Ctx) error {
	sess, err := m.r.Session().Get(context)
	if err != nil {
		return err
	}

	if sess.Get("user_id") != nil {
		sess.Delete("customer_id")
	} else {
		if err := sess.Destroy(); err != nil {
			return err
		}
	}

	return context.SendStatus(fiber.StatusOK)
}

func (m *Auth) GetSession(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(&result)
}

func (m *Auth) GetTocken(context fiber.Ctx) error {
	req, err := api.BindCreate[types.CreateAuth](context, m.r.Validator())
	if err != nil {
		return err
	}

	result := m.r.AuthService().SetContext(context.Context()).AuthenticateCustomer(req.Email, req.Password)
	if result.Success && result.Customer != nil {
		tocken, err := m.r.TockenService().SignToken(map[string]interface{}{
			"customer_id": result.Customer.Id,
			"domain":      "store",
		})
		if err != nil {
			return err
		}

		return context.Status(fiber.StatusOK).JSON(tocken)
	} else {
		return context.SendStatus(fiber.StatusUnauthorized)
	}
}

func (m *Auth) Exist(context fiber.Ctx) error {
	email := context.Params("email")

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveRegisteredByEmail(email, &sql.Options{Selects: []string{"id", "has_account"}})
	if err != nil {
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"exists": false,
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"exists": result.HasAccount,
	})
}
