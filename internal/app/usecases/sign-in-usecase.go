package usecases

import (
	"github.com/proyectum/ms-auth/internal/adapters/out/persistence"
	"github.com/proyectum/ms-auth/internal/app/services"
	"github.com/proyectum/ms-auth/internal/domain/entities"
	"github.com/proyectum/ms-auth/internal/domain/errors"
	"github.com/proyectum/ms-auth/internal/domain/repository"
	"github.com/proyectum/ms-auth/internal/domain/usecases"
	"log"
)

type signInUseCaseImpl struct {
	hashService       services.HashPasswordService
	getUserRepository repository.GetUserRepository
	existsUseCase     usecases.ExistsUserUseCase
	jwtService        services.JWTService
}

func (s *signInUseCaseImpl) SignIn(username, password string) (*entities.SignInUser, error) {
	err := s.checkParameters(username, password)

	if err != nil {
		return nil, err
	}

	err = s.checkIfExists(username)

	if err != nil {
		return nil, err
	}

	user, err := s.getUserRepository.GetByUsernameOrEmail(username, username)

	if err != nil {
		log.Printf("error %v\n", err)
		return nil, errors.NewInvalidCredentialsError("invalid credentials")
	}

	err = s.checkCredentials(user, password)

	if err != nil {
		return nil, err
	}

	return s.getSignInUser(user)
}

func (s *signInUseCaseImpl) getSignInUser(user *entities.User) (*entities.SignInUser, error) {
	token, err := s.jwtService.GetToken(user)

	if err != nil {
		log.Printf("error %v\n", err)
		return nil, errors.NewInvalidCredentialsError("invalid credentials")
	}

	return &entities.SignInUser{
		Token: token,
	}, nil
}

func (s *signInUseCaseImpl) checkCredentials(user *entities.User, password string) error {

	valid := s.hashService.Check(password, user.Password)

	if !valid {
		return errors.NewInvalidCredentialsError("invalid credentials")
	}

	return nil
}

func (s *signInUseCaseImpl) checkIfExists(username string) error {
	exists, err := s.existsUseCase.ExistsByUser(username, username)

	if err != nil {
		log.Printf("error %v\n", err)
		return errors.NewInvalidCredentialsError("invalid credentials")
	}

	if !exists {
		return errors.NewInvalidEmailError("invalid credentials")
	}

	return nil
}

func NewSignInUseCase() usecases.SignInUseCase {
	return &signInUseCaseImpl{
		hashService:       services.NewHashPasswordService(),
		getUserRepository: persistence.NewGetUserRepository(),
		existsUseCase:     NewExistsUseCase(),
		jwtService:        services.NewJWTService(),
	}
}

func (s *signInUseCaseImpl) checkParameters(username string, password string) error {
	if len(username) == 0 {
		return errors.NewInvalidUsernameError("username can't be empty")
	}

	if len(password) == 0 {
		return errors.NewInvalidPasswordError("password can't be empty")
	}

	return nil
}
