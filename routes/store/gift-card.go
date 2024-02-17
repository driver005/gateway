package store

import (
	"github.com/driver005/gateway/sql"
	"github.com/gofiber/fiber/v3"
)

type GiftCard struct {
	r Registry
}

func NewGiftCard(r Registry) *GiftCard {
	m := GiftCard{r: r}
	return &m
}

func (m *GiftCard) SetRoutes(router fiber.Router) {
	route := router.Group("/gift-cards")
	route.Get("/:code", m.Get)
}

// @oas:path [get] /store/gift-cards/{code}
// operationId: "GetGiftCardsCode"
// summary: "Get Gift Card by Code"
// description: "Retrieve a Gift Card's details by its associated unique code."
// parameters:
//   - (path) code=* {string} The unique Gift Card code.
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
//     medusa.giftCards.retrieve(code)
//     .then(({ gift_card }) => {
//     console.log(gift_card.id);
//     })
//
//   - lang: tsx
//     label: Medusa React
//     source: |
//     import React from "react"
//     import { useGiftCard } from "medusa-react"
//
//     type Props = {
//     giftCardCode: string
//     }
//
//     const GiftCard = ({ giftCardCode }: Props) => {
//     const { gift_card, isLoading, isError } = useGiftCard(
//     giftCardCode
//     )
//
//     return (
//     <div>
//     {isLoading && <span>Loading...</span>}
//     {gift_card && <span>{gift_card.value}</span>}
//     {isError && <span>Gift Card does not exist</span>}
//     </div>
//     )
//     }
//
//     export default GiftCard
//
//   - lang: Shell
//     label: cURL
//     source: |
//     curl '{backend_url}/store/gift-cards/{code}'
//
// tags:
//   - Gift Cards
//
// responses:
//
//	200:
//	  description: OK
//	  content:
//	    application/json:
//	      schema:
//	        $ref: "#/components/schemas/StoreGiftCardsRes"
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
func (m *GiftCard) Get(context fiber.Ctx) error {
	code := context.Params("code")

	result, err := m.r.GiftCardService().SetContext(context.Context()).RetrieveByCode(code, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
