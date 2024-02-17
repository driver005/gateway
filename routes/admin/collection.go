package admin

import (
	"github.com/driver005/gateway/api"
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
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/products", m.AddProducts)
	route.Delete("/:id/products", m.RemoveProducts)
}

// @oas:path [get] /admin/collections/{id}
// operationId: "GetCollectionsCollection"
// summary: "Get a Collection"
// description: "Retrieve a Product Collection by its ID. The products associated with it are expanded and returned as well."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product Collection
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
//     medusa.admin.collections.retrieve(collectionId)
//     .then(({ collection }) => {
//     console.log(collection.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCollection } from "medusa-react"
//
//     type Props = {
//     collectionId: string
//     }
//
//     const Collection = ({ collectionId }: Props) => {
//     const { collection, isLoading } = useAdminCollection(collectionId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {collection && <span>{collection.title}</span>}
//     </div>
//     )
//     }
//
//     export default Collection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/collections/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Collections
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCollectionsRes"
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
func (m *Collection) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ProductCollectionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [get] /admin/collections
// operationId: "GetCollections"
// summary: "List Collections"
// description: "Retrieve a list of Product Collection. The product collections can be filtered by fields such as `handle` or `title`. The collections can also be sorted or paginated."
// x-authenticated: true
// parameters:
//   - (query) limit=10 {integer} The number of collections to return.
//   - (query) offset=0 {integer} The number of collections to skip when retrieving the collections.
//   - (query) title {string} Filter collections by their title.
//   - (query) handle {string} Filter collections by their handle.
//   - (query) q {string} a term to search collections by their title or handle.
//   - (query) discount_condition_id {string} Filter collections by a discount condition ID associated with them.
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
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetCollectionsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.collections.list()
//     .then(({ collections, limit, offset, count }) => {
//     console.log(collections.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCollections } from "medusa-react"
//
//     const Collections = () => {
//     const { collections, isLoading } = useAdminCollections()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {collections && !collections.length && <span>
//     No Product Collections
//     </span>}
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
//     export default Collections
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/collections' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Collections
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCollectionsListRes"
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

// @oas:path [post] /admin/collections
// operationId: "PostCollections"
// summary: "Create a Collection"
// description: "Create a Product Collection."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostCollectionsReq"
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
//     medusa.admin.collections.create({
//     title: "New Collection"
//     })
//     .then(({ collection }) => {
//     console.log(collection.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateCollection } from "medusa-react"
//
//     const CreateCollection = () => {
//     const createCollection = useAdminCreateCollection()
//     // ...
//
//     const handleCreate = (title: string) => {
//     createCollection.mutate({
//     title
//     }, {
//     onSuccess: ({ collection }) => {
//     console.log(collection.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateCollection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/collections' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "New Collection"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Collections
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCollectionsRes"
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
func (m *Collection) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateProductCollection](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCollectionService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [post] /admin/collections/{id}
// operationId: "PostCollectionsCollection"
// summary: "Update a Collection"
// description: "Update a Product Collection's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Collection.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostCollectionsCollectionReq"
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
//     medusa.admin.collections.update(collectionId, {
//     title: "New Collection"
//     })
//     .then(({ collection }) => {
//     console.log(collection.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateCollection } from "medusa-react"
//
//     type Props = {
//     collectionId: string
//     }
//
//     const Collection = ({ collectionId }: Props) => {
//     const updateCollection = useAdminUpdateCollection(collectionId)
//     // ...
//
//     const handleUpdate = (title: string) => {
//     updateCollection.mutate({
//     title
//     }, {
//     onSuccess: ({ collection }) => {
//     console.log(collection.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Collection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/collections/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "New Collection"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Collections
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCollectionsRes"
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
func (m *Collection) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateProductCollection](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCollectionService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/collections/{id}
// operationId: "DeleteCollectionsCollection"
// summary: "Delete a Collection"
// description: "Delete a Product Collection. This does not delete associated products."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Collection.
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
//     medusa.admin.collections.delete(collectionId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteCollection } from "medusa-react"
//
//     type Props = {
//     collectionId: string
//     }
//
//     const Collection = ({ collectionId }: Props) => {
//     const deleteCollection = useAdminDeleteCollection(collectionId)
//     // ...
//
//     const handleDelete = (title: string) => {
//     deleteCollection.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Collection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/collections/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Collections
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCollectionsDeleteRes"
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
func (m *Collection) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ProductCollectionService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "product-collection",
		"deleted": true,
	})
}

// @oas:path [post] /admin/collections/{id}/products/batch
// operationId: "PostProductsToCollection"
// summary: "Add Products to Collection"
// description: "Add products to a product collection."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the product collection.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostProductsToCollectionReq"
//
// x-codegen:
//
//	method: addProducts
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.collections.addProducts(collectionId, {
//     product_ids: [
//     productId1,
//     productId2
//     ]
//     })
//     .then(({ collection }) => {
//     console.log(collection.products)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminAddProductsToCollection } from "medusa-react"
//
//     type Props = {
//     collectionId: string
//     }
//
//     const Collection = ({ collectionId }: Props) => {
//     const addProducts = useAdminAddProductsToCollection(collectionId)
//     // ...
//
//     const handleAddProducts = (productIds: string[]) => {
//     addProducts.mutate({
//     product_ids: productIds
//     }, {
//     onSuccess: ({ collection }) => {
//     console.log(collection.products)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Collection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/collections/{id}/products/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "product_ids": [
//     "prod_01G1G5V2MBA328390B5AXJ610F"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Collections
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminCollectionsRes"
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
func (m *Collection) AddProducts(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddProductsToCollectionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCollectionService().AddProducts(id, model.ProductIds)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

// @oas:path [delete] /admin/collections/{id}/products/batch
// operationId: "DeleteProductsFromCollection"
// summary: "Remove Products from Collection"
// description: "Remove a list of products from a collection. This would not delete the product, only the association between the product and the collection."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Product Collection.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeleteProductsFromCollectionReq"
//
// x-codegen:
//
//	method: removeProducts
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.collections.removeProducts(collectionId, {
//     product_ids: [
//     productId1,
//     productId2
//     ]
//     })
//     .then(({ id, object, removed_products }) => {
//     console.log(removed_products)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminRemoveProductsFromCollection } from "medusa-react"
//
//     type Props = {
//     collectionId: string
//     }
//
//     const Collection = ({ collectionId }: Props) => {
//     const removeProducts = useAdminRemoveProductsFromCollection(collectionId)
//     // ...
//
//     const handleRemoveProducts = (productIds: string[]) => {
//     removeProducts.mutate({
//     product_ids: productIds
//     }, {
//     onSuccess: ({ id, object, removed_products }) => {
//     console.log(removed_products)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default Collection
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/collections/{id}/products/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "product_ids": [
//     "prod_01G1G5V2MBA328390B5AXJ610F"
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Product Collections
//
// responses:
//
//	"200":
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminDeleteProductsFromCollectionRes"
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
func (m *Collection) RemoveProducts(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddProductsToCollectionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.ProductCollectionService().RemoveProducts(id, model.ProductIds); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":               id,
		"object":           "product-collection",
		"removed_products": model.ProductIds,
	})
}
