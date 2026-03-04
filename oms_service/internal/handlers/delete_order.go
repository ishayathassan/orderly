package handlers

import (
	"net/http"
	"orderly/oms-service/internal/repositories"

	"github.com/gin-gonic/gin"
)

func DeleteOrder(c *gin.Context){
	orderID := c.MustGet("orderID").(uint)

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