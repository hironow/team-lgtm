package memory

import (
	"context"
	"fmt"
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
