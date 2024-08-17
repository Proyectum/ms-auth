package usecases

import (
	"github.com/proyectum/ms-auth/internal/domain/entities"
)

type SignUpUseCase interface {
	SignUp(username, password, email string) error
}

type ExistsUserUseCase interface {
	ExistsByUser(username, email string) (bool, error)
}

type SignInUseCase interface {
	SignIn(username, password string) (*entities.SignInUser, error)
}

type GetUserUseCase interface {
	GetByUsername(username string) (entities.User, error)
}

type SaveUserUseCase interface {
	Save(user entities.User) error
}
