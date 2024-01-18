package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
)

type AdminCreateUserRequest struct {
	Email     string          `validate:"email"`
	FirstName string          `validate:"omitempty"`
	LastName  string          `validate:"omitempty"`
	Role      models.UserRole `validate:"omitempty"`
	Password  string          `validate:"required"`
}

type AdminUpdateUserRequest struct {
	FirstName string          `json:"first_name,omitempty"`
	LastName  string          `json:"last_name,omitempty"`
	Role      models.UserRole `json:"role,omitempty"`
	ApiToken  string          `json:"api_token,omitempty"`
	Metadata  core.JSONB      `json:"metadata,omitempty"`
}

type User struct {
	r Registry
}

func NewUser(r Registry) *User {
	m := User{r: r}
	return &m
}

func (m *User) SetRoutes(router fiber.Router) {
	route := router.Group("/users")
	route.Get("/:user_id", m.Get)
	route.Post("/", m.Create)
	route.Post("/:user_id", m.Update)
	route.Delete("/:user_id", m.Delete)
	route.Get("/", m.List)

	route.Post("/password-tocken", m.ResetPasswordTocken)
	route.Post("/reste-password", m.ResetPassword)
}

func (m *User) Get(context fiber.Ctx) error {
	Id, err := utils.ParseUUID(context.Params("user_id"))
	if err != nil {
		return err
	}

	user, err := m.r.UserService().SetContext(context.Context()).Retrieve(Id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(user)
}

func (m *User) List(context fiber.Ctx) error {
	user, err := m.r.UserService().SetContext(context.Context()).List(types.FilterableUser{}, &sql.Options{})
	if err != nil {
		return err
	}
	return context.Status(fiber.StatusOK).JSON(user)
}

func (m *User) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateUserInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.UserService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	result.PasswordHash = ""
	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *User) Update(context fiber.Ctx) error {
	Id, err := utils.ParseUUID(context.Params("user_id"))
	if err != nil {
		return err
	}

	model, err := api.BindCreate[types.UpdateUserInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.UserService().SetContext(context.Context()).Update(Id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *User) Delete(context fiber.Ctx) error {
	Id, err := utils.ParseUUID(context.Params("user_id"))
	if err != nil {
		return err
	}

	if err := m.r.UserService().SetContext(context.Context()).Delete(Id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"id":      Id,
		"object":  "user",
		"deleted": true,
	})
}

func (m *User) ResetPassword(context fiber.Ctx) error {
	return nil
}

func (m *User) ResetPasswordTocken(context fiber.Ctx) error {
	return nil
}
