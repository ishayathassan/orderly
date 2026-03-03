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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Order ID"})
		return
	}
	orderID := uint(id64)
	for idx, order := range repositories.Orders {
		if order.ID == orderID {
			repositories.Orders = append(repositories.Orders[:idx],repositories.Orders[idx+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Order ID"})
}