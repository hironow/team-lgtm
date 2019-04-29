package user

import "github.com/google/uuid"

type User struct {
	ID   string
	Name string
}

func NewUser() *User {
	return &User{
		ID: generateID(),
	}
}

func generateID() string {
	return uuid.New().String()
}
