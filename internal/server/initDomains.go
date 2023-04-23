package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	authHandler "github.com/DanielKirkwood/youchooseserver/internal/domain/auth/handler"
	authRepository "github.com/DanielKirkwood/youchooseserver/internal/domain/auth/repository"
	authUseCase "github.com/DanielKirkwood/youchooseserver/internal/domain/auth/usecase"
	friendshipHandler "github.com/DanielKirkwood/youchooseserver/internal/domain/friendship/handler"
	friendshipRepository "github.com/DanielKirkwood/youchooseserver/internal/domain/friendship/repository"
	friendshipUseCase "github.com/DanielKirkwood/youchooseserver/internal/domain/friendship/usecase"
	"github.com/DanielKirkwood/youchooseserver/internal/middleware"
	"github.com/DanielKirkwood/youchooseserver/internal/util/respond"
)

type Domain struct {
	Auth *authHandler.Handler
	Friendship *friendshipHandler.Handler
}

func (s *Server) InitDomains() {
	s.initVersion()
	s.initAuth()
	s.initFriendship()
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

func (s *Server) initFriendship() {
	newFriendshipRepo := friendshipRepository.New(s.ent)
	newFriendshipUseCase := friendshipUseCase.New(newFriendshipRepo)
	s.Domain.Friendship = friendshipHandler.RegisterHTTPEndPoints(s.router, newFriendshipUseCase)
}
