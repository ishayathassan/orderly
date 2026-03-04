package handlers

import (
	"net/http"
	"orderly/oms-service/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetOrderByID(c *gin.Context) {
	orderID := c.MustGet("orderID").(uint)

	order, err := repositories.GetByID(orderID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order id not found",
		})
		return
	}
	c.JSON(http.StatusOK, order)
}