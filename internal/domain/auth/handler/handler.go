package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/DanielKirkwood/youchooseserver/internal/domain/auth"
	"github.com/DanielKirkwood/youchooseserver/internal/domain/auth/usecase"
	"github.com/DanielKirkwood/youchooseserver/internal/util/query"
	"github.com/DanielKirkwood/youchooseserver/internal/util/respond"
)

type Handler struct {
	useCase usecase.Auth
}

// NewHandler returns a instance of a Handler with
// the given auth useCase
func NewHandler(useCase usecase.Auth) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req auth.LoginRequest
		err error
	)
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	err = h.useCase.Create(r.Context(), &req)
	if err != nil {
		log.Println(err)
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, nil)
}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	otp := query.String(r, "opt")
	if len(otp) == 0 {
		respond.Error(w, http.StatusBadRequest, errors.New("otp is required"))
		return
	}

	if len(otp) != 36 {
		respond.Error(w, http.StatusBadRequest, errors.New("otp is malformed"))
		return
	}

	userid, err := query.Int(r, "userid")
	if userid == 0 || err != nil {
		respond.Error(w, http.StatusBadRequest, errors.New("userid is required"))
		return
	}

	req := &auth.VerifyRequest{
		UserID: userid,
		OTP:    otp,
	}

	tokenString, err := h.useCase.Verify(r.Context(), req)
	if err != nil {
		respond.Error(w, http.StatusUnauthorized, err)
	}

	respond.Json(w, http.StatusOK, map[string]string{"token": tokenString})
}
