package clouddatastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/hironow/team-lgtm/backend/todo"
	"github.com/hironow/team-lgtm/backend/user"
	"google.golang.org/api/iterator"
)

type todoRepository struct {
	dsClient *datastore.Client
}

func NewTodoRepository(dsClient *datastore.Client) (todo.Repository, error) {
	return &todoRepository{dsClient: dsClient}, nil
}

func (repo *todoRepository) parentKey(u *user.User) *datastore.Key {
	return datastore.NameKey("User", u.ID, nil)
}

func (repo *todoRepository) key(id string, u *user.User) *datastore.Key {
	return datastore.NameKey("Todo", id, repo.parentKey(u))
}

func (repo *todoRepository) Get(ctx context.Context, id string, user *user.User) (*todo.Todo, error) {
	log.Printf("clouddatastore todoRepository Get")
	dst := &todo.Todo{}
	err := repo.dsClient.Get(ctx, repo.key(id, user), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (repo *todoRepository) Put(ctx context.Context, t *todo.Todo, user *user.User) error {
	log.Printf("clouddatastore todoRepository Put")
	_, err := repo.dsClient.Put(ctx, repo.key(t.ID, user), t)
	if err != nil {
		return err
	}
	return nil
}

func (repo *todoRepository) List(ctx context.Context, cursor string, limit int, u *user.User) ([]*todo.Todo, string, error) {
	log.Printf("clouddatastore todoRepository List")
	q := datastore.NewQuery("Todo").Ancestor(repo.parentKey(u))
	if cursor != "" {
		dsCursor, err := datastore.DecodeCursor(cursor)
		if err != nil {
			return nil, "", err
		}
		q = q.Start(dsCursor)
	}
	q = q.Limit(limit)

	var el []*todo.Todo
	it := repo.dsClient.Run(ctx, q)
	for {
		var e todo.Todo
		if _, err := it.Next(&e); err == iterator.Done {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		el = append(el, &e)
	}
	nextCursor, err := it.Cursor()
	if err != nil {
		return nil, "", err
	}

	return el, nextCursor.String(), nil
}
