package admin

type PaymentCollection struct {
	r Registry
}

func NewPaymentCollection(r Registry) *PaymentCollection {
	m := PaymentCollection{r: r}
	return &m
}
