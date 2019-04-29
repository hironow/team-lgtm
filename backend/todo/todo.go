package todo

import (
	"github.com/google/uuid"
	"github.com/hironow/team-lgtm/backend/user"
)

type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

func NewTodo(user *user.User) *Todo {
	return &Todo{
		ID:     generateID(),
		UserID: user.ID,
	}
}

func generateID() string {
	return uuid.New().String()
}
