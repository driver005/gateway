package admin

import "github.com/gofiber/fiber/v3"

type Reservation struct {
	r Registry
}

func NewReservation(r Registry) *Reservation {
	m := Reservation{r: r}
	return &m
}

func (m *Reservation) SetRoutes(router fiber.Router) {
	route := router.Group("/reservations")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
}

func (m *Reservation) Get(context fiber.Ctx) error {
	return nil
}

func (m *Reservation) List(context fiber.Ctx) error {
	return nil
}

func (m *Reservation) Create(context fiber.Ctx) error {
	return nil
}

func (m *Reservation) Update(context fiber.Ctx) error {
	return nil
}

func (m *Reservation) Delete(context fiber.Ctx) error {
	return nil
}
