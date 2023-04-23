package handler

import (
	"github.com/DanielKirkwood/youchooseserver/internal/domain/friendship/usecase"
	"github.com/go-chi/chi/v5"
)

func RegisterHTTPEndPoints(router *chi.Mux, useCase usecase.Friendship) *Handler {
	h := NewHandler(useCase)

	router.Route("/api/v1/friendship", func(r chi.Router) {
		r.Post("/create", h.Create)
		r.Patch("/update", h.Create)
	})
	return h
}
