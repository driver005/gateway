package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type App struct {
	r Registry
}

func NewApp(r Registry) *App {
	m := App{r: r}
	return &m
}

func (m *App) SetRoutes(router fiber.Router) {
	route := router.Group("/apps")
	route.Get("", m.List)
	route.Post("/authorizations", m.Authorize)
}

func (m *App) List(context fiber.Ctx) error {
	result, err := m.r.OAuthService().SetContext(context.Context()).List(&models.OAuth{}, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *App) Authorize(context fiber.Ctx) error {
	model, err := api.BindCreate[types.OauthAuthorize](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.OAuthService().SetContext(context.Context()).GenerateToken(model.ApplicationName, model.Code, model.State)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
