package admin

type OrderEdit struct {
	r Registry
}

func NewOrderEdit(r Registry) *OrderEdit {
	m := OrderEdit{r: r}
	return &m
}
