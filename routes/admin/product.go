package admin

type Product struct {
	r Registry
}

func NewProduct(r Registry) *Product {
	m := Product{r: r}
	return &m
}
