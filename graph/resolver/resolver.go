package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/jeevandhakal/todos/graph/generated"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ db *gorm.DB }

func NewSchema(db *gorm.DB) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{db},
	})
}
