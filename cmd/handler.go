package cmd

import (
	"context"

	"github.com/driver005/gateway/registry"
)

type Handler struct {
}

func (h *Handler) GenerateRegistry(ctx context.Context) *registry.Base {
	return registry.New(ctx)
}

func NewHandler() *Handler {
	return &Handler{}
}
