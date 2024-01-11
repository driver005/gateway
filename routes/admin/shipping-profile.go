package admin

type ShippingProfile struct {
	r Registry
}

func NewShippingProfile(r Registry) *ShippingProfile {
	m := ShippingProfile{r: r}
	return &m
}
