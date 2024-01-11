package admin

type ProductType struct {
	r Registry
}

func NewProductType(r Registry) *ProductType {
	m := ProductType{r: r}
	return &m
}
