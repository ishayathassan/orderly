package handlers

import (
	"net/http"
	"orderly/oms-service/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteOrder(c *gin.Context){
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 32) // ParseUint always returns uint64, bit size=32 for valid uint32 checking
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Order ID",
		})
		return
	}
	orderID := uint(id64)

	if _,err := repositories.GetByID(orderID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}

	if err := repositories.Delete(orderID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete order",
		})
		return
	}

	c.Status(http.StatusNoContent)

}