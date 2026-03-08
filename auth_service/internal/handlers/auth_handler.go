package handlers

import (
	"net/http"
	"orderly/auth-service/internal/auth"
	"orderly/auth-service/internal/database"
	"orderly/auth-service/internal/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req models.LoginRequest
	var user models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
		})
		return
	}

	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username",
		})
		return
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	token, err := auth.GenerateToken(user.ID.String(), user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}