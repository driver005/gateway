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

// @oas:path [post] /store/auth
// operationId: "PostAuth"
// summary: "Customer Login"
// description: "Log a customer in and includes the Cookie session in the response header. The cookie session can be used in subsequent requests to authenticate the customer.
// When using Medusa's JS or Medusa React clients, the cookie is automatically attached to subsequent requests."
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/PostAuthReq"
//
// x-codegen:
//
//	method: authenticate
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.auth.authenticate({
//     email: "user@example.com",
//     password: "user@example.com"
//     })
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/auth' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com",
//     "password": "supersecret"
//     }'
//
// tags:
//   - Auth
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreAuthRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  $ref: "#/components/responses/incorrect_credentials"
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

// @oas:path [delete] /store/auth
// operationId: "DeleteAuth"
// summary: "Customer Log out"
// description: "Delete the current session for the logged in customer."
// x-authenticated: true
// x-codegen:
//
//	method: deleteSession
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.auth.deleteSession()
//     .then(() => {
//     // customer logged out successfully
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '{backend_url}/store/auth' \
//     -H 'Authorization: Bearer {access_token}'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Auth
//
// responses:
//
//	"200":
//	  description: OK
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  $ref: "#/components/responses/unauthorized"
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

// @oas:path [get] /store/auth
// operationId: "GetAuth"
// summary: "Get Current Customer"
// description: "Retrieve the currently logged in Customer's details."
// x-authenticated: true
// x-codegen:
//
//	method: getSession
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged
//     medusa.auth.getSession()
//     .then(({ customer }) => {
//     console.log(customer.id);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/auth' \
//     -H 'Authorization: Bearer {access_token}'
//
// security:
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Auth
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreAuthRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  $ref: "#/components/responses/unauthorized"
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
func (m *Auth) GetSession(context fiber.Ctx) error {
	id := context.Locals("customer_id").(uuid.UUID)

	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(&result)
}

// @oas:path [post] /store/auth/token
// operationId: "PostToken"
// summary: "Customer Login (JWT)"
// x-authenticated: false
// description: "After a successful login, a JWT token is returned, which can be used to send authenticated requests."
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/StorePostAuthReq"
//
// x-codegen:
//
//	method: getToken
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.auth.getToken({
//     email: 'user@example.com',
//     password: 'supersecret'
//     })
//     .then(({ access_token }) => {
//     console.log(access_token);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/store/auth/token' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com",
//     "password": "supersecret"
//     }'
//
// tags:
//   - Auth
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreBearerAuthRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
//	"401":
//	  $ref: "#/components/responses/incorrect_credentials"
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

// @oas:path [get] /store/auth/{email}
// operationId: "GetAuthEmail"
// summary: "Check if Email Exists"
// description: "Check if there's a customer already registered with the provided email."
// parameters:
//   - in: path
//     name: email
//     schema:
//     type: string
//     format: email
//     required: true
//     description: The email to check.
//
// x-codegen:
//
//	method: exists
//
// x-codeSamples:
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.auth.exists("user@example.com")
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/auth/user@example.com'
//
// tags:
//   - Auth
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreGetAuthEmailRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
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
