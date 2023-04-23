package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/DanielKirkwood/youchooseserver/internal/domain/friendship"
	"github.com/DanielKirkwood/youchooseserver/internal/domain/friendship/usecase"
	"github.com/DanielKirkwood/youchooseserver/internal/util/respond"
)

type Handler struct {
	useCase usecase.Friendship
}

func NewHandler(useCase usecase.Friendship) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req friendship.CreateRequest
		err error
	)
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	friendshipSchema, err := h.useCase.Create(r.Context(), &req)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, fmt.Errorf("error occurred creating friendship: %w", err))
		return
	}

	respond.Json(w, http.StatusCreated, friendshipSchema)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var (
		req friendship.UpdateRequest
		err error
	)
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	if req.Status == "requested" {
		respond.Error(w, http.StatusBadRequest, errors.New("cannot change status to requested"))
		return
	}

	updated, err := h.useCase.Update(r.Context(), &req)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, fmt.Errorf("error occurred updating friendship: %w", err))
		return
	}

	respond.Json(w, http.StatusOK, updated)
}
