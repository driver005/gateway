package admin

type ProductCategory struct {
	r Registry
}

func NewProductCategory(r Registry) *ProductCategory {
	m := ProductCategory{r: r}
	return &m
}
