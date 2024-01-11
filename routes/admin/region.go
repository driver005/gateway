package admin

type Region struct {
	r Registry
}

func NewRegion(r Registry) *Region {
	m := Region{r: r}
	return &m
}
