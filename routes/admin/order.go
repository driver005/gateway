package admin

type Order struct {
	r Registry
}

func NewOrder(r Registry) *Order {
	m := Order{r: r}
	return &m
}
