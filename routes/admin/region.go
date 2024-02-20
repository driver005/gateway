package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Region struct {
	r    Registry
	name string
}

func NewRegion(r Registry) *Region {
	m := Region{r: r, name: "region"}
	return &m
}

func (m *Region) SetRoutes(router fiber.Router) {
	route := router.Group("/regions")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/:id/fulfillment-options", m.GetFulfillmentOptions)
	route.Post("/:id/countries", m.AddCountry)
	route.Delete("/:id/countries:country_code", m.RemoveCountry)
	route.Post("/:id/payment-providers", m.AddPaymentProvider)
	route.Delete("/:id/payment-providers/:provider_id", m.RemovePaymentProvider)
	route.Post("/:id/fulfillment-providers", m.AddFullfilmentProvider)
	route.Delete("/:id/fulfillment-providers/:provider_id", m.RemoveFullfilmentProvider)
}

// @oas:path [get] /admin/regions/{id}
// operationId: "GetRegionsRegion"
// summary: "Get a Region"
// description: "Retrieve a Region's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Region.
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
//     medusa.admin.regions.retrieve(regionId)
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminRegion } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const { region, isLoading } = useAdminRegion(
//     regionId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {region && <span>{region.name}</span>}
//     </div>
//     )
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/regions/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsRes"
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
func (m *Region) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/regions
// operationId: "GetRegions"
// summary: "List Regions"
// description: "Retrieve a list of Regions. The regions can be filtered by fields such as `created_at`. The regions can also be paginated."
// x-authenticated: true
// parameters:
//   - (query) q {string} Term used to search regions' name.
//   - (query) order {string} A field to sort-order the retrieved regions by.
//   - in: query
//     name: limit
//     schema:
//     type: integer
//     default: 50
//     required: false
//     description: Limit the number of regions returned.
//   - in: query
//     name: offset
//     schema:
//     type: integer
//     default: 0
//     required: false
//     description: The number of regions to skip when retrieving the regions.
//   - in: query
//     name: created_at
//     required: false
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
//     required: false
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
//     required: false
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
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetRegionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.regions.list()
//     .then(({ regions, limit, offset, count }) => {
//     console.log(regions.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminRegions } from "medusa-react"
//
//     const Regions = () => {
//     const { regions, isLoading } = useAdminRegions()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {regions && !regions.length && <span>No Regions</span>}
//     {regions && regions.length > 0 && (
//     <ul>
//     {regions.map((region) => (
//     <li key={region.id}>{region.name}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Regions
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/regions' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsListRes"
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
func (m *Region) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableRegion](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.RegionService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"regions": result,
		"count":   count,
		"offset":  config.Skip,
		"limit":   config.Take,
	})
}

// @oas:path [post] /admin/regions
// operationId: "PostRegions"
// summary: "Create a Region"
// description: "Create a Region."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostRegionsReq"
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
//     medusa.admin.regions.create({
//     name: "Europe",
//     currency_code: "eur",
//     tax_rate: 0,
//     payment_providers: [
//     "manual"
//     ],
//     fulfillment_providers: [
//     "manual"
//     ],
//     countries: [
//     "DK"
//     ]
//     })
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateRegion } from "medusa-react"
//
//     type CreateData = {
//     name: string
//     currency_code: string
//     tax_rate: number
//     payment_providers: string[]
//     fulfillment_providers: string[]
//     countries: string[]
//     }
//
//     const CreateRegion = () => {
//     const createRegion = useAdminCreateRegion()
//     // ...
//
//     const handleCreate = (regionData: CreateData) => {
//     createRegion.mutate(regionData, {
//     onSuccess: ({ region }) => {
//     console.log(region.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateRegion
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/regions' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "Europe",
//     "currency_code": "eur",
//     "tax_rate": 0,
//     "payment_providers": [
//     "manual"
//     ],
//     "fulfillment_providers": [
//     "manual"
//     ],
//     "countries": [
//     "DK"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsRes"
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
func (m *Region) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateRegionInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/regions/{id}
// operationId: "PostRegionsRegion"
// summary: "Update a Region"
// description: "Update a Region's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Region.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostRegionsRegionReq"
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
//     medusa.admin.regions.update(regionId, {
//     name: "Europe"
//     })
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateRegion } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const updateRegion = useAdminUpdateRegion(regionId)
//     // ...
//
//     const handleUpdate = (
//     countries: string[]
//     ) => {
//     updateRegion.mutate({
//     countries,
//     }, {
//     onSuccess: ({ region }) => {
//     console.log(region.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/regions/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "Europe"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsRes"
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
func (m *Region) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateRegionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/regions/{id}
// operationId: "DeleteRegionsRegion"
// summary: "Delete a Region"
// description: "Delete a Region. Associated resources, such as providers or currencies are not deleted. Associated tax rates are deleted."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Region.
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
//     medusa.admin.regions.delete(regionId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteRegion } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const deleteRegion = useAdminDeleteRegion(regionId)
//     // ...
//
//     const handleDelete = () => {
//     deleteRegion.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/regions/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsDeleteRes"
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
func (m *Region) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.RegionService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "region",
		"deleted": true,
	})
}

