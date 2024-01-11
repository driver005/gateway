package admin

type PriceList struct {
	r Registry
}

func NewPriceList(r Registry) *PriceList {
	m := PriceList{r: r}
	return &m
}
