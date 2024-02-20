package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type PublishableApiKey struct {
	r    Registry
	name string
}

func NewPublishableApiKey(r Registry) *PublishableApiKey {
	m := PublishableApiKey{r: r, name: "publishable_api_key"}
	return &m
}

func (m *PublishableApiKey) SetRoutes(router fiber.Router) {
	route := router.Group("/publishable-api-keys")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/revoke", m.Revoke)
	route.Get("/:id/sales-channels", m.ListChannels)
	route.Post("/:id/sales-channels/batch", m.AddChannelsBatch)
	route.Delete("/:id/sales-channels/batch", m.DeleteChannelsBatch)
}

// @oas:path [get] /admin/publishable-api-keys/{id}
// operationId: "GetPublishableApiKeysPublishableApiKey"
// summary: "Get a Publishable API Key"
// description: "Retrieve a publishable API key's details."
// parameters:
//   - (path) id=* {string} The ID of the Publishable API Key.
//
// x-authenticated: true
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
//     medusa.admin.publishableApiKeys.retrieve(publishableApiKeyId)
//     .then(({ publishable_api_key }) => {
//     console.log(publishable_api_key.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeletePublishableApiKey } from "medusa-react"
//
//     type Props = {
//     publishableApiKeyId: string
//     }
//
//     const PublishableApiKey = ({
//     publishableApiKeyId
//     }: Props) => {
//     const deleteKey = useAdminDeletePublishableApiKey(
//     publishableApiKeyId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteKey.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PublishableApiKey
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/publishable-api-keys/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Publishable Api Keys
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPublishableApiKeysRes"
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
func (m *PublishableApiKey) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/publishable-api-keys
// operationId: "GetPublishableApiKeys"
// summary: "List Publishable API keys"
// description: "Retrieve a list of publishable API keys. The publishable API keys can be filtered by fields such as `q`. The publishable API keys can also be paginated."
// x-authenticated: true
// parameters:
//   - (query) q {string} term to search publishable API keys' titles.
//   - (query) limit=20 {number} Limit the number of publishable API keys returned.
//   - (query) offset=0 {number} The number of publishable API keys to skip when retrieving the publishable API keys.
//   - (query) expand {string} Comma-separated relations that should be expanded in the returned publishable API keys.
//   - (query) fields {string} Comma-separated fields that should be included in the returned publishable API keys.
//
// x-codegen:
//
//	method: list
//	queryParams: GetPublishableApiKeysParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.publishableApiKeys.list()
//     .then(({ publishable_api_keys, count, limit, offset }) => {
//     console.log(publishable_api_keys)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { PublishableApiKey } from "@medusajs/medusa"
//     import { useAdminPublishableApiKeys } from "medusa-react"
//
//     const PublishableApiKeys = () => {
//     const { publishable_api_keys, isLoading } =
//     useAdminPublishableApiKeys()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {publishable_api_keys && !publishable_api_keys.length && (
//     <span>No Publishable API Keys</span>
//     )}
//     {publishable_api_keys &&
//     publishable_api_keys.length > 0 && (
//     <ul>
//     {publishable_api_keys.map(
//     (publishableApiKey: PublishableApiKey) => (
//     <li key={publishableApiKey.id}>
//     {publishableApiKey.title}
//     </li>
//     )
//     )}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default PublishableApiKeys
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/publishable-api-keys' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Publishable Api Keys
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPublishableApiKeysListRes"
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
func (m *PublishableApiKey) List(context fiber.Ctx) error {
	model, config, err := api.BindList[models.PublishableApiKey](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.PublishableApiKeyService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"publishable_api_keys": result,
		"count":                count,
		"offset":               config.Skip,
		"limit":                config.Take,
	})
}

