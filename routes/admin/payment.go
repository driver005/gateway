package admin

type Payment struct {
	r Registry
}

func NewPayment(r Registry) *Payment {
	m := Payment{r: r}
	return &m
}
