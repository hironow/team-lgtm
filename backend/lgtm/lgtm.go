package lgtm

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hironow/team-lgtm/backend/user"
)

type LGTM struct {
	ID          string
	Description string
	UserID      string
}

func NewLGTM(u *user.User) *LGTM {
	return &LGTM{
		ID:     generateID(u),
		UserID: u.ID,
	}
}

func generateID(u *user.User) string {
	return fmt.Sprintf("%s_%s", u.ID, uuid.New().String())
}
