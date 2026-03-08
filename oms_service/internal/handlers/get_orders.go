package handlers

import (
	"net/http"
	"orderly/oms-service/internal/services"
	"orderly/oms-service/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get all orders
// @Description Retrieve all orders
// @Tags orders
// @Produce json
// @Success 200 {array} models.Order "Example: [{\"id\":1,\"user_id\":\"123\",\"item_name\":\"Laptop\",\"quantity\":2,\"amount\":2500.50,\"status\":\"pending\",\"created_at\":\"2026-03-08T12:00:00Z\",\"updated_at\":\"2026-03-08T12:00:00Z\"}]"
// @Failure 500 {object} utils.ErrorResponse "Failed to retrieve orders"
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	orders, err := services.GetAllOrders()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to retrieve orders")
		return
	}

	c.JSON(http.StatusOK, orders)
}