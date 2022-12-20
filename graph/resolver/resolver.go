package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/jeevandhakal/todos/graph/generated"
	"github.com/jeevandhakal/todos/pkg/adapter/controller"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	controller.Controller
}

func NewSchema(controller controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{controller},
	})
}
