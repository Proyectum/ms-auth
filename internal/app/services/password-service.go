package services

import (
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type HashPasswordService interface {
	Hash(password string) (string, error)
	Check(raw, hash string) bool
}

type hashPasswordService struct{}

func (h *hashPasswordService) Hash(password string) (string, error) {
	cost := viper.GetInt("security.password.cost")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func (h *hashPasswordService) Check(raw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	return err == nil
}

func NewHashPasswordService() HashPasswordService {
	return &hashPasswordService{}
}
