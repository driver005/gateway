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

// @oas:path [get] /admin/users/{id}
// operationId: "GetUsersUser"
// summary: "Get a User"
// description: "Retrieve an admin user's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the User.
//
// x-codegen:
//
//	method: retrieve
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.users.retrieve(userId)
//     .then(({ user }) => {
//     console.log(user.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUser } from "medusa-react"
//
//     type Props = {
//     userId: string
//     }
//
//     const User = ({ userId }: Props) => {
//     const { user, isLoading } = useAdminUser(
//     userId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {user && <span>{user.first_name} {user.last_name}</span>}
//     </div>
//     )
//     }
//
//     export default User
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/users/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Users
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminUserRes"
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
func (m *User) Get(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	user, err := m.r.UserService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(user)
}

// @oas:path [get] /admin/users
// operationId: "GetUsers"
// summary: "List Users"
// description: "Retrieves a list of users. The users can be filtered by fields such as `q` or `email`. The users can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) email {string} Filter by email.
//   - (query) first_name {string} Filter by first name.
//   - (query) last_name {string} Filter by last name.
//   - (query) q {string} Term used to search users' first name, last name, and email.
//   - (query) order {string} A user field to sort-order the retrieved users by.
//   - in: query
//     name: id
//     style: form
//     explode: false
//     description: Filter by user IDs.
//     schema:
//     oneOf:
//   - type: string
//     description: ID of the user.
//   - type: array
//     items:
//     type: string
//     description: ID of a user.
//   - in: query
//     name: created_at
//     description: Filter by a creation date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: updated_at
//     description: Filter by an update date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - in: query
//     name: deleted_at
//     description: Filter by a deletion date range.
//     schema:
//     type: object
//     properties:
//     lt:
//     type: string
//     description: filter by dates less than this date
//     format: date
//     gt:
//     type: string
//     description: filter by dates greater than this date
//     format: date
//     lte:
//     type: string
//     description: filter by dates less than or equal to this date
//     format: date
//     gte:
//     type: string
//     description: filter by dates greater than or equal to this date
//     format: date
//   - (query) offset=0 {integer} The number of users to skip when retrieving the users.
//   - (query) limit=20 {integer} Limit the number of users returned.
//   - (query) fields {string} Comma-separated fields that should be included in the returned users.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetUsersParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.users.list()
//     .then(({ users, limit, offset, count }) => {
//     console.log(users.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUsers } from "medusa-react"
//
//     const Users = () => {
//     const { users, isLoading } = useAdminUsers()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {users && !users.length && <span>No Users</span>}
//     {users && users.length > 0 && (
//     <ul>
//     {users.map((user) => (
//     <li key={user.id}>{user.email}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Users
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/users' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Users
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminUsersListRes"
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

