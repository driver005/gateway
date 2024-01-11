package admin

type SalesChannel struct {
	r Registry
}

func NewSalesChannel(r Registry) *SalesChannel {
	m := SalesChannel{r: r}
	return &m
}
