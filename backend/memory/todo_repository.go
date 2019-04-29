package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/hironow/team-lgtm/backend/todo"
	"github.com/hironow/team-lgtm/backend/user"
)

type todoRepository struct {
	mtx   sync.RWMutex
	todos map[string]*todo.Todo
}

func NewTodoRepository() todo.Repository {
	return &todoRepository{
		todos: make(map[string]*todo.Todo),
	}
}

func (repo *todoRepository) Get(ctx context.Context, id string, user *user.User) (*todo.Todo, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	for _, val := range repo.todos {
		if val.ID == id {
			return val, nil
		}
	}
	return nil, fmt.Errorf("entity not found")
}

func (repo *todoRepository) Put(ctx context.Context, t *todo.Todo, user *user.User) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	repo.todos[t.ID] = t
	return nil
}
