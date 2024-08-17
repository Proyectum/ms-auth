package usecases

import (
	"errors"
	"github.com/proyectum/ms-auth/internal/adapters/out/persistence"
	domainErrors "github.com/proyectum/ms-auth/internal/domain/errors"
	"github.com/proyectum/ms-auth/internal/domain/repository"
	"github.com/proyectum/ms-auth/internal/domain/usecases"
)

type existsUseCaseImpl struct {
	getUserRepository repository.GetUserRepository
}

func (e *existsUseCaseImpl) ExistsByUser(username, email string) (bool, error) {
	user, err := e.getUserRepository.GetByUsernameOrEmail(username, email)
	if err != nil {
		if errors.Is(err, &domainErrors.UserNotFoundError{}) {
			return false, nil
		}
		return false, err
	}

	return user != nil, nil
}

func NewExistsUseCase() usecases.ExistsUserUseCase {
	return &existsUseCaseImpl{
		getUserRepository: persistence.NewGetUserRepository(),
	}
}
