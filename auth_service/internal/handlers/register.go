package handlers

import (
	"net/http"
	"orderly/auth-service/internal/auth"
	"orderly/auth-service/internal/database"
	"orderly/auth-service/internal/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
		})
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not hash password",
		})
	}

	newUser := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Role: "user",
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Username already exists",
		})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}