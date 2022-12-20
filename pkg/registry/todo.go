package registry

import (
	"github.com/jeevandhakal/todos/pkg/adapter/controller"
	"github.com/jeevandhakal/todos/pkg/adapter/repository"
	"github.com/jeevandhakal/todos/pkg/usecase"
)

func (r *registry) NewTodoController() controller.TodoController {
	tr := repository.NewTodoRepository(r.db)
	uc := usecase.NewTodoUsecase(tr)
	return controller.NewTodoController(uc)
}
