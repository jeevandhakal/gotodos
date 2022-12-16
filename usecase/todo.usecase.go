package usecase

import (
	"github.com/jeevandhakal/todos/graph/model"
	"github.com/jeevandhakal/todos/registry"
)

type TodoUsecase struct {
	todoRegistry registry.TodoUsecaseRegistry
}

func NewTodoUsecase(todoregistry registry.TodoUsecaseRegistry) registry.TodoUsecaseRegistry {
	return &TodoUsecase{
		todoRegistry: todoregistry,
	}
}

func (todousecase TodoUsecase) GetTodo(id string) *model.Todo {
	return todousecase.todoRegistry.GetTodo(id)
}

func (todousecase TodoUsecase) GetTodos() []*model.Todo {
	return todousecase.todoRegistry.GetTodos()
}
