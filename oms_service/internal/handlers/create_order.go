package handlers

import (
	"fmt"
	"net/http"
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var newOrder models.Order
	if err := c.BindJSON(&newOrder); err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Order format"})
		return
	}
	repositories.OrderCount += 1
	newOrder.ID = repositories.OrderCount
	newOrder.Status = "pending"
	newOrder.CreatedAt = time.Now()
	repositories.Orders = append(repositories.Orders, newOrder)

	c.IndentedJSON(http.StatusCreated, newOrder)
}