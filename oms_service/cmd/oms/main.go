// @title Orderly OMS API
// @version 1.0
// @description This is the Order Management Service for Orderly microservices.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/contact
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
package main

import (
	"orderly/oms-service/internal/database"
	"orderly/oms-service/internal/handlers"
	"orderly/oms-service/internal/metrics"
	"orderly/oms-service/internal/middlewares"

	_ "orderly/oms-service/docs"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	database.InitDB()

	metrics.RegisterMetrics()

	router := gin.Default()



	router.Use(middlewares.MetricsMiddleware())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// OMS Routes
	router.GET("/orders", handlers.GetOrders)
	router.POST("/orders", handlers.CreateOrder)
	router.GET("/orders/:id", middlewares.ValidateID(), handlers.GetOrderByID)
	router.DELETE("/orders/:id", middlewares.ValidateID(), handlers.DeleteOrder)
	router.PUT("/orders/:id",middlewares.ValidateID(), handlers.UpdateOrder)

	router.Run("localhost:8080")
}