// @oas:path [post] /admin/regions/{id}/countries
// operationId: "PostRegionsRegionCountries"
// summary: "Add Country"
// description: "Add a Country to the list of Countries in a Region."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Region.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostRegionsRegionCountriesReq"
//
// x-codegen:
//
//	method: addCountry
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.regions.addCountry(regionId, {
//     country_code: "dk"
//     })
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminRegionAddCountry } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const addCountry = useAdminRegionAddCountry(regionId)
//     // ...
//
//     const handleAddCountry = (
//     countryCode: string
//     ) => {
//     addCountry.mutate({
//     country_code: countryCode
//     }, {
//     onSuccess: ({ region }) => {
//     console.log(region.countries)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/regions/{region_id}/countries' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "country_code": "dk"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsRes"
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
func (m *Region) AddCountry(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.RegionCountries](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).AddCountry(id, model.CountryCode); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/regions/{id}/fulfillment-providers
// operationId: "PostRegionsRegionFulfillmentProviders"
// summary: "Add Fulfillment Provider"
// description: "Add a Fulfillment Provider to the list of fulfullment providers in a Region."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Region.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostRegionsRegionFulfillmentProvidersReq"
//
// x-codegen:
//
//	method: addFulfillmentProvider
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.regions.addFulfillmentProvider(regionId, {
//     provider_id: "manual"
//     })
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminRegionAddFulfillmentProvider
//     } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const addFulfillmentProvider =
//     useAdminRegionAddFulfillmentProvider(regionId)
//     // ...
//
//     const handleAddFulfillmentProvider = (
//     providerId: string
//     ) => {
//     addFulfillmentProvider.mutate({
//     provider_id: providerId
//     }, {
//     onSuccess: ({ region }) => {
//     console.log(region.fulfillment_providers)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/regions/{id}/fulfillment-providers' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "provider_id": "manual"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsRes"
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
func (m *Region) AddFullfilmentProvider(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.RegionFulfillmentProvider](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).AddFulfillmentProvider(id, model.ProviderId); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/regions/{id}/payment-providers
// operationId: "PostRegionsRegionPaymentProviders"
// summary: "Add Payment Provider"
// description: "Add a Payment Provider to the list of payment providers in a Region."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Region.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostRegionsRegionPaymentProvidersReq"
//
// x-codegen:
//
//	method: addPaymentProvider
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.regions.addPaymentProvider(regionId, {
//     provider_id: "manual"
//     })
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminRegionAddPaymentProvider
//     } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const addPaymentProvider =
//     useAdminRegionAddPaymentProvider(regionId)
//     // ...
//
//     const handleAddPaymentProvider = (
//     providerId: string
//     ) => {
//     addPaymentProvider.mutate({
//     provider_id: providerId
//     }, {
//     onSuccess: ({ region }) => {
//     console.log(region.payment_providers)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/regions/{id}/payment-providers' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "provider_id": "manual"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsRes"
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
func (m *Region) AddPaymentProvider(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.RegionPaymentProvider](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).AddPaymentProvider(id, model.ProviderId); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/regions/{id}/fulfillment-options
// operationId: "GetRegionsRegionFulfillmentOptions"
// summary: "List Fulfillment Options"
// description: "Retrieve a list of fulfillment options available in a Region."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Region.
//
// x-codegen:
//
//	method: retrieveFulfillmentOptions
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.regions.retrieveFulfillmentOptions(regionId)
//     .then(({ fulfillment_options }) => {
//     console.log(fulfillment_options.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminRegionFulfillmentOptions } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const {
//     fulfillment_options,
//     isLoading
//     } = useAdminRegionFulfillmentOptions(
//     regionId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {fulfillment_options && !fulfillment_options.length && (
//     <span>No Regions</span>
//     )}
//     {fulfillment_options &&
//     fulfillment_options.length > 0 && (
//     <ul>
//     {fulfillment_options.map((option) => (
//     <li key={option.provider_id}>
//     {option.provider_id}
//     </li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/regions/{id}/fulfillment-options' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminGetRegionsRegionFulfillmentOptionsRes"
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
func (m *Region) GetFulfillmentOptions(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	region, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	var fpsIds uuid.UUIDs
	for _, f := range region.FulfillmentProviders {
		fpsIds = append(fpsIds, f.Id)
	}

	result, err := m.r.FulfillmentProviderService().SetContext(context.Context()).ListFulfillmentOptions(fpsIds)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/regions/{id}/countries/{country_code}
// operationId: "PostRegionsRegionCountriesCountry"
// summary: "Remove Country"
// x-authenticated: true
// description: "Remove a Country from the list of Countries in a Region. The country will still be available in the system, and it can be used in other regions."
// parameters:
//   - (path) id=* {string} The ID of the Region.
//   - in: path
//     name: country_code
//     description: The 2 character ISO code for the Country.
//     required: true
//     schema:
//     type: string
//     externalDocs:
//     url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//     description: See a list of codes.
//
// x-codegen:
//
//	method: deleteCountry
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.regions.deleteCountry(regionId, "dk")
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminRegionRemoveCountry } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const removeCountry = useAdminRegionRemoveCountry(regionId)
//     // ...
//
//     const handleRemoveCountry = (
//     countryCode: string
//     ) => {
//     removeCountry.mutate(countryCode, {
//     onSuccess: ({ region }) => {
//     console.log(region.countries)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/regions/{id}/countries/{country_code}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsRes"
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
func (m *Region) RemoveCountry(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	countryCode := context.Params("country_code")

	if _, err := m.r.RegionService().SetContext(context.Context()).RemoveCountry(id, countryCode); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/regions/{id}/fulfillment-providers/{provider_id}
// operationId: "PostRegionsRegionFulfillmentProvidersProvider"
// summary: "Remove Fulfillment Provider"
// description: "Remove a Fulfillment Provider from a Region. The fulfillment provider will still be available for usage in other regions."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Region.
//   - (path) provider_id=* {string} The ID of the Fulfillment Provider.
//
// x-codegen:
//
//	method: deleteFulfillmentProvider
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.regions.deleteFulfillmentProvider(regionId, "manual")
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminRegionDeleteFulfillmentProvider
//     } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const removeFulfillmentProvider =
//     useAdminRegionDeleteFulfillmentProvider(regionId)
//     // ...
//
//     const handleRemoveFulfillmentProvider = (
//     providerId: string
//     ) => {
//     removeFulfillmentProvider.mutate(providerId, {
//     onSuccess: ({ region }) => {
//     console.log(region.fulfillment_providers)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/regions/{id}/fulfillment-providers/{provider_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsRes"
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
func (m *Region) RemoveFullfilmentProvider(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).RemoveFulfillmentProvider(id, providerId); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/regions/{id}/payment-providers/{provider_id}
// operationId: "PostRegionsRegionPaymentProvidersProvider"
// summary: "Remove Payment Provider"
// description: "Remove a Payment Provider from a Region. The payment provider will still be available for usage in other regions."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Region.
//   - (path) provider_id=* {string} The ID of the Payment Provider.
//
// x-codegen:
//
//	method: deletePaymentProvider
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.regions.deletePaymentProvider(regionId, "manual")
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminRegionDeletePaymentProvider
//     } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({
//     regionId
//     }: Props) => {
//     const removePaymentProvider =
//     useAdminRegionDeletePaymentProvider(regionId)
//     // ...
//
//     const handleRemovePaymentProvider = (
//     providerId: string
//     ) => {
//     removePaymentProvider.mutate(providerId, {
//     onSuccess: ({ region }) => {
//     console.log(region.payment_providers)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Region
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/regions/{id}/payment-providers/{provider_id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminRegionsRes"
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
func (m *Region) RemovePaymentProvider(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).RemovePaymentProvider(id, providerId); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}
