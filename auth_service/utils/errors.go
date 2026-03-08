package utils

import (
	"errors"
	"orderly/auth-service/internal/models"

	"github.com/gin-gonic/gin"
)

var ErrUsernameExists = errors.New("username already exists")

func RespondError(c *gin.Context, status int, code, message string) {
	c.AbortWithStatusJSON(status, models.ErrorResponse{
		Code:    code,
		Message: message,
	})
}