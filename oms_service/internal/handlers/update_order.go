package handlers

import (
	"fmt"
	"net/http"
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateOrder(c *gin.Context) {
	orderID := c.MustGet("orderID").(uint)

	var updatedOrder models.Order
	if err := c.BindJSON(&updatedOrder); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Order format",
		})
		return
	}
	prevOrder, err := repositories.GetByID(orderID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
	}

	updatedOrder.ID = orderID
	updatedOrder.CreatedAt = prevOrder.CreatedAt
	updatedOrder.UpdatedAt = time.Now() 

	if err := repositories.Update(updatedOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update order",
		})
	}

	c.JSON(http.StatusOK, updatedOrder)
}