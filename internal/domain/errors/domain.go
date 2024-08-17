package errors

import "errors"

type InvalidCredentialsError struct {
	message string
}

func (i InvalidCredentialsError) Error() string {
	return i.message
}

func (i InvalidCredentialsError) Is(target error) bool {
	var err *InvalidCredentialsError
	ok := errors.As(target, &err)
	return ok
}

func NewInvalidCredentialsError(message string) error {
	return &InvalidCredentialsError{
		message: message,
	}
}

type InvalidUsernameError struct {
	message string
}

func (i InvalidUsernameError) Error() string {
	return i.message
}

func (i InvalidUsernameError) Is(target error) bool {
	var err *InvalidUsernameError
	ok := errors.As(target, &err)
	return ok
}

func NewInvalidUsernameError(message string) error {
	return &InvalidUsernameError{
		message: message,
	}
}

type UserAlreadyExistsError struct {
	message string
}

func (i UserAlreadyExistsError) Error() string {
	return i.message
}

func (i UserAlreadyExistsError) Is(target error) bool {
	var err *UserAlreadyExistsError
	ok := errors.As(target, &err)
	return ok
}

func NewUserAlreadyExistsError(message string) error {
	return &UserAlreadyExistsError{
		message: message,
	}
}

type UserNotFoundError struct {
	message string
}

func (i UserNotFoundError) Error() string {
	return i.message
}

func (i UserNotFoundError) Is(target error) bool {
	var err *UserNotFoundError
	ok := errors.As(target, &err)
	return ok
}

func NewUserNotFoundError(message string) error {
	return &UserNotFoundError{
		message: message,
	}
}

type InvalidPasswordError struct {
	message string
}

func (i InvalidPasswordError) Error() string {
	return i.message
}

func (i InvalidPasswordError) Is(target error) bool {
	var err *InvalidPasswordError
	ok := errors.As(target, &err)
	return ok
}

func NewInvalidPasswordError(message string) error {
	return &InvalidUsernameError{
		message: message,
	}
}

type InvalidEmailError struct {
	message string
}

func (i InvalidEmailError) Error() string {
	return i.message
}

func (i InvalidEmailError) Is(target error) bool {
	var err *InvalidEmailError
	ok := errors.As(target, &err)
	return ok
}

func NewInvalidEmailError(message string) error {
	return &InvalidUsernameError{
		message: message,
	}
}
