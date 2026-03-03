package handlers

import (
	"fmt"
	"net/http"
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/repositories"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var newOrder models.Order

	if err := c.BindJSON(&newOrder); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid order format",
		})
		return
	}

	if err := repositories.Create(&newOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save order",
		})
		return
	}

	c.JSON(http.StatusCreated, newOrder)
}