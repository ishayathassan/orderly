package handlers

import (
	"fmt"
	"net/http"
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateOrder(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 32) // ParseUint always returns uint64, bit size=32 for valid uint32 checking
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid order id format"})
		return
	}
	orderID := uint(id64)

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