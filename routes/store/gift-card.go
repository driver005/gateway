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

func (m *GiftCard) Get(context fiber.Ctx) error {
	code := context.Params("code")

	result, err := m.r.GiftCardService().SetContext(context.Context()).RetrieveByCode(code, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
