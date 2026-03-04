package main

import (
	"orderly/oms-service/internal/database"
	"orderly/oms-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()

	router := gin.Default()
	router.GET("/orders", handlers.GetOrders)
	router.POST("/orders", handlers.CreateOrder)
	router.GET("/orders/:id", handlers.ValidateID(), handlers.GetOrderByID)
	router.DELETE("/orders/:id", handlers.ValidateID(), handlers.DeleteOrder)
	router.PUT("/orders/:id",handlers.ValidateID(), handlers.UpdateOrder)


	router.Run("localhost:8080")
}

func initDB() {
	panic("unimplemented")
}