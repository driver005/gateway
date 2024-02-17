package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Region struct {
	r Registry
}

func NewRegion(r Registry) *Region {
	m := Region{r: r}
	return &m
}

func (m *Region) SetRoutes(router fiber.Router) {
	route := router.Group("/regions")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

// @oas:path [get] /store/regions/{id}
// operationId: GetRegionsRegion
// summary: Get a Region
// description: "Retrieve a Region's details."
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
//     medusa.regions.retrieve(regionId)
//     .then(({ region }) => {
//     console.log(region.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useRegion } from "medusa-react"
//
//     type Props = {
//     regionId: string
//     }
//
//     const Region = ({ regionId }: Props) => {
//     const { region, isLoading } = useRegion(
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
//     curl '{backend_url}/store/regions/{id}'
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreRegionsRes"
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
func (m *Region) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /store/regions
// operationId: GetRegions
// summary: List Regions
// description: "Retrieve a list of regions. The regions can be filtered by fields such as `created_at`. The regions can also be paginated. This API Route is useful to
//
//	show the customer all available regions to choose from."
//
// externalDocs:
//
//	description: "How to use regions in a storefront"
//	url: "https://docs.medusajs.com/modules/regions-and-currencies/storefront/use-regions"
//
// parameters:
//   - (query) offset=0 {integer} The number of regions to skip when retrieving the regions.
//   - (query) limit=100 {integer} Limit the number of regions returned.
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
//
// x-codegen:
//
//	method: list
//	queryParams: StoreGetRegionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.regions.list()
//     .then(({ regions, count, limit, offset }) => {
//     console.log(regions.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useRegions } from "medusa-react"
//
//     const Regions = () => {
//     const { regions, isLoading } = useRegions()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {regions?.length && (
//     <ul>
//     {regions.map((region) => (
//     <li key={region.id}>
//     {region.name}
//     </li>
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
//     curl '{backend_url}/store/regions'
//
// tags:
//   - Regions
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreRegionsListRes"
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
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}
