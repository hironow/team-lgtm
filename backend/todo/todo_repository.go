package todo

import (
	"context"

	"github.com/hironow/team-lgtm/backend/user"
)

type Repository interface {
	Get(ctx context.Context, id string, u *user.User) (*Todo, error)
	Put(ctx context.Context, t *Todo, u *user.User) error
	List(ctx context.Context, cursor string, limit int, u *user.User) ([]*Todo, string, error)
}
