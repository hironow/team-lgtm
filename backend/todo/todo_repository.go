package todo

import (
	"context"

	"github.com/hironow/team-lgtm/backend/user"
)

type Repository interface {
	Get(ctx context.Context, id string, user *user.User) (*Todo, error)
	Put(ctx context.Context, t *Todo, user *user.User) error
}
