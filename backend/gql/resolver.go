//go:generate go run github.com/99designs/gqlgen
package gql

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/hironow/team-lgtm/backend/clouddatastore"
	"github.com/hironow/team-lgtm/backend/memory"
	"github.com/hironow/team-lgtm/backend/todo"
	"github.com/hironow/team-lgtm/backend/user"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	userRepository user.Repository
	todoRepository todo.Repository
}

func NewResolver(dsClient *datastore.Client) (ResolverRoot, error) {
	var (
		userRepository user.Repository
		todoRepository todo.Repository
	)
	if dsClient != nil {
		var err error
		userRepository, err = clouddatastore.NewUserRepository(dsClient)
		if err != nil {
			return nil, err
		}
		todoRepository, err = clouddatastore.NewTodoRepository(dsClient)
		if err != nil {
			return nil, err
		}
	} else {
		userRepository = memory.NewUserRepository()
		todoRepository = memory.NewTodoRepository()
	}

	// check user
	{
		users, _,  err := userRepository.List(context.Background(), "", 10)
		if err != nil {
			panic(err)
		}
		log.Printf("%+v", users)
	}

	// dummy user
	{
		u := &user.User{ID: "user1", Name: "ユーザ壱"}
		if err := userRepository.Put(context.Background(), u); err != nil {
			panic(err)
		}
	}

	return &Resolver{
		userRepository: userRepository,
		todoRepository: todoRepository,
	}, nil
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

func (r *queryResolver) Todos(ctx context.Context, cursor *string) (*TodosReply, error) {
	// return r.todos, nil

	log.Printf("cursor: %s", *cursor)

	userID:= "user1"
	u, err := r.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	log.Printf("%+v", u)

	todos, nextCursor, err := r.todoRepository.List(ctx, "", 5, u)
	if err != nil {
		return nil, err
	}
	log.Printf("next cursor: %s", nextCursor)

	return &TodosReply{Todos: todos, Cursor: nextCursor}, nil
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
