package admin

type Return struct {
	r Registry
}

func NewReturn(r Registry) *Return {
	m := Return{r: r}
	return &m
}
