//go:generate go run github.com/99designs/gqlgen
package gql

import (
	"context"

	"github.com/hironow/team-lgtm/backend/memory"
	"github.com/hironow/team-lgtm/backend/todo"
	"github.com/hironow/team-lgtm/backend/user"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	userRepository user.Repository
	todoRepository todo.Repository
}

func NewResolver() ResolverRoot {
	userRepository := memory.NewUserRepository()
	todoRepository := memory.NewTodoRepository()

	// dummy user
	u := &user.User{ID: "user1"}
	ctx := context.Background()
	if err := userRepository.Put(ctx, u); err != nil {
		panic(err)
	}

	return &Resolver{
		userRepository: userRepository,
		todoRepository: todoRepository,
	}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*todo.Todo, error) {
	// t := &todo.Todo{
	// 	Text:   input.Text,
	// 	ID:     fmt.Sprintf("T%d", rand.Int()),
	// 	UserID: input.UserID,
	// }
	// r.todos = append(r.todos, *t)
	// return t, nil

	u, err := r.userRepository.Get(ctx, input.UserID)
	if err != nil {
		return nil, err
	}
	t := todo.NewTodo(u)
	t.Text = input.Text
	if err := r.todoRepository.Put(ctx, t, u); err != nil {
		return nil, err
	}
	return t, nil
}
func (r *mutationResolver) SignUp(ctx context.Context, input NewSignUp) (*user.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) SignIn(ctx context.Context, input NewSignIn) (*user.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]todo.Todo, error) {
	// return r.todos, nil

	return nil, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *todo.Todo) (*user.User, error) {
	// return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil

	u, err := r.userRepository.Get(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}
	return u, nil
}
