package admin

type ProductTag struct {
	r Registry
}

func NewProductTag(r Registry) *ProductTag {
	m := ProductTag{r: r}
	return &m
}
