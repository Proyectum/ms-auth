package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/proyectum/ms-auth/internal/domain/entities"
	"github.com/spf13/viper"
	"time"
)

type JWTService interface {
	GetToken(user *entities.User) (string, error)
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type jwtService struct {
}

func (j *jwtService) GetToken(user *entities.User) (string, error) {
	jwtSecret := viper.GetString("security.jwt.secret")
	expirationConf := viper.GetDuration("security.jwt.expiration")
	expirationTime := time.Now().Add(expirationConf * time.Hour)

	claims := &Claims{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewJWTService() JWTService {
	return &jwtService{}
}
