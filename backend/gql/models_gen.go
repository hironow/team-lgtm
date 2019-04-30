// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"github.com/hironow/team-lgtm/backend/todo"
)

type NewSignIn struct {
	Name *string `json:"name"`
}

type NewSignUp struct {
	Name string `json:"name"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type TodosReply struct {
	Todos  []todo.Todo `json:"todos"`
	Cursor string      `json:"cursor"`
}
