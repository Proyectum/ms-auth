package usecases

import (
	"github.com/google/uuid"
	"github.com/proyectum/ms-auth/internal/adapters/out/persistence"
	"github.com/proyectum/ms-auth/internal/app/services"
	"github.com/proyectum/ms-auth/internal/domain/entities"
	"github.com/proyectum/ms-auth/internal/domain/errors"
	"github.com/proyectum/ms-auth/internal/domain/repository"
	"github.com/proyectum/ms-auth/internal/domain/usecases"
	"regexp"
)

type signUpUseCaseImpl struct {
	hashUseCase    services.HashPasswordService
	saveRepository repository.SaveUserRepository
	existsUseCase  usecases.ExistsUserUseCase
}

func NewSignUpUseCase() usecases.SignUpUseCase {
	return &signUpUseCaseImpl{
		hashUseCase:    services.NewHashPasswordService(),
		saveRepository: persistence.NewSaveUserRepository(),
		existsUseCase:  NewExistsUseCase(),
	}
}

func (s *signUpUseCaseImpl) SignUp(username string, password string, email string) error {

	if err := s.checkParameters(username, password, email); err != nil {
		return err
	}

	exists, err := s.existsUseCase.ExistsByUser(username, email)

	if err != nil {
		return err
	}

	if exists {
		return errors.NewUserAlreadyExistsError("user already exists")
	}
	passwordHashed, err := s.hashUseCase.Hash(password)

	if err != nil {
		return err
	}

	user := entities.User{
		ID:       uuid.New(),
		Username: username,
		Password: passwordHashed,
		Email:    email,
	}

	if err = s.saveRepository.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *signUpUseCaseImpl) checkParameters(username string, password string, email string) error {
	if len(username) == 0 {
		return errors.NewInvalidUsernameError("username can't be empty")
	}

	if len(username) < 3 || len(username) > 8 {
		return errors.NewInvalidUsernameError("username should be min 3 and max 8")
	}

	if len(password) == 0 {
		return errors.NewInvalidPasswordError("password can't be empty")
	}

	if len(password) < 6 || len(password) > 14 {
		return errors.NewInvalidPasswordError("password should be min 6 and max 14")
	}

	if len(email) == 0 {
		return errors.NewInvalidEmailError("email can't be empty")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return errors.NewInvalidEmailError("invalid email format")
	}

	return nil
}
