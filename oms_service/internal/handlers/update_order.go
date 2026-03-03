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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Order ID format"})
		return
	}
	orderID := uint(id64)

	var updatedOrder models.Order
	if err := c.BindJSON(&updatedOrder); err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Order format"})
		return
	}
	
	for idx, order := range repositories.Orders {
		if order.ID == orderID {
			repositories.Orders[idx] = updatedOrder

			repositories.Orders[idx].ID = orderID
			repositories.Orders[idx].CreatedAt = order.CreatedAt
			repositories.Orders[idx].UpdatedAt = time.Now()

			c.JSON(http.StatusOK, order)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Order not found"})

}