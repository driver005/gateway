package cmd

type Handler struct {
	r Registry
}

func NewHandler(r Registry) *Handler {
	return &Handler{
		r,
	}
}
