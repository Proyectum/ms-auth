package usecases

import "github.com/proyectum/ms-auth/internal/domain/entities"

type ValidationUseCase interface {
	Validate(token string) entities.ValidationResult
}
