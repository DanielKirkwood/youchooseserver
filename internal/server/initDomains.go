package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	authHandler "github.com/DanielKirkwood/youchooseserver/internal/domain/auth/handler"
	authRepository "github.com/DanielKirkwood/youchooseserver/internal/domain/auth/repository"
	authUseCase "github.com/DanielKirkwood/youchooseserver/internal/domain/auth/usecase"
	"github.com/DanielKirkwood/youchooseserver/internal/middleware"
	"github.com/DanielKirkwood/youchooseserver/internal/util/respond"
)

type Domain struct {
	Auth *authHandler.Handler
}

func (s *Server) InitDomains() {
	s.initVersion()
	s.initAuth()
}

func (s *Server) initVersion() {
	s.router.Route("/version", func(router chi.Router) {
		router.Use(middleware.Json)

		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			respond.Json(w, http.StatusOK, map[string]string{"version": s.Version})
		})
	})
}

func (s *Server) initAuth() {
	newAuthRepo := authRepository.New(s.ent, s.email, s.cfg.Api.TokenSecret)
	newAuthUseCase := authUseCase.New(newAuthRepo)
	s.Domain.Auth = authHandler.RegisterHTTPEndPoints(s.router, newAuthUseCase)
}
