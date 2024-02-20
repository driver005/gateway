package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ShippingProfile struct {
	r    Registry
	name string
}

func NewShippingProfile(r Registry) *ShippingProfile {
	m := ShippingProfile{r: r, name: "shipping_profile"}
	return &m
}

func (m *ShippingProfile) SetRoutes(router fiber.Router) {
	route := router.Group("/shipping-profiles")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
}

// @oas:path [get] /admin/shipping-profiles/{id}
// operationId: "GetShippingProfilesProfile"
// summary: "Get a Shipping Profile"
// description: "Retrieve a Shipping Profile's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Shipping Profile.
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
//     medusa.admin.shippingProfiles.retrieve(profileId)
//     .then(({ shipping_profile }) => {
//     console.log(shipping_profile.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminShippingProfile } from "medusa-react"
//
//     type Props = {
//     shippingProfileId: string
//     }
//
//     const ShippingProfile = ({ shippingProfileId }: Props) => {
//     const {
//     shipping_profile,
//     isLoading
//     } = useAdminShippingProfile(
//     shippingProfileId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {shipping_profile && (
//     <span>{shipping_profile.name}</span>
//     )}
//     </div>
//     )
//     }
//
//     export default ShippingProfile
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/shipping-profiles/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Shipping Profiles
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminShippingProfilesRes"
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
func (m *ShippingProfile) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ShippingProfileService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/shipping-profiles
// operationId: "GetShippingProfiles"
// summary: "List Shipping Profiles"
// description: "Retrieve a list of Shipping Profiles."
// x-authenticated: true
// x-codegen:
//
//	method: list
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.shippingProfiles.list()
//     .then(({ shipping_profiles }) => {
//     console.log(shipping_profiles.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminShippingProfiles } from "medusa-react"
//
//     const ShippingProfiles = () => {
//     const {
//     shipping_profiles,
//     isLoading
//     } = useAdminShippingProfiles()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {shipping_profiles && !shipping_profiles.length && (
//     <span>No Shipping Profiles</span>
//     )}
//     {shipping_profiles && shipping_profiles.length > 0 && (
//     <ul>
//     {shipping_profiles.map((profile) => (
//     <li key={profile.id}>{profile.name}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default ShippingProfiles
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/shipping-profiles' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Shipping Profiles
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminShippingProfilesListRes"
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
func (m *ShippingProfile) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableShippingProfile](context)
	if err != nil {
		return err
	}
	result, err := m.r.ShippingProfileService().SetContext(context.Context()).List(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"shipping_profiles": result,
	})
}

// @oas:path [post] /admin/shipping-profiles
// operationId: "PostShippingProfiles"
// summary: "Create a Shipping Profile"
// description: "Create a Shipping Profile."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostShippingProfilesReq"
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
//     medusa.admin.shippingProfiles.create({
//     name: "Large Products"
//     })
//     .then(({ shipping_profile }) => {
//     console.log(shipping_profile.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { ShippingProfileType } from "@medusajs/medusa"
//     import { useAdminCreateShippingProfile } from "medusa-react"
//
//     const CreateShippingProfile = () => {
//     const createShippingProfile = useAdminCreateShippingProfile()
//     // ...
//
//     const handleCreate = (
//     name: string,
//     type: ShippingProfileType
//     ) => {
//     createShippingProfile.mutate({
//     name,
//     type
//     }, {
//     onSuccess: ({ shipping_profile }) => {
//     console.log(shipping_profile.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateShippingProfile
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/shipping-profiles' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "Large Products"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Shipping Profiles
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminShippingProfilesRes"
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
func (m *ShippingProfile) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateShippingProfile](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingProfileService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/shipping-profiles/{id}
// operationId: "PostShippingProfilesProfile"
// summary: "Update a Shipping Profile"
// description: "Update a Shipping Profile's details."
// parameters:
//   - (path) id=* {string} The ID of the Shipping Profile.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostShippingProfilesProfileReq"
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
//     medusa.admin.shippingProfiles.update(shippingProfileId, {
//     name: 'Large Products'
//     })
//     .then(({ shipping_profile }) => {
//     console.log(shipping_profile.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { ShippingProfileType } from "@medusajs/medusa"
//     import { useAdminUpdateShippingProfile } from "medusa-react"
//
//     type Props = {
//     shippingProfileId: string
//     }
//
//     const ShippingProfile = ({ shippingProfileId }: Props) => {
//     const updateShippingProfile = useAdminUpdateShippingProfile(
//     shippingProfileId
//     )
//     // ...
//
//     const handleUpdate = (
//     name: string,
//     type: ShippingProfileType
//     ) => {
//     updateShippingProfile.mutate({
//     name,
//     type
//     }, {
//     onSuccess: ({ shipping_profile }) => {
//     console.log(shipping_profile.name)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ShippingProfile
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/shipping-profiles/{id} \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "Large Products"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Shipping Profiles
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminShippingProfilesRes"
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
func (m *ShippingProfile) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateShippingProfile](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingProfileService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/shipping-profiles/{id}
// operationId: "DeleteShippingProfilesProfile"
// summary: "Delete a Shipping Profile"
// description: "Delete a Shipping Profile. Associated shipping options are deleted as well."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Shipping Profile.
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
//     medusa.admin.shippingProfiles.delete(profileId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteShippingProfile } from "medusa-react"
//
//     type Props = {
//     shippingProfileId: string
//     }
//
//     const ShippingProfile = ({ shippingProfileId }: Props) => {
//     const deleteShippingProfile = useAdminDeleteShippingProfile(
//     shippingProfileId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteShippingProfile.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default ShippingProfile
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/shipping-profiles/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Shipping Profiles
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDeleteShippingProfileRes"
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
func (m *ShippingProfile) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ShippingProfileService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "shipping-profile",
		"deleted": true,
	})
}
