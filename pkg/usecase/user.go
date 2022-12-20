package usecase

import (
	"context"

	"github.com/jeevandhakal/todos/graph/model"
	"github.com/jeevandhakal/todos/pkg/adapter/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	Create(ctx context.Context, input model.NewUser) (string, error)
	Login(ctx context.Context, input model.Login) (string, error)
	RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error)
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu userUsecase) Create(ctx context.Context, input model.NewUser) (string, error) {
	return uu.userRepository.Create(ctx, input)
}

func (uu userUsecase) Login(ctx context.Context, input model.Login) (string, error) {
	return uu.userRepository.Login(ctx, input)
}

func (uu userUsecase) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	return "", nil
}
