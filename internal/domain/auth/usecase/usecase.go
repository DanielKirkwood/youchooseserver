package usecase

import (
	"context"

	"github.com/DanielKirkwood/youchooseserver/internal/domain/auth"
	"github.com/DanielKirkwood/youchooseserver/internal/domain/auth/repository"
)

type AuthUseCase struct {
	repo repository.Auth
}

type Auth interface {
	Create(ctx context.Context, req *auth.LoginRequest) error
	Verify(ctx context.Context, req *auth.VerifyRequest) (string, error)
}

func New(repo repository.Auth) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

func (u *AuthUseCase) Create(ctx context.Context, req *auth.LoginRequest) error {
	return u.repo.Create(ctx, req)
}

func (u *AuthUseCase) Verify(ctx context.Context, req *auth.VerifyRequest) (string, error) {
	return u.repo.Verify(ctx, req)
}
