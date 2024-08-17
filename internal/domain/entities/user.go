package entities

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	Email    string
}

type SignInUser struct {
	Token string
}
