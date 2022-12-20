package registry

import (
	"github.com/jeevandhakal/todos/pkg/adapter/controller"
	"github.com/jeevandhakal/todos/pkg/adapter/repository"
	"github.com/jeevandhakal/todos/pkg/usecase"
)

func (r *registry) NewUserController() controller.UserController {
	ur := repository.NewUserRepository(r.db)
	uc := usecase.NewUserUsecase(ur)
	return controller.NewUserController(uc)
}
