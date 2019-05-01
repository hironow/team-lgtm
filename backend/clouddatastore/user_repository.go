package clouddatastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/hironow/team-lgtm/backend/user"
	"google.golang.org/api/iterator"
)

type userRepository struct {
	dsClient *datastore.Client
}

func NewUserRepository(dsClient *datastore.Client) (user.Repository, error) {
	return &userRepository{dsClient: dsClient}, nil
}

func (repo *userRepository) key(id string) *datastore.Key {
	return datastore.NameKey("User", id, nil)
}

func (repo *userRepository) Get(ctx context.Context, id string) (*user.User, error) {
	log.Printf("clouddatastore userRepository Get")
	dst := &user.User{}
	err := repo.dsClient.Get(ctx, repo.key(id), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (repo *userRepository) Put(ctx context.Context, u *user.User) error {
	log.Printf("clouddatastore userRepository Put")
	_, err := repo.dsClient.Put(ctx, repo.key(u.ID), u)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) List(ctx context.Context, cursor string, limit int) ([]*user.User, string, error) {
	log.Printf("clouddatastore userRepository List")
	q := datastore.NewQuery("User")
	if cursor != "" {
		dsCursor, err := datastore.DecodeCursor(cursor)
		if err != nil {
			return nil, "", err
		}
		q = q.Start(dsCursor)
	}
	q = q.Limit(limit)

	var el []*user.User
	it := repo.dsClient.Run(ctx, q)
	for {
		var e user.User
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
