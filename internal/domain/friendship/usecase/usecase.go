package usecase

import (
	"context"

	"github.com/DanielKirkwood/youchooseserver/internal/domain/friendship"
	"github.com/DanielKirkwood/youchooseserver/internal/domain/friendship/repository"
)


type FriendshipUseCase struct {
	repo repository.Friendship
}

type Friendship interface {
	Create(ctx context.Context, req *friendship.CreateRequest) (*friendship.Schema, error)
	Update(ctx context.Context, req *friendship.UpdateRequest) (*friendship.Schema, error)
}

func New(repo repository.Friendship) *FriendshipUseCase {
	return &FriendshipUseCase{
		repo: repo,
	}
}

func (u *FriendshipUseCase) Create(ctx context.Context, req *friendship.CreateRequest) (*friendship.Schema, error) {
	return u.repo.Create(ctx, req)
}

func (u *FriendshipUseCase) Update(ctx context.Context, req *friendship.UpdateRequest) (*friendship.Schema, error) {
	return u.repo.Update(ctx, req)
}
