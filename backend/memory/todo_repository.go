package memory

import (
	"fmt"
	"sync"

	"github.com/hironow/team-lgtm/backend/todo"
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

func (repo *todoRepository) Get(id string) (*todo.Todo, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	for _, val := range repo.todos {
		if val.ID == id {
			return val, nil
		}
	}
	return nil, fmt.Errorf("entity not found")
}

func (repo *todoRepository) Put(t *todo.Todo) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	repo.todos[t.ID] = t
	return nil
}