// @oas:path [post] /admin/users
// operationId: "PostUsers"
// summary: "Create a User"
// description: "Create an admin User. The user has the same privileges as all admin users, and will be able to authenticate and perform admin functionalities right after creation."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminCreateUserRequest"
//
// x-codegen:
//
//	method: create
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.users.create({
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
//     import { useAdminCreateUser } from "medusa-react"
//
//     const CreateUser = () => {
//     const createUser = useAdminCreateUser()
//     // ...
//
//     const handleCreateUser = () => {
//     createUser.mutate({
//     email: "user@example.com",
//     password: "supersecret",
//     }, {
//     onSuccess: ({ user }) => {
//     console.log(user.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateUser
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/users' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com",
//     "password": "supersecret"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Users
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminUserRes"
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

// @oas:path [post] /admin/users/{id}
// operationId: "PostUsersUser"
// summary: "Update a User"
// description: "Update an admin user's details."
// parameters:
//   - (path) id=* {string} The ID of the User.
//
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminUpdateUserRequest"
//
// x-codegen:
//
//	method: update
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.users.update(userId, {
//     first_name: "Marcellus"
//     })
//     .then(({ user }) => {
//     console.log(user.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateUser } from "medusa-react"
//
//     type Props = {
//     userId: string
//     }
//
//     const User = ({ userId }: Props) => {
//     const updateUser = useAdminUpdateUser(userId)
//     // ...
//
//     const handleUpdateUser = (
//     firstName: string
//     ) => {
//     updateUser.mutate({
//     first_name: firstName,
//     }, {
//     onSuccess: ({ user }) => {
//     console.log(user.first_name)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default User
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/users/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "first_name": "Marcellus"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Users
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminUserRes"
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

// @oas:path [delete] /admin/users/{id}
// operationId: "DeleteUsersUser"
// summary: "Delete a User"
// description: "Delete a User. Once deleted, the user will not be able to authenticate or perform admin functionalities."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the User.
//
// x-codegen:
//
//	method: delete
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.users.delete(userId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteUser } from "medusa-react"
//
//     type Props = {
//     userId: string
//     }
//
//     const User = ({ userId }: Props) => {
//     const deleteUser = useAdminDeleteUser(userId)
//     // ...
//
//     const handleDeleteUser = () => {
//     deleteUser.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default User
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/users/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Users
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDeleteUserRes"
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

// @oas:path [post] /admin/users/reset-password
// operationId: "PostUsersUserPassword"
// summary: "Reset Password"
// description: "Reset the password of an admin User using their reset password token. A user must request to reset their password first before attempting to reset their
// password with this request."
// externalDocs:
//
//	description: How to reset a user's password
//	url: https://docs.medusajs.com/modules/users/admin/manage-profile#reset-password
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminResetPasswordRequest"
//
// x-codegen:
//
//	method: resetPassword
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.users.resetPassword({
//     token: "supersecrettoken",
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
//     import { useAdminResetPassword } from "medusa-react"
//
//     const ResetPassword = () => {
//     const resetPassword = useAdminResetPassword()
//     // ...
//
//     const handleResetPassword = (
//     token: string,
//     password: string
//     ) => {
//     resetPassword.mutate({
//     token,
//     password,
//     }, {
//     onSuccess: ({ user }) => {
//     console.log(user.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ResetPassword
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/users/reset-password' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "token": "supersecrettoken",
//     "password": "supersecret"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Users
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminUserRes"
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
		return utils.NewApplictaionError(
			utils.UNAUTHORIZED,
			"Invalid or expired password reset token",
		)
	}

	result, err := m.r.UserService().SetContext(context.Context()).SetPassword(user.Id, model.Password)
	if err != nil {
		return err
	}

	result.PasswordHash = ""
	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/users/password-token
// operationId: "PostUsersUserPasswordToken"
// summary: "Request Password Reset"
// description: "Generate a password token for an admin user with a given email. This also triggers the `user.password_reset` event. So, if you have a Notification Service installed
// that can handle this event, a notification, such as an email, will be sent to the user. The token is triggered as part of the `user.password_reset` event's payload.
// That token must be used later to reset the password using the [Reset Password](https://docs.medusajs.com/api/admin#users_postusersuserpassword) API Route."
// externalDocs:
//
//	description: How to reset a user's password
//	url: https://docs.medusajs.com/modules/users/admin/manage-profile#reset-password
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminResetPasswordTokenRequest"
//
// x-codegen:
//
//	method: sendResetPasswordToken
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.users.sendResetPasswordToken({
//     email: "user@example.com"
//     })
//     .then(() => {
//     // successful
//     })
//     .catch(() => {
//     // error occurred
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminSendResetPasswordToken } from "medusa-react"
//
//     const Login = () => {
//     const requestPasswordReset = useAdminSendResetPasswordToken()
//     // ...
//
//     const handleResetPassword = (
//     email: string
//     ) => {
//     requestPasswordReset.mutate({
//     email
//     }, {
//     onSuccess: () => {
//     // successful
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
//     curl -X POST '"{backend_url}"/admin/users/password-token' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "email": "user@example.com"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Users
//
// responses:
//
//	204:
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
