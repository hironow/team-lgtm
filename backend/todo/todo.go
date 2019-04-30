package todo

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hironow/team-lgtm/backend/user"
)

type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

func NewTodo(u *user.User) *Todo {
	return &Todo{
		ID:     generateID(u),
		UserID: u.ID,
	}
}

func generateID(u *user.User) string {
	return fmt.Sprintf("%s_%s", u.ID, uuid.New().String())
}
