package repository

import (
	"context"

	"github.com/jeevandhakal/todos/graph/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, input model.NewUser) (string, error)
	Login(ctx context.Context, input model.Login) (string, error)
}

type userrepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userrepository{
		db: db,
	}
}

func (ur userrepository) Create(ctx context.Context, input model.NewUser) (string, error) {
	return "", nil
}

func (ur userrepository) Login(ctx context.Context, input model.Login) (string, error) {
	return "", nil
}
