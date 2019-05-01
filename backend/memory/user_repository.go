package memory

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/hironow/team-lgtm/backend/user"
)

type userRepository struct {
	mtx   sync.RWMutex
	users map[string]*user.User
}

func NewUserRepository() user.Repository {
	return &userRepository{
		users: make(map[string]*user.User),
	}
}

func (repo *userRepository) Get(ctx context.Context, id string) (*user.User, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	for _, val := range repo.users {
		if val.ID == id {
			return val, nil
		}
	}
	return nil, fmt.Errorf("entity not found")
}

func (repo *userRepository) Put(ctx context.Context, u *user.User) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	repo.users[u.ID] = u
	return nil
}

func (repo *userRepository) List(ctx context.Context, cursor string, limit int) ([]*user.User, string, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	var users []*user.User
	var nextCursor string
	for _, val := range repo.users {
		if len(users) > limit+1 {
			// 次のIDをcursorとして扱う
			nextCursor = val.ID
			break
		}
		if cursor != "" {
			if val.ID == cursor {
				// cursor以降を取得対象にする
				cursor = ""
				users = append(users, val)
			}
			log.Printf("skip %s", val.ID)
		} else {
			users = append(users, val)
		}
	}

	return users, nextCursor, nil
}
