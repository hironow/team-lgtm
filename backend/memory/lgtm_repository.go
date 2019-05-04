package memory

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/hironow/team-lgtm/backend/lgtm"
	"github.com/hironow/team-lgtm/backend/user"
)

type lgtmRepository struct {
	mtx   sync.RWMutex
	lgtms map[string]*lgtm.LGTM
}

func NewLGTMRepository() lgtm.Repository {
	return &lgtmRepository{
		lgtms: make(map[string]*lgtm.LGTM),
	}
}

func (repo *lgtmRepository) Get(ctx context.Context, id string, u *user.User) (*lgtm.LGTM, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	for _, val := range repo.lgtms {
		if val.ID == id {
			return val, nil
		}
	}
	return nil, fmt.Errorf("entity not found")
}

func (repo *lgtmRepository) Put(ctx context.Context, t *lgtm.LGTM, u *user.User) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	repo.lgtms[t.ID] = t
	return nil
}

func (repo *lgtmRepository) List(ctx context.Context, cursor string, limit int, u *user.User) ([]*lgtm.LGTM, string, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	var lgtms []*lgtm.LGTM
	var nextCursor string
	for _, val := range repo.lgtms {
		if len(lgtms) > limit+1 {
			// 次のIDをcursorとして扱う
			nextCursor = val.ID
			break
		}
		if cursor != "" {
			if val.ID == cursor {
				// cursor以降を取得対象にする
				cursor = ""
				lgtms = append(lgtms, val)
			}
			log.Printf("skip %s", val.ID)
		} else {
			if val.UserID == u.ID {
				lgtms = append(lgtms, val)
			}
		}
	}

	return lgtms, nextCursor, nil
}
