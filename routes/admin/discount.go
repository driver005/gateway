package admin

type Discount struct {
	r Registry
}

func NewDiscount(r Registry) *Discount {
	m := Discount{r: r}
	return &m
}
