package middlewares

import (
	"github.com/driver005/gateway/services"
)

type Handler struct {
	r Registry
}

type Registry interface {
	TockenService() *services.TockenService
	AuthService() *services.AuthService
}

func New(r Registry) *Handler {
	return &Handler{
		r: r,
	}
}
