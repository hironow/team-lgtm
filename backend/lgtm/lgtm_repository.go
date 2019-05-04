package lgtm

import (
	"context"

	"github.com/hironow/team-lgtm/backend/user"
)

type Repository interface {
	Get(ctx context.Context, id string, u *user.User) (*LGTM, error)
	Put(ctx context.Context, t *LGTM, u *user.User) error
	List(ctx context.Context, cursor string, limit int, u *user.User) ([]*LGTM, string, error)
}
