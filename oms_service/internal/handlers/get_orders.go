package handlers

import (
	"net/http"
	"orderly/oms-service/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {

	orders, err := repositories.GetAll()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve orders from database",
		})
	}

	c.JSON(http.StatusOK, orders)
}