// @oas:path [post] /admin/publishable-api-keys
// operationId: "PostPublishableApiKeys"
// summary: "Create Publishable API Key"
// description: "Create a Publishable API Key."
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostPublishableApiKeysReq"
//
// x-authenticated: true
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
//     medusa.admin.publishableApiKeys.create({
//     title
//     })
//     .then(({ publishable_api_key }) => {
//     console.log(publishable_api_key.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreatePublishableApiKey } from "medusa-react"
//
//     const CreatePublishableApiKey = () => {
//     const createKey = useAdminCreatePublishableApiKey()
//     // ...
//
//     const handleCreate = (title: string) => {
//     createKey.mutate({
//     title,
//     }, {
//     onSuccess: ({ publishable_api_key }) => {
//     console.log(publishable_api_key.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreatePublishableApiKey
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/publishable-api-keys' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "Web API Key"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Publishable Api Keys
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPublishableApiKeysRes"
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
func (m *PublishableApiKey) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreatePublishableApiKeyInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).Create(model, uuid.Nil)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/publishable-api-keys/{id}
// operationId: "PostPublishableApiKysPublishableApiKey"
// summary: "Update Publishable API Key"
// description: "Update a Publishable API Key's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Publishable API Key.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostPublishableApiKeysPublishableApiKeyReq"
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
//     medusa.admin.publishableApiKeys.update(publishableApiKeyId, {
//     title: "new title"
//     })
//     .then(({ publishable_api_key }) => {
//     console.log(publishable_api_key.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdatePublishableApiKey } from "medusa-react"
//
//     type Props = {
//     publishableApiKeyId: string
//     }
//
//     const PublishableApiKey = ({
//     publishableApiKeyId
//     }: Props) => {
//     const updateKey = useAdminUpdatePublishableApiKey(
//     publishableApiKeyId
//     )
//     // ...
//
//     const handleUpdate = (title: string) => {
//     updateKey.mutate({
//     title,
//     }, {
//     onSuccess: ({ publishable_api_key }) => {
//     console.log(publishable_api_key.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PublishableApiKey
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/publishable-api-key/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "title": "new title"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Publishable Api Keys
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPublishableApiKeysRes"
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
func (m *PublishableApiKey) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdatePublishableApiKeyInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/publishable-api-keys/{id}
// operationId: "DeletePublishableApiKeysPublishableApiKey"
// summary: "Delete Publishable API Key"
// description: "Delete a Publishable API Key. Associated resources, such as sales channels, are not deleted."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Publishable API Key to delete.
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
//     medusa.admin.publishableApiKeys.delete(publishableApiKeyId)
//     .then(({ id, object, deleted }) => {
//     console.log(id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeletePublishableApiKey } from "medusa-react"
//
//     type Props = {
//     publishableApiKeyId: string
//     }
//
//     const PublishableApiKey = ({
//     publishableApiKeyId
//     }: Props) => {
//     const deleteKey = useAdminDeletePublishableApiKey(
//     publishableApiKeyId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteKey.mutate(void 0, {
//     onSuccess: ({ id, object, deleted }) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PublishableApiKey
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/publishable-api-key/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Publishable Api Keys
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPublishableApiKeyDeleteRes"
//	"400":
//	  description: "Bad Request"
//	  content:
//	    application/json:
//	      schema:
//	        $ref:  "#/components/responses/400_error"
func (m *PublishableApiKey) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.PublishableApiKeyService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "publishable-api-key",
		"deleted": true,
	})
}

