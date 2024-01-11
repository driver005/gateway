package admin

type Note struct {
	r Registry
}

func NewNote(r Registry) *Note {
	m := Note{r: r}
	return &m
}
