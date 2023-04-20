package handler

import (
	"github.com/DanielKirkwood/youchooseserver/internal/domain/auth/usecase"
	"github.com/go-chi/chi/v5"
)

// RegisterHTTPEndPoints connects the handler functions
// to a chi route
func RegisterHTTPEndPoints(router *chi.Mux, useCase usecase.Auth) *Handler {
	h := NewHandler(useCase)

	router.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/passwordless/login", h.Create)
		r.Post("/passwordless/verify", h.Verify)
	})

	return h
}
