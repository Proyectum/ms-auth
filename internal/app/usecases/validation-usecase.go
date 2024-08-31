package usecases

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/proyectum/ms-auth/internal/boot"
	"github.com/proyectum/ms-auth/internal/domain/entities"
	"github.com/proyectum/ms-auth/internal/domain/usecases"
	"log"
	"time"
)

type validationUseCaseImpl struct{}

func NewValidationUseCase() usecases.ValidationUseCase {
	return &validationUseCaseImpl{}
}

func (uc *validationUseCaseImpl) Validate(tokenStr string) entities.ValidationResult {
	if len(tokenStr) == 0 {
		log.Println("token is empty")
		return entities.ValidationResult{
			Status: entities.AuthUnauthorized,
		}
	}

	jwtSecret := boot.CONFIG.Security.JWT.Secret

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown sign: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		log.Println("validating error", err)
		return entities.ValidationResult{
			Status: entities.AuthError,
		}
	}

	if !token.Valid {
		log.Println("invalid token")
		return entities.ValidationResult{
			Status: entities.AuthUnauthorized,
		}
	}

	expiration, err := token.Claims.GetExpirationTime()

	if err != nil {
		log.Println("error getting expiration time", err)
		return entities.ValidationResult{
			Status: entities.AuthError,
		}
	}

	if expiration.Before(time.Now()) {
		return entities.ValidationResult{
			Status: entities.AuthUnauthorized,
		}
	}

	claims, parsedClaims := token.Claims.(jwt.MapClaims)

	if !parsedClaims {
		log.Println("error parsed claims")
		return entities.ValidationResult{
			Status: entities.AuthError,
		}
	}

	username := claims["username"].(string)
	email := claims["email"].(string)

	return entities.ValidationResult{
		Status:   entities.AuthGranted,
		Scopes:   []entities.AuthScope{entities.ScopeRead, entities.ScopeWrite}, // TODO: pending add to jwt
		Username: username,
		Email:    email,
	}
}
