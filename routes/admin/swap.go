package admin

type Swap struct {
	r Registry
}

func NewSwap(r Registry) *Swap {
	m := Swap{r: r}
	return &m
}
