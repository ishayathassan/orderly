package main

import (
	"orderly/oms-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/orders", handlers.GetOrders)
	router.POST("/orders", handlers.CreateOrder)
	router.GET("/orders/:id", handlers.GetOrderByID)

	router.Run("localhost:8080")
}