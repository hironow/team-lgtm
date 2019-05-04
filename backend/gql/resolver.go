//go:generate go run github.com/99designs/gqlgen
package gql

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/hironow/team-lgtm/backend/clouddatastore"
	"github.com/hironow/team-lgtm/backend/gql/middleware"
	"github.com/hironow/team-lgtm/backend/lgtm"
	"github.com/hironow/team-lgtm/backend/memory"
	"github.com/hironow/team-lgtm/backend/todo"
	"github.com/hironow/team-lgtm/backend/user"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	userRepository user.Repository
	todoRepository todo.Repository
	lgtmRepository lgtm.Repository
}

func NewResolver(dsClient *datastore.Client) (ResolverRoot, error) {
	var (
		userRepository user.Repository
		todoRepository todo.Repository
		lgtmRepository lgtm.Repository
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
		lgtmRepository, err = clouddatastore.NewLGTMRepository(dsClient)
		if err != nil {
			return nil, err
		}
	} else {
		userRepository = memory.NewUserRepository()
		todoRepository = memory.NewTodoRepository()
		lgtmRepository = memory.NewLGTMRepository()
	}

	// check user
	{
		users, _, err := userRepository.List(context.Background(), "", 10)
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
		lgtmRepository: lgtmRepository,
	}, nil
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SignUp(ctx context.Context, input NewSignUp) (*user.User, error) {
	uID := middleware.ForContext(ctx)
	log.Printf("uID: %s", uID)

	// TODO: uID -> userID

	u := user.NewUser()
	u.Name = input.Name
	if err := r.userRepository.Put(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *mutationResolver) SignIn(ctx context.Context, input NewSignIn) (*user.User, error) {
	uID := middleware.ForContext(ctx)
	log.Printf("uID: %s", uID)

	// TODO: uID -> userID

	userID := "user1"
	u, err := r.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*todo.Todo, error) {
	uID := middleware.ForContext(ctx)
	log.Printf("uID: %s", uID)

	// TODO: uID -> userID

	userID := "user1"
	u, err := r.userRepository.Get(ctx, userID)
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


func (r *mutationResolver) CreateLgtm(ctx context.Context, input NewLgtm) (*lgtm.LGTM, error) {
	uID := middleware.ForContext(ctx)
	log.Printf("uID: %s", uID)

	// TODO: uID -> userID

	userID := "user1"
	u, err := r.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	t := lgtm.NewLGTM(u)
	t.Description = input.Description
	if err := r.lgtmRepository.Put(ctx, t, u); err != nil {
		return nil, err
	}
	return t, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (*user.User, error) {
	uID := middleware.ForContext(ctx)
	log.Printf("uID: %s", uID)

	// TODO: uID -> userID

	userID := "user1"
	u, err := r.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *queryResolver) Todos(ctx context.Context, cursor *string) (*TodoConnection, error) {
	log.Printf("cursor: %s", *cursor)

	uID := middleware.ForContext(ctx)
	log.Printf("uID: %s", uID)

	// TODO: uID -> userID

	userID := "user1"
	u, err := r.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	todos, nextCursor, err := r.todoRepository.List(ctx, *cursor, 3, u)
	if err != nil {
		return nil, err
	}
	log.Printf("next cursor: %s", nextCursor)

	return &TodoConnection{Todos: todos, HasMore: len(todos) != 0, Cursor: nextCursor}, nil
}

func (r *queryResolver) Lgtms(ctx context.Context, cursor *string) (*LGTMConnection, error) {
	log.Printf("cursor: %s", *cursor)

	uID := middleware.ForContext(ctx)
	log.Printf("uID: %s", uID)

	// TODO: uID -> userID

	userID := "user1"
	u, err := r.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	lgtms, nextCursor, err := r.lgtmRepository.List(ctx, *cursor, 3, u)
	if err != nil {
		return nil, err
	}
	log.Printf("next cursor: %s", nextCursor)

	return &LGTMConnection{Lgtms: lgtms, HasMore: len(lgtms) != 0, Cursor: nextCursor}, nil
}