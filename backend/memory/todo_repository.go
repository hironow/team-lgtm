package memory

import (
	"context"
	"fmt"
	"log"
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

func (repo *todoRepository) Get(ctx context.Context, id string, u *user.User) (*todo.Todo, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	for _, val := range repo.todos {
		if val.ID == id {
			return val, nil
		}
	}
	return nil, fmt.Errorf("entity not found")
}

func (repo *todoRepository) Put(ctx context.Context, t *todo.Todo, u *user.User) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	repo.todos[t.ID] = t
	return nil
}

func (repo *todoRepository) List(ctx context.Context, cursor string, limit int, u *user.User) ([]*todo.Todo, string, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	var todos []*todo.Todo
	var nextCursor string
	for _, val := range repo.todos {
		if len(todos) > limit+1 {
			// 次のIDをcursorとして扱う
			nextCursor = val.ID
			break
		}
		if cursor != "" {
			if val.ID == cursor {
				// cursor以降を取得対象にする
				cursor = ""
				todos = append(todos, val)
			}
			log.Printf("skip %s", val.ID)
		} else {
			if val.UserID == u.ID {
				todos = append(todos, val)
			}
		}
	}

	return todos, nextCursor, nil
}
