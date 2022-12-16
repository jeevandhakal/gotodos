package registry

import "github.com/jeevandhakal/todos/graph/model"

type TodoRepositoryRegistry interface {
	GetTodo(id string) *model.Todo
	GetTodos() []*model.Todo
}

type TodoUsecaseRegistry interface {
	GetTodo(id string) *model.Todo
	GetTodos() []*model.Todo
}
