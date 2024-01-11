package admin

type ShippingOption struct {
	r Registry
}

func NewShippingOption(r Registry) *ShippingOption {
	m := ShippingOption{r: r}
	return &m
}
