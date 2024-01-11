package main

import (
	"context"

	"github.com/driver005/gateway/registry"
	// "github.com/gofiber/fiber/v3"
	// "github.com/gofiber/fiber/v3/middleware/cors"
	// "github.com/gofiber/fiber/v3/middleware/logger"
)

//Swagger: swag init -g main.go  --parseDependency --parseInternal --parseDepth 1  --output docs/

var ctx = context.Background()

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /api
// @schemes http
func main() {
	r := registry.New(ctx)

	r.Setup()
}
