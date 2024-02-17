package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Collection struct {
	r Registry
}

func NewCollection(r Registry) *Collection {
	m := Collection{r: r}
	return &m
}

func (m *Collection) SetRoutes(router fiber.Router) {
	route := router.Group("/collections")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

// @oas:path [get] /store/collections/{id}
// operationId: "GetCollectionsCollection"
// summary: "Get a Collection"
// description: "Retrieve a Product Collection's details."
// parameters:
//   - (path) id=* {string} The id of the Product Collection
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
//     medusa.collections.retrieve(collectionId)
//     .then(({ collection }) => {
//     console.log(collection.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCollection } from "medusa-react"
//
//     type Props = {
//     collectionId: string
//     }
//
//     const ProductCollection = ({ collectionId }: Props) => {
//     const { collection, isLoading } = useCollection(collectionId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {collection && <span>{collection.title}</span>}
//     </div>
//     )
//     }
//
//     export default ProductCollection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/collections/{id}'
//
// tags:
//   - Product Collections
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCollectionsRes"
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
func (m *Collection) Get(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.ProductCollectionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /store/collections
// operationId: "GetCollections"
// summary: "List Collections"
// description: "Retrieve a list of product collections. The product collections can be filtered by fields such as `handle` or `created_at`. The product collections can also be paginated."
// parameters:
//   - (query) offset=0 {integer} The number of product collections to skip when retrieving the product collections.
//   - (query) limit=10 {integer} Limit the number of product collections returned.
//   - in: query
//     name: handle
//     style: form
//     explode: false
//     description: Filter by handles
//     schema:
//     type: array
//     items:
//     type: string
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
//	queryParams: StoreGetCollectionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     medusa.collections.list()
//     .then(({ collections, limit, offset, count }) => {
//     console.log(collections.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useCollections } from "medusa-react"
//
//     const ProductCollections = () => {
//     const { collections, isLoading } = useCollections()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {collections && collections.length === 0 && (
//     <span>No Product Collections</span>
//     )}
//     {collections && collections.length > 0 && (
//     <ul>
//     {collections.map((collection) => (
//     <li key={collection.id}>{collection.title}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default ProductCollections
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/collections'
//
// tags:
//   - Product Collections
//
// responses:
//
//	"200":
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreCollectionsListRes"
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
func (m *Collection) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableCollection](context)
	if err != nil {
		return err
	}

	result, count, err := m.r.ProductCollectionService().SetContext(context.Context()).ListAndCount(model, config)
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
