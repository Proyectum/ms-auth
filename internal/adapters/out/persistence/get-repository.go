package persistence

import (
	"errors"
	"github.com/proyectum/ms-auth/internal/domain/entities"
	domainErrors "github.com/proyectum/ms-auth/internal/domain/errors"
	"github.com/proyectum/ms-auth/internal/domain/repository"
	"gorm.io/gorm"
)

type getUserRepositoryImpl struct {
	ds *gorm.DB
}

func (g *getUserRepositoryImpl) GetByUsernameOrEmail(username, email string) (*entities.User, error) {
	var entity UserEntity
	result := g.ds.Where("email = ? or username = ?", email, username).First(&entity)

	return g.handleResult(&entity, result)
}

func (g *getUserRepositoryImpl) GetByEmail(email string) (*entities.User, error) {
	var entity UserEntity
	result := g.ds.Where("email = ?", email).First(&entity)

	return g.handleResult(&entity, result)
}

func (g *getUserRepositoryImpl) GetByUsername(username string) (*entities.User, error) {
	var entity UserEntity
	result := g.ds.Where("username = ?", username).First(&entity)

	return g.handleResult(&entity, result)
}

func (g *getUserRepositoryImpl) handleResult(entity *UserEntity, result *gorm.DB) (*entities.User, error) {
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domainErrors.NewUserNotFoundError("user not found")
		}
		return nil, result.Error
	}

	user := entities.User{
		ID:       entity.ID,
		Email:    entity.Email,
		Username: entity.Username,
		Password: entity.Password,
	}

	return &user, nil
}

func NewGetUserRepository() repository.GetUserRepository {
	return &getUserRepositoryImpl{
		ds: getDatasource(),
	}
}
