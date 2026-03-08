// @title Orderly Auth Service API
// @version 1.0
// @description This is the auth service for Orderly microservices.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/contact
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8081
// @BasePath /
package main

import (
	_ "orderly/auth-service/docs" // import generated docs
	"orderly/auth-service/internal/database"
	"orderly/auth-service/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.InitDB()
	
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)


	router.Run("localhost:8081")
}