package usecases

import (
	"github.com/proyectum/ms-auth/internal/domain/entities"
	"github.com/proyectum/ms-auth/internal/domain/usecases"
)

type getUserUseCaseImpl struct {
}

func (g *getUserUseCaseImpl) GetByUsername(username string) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewGetUserUseCase() usecases.GetUserUseCase {
	return &getUserUseCaseImpl{}
}
