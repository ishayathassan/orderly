package handlers

import (
	"net/http"
	"orderly/oms-service/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repositories.Orders)
}