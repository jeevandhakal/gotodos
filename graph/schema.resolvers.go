package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jeevandhakal/todos/auth"
	"github.com/jeevandhakal/todos/graph/generated"
	"github.com/jeevandhakal/todos/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Todo{}, fmt.Errorf("access denied")
	}
	var todo model.Todo
	todo.Title = input.Title
	todo.Description = input.Description
	todo.Completed = input.Completed
	todo.User = user
	todoID := todo.Create()

	grahpqlUser := &model.User{
		ID:       user.ID,
		Username: user.Username,
	}
	return &model.Todo{ID: todoID, Title: todo.Title, Description: todo.Description, Completed: todo.Completed, User: grahpqlUser}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user model.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user model.User
	user.Username = input.Username
	user.Password = input.Password

	correct := user.Authenticate()
	if !correct {
		return "", &model.WrongUsernameOrPasswordError{}
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := auth.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := auth.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	user := auth.ForContext(ctx)
	fmt.Println(user.Username)
	var resultTodos []*model.Todo
	dbTodos := model.GetTodos(user.Username)
	for _, todo := range dbTodos {
		resultTodos = append(resultTodos, &model.Todo{ID: todo.ID, Title: todo.Title, Description: todo.Description, Completed: todo.Completed})
	}
	return resultTodos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
