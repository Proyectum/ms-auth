package persistence

import (
	"github.com/proyectum/ms-auth/internal/domain/entities"
	"github.com/proyectum/ms-auth/internal/domain/repository"
	"gorm.io/gorm"
	"time"
)

type saveUserRepositoryImpl struct {
	ds *gorm.DB
}

func (s *saveUserRepositoryImpl) Save(user entities.User) error {
	entity := &UserEntity{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	result := s.ds.Create(entity)

	return result.Error
}

func NewSaveUserRepository() repository.SaveUserRepository {
	return &saveUserRepositoryImpl{
		ds: getDatasource(),
	}
}
