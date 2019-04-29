package clouddatastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/hironow/team-lgtm/backend/user"
)

type userRepository struct {
	dsClient *datastore.Client
}

func NewUserRepository(ctx context.Context) (user.Repository, error) {
	client, err := datastore.NewClient(ctx, "project-id")
	if err != nil {
		return nil, err
	}
	return &userRepository{dsClient: client}, nil
}

func (repo *userRepository) key(id string) *datastore.Key {
	return datastore.NameKey("User", id, nil)
}

func (repo *userRepository) Get(ctx context.Context, id string) (*user.User, error) {
	dst := &user.User{}
	err := repo.dsClient.Get(ctx, repo.key(id), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (repo *userRepository) Put(ctx context.Context, u *user.User) error {
	_, err := repo.dsClient.Put(ctx, repo.key(u.ID), u)
	if err != nil {
		return err
	}
	return nil
}
