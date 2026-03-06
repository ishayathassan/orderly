package main

import (
	"orderly/oms-service/internal/database"
	"orderly/oms-service/internal/handlers"
	"orderly/oms-service/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()

	router := gin.Default()
	router.GET("/orders", handlers.GetOrders)
	router.POST("/orders", handlers.CreateOrder)
	router.GET("/orders/:id", middlewares.ValidateID(), handlers.GetOrderByID)
	router.DELETE("/orders/:id", middlewares.ValidateID(), handlers.DeleteOrder)
	router.PUT("/orders/:id",middlewares.ValidateID(), handlers.UpdateOrder)

	router.Run("localhost:8080")
}
