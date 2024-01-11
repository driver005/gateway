package repository

type Handler struct {
	r Registry
}

type Registry interface {
}

func NewHandler(r Registry) *Handler {
	return &Handler{r: r}
}
