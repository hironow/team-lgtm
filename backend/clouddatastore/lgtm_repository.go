package clouddatastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/hironow/team-lgtm/backend/lgtm"
	"github.com/hironow/team-lgtm/backend/user"
	"google.golang.org/api/iterator"
)

type lgtmRepository struct {
	dsClient *datastore.Client
}

func NewLGTMRepository(dsClient *datastore.Client) (lgtm.Repository, error) {
	return &lgtmRepository{dsClient: dsClient}, nil
}

func (repo *lgtmRepository) parentKey(u *user.User) *datastore.Key {
	return datastore.NameKey("User", u.ID, nil)
}

func (repo *lgtmRepository) key(id string, u *user.User) *datastore.Key {
	return datastore.NameKey("LGTM", id, repo.parentKey(u))
}

func (repo *lgtmRepository) Get(ctx context.Context, id string, user *user.User) (*lgtm.LGTM, error) {
	log.Printf("clouddatastore lgtmRepository Get")
	dst := &lgtm.LGTM{}
	err := repo.dsClient.Get(ctx, repo.key(id, user), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (repo *lgtmRepository) Put(ctx context.Context, t *lgtm.LGTM, user *user.User) error {
	log.Printf("clouddatastore lgtmRepository Put")
	_, err := repo.dsClient.Put(ctx, repo.key(t.ID, user), t)
	if err != nil {
		return err
	}
	return nil
}

func (repo *lgtmRepository) List(ctx context.Context, cursor string, limit int, u *user.User) ([]*lgtm.LGTM, string, error) {
	log.Printf("clouddatastore lgtmRepository List")
	q := datastore.NewQuery("LGTM").Ancestor(repo.parentKey(u))
	if cursor != "" {
		dsCursor, err := datastore.DecodeCursor(cursor)
		if err != nil {
			return nil, "", err
		}
		q = q.Start(dsCursor)
	}
	q = q.Limit(limit)

	var el []*lgtm.LGTM
	it := repo.dsClient.Run(ctx, q)
	for {
		var e lgtm.LGTM
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
