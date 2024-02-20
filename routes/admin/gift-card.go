package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type GiftCard struct {
	r    Registry
	name string
}

func NewGiftCard(r Registry) *GiftCard {
	m := GiftCard{r: r, name: "gift_card"}
	return &m
}

func (m *GiftCard) SetRoutes(router fiber.Router) {
	route := router.Group("/gift-cards")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
}

// @oas:path [get] /admin/gift-cards/{id}
// operationId: "GetGiftCardsGiftCard"
// summary: "Get a Gift Card"
// description: "Retrieve a Gift Card's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Gift Card.
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
//     medusa.admin.giftCards.retrieve(giftCardId)
//     .then(({ gift_card }) => {
//     console.log(gift_card.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminGiftCard } from "medusa-react"
//
//     type Props = {
//     giftCardId: string
//     }
//
//     const CustomGiftCard = ({ giftCardId }: Props) => {
//     const { gift_card, isLoading } = useAdminGiftCard(giftCardId)
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {gift_card && <span>{gift_card.code}</span>}
//     </div>
//     )
//     }
//
//     export default CustomGiftCard
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/gift-cards/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Gift Cards
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminGiftCardsRes"
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
func (m *GiftCard) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.GiftCardService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [get] /admin/gift-cards
// operationId: "GetGiftCards"
// summary: "List Gift Cards"
// description: "Retrieve a list of Gift Cards. The gift cards can be filtered by fields such as `q`. The gift cards can also paginated."
// x-authenticated: true
// parameters:
//   - (query) offset=0 {number} The number of gift cards to skip when retrieving the gift cards.
//   - (query) limit=50 {number} Limit the number of gift cards returned.
//   - (query) q {string} a term to search gift cards' code or display ID
//
// x-codegen:
//
//	method: list
//	queryParams: AdminGetGiftCardsParams
//
// x-codeSamples:
//
//   - lang: JavaScript
//     label: JS Client
//     source: |
//     import Medusa from "@medusajs/medusa-js"
//     const medusa = new Medusa({ baseUrl: MEDUSA_BACKEND_URL, maxRetries: 3 })
//     // must be previously logged in or use api token
//     medusa.admin.giftCards.list()
//     .then(({ gift_cards, limit, offset, count }) => {
//     console.log(gift_cards.length);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { GiftCard } from "@medusajs/medusa"
//     import { useAdminGiftCards } from "medusa-react"
//
//     const CustomGiftCards = () => {
//     const { gift_cards, isLoading } = useAdminGiftCards()
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {gift_cards && !gift_cards.length && (
//     <span>No custom gift cards...</span>
//     )}
//     {gift_cards && gift_cards.length > 0 && (
//     <ul>
//     {gift_cards.map((giftCard: GiftCard) => (
//     <li key={giftCard.id}>{giftCard.code}</li>
//     ))}
//     </ul>
//     )}
//     </div>
//     )
//     }
//
//     export default CustomGiftCards
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '"{backend_url}"/admin/gift-cards' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Gift Cards
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminGiftCardsListRes"
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
func (m *GiftCard) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableGiftCard](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.GiftCardService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"gift_cards": result,
		"count":      count,
		"offset":     config.Skip,
		"limit":      config.Take,
	})
}

// @oas:path [post] /admin/gift-cards
// operationId: "PostGiftCards"
// summary: "Create a Gift Card"
// description: "Create a Gift Card that can redeemed by its unique code. The Gift Card is only valid within 1 region."
// x-authenticated: true
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostGiftCardsReq"
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
//     medusa.admin.giftCards.create({
//     region_id
//     })
//     .then(({ gift_card }) => {
//     console.log(gift_card.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminCreateGiftCard } from "medusa-react"
//
//     const CreateCustomGiftCards = () => {
//     const createGiftCard = useAdminCreateGiftCard()
//     // ...
//
//     const handleCreate = (
//     regionId: string,
//     value: number
//     ) => {
//     createGiftCard.mutate({
//     region_id: regionId,
//     value,
//     }, {
//     onSuccess: ({ gift_card }) => {
//     console.log(gift_card.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CreateCustomGiftCards
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/gift-cards' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "region_id": "{region_id}"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Gift Cards
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminGiftCardsRes"
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
func (m *GiftCard) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateGiftCardInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.GiftCardService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [post] /admin/gift-cards/{id}
// operationId: "PostGiftCardsGiftCard"
// summary: "Update a Gift Card"
// description: "Update a Gift Card's details."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Gift Card.
//
// requestBody:
//
//	content:
//	  application/json:
//	    schema:
//	      $ref: "#/components/schemas/AdminPostGiftCardsGiftCardReq"
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
//     medusa.admin.giftCards.update(giftCardId, {
//     region_id
//     })
//     .then(({ gift_card }) => {
//     console.log(gift_card.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminUpdateGiftCard } from "medusa-react"
//
//     type Props = {
//     customGiftCardId: string
//     }
//
//     const CustomGiftCard = ({ customGiftCardId }: Props) => {
//     const updateGiftCard = useAdminUpdateGiftCard(
//     customGiftCardId
//     )
//     // ...
//
//     const handleUpdate = (regionId: string) => {
//     updateGiftCard.mutate({
//     region_id: regionId,
//     }, {
//     onSuccess: ({ gift_card }) => {
//     console.log(gift_card.id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CustomGiftCard
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X POST '"{backend_url}"/admin/gift-cards/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"' \
//     -H 'Content-Type: application/json' \
//     --data-raw '{
//     "region_id": "{region_id}"
//     }'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Gift Cards
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminGiftCardsRes"
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
func (m *GiftCard) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateGiftCardInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.GiftCardService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		(m.name): result,
	})
}

// @oas:path [delete] /admin/gift-cards/{id}
// operationId: "DeleteGiftCardsGiftCard"
// summary: "Delete a Gift Card"
// description: "Delete a Gift Card. Once deleted, it can't be used by customers."
// x-authenticated: true
// parameters:
//   - (path) id=* {string} The ID of the Gift Card to delete.
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
//     medusa.admin.giftCards.delete(giftCardId)
//     .then(({ id, object, deleted }) => {
//     console.log(id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useAdminDeleteGiftCard } from "medusa-react"
//
//     type Props = {
//     customGiftCardId: string
//     }
//
//     const CustomGiftCard = ({ customGiftCardId }: Props) => {
//     const deleteGiftCard = useAdminDeleteGiftCard(
//     customGiftCardId
//     )
//     // ...
//
//     const handleDelete = () => {
//     deleteGiftCard.mutate(void 0, {
//     onSuccess: ({ id, object, deleted}) => {
//     console.log(id)
//     }
//     })
//     }
//
//     // ...
//     }
//
//     export default CustomGiftCard
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl -X DELETE '"{backend_url}"/admin/gift-cards/{id}' \
//     -H 'x-medusa-access-token: "{api_token}"'
//
// security:
//   - api_token: []
//   - cookie_auth: []
//   - jwt_token: []
//
// tags:
//   - Gift Cards
//
// responses:
//
//	200:
//	  description: "OK"
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/AdminGiftCardsDeleteRes"
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
func (m *GiftCard) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.GiftCardService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "gift-card",
		"deleted": true,
	})
}
