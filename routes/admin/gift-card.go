package admin

import "github.com/gofiber/fiber/v3"

type GiftCard struct {
	r Registry
}

func NewGiftCard(r Registry) *GiftCard {
	m := GiftCard{r: r}
	return &m
}

func (m *GiftCard) Get(context fiber.Ctx) error {
	return nil
}

func (m *GiftCard) List(context fiber.Ctx) error {
	return nil
}

func (m *GiftCard) Create(context fiber.Ctx) error {
	return nil
}

func (m *GiftCard) Update(context fiber.Ctx) error {
	return nil
}

func (m *GiftCard) Delete(context fiber.Ctx) error {
	return nil
}
