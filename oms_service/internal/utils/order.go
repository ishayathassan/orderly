package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)


var ErrOrderNotFound = errors.New("order not found")


type ErrorResponse struct {
	Code    string `json:"code,omitempty" example:"INVALID_REQUEST"`
	Message string `json:"message" example:"Something went wrong"`
}

func RespondError(c *gin.Context, status int, code, message string) {
	c.AbortWithStatusJSON(status, ErrorResponse{
		Code:    code,
		Message: message,
	})
}