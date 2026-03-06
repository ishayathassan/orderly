package main

import (
	"orderly/auth-service/internal/database"
	"orderly/auth-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	
	router := gin.Default()

	router.POST("/register", handlers.Register)
	router.GET("/login", handlers.Login)


	router.Run("localhost:8081")
}