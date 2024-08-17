package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/proyectum/ms-auth/internal/adapters/in/http/api"
	appUseCases "github.com/proyectum/ms-auth/internal/app/usecases"
	domainErrors "github.com/proyectum/ms-auth/internal/domain/errors"
	"github.com/proyectum/ms-auth/internal/domain/usecases"
	"net/http"
)

type authRoutes struct {
	signUpUseCase usecases.SignUpUseCase
	signInUseCase usecases.SignInUseCase
}

func RegisterRoutes(r *gin.Engine) {
	api.RegisterHandlers(r, &authRoutes{
		signUpUseCase: appUseCases.NewSignUpUseCase(),
		signInUseCase: appUseCases.NewSignInUseCase(),
	})
}

func (ar *authRoutes) SignUp(c *gin.Context) {
	var request api.SignUpRequest
	if err := c.Bind(&request); err != nil {
		handleBadRequestError(err, c)
		return
	}

	err := ar.signUpUseCase.SignUp(request.Username, request.Password, string(request.Email))

	if err != nil {
		ar.handleError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (ar *authRoutes) SignIn(c *gin.Context) {
	var request api.SignInRequest
	if err := c.Bind(&request); err != nil {
		handleBadRequestError(err, c)
		return
	}

	signIn, err := ar.signInUseCase.SignIn(request.Username, request.Password)

	if err != nil {
		ar.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, api.SignInResponse{
		Token: &signIn.Token,
	})
}

func (ar *authRoutes) handleError(c *gin.Context, err error) {
	if errors.Is(err, &domainErrors.InvalidUsernameError{}) ||
		errors.Is(err, &domainErrors.InvalidPasswordError{}) ||
		errors.Is(err, &domainErrors.InvalidEmailError{}) {
		handleBadRequestError(err, c)
		return
	}

	if errors.Is(err, &domainErrors.UserAlreadyExistsError{}) {
		handleConflictError(err, c)
		return
	}

	if errors.Is(err, &domainErrors.InvalidCredentialsError{}) {
		handleUnauthorizedError(err, c)
		return
	}

	handleInternalError(err, c)
}

func handleUnauthorizedError(err error, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusText(http.StatusUnauthorized),
		"message": err.Error(),
	})
}

func handleConflictError(err error, c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{
		"code":    http.StatusText(http.StatusConflict),
		"message": err.Error(),
	})
}

func handleInternalError(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusText(http.StatusInternalServerError),
		"message": err.Error(),
	})
}

func handleBadRequestError(err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusText(http.StatusBadRequest),
		"message": err.Error(),
	})
}
