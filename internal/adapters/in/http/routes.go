package http

import (
	"github.com/gin-gonic/gin"
	"github.com/proyectum/ms-auth/internal/adapters/in/http/api"
	"net/http"
)

type authRoutes struct{}

func RegisterRoutes(r *gin.Engine) {
	api.RegisterHandlers(r, &authRoutes{})
}

func (ar *authRoutes) SignUp(c *gin.Context) {
	var request api.SignUpRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusText(http.StatusBadRequest),
			"message": "Invalid input payload",
		})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (ar *authRoutes) SignIn(c *gin.Context) {
	var request api.SignInRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusText(http.StatusBadRequest),
			"message": "Invalid input payload",
		})
		return
	}
	token := "jwt-token"
	c.JSON(http.StatusOK, api.SignInResponse{
		Token: &token,
	})
}
