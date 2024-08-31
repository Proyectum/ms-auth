package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/proyectum/ms-auth/internal/adapters/in/http/api"
	appUseCases "github.com/proyectum/ms-auth/internal/app/usecases"
	"github.com/proyectum/ms-auth/internal/domain/entities"
	domainErrors "github.com/proyectum/ms-auth/internal/domain/errors"
	"github.com/proyectum/ms-auth/internal/domain/usecases"
	"net/http"
	"strings"
)

const authHeader = "Authorization"
const bearerPrefix = "Bearer"
const space = " "

type authRoutes struct {
	signUpUseCase     usecases.SignUpUseCase
	signInUseCase     usecases.SignInUseCase
	validationUseCase usecases.ValidationUseCase
}

func RegisterRoutes(r *gin.Engine) {
	api.RegisterHandlers(r, &authRoutes{
		signUpUseCase:     appUseCases.NewSignUpUseCase(),
		signInUseCase:     appUseCases.NewSignInUseCase(),
		validationUseCase: appUseCases.NewValidationUseCase(),
	})
}

func (r *authRoutes) Validation(c *gin.Context) {
	token := r.getToken(c)

	if c.IsAborted() {
		return
	}

	validation := r.validationUseCase.Validate(token)

	if validation.Status != entities.AuthGranted {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.Header("X-Auth-User", validation.Username)
	c.Header("X-Auth-Email", validation.Email)
	c.Header("X-Auth-Scopes", validation.JoinScopes())
	c.Status(http.StatusOK)
}

func (r *authRoutes) SignUp(c *gin.Context) {
	var request api.SignUpRequest
	if err := c.Bind(&request); err != nil {
		handleBadRequestError(err, c)
		return
	}

	err := r.signUpUseCase.SignUp(request.Username, request.Password, string(request.Email))

	if err != nil {
		r.handleError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (r *authRoutes) SignIn(c *gin.Context) {
	var request api.SignInRequest
	if err := c.Bind(&request); err != nil {
		handleBadRequestError(err, c)
		return
	}

	signIn, err := r.signInUseCase.SignIn(request.Username, request.Password)

	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, api.SignInResponse{
		Token: &signIn.Token,
	})
}

func (r *authRoutes) handleError(c *gin.Context, err error) {
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

func (r *authRoutes) getToken(c *gin.Context) string {
	header := c.GetHeader(authHeader)

	if len(header) == 0 {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return ""
	}

	parts := strings.Split(header, space)
	if len(parts) != 2 || parts[0] != bearerPrefix {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return ""
	}

	return parts[1]
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
