package user

import "context"

type Repository interface {
	Get(ctx context.Context, id string) (*User, error)
	Put(ctx context.Context, t *User) error
}
