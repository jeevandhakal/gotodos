package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jeevandhakal/todos/graph/generated"
	"github.com/jeevandhakal/todos/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return r.Todo.Create(ctx, input)
}

// ChangeCompletedStatus is the resolver for the changeCompletedStatus field.
func (r *mutationResolver) ChangeCompletedStatus(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error) {
	return r.Todo.Update(ctx, input)
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) ([]*model.Todo, error) {
	_, err := r.Todo.Delete(ctx, id)
	return []*model.Todo{}, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.Todo.List(ctx)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
