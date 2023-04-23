package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/DanielKirkwood/youchooseserver/ent"
	"github.com/DanielKirkwood/youchooseserver/ent/friendship"

	domainFriendship "github.com/DanielKirkwood/youchooseserver/internal/domain/friendship"
)

type repository struct {
	ent *ent.Client
}

type Friendship interface {
	Create(ctx context.Context, req *domainFriendship.CreateRequest) (*domainFriendship.Schema, error)
	Update(ctx context.Context, req *domainFriendship.UpdateRequest) (*domainFriendship.Schema, error)
}

func New(ent *ent.Client) *repository {
	return &repository{
		ent: ent,
	}
}

func (r *repository) Create(ctx context.Context, req *domainFriendship.CreateRequest) (*domainFriendship.Schema, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	create, err := r.ent.Friendship.Create().
		SetUserID(req.UserID).
		SetFriendID(req.FriendID).
		SetStatus(friendship.Status("requested")).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("friendship.repository.Create: %w", err)
	}

	// Both created_at and updated_at are created database-side instead of ent.
	// So ent does not return both.
	created, err := r.ent.Friendship.Get(ctx, create.ID)
	if err != nil {
		return nil, fmt.Errorf("friendship not found: %w", err)
	}

	res := &domainFriendship.Schema{
		ID:        created.ID,
		UserID:    created.UserID,
		FriendID:  created.FriendID,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.CreatedAt,
	}

	return res, nil
}

func (r *repository) Update(ctx context.Context, req *domainFriendship.UpdateRequest) (*domainFriendship.Schema, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	found, err := r.ent.Friendship.Query().
		Where(friendship.UserID(req.UserID)).
		Where(friendship.FriendID(req.FriendID)).
		Only(ctx)
	switch {
	// If the entity does not meet a specific condition,
	// the operation will return an "ent.NotFoundError".
	case ent.IsNotFound(err):
		return nil, fmt.Errorf("friendship item was not found")
	// Any other error.
	case err != nil:
		return nil, fmt.Errorf("query error: %w", err)
	}

	if req.Status == "decline" {
		err := r.ent.Friendship.DeleteOne(found).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("delete error: %w", err)
		}

		res := &domainFriendship.Schema{
			ID:        found.ID,
			UserID:    found.UserID,
			FriendID:  found.FriendID,
			CreatedAt: found.CreatedAt,
			UpdatedAt: found.CreatedAt,
		}

		return res, nil
	}

	updated, err := r.ent.Friendship.
		UpdateOne(found).
		SetStatus(friendship.Status(req.Status)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}

	res := &domainFriendship.Schema{
		ID:        updated.ID,
		UserID:    updated.UserID,
		FriendID:  updated.FriendID,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.CreatedAt,
	}

	return res, nil
}