// @oas:path [post] /admin/publishable-api-keys/{id}/sales-channels/batch
// operationId: "PostPublishableApiKeySalesChannelsChannelsBatch"
// summary: "Add Sales Channels"
// description: "Add a list of sales channels to a publishable API key."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Publishable Api Key.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostPublishableApiKeySalesChannelsBatchReq"
//
// x-codegen:
//
//	method: addSalesChannelsBatch
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.publishableApiKeys.addSalesChannelsBatch(publishableApiKeyId, {
//     sales_channel_ids: [
//     {
//     id: channelId
//     }
//     ]
//     })
//     .then(({ publishable_api_key }) => {
//     console.log(publishable_api_key.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminAddPublishableKeySalesChannelsBatch,
//     } from "medusa-react"
//
//     type Props = {
//     publishableApiKeyId: string
//     }
//
//     const PublishableApiKey = ({
//     publishableApiKeyId
//     }: Props) => {
//     const addSalesChannels =
//     useAdminAddPublishableKeySalesChannelsBatch(
//     publishableApiKeyId
//     )
//     // ...
//
//     const handleAdd = (salesChannelId: string) => {
//     addSalesChannels.mutate({
//     sales_channel_ids: [
//     {
//     id: salesChannelId,
//     },
//     ],
//     }, {
//     onSuccess: ({ publishable_api_key }) => {
//     console.log(publishable_api_key.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PublishableApiKey
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/publishable-api-keys/{pak_id}/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "sales_channel_ids": [
//     {
//     "id": "{sales_channel_id}"
//     }
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Publishable Api Keys
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPublishableApiKeysRes"
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
func (m *PublishableApiKey) AddChannelsBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.PublishableApiKeySalesChannelsBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	var ids uuid.UUIDs
	for _, f := range model.SalesChannelIds {
		ids = append(ids, f.Id)
	}

	if err := m.r.PublishableApiKeyService().SetContext(context.Context()).AddSalesChannels(id, ids); err != nil {
		return err
	}

	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/publishable-api-keys/{id}/sales-channels/batch
// operationId: "DeletePublishableApiKeySalesChannelsChannelsBatch"
// summary: "Remove Sales Channels"
// description: "Remove a list of sales channels from a publishable API key. This doesn't delete the sales channels and only removes the association between them and the publishable API key."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Publishable API Key.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminDeletePublishableApiKeySalesChannelsBatchReq"
//
// x-codegen:
//
//	method: deleteSalesChannelsBatch
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.publishableApiKeys.deleteSalesChannelsBatch(publishableApiKeyId, {
//     sales_channel_ids: [
//     {
//     id: channelId
//     }
//     ]
//     })
//     .then(({ publishable_api_key }) => {
//     console.log(publishable_api_key.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminRemovePublishableKeySalesChannelsBatch,
//     } from "medusa-react"
//
//     type Props = {
//     publishableApiKeyId: string
//     }
//
//     const PublishableApiKey = ({
//     publishableApiKeyId
//     }: Props) => {
//     const deleteSalesChannels =
//     useAdminRemovePublishableKeySalesChannelsBatch(
//     publishableApiKeyId
//     )
//     // ...
//
//     const handleDelete = (salesChannelId: string) => {
//     deleteSalesChannels.mutate({
//     sales_channel_ids: [
//     {
//     id: salesChannelId,
//     },
//     ],
//     }, {
//     onSuccess: ({ publishable_api_key }) => {
//     console.log(publishable_api_key.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PublishableApiKey
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/publishable-api-keys/{id}/batch' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "sales_channel_ids": [
//     {
//     "id": "{sales_channel_id}"
//     }
//     ]
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Publishable Api Keys
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPublishableApiKeysRes"
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
func (m *PublishableApiKey) DeleteChannelsBatch(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.PublishableApiKeySalesChannelsBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	var ids uuid.UUIDs
	for _, f := range model.SalesChannelIds {
		ids = append(ids, f.Id)
	}

	if err := m.r.PublishableApiKeyService().SetContext(context.Context()).RemoveSalesChannels(id, ids); err != nil {
		return err
	}

	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/publishable-api-keys/{id}/sales-channels
// operationId: "GetPublishableApiKeySalesChannels"
// summary: "List Sales Channels"
// description: "List the sales channels associated with a publishable API key. The sales channels can be filtered by fields such as `q`."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the publishable API key.
//   - (query) q {string} query to search sales channels' names and descriptions.
//
// x-codegen:
//
//	method: listSalesChannels
//	queryParams: GetPublishableApiKeySalesChannelsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.publishableApiKeys.listSalesChannels()
//     .then(({ sales_channels }) => {
//     console.log(sales_channels.length)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import {
//     useAdminPublishableApiKeySalesChannels,
//     } from "medusa-react"
//
//     type Props = {
//     publishableApiKeyId: string
//     }
//
//     const SalesChannels = ({
//     publishableApiKeyId
//     }: Props) => {
//     const { sales_channels, isLoading } =
//     useAdminPublishableApiKeySalesChannels(
//     publishableApiKeyId
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {sales_channels && !sales_channels.length && (
//     <span>No Sales Channels</span>
//     )}
//     {sales_channels && sales_channels.length > 0 && (
//     <ul>
//     {sales_channels.map((salesChannel) => (
//     <li key={salesChannel.id}>{salesChannel.name}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default SalesChannels
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/publishable-api-keys/{id}/sales-channels' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Publishable Api Keys
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPublishableApiKeysListSalesChannelsRes"
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
func (m *PublishableApiKey) ListChannels(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).ListSalesChannels(id, &config.Q)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"sales_channels": result,
	})
}

// @oas:path [post] /admin/publishable-api-keys/{id}/revoke
// operationId: "PostPublishableApiKeysPublishableApiKeyRevoke"
// summary: "Revoke a Publishable API Key"
// description: "Revoke a Publishable API Key. Revoking the publishable API Key can't be undone, and the key can't be used in future requests."
// parameters:
//   - (path) id=* {string} The ID of the Publishable API Key.
//
// x-authenticated: true
// x-codegen:
//
//	method: revoke
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.publishableApiKeys.revoke(publishableApiKeyId)
//     .then(({ publishable_api_key }) => {
//     console.log(publishable_api_key.id)
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminRevokePublishableApiKey } from "medusa-react"
//
//     type Props = {
//     publishableApiKeyId: string
//     }
//
//     const PublishableApiKey = ({
//     publishableApiKeyId
//     }: Props) => {
//     const revokeKey = useAdminRevokePublishableApiKey(
//     publishableApiKeyId
//     )
//     // ...
//
//     const handleRevoke = () => {
//     revokeKey.mutate(void 0, {
//     onSuccess: ({ publishable_api_key }) => {
//     console.log(publishable_api_key.revoked_at)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default PublishableApiKey
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/publishable-api-keys/{id}/revoke' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Publishable Api Keys
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminPublishableApiKeysRes"
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
func (m *PublishableApiKey) Revoke(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	user := api.GetUser(context)

	if err := m.r.PublishableApiKeyService().SetContext(context.Context()).Revoke(id, user); err != nil {
		return err
	}

	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})

}
