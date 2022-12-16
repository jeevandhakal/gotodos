package repository

import (
	"github.com/jeevandhakal/todos/graph/model"
	"github.com/jeevandhakal/todos/registry"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) registry.TodoRepositoryRegistry {
	return &TodoRepository{
		db: db,
	}
}

func (todorepository *TodoRepository) GetTodo(id string) *model.Todo {
	return &model.Todo{}
}

func (todorepository *TodoRepository) GetTodos() []*model.Todo {
	return []*model.Todo{}
}
