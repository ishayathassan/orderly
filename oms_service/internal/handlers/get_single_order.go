package handlers

import (
	"net/http"
	"orderly/oms-service/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 32) // ParseUint always returns uint64, bit size=32 for valid uint32 checking
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid order id",
		})
		return
	}
	orderID := uint(id64)

	order, err := repositories.GetByID(orderID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order id not found",
		})
		return
	}
	c.JSON(http.StatusOK, order)
}