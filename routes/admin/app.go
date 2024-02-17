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

// @oas:path [get] /admin/apps
// operationId: "GetApps"
// summary: "List Applications"
// description: "Retrieve a list of applications registered in the Medusa backend."
// x-authenticated: true
// x-codegen:
//
//	method: list
//
// x-codeSamples:
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/apps' \
//     -H x-medusa-access-token: "{api_token}"
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Apps Oauth
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminAppsListRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description:  "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
func (m *App) List(context fiber.Ctx) error {
	result, err := m.r.OAuthService().SetContext(context.Context()).List(&models.OAuth{}, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/apps/authorizations
// operationId: "PostApps"
// summary: "Generate Token for App"
// description: "Use an app's Oauth provider to generate and store a new token for authentication."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostAppsReq"
//
// x-codegen:
//
//	method: authorize
//
// x-codeSamples:
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/apps/authorizations' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "application_name": "example",
//     "state": "ready",
//     "code": "token"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Apps Oauth
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminAppsRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  description: "Unauthorized"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/unauthorized"
//	"404":
//	  description: "Not Found"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/not_found_error"
//	"409":
//	  description: "Invalid State"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_state_error"
//	"422":
//	  description: "Invalid Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/invalid_request_error"
//	"500":
//	  description: "Internal Server"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/responses/500_error"
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
