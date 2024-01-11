package admin

type TaxRate struct {
	r Registry
}

func NewTaxRate(r Registry) *TaxRate {
	m := TaxRate{r: r}
	return &m
}
