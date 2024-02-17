package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/interfaces"
	"github.com/gofiber/fiber/v3"
)

type StockLocation struct {
	r Registry
}

func NewStockLocation(r Registry) *StockLocation {
	m := StockLocation{r: r}
	return &m
}

func (m *StockLocation) SetRoutes(router fiber.Router) {
	route := router.Group("/stock-locations")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
}

// @oas:path [get] /admin/stock-locations/{id}
// operationId: "GetStockLocationsStockLocation"
// summary: "Get a Stock Location"
// description: "Retrieve a Stock Location's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Stock Location.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned stock location.
//   - (query) fields {string} Comma-separated fields that should be included in the returned stock location.
//
// x-codegen:
//
//	method: retrieve
//	queryParams: AdminGetStockLocationsLocationParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.stockLocations.retrieve(stockLocationId)
//     .then(({ stock_location }) => {
//     console.log(stock_location.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminStockLocation } from "medusa-react"
//
//     type Props = {
//     stockLocationId: string
//     }
//
//     const StockLocation = ({ stockLocationId }: Props) => {
//     const {
//     stock_location,
//     isLoading
//     } = useAdminStockLocation(stockLocationId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {stock_location && (
//     <span>{stock_location.name}</span>
//     )}
//     </div>
//     )
//     }
//
//     export default StockLocation
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/stock-locations/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Stock Locations
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminStockLocationsRes"
func (m *StockLocation) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.StockLocationService().Retrieve(context.Context(), id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /admin/stock-locations
// operationId: "GetStockLocations"
// summary: "List Stock Locations"
// description: "Retrieve a list of stock locations. The stock locations can be filtered by fields such as `name` or `created_at`. The stock locations can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) id {string} Filter by ID.
//   - (query) name {string} Filter by name.
//   - (query) order {string} A stock-location field to sort-order the retrieved stock locations by.
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
//   - (query) offset=0 {integer} The number of stock locations to skip when retrieving the stock locations.
//   - (query) limit=20 {integer} Limit the number of stock locations returned.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned stock locations.
//   - (query) fields {string} Comma-separated fields that should be included in the returned stock locations.
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetStockLocationsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.stockLocations.list()
//     .then(({ stock_locations, limit, offset, count }) => {
//     console.log(stock_locations.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminStockLocations } from "medusa-react"
//
//     function StockLocations() {
//     const {
//     stock_locations,
//     isLoading
//     } = useAdminStockLocations()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {stock_locations && !stock_locations.length && (
//     <span>No Locations</span>
//     )}
//     {stock_locations && stock_locations.length > 0 && (
//     <ul>
//     {stock_locations.map(
//     (location) => (
//     <li key={location.id}>{location.name}</li>
//     )
//     )}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default StockLocations
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/stock-locations' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Stock Locations
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminStockLocationsListRes"
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
func (m *StockLocation) List(context fiber.Ctx) error {
	model, config, err := api.BindList[interfaces.FilterableStockLocation](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.StockLocationService().ListAndCount(context.Context(), *model, config)
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

// @oas:path [post] /admin/stock-locations
// operationId: "PostStockLocations"
// summary: "Create a Stock Location"
// description: "Create a Stock Location."
// x-authenticated: true
// parameters:
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned stock location.
//   - (query) fields {string} Comma-separated fields that should be included in the returned stock location.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostStockLocationsReq"
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
//     medusa.admin.stockLocations.create({
//     name: "Main Warehouse",
//     })
//     .then(({ stock_location }) => {
//     console.log(stock_location.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateStockLocation } from "medusa-react"
//
//     const CreateStockLocation = () => {
//     const createStockLocation = useAdminCreateStockLocation()
//     // ...
//
//     const handleCreate = (name: string) => {
//     createStockLocation.mutate({
//     name,
//     }, {
//     onSuccess: ({ stock_location }) => {
//     console.log(stock_location.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateStockLocation
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/stock-locations' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "App"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Stock Locations
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminStockLocationsRes"
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
func (m *StockLocation) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[interfaces.CreateStockLocationInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.StockLocationService().Create(context.Context(), *model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/stock-locations/{id}
// operationId: "PostStockLocationsStockLocation"
// summary: "Update a Stock Location"
// description: "Update a Stock Location's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Stock Location.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned stock location.
//   - (query) fields {string} Comma-separated fields that should be included in the returned stock location.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostStockLocationsLocationReq"
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
//     medusa.admin.stockLocations.update(stockLocationId, {
//     name: 'Main Warehouse'
//     })
//     .then(({ stock_location }) => {
//     console.log(stock_location.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateStockLocation } from "medusa-react"
//
//     type Props = {
//     stockLocationId: string
//     }
//
//     const StockLocation = ({ stockLocationId }: Props) => {
//     const updateLocation = useAdminUpdateStockLocation(
//     stockLocationId
//     )
//     // ...
//
//     const handleUpdate = (
//     name: string
//     ) => {
//     updateLocation.mutate({
//     name
//     }, {
//     onSuccess: ({ stock_location }) => {
//     console.log(stock_location.name)
//     }
//     })
//     }
//     }
//
//     export default StockLocation
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/stock-locations/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "name": "Main Warehouse"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Stock Locations
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminStockLocationsRes"
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
func (m *StockLocation) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[interfaces.UpdateStockLocationInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.StockLocationService().Update(context.Context(), id, *model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/stock-locations/{id}
// operationId: "DeleteStockLocationsStockLocation"
// summary: "Delete a Stock Location"
// description: "Delete a Stock Location."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Stock Location.
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.stockLocations.delete(stockLocationId)
//     .then(({ id, object, deleted }) => {
//     console.log(id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteStockLocation } from "medusa-react"
//
//     type Props = {
//     stockLocationId: string
//     }
//
//     const StockLocation = ({ stockLocationId }: Props) => {
//     const deleteLocation = useAdminDeleteStockLocation(
//     stockLocationId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteLocation.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//     }
//
//     export default StockLocation
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/stock-locations/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Stock Locations
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminStockLocationsDeleteRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
func (m *StockLocation) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.StockLocationService().Delete(context.Context(), id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "stock-location",
		"deleted": true,
	})
}
