package admin

type Store struct {
	r Registry
}

func NewStore(r Registry) *Store {
	m := Store{r: r}
	return &m
}
