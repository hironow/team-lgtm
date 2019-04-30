package clouddatastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/hironow/team-lgtm/backend/todo"
	"github.com/hironow/team-lgtm/backend/user"
)

type todoRepository struct {
	dsClient *datastore.Client
}

func NewTodoRepository(dsClient *datastore.Client) (todo.Repository, error) {
	return &todoRepository{dsClient: dsClient}, nil
}

func (repo *todoRepository) key(id string, user *user.User) *datastore.Key {
	userKey := datastore.NameKey("User", user.ID, nil)
	return datastore.NameKey("Todo", id, userKey)
}

func (repo *todoRepository) Get(ctx context.Context, id string, user *user.User) (*todo.Todo, error) {
	dst := &todo.Todo{}
	err := repo.dsClient.Get(ctx, repo.key(id, user), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (repo *todoRepository) Put(ctx context.Context, t *todo.Todo, user *user.User) error {
	_, err := repo.dsClient.Put(ctx, repo.key(t.ID, user), t)
	if err != nil {
		return err
	}
	return nil
}
