package repository

import "github.com/proyectum/ms-auth/internal/domain/entities"

type GetUserRepository interface {
	GetByUsername(username string) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	GetByUsernameOrEmail(username, email string) (*entities.User, error)
}

type SaveUserRepository interface {
	Save(user entities.User) error
}
