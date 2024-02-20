package admin

import (
	"fmt"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
)

type Auth struct {
	r    Registry
	name string
}

func NewAuth(r Registry) *Auth {
	m := Auth{r: r, name: "user"}
	return &m
}

func (m *Auth) SetRoutes(router fiber.Router) {
	route := router.Group("/auth")
	route.Get("", m.GetSession, m.r.Middleware().Authenticate())
	route.Post("", m.CreateSession)
	route.Delete("", m.DeleteSession, m.r.Middleware().Authenticate())
	route.Post("/tocken", m.GetTocken, m.r.Middleware().Authenticate())
}

// @oas:path [post] /admin/auth
// operationId: "PostAuth"
// summary: "User Login"
// x-authenticated: false
// description: "Log a User in and includes the Cookie session in the response header. The cookie session can be used in subsequent requests to authorize the user to perform admin functionalities.
// When using Medusa's JS or Medusa React clients, the cookie is automatically attached to subsequent requests."
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostAuthReq"
//
// x-codegen:
//
//	method: createSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.admin.auth.createSession({
//     email: "user@example.com",
//     password: "supersecret"
//     })
//     .then(({ user }) => {
//     console.log(user.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminLogin } from "medusa-react"
//
//     const Login = () => {
//     const adminLogin = useAdminLogin()
//     // ...
//
//     const handleLogin = () => {
//     adminLogin.mutate({
//     email: "user@example.com",
//     password: "supersecret",
//     }, {
//     onSuccess: ({ user }) => {
//     console.log(user)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Login
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/auth' \
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
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminAuthRes"
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
//	        $ref:  "#/components/responses/incorrect_credentials"
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
		fmt.Println(er)
		return er
	}

	req, err := api.BindCreate[types.CreateAuth](context, m.r.Validator())
	if err != nil {
		return err
	}

	result := m.r.AuthService().SetContext(context.Context()).Authenticate(req.Email, req.Password)
	if result.Success && result.User != nil {
		sess.Set("user_id", result.User.Id.String())
		if err := sess.Save(); err != nil {
			return err
		}
		result.User.PasswordHash = ""
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			(m.name): result.User,
		})
	} else {
		return context.SendStatus(fiber.StatusUnauthorized)
	}
}

// @oas:path [delete] /admin/auth
// operationId: "DeleteAuth"
// summary: "User Logout"
// x-authenticated: true
// description: "Delete the current session for the logged in user. This will only work if you're using Cookie session for authentication. If the API token is still passed in the header,
// the user is still authorized to perform admin functionalities in other API Routes."
// x-codegen:
//
//	method: deleteSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in
//     medusa.admin.auth.deleteSession()
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteSession } from "medusa-react"
//
//     const Logout = () => {
//     const adminLogout = useAdminDeleteSession()
//     // ...
//
//     const handleLogout = () => {
//     adminLogout.mutate(undefined, {
//     onSuccess: () => {
//     // user logged out.
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Logout
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/auth' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Auth
//
// responses:
//
//	"200":
//	  description: "OK"
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

// @oas:path [get] /admin/auth
// operationId: "GetAuth"
// summary: "Get Current User"
// x-authenticated: true
// description: "Get the currently logged in user's details."
// x-codegen:
//
//	method: getSession
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.auth.getSession()
//     .then(({ user }) => {
//     console.log(user.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminGetSession } from "medusa-react"
//
//     const Profile = () => {
//     const { user, isLoading } = useAdminGetSession()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {user && <span>{user.email}</span>}
//     </div>
//     )
//     }
//
//     export default Profile
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/auth' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Auth
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminAuthRes"
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
func (m *Auth) GetSession(context fiber.Ctx) error {
	userId := utils.GetUser(context)

	result, err := m.r.UserService().SetContext(context.Context()).Retrieve(userId, &sql.Options{})
	if err != nil {
		return err
	}

	result.PasswordHash = ""
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/auth/token
// operationId: "PostToken"
// summary: "User Login (JWT)"
// x-authenticated: false
// description: "After a successful login, a JWT token is returned, which can be used to send authenticated requests."
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostAuthReq"
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
//     medusa.admin.auth.getToken({
//     email: 'user@example.com',
//     password: 'supersecret'
//     })
//     .then(({ access_token }) => {
//     console.log(access_token);
//     })
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '{backend_url}/admin/auth/token' \
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
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminBearerAuthRes"
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
//	        $ref:  "#/components/responses/incorrect_credentials"
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

	result := m.r.AuthService().SetContext(context.Context()).Authenticate(req.Email, req.Password)
	if result.Success && result.User != nil {
		tocken, err := m.r.TockenService().SignToken(map[string]interface{}{
			"user_id": result.User.Id,
			"domain":  "admin",
		})
		if err != nil {
			return err
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"access_token": tocken,
		})
	} else {
		return context.SendStatus(fiber.StatusUnauthorized)
	}
}
