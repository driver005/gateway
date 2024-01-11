package routes

import (
	"github.com/driver005/gateway/routes/admin"
	"github.com/gofiber/fiber/v3"
)

type Registry interface {
	AdminRouter() fiber.Router
	StoreRouter() fiber.Router

	//Admin
	AdminAuth() *admin.Auth
}
