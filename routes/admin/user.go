package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
)

type User struct {
	r Registry
}

func NewUser(r Registry) *User {
	m := User{r: r}
	return &m
}

func (m *User) UnauthenticatedUserRoutes(router fiber.Router) {
	route := router.Group("/users")

	route.Post("/password-tocken", m.ResetPasswordTocken)
	route.Post("/reste-password", m.ResetPassword)
}

func (m *User) SetRoutes(router fiber.Router) {
	route := router.Group("/users")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
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
	model, config, err := api.BindList[types.FilterableUser](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.UserService().SetContext(context.Context()).ListAndCount(model, config)
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
	model, id, err := api.BindUpdate[types.UpdateUserInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.UserService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *User) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.CustomerGroupService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "user",
		"deleted": true,
	})
}

func (m *User) ResetPassword(context fiber.Ctx) error {
	model, err := api.BindCreate[types.UserResetPasswordRequest](context, m.r.Validator())
	if err != nil {
		return err
	}

	user, err := m.r.UserService().SetContext(context.Context()).RetrieveByEmail(model.Email, &sql.Options{Selects: []string{"id", "password_hash"}})
	if err != nil {
		return err
	}

	tocken, claims, er := m.r.TockenService().VerifyTokenWithSecret(model.Token, []byte(user.PasswordHash))
	if er != nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			er.Error(),
		)
	}

	if tocken == nil || claims["user_id"] != user.Id {
		return context.Status(fiber.StatusUnauthorized).SendString("Invalid or expired password reset token")
	}

	result, err := m.r.UserService().SetContext(context.Context()).SetPassword(user.Id, model.Password)
	if err != nil {
		return err
	}

	result.PasswordHash = ""
	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *User) ResetPasswordTocken(context fiber.Ctx) error {
	model, err := api.BindCreate[types.UserResetPasswordToken](context, m.r.Validator())
	if err != nil {
		return err
	}

	user, err := m.r.UserService().SetContext(context.Context()).RetrieveByEmail(model.Email, &sql.Options{})
	if err != nil {
		return err
	}

	if user != nil {
		if _, err := m.r.UserService().SetContext(context.Context()).GenerateResetPasswordToken(user.Id); err != nil {
			return err
		}
	}

	return context.SendStatus(204)
}
