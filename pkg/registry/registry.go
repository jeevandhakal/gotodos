package registry

import (
	"github.com/jeevandhakal/todos/pkg/adapter/controller"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

// Registry is an interface of registry
type Registry interface {
	NewController() controller.Controller
}

// New registers entire controller with dependencies
func New(db *gorm.DB) Registry {
	return &registry{
		db: db,
	}
}

// NewController generates controllers
func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		User: r.NewUserController(),
		Todo: r.NewTodoController(),
	}
}
