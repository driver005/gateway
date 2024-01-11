package admin

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AdminPostAuthReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Auth struct {
	r Registry
}

func NewAuth(r Registry) *Auth {
	m := Auth{r: r}
	return &m
}

func (m *Auth) SetRoutes(router fiber.Router) {
	route := router.Group("/auth")
	route.Get("/", m.GetSession, m.r.Middleware().Authenticate()...)
	route.Post("/", m.CreateSession)
	route.Delete("/", m.DeleteSession, m.r.Middleware().Authenticate()...)
	route.Post("/tocken", m.GetTocken, m.r.Middleware().Authenticate()...)
}

func (m *Auth) CreateSession(context fiber.Ctx) error {
	sess, err := m.r.Session().Get(context)
	if err != nil {
		return err
	}

	var req AdminPostAuthReq
	if err := context.Bind().Body(&req); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}

	result := m.r.AuthService().SetContext(context.Context()).Authenticate(req.Email, req.Password)
	if result.Success && result.User != nil {
		sess.Set("user_id", result.User.Id)
		if err := sess.Save(); err != nil {
			return err
		}
		result.User.PasswordHash = ""
		return context.Status(fiber.StatusOK).JSON(&result.User)
	} else {
		return context.Status(fiber.StatusUnauthorized).JSON(result)
	}
}

func (m *Auth) DeleteSession(context fiber.Ctx) error {
	sess, err := m.r.Session().Get(context)
	if err != nil {
		return err
	}

	if sess.Get("customer_id") != nil {
		sess.Delete("user_id")
	} else {
		if err := sess.Destroy(); err != nil {
			return err
		}
	}

	return context.Status(fiber.StatusOK).Send(nil)
}

func (m *Auth) GetSession(context fiber.Ctx) error {
	var userId uuid.UUID
	user, ok := context.Locals("user").(*models.User)
	if !ok {
		userId = context.Locals("user_id").(uuid.UUID)
	} else {
		userId = user.Id
	}

	user, err := m.r.UserService().SetContext(context.Context()).Retrieve(userId, sql.Options{})
	if err != nil {
		return err
	}

	user.PasswordHash = ""
	return context.Status(fiber.StatusOK).JSON(&user)
}

func (m *Auth) GetTocken(context fiber.Ctx) error {
	var req AdminPostAuthReq
	if err := context.Bind().Body(&req); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}

	result := m.r.AuthService().SetContext(context.Context()).Authenticate(req.Email, req.Password)
	if result.Success && result.User != nil {
		tocken, err := m.r.TockenService().SignToken(map[string]interface{}{
			"user_id": result.User.Id,
			"domain":  "admin",
		})
		if err != nil {
			return err
		}

		return context.Status(fiber.StatusOK).JSON(tocken)
	} else {
		return context.Status(fiber.StatusUnauthorized).JSON(result)
	}
}
