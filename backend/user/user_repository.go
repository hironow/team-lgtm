package user

import "context"

type Repository interface {
	Get(ctx context.Context, id string) (*User, error)
	Put(ctx context.Context, u *User) error
	List(ctx context.Context, cursor string, limit int) ([]*User, string, error)
}
