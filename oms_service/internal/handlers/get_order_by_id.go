package handlers

import (
	"net/http"
	"orderly/oms-service/internal/services"
	"orderly/oms-service/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get order by ID
// @Description Retrieve a single order by its ID
// @Tags orders
// @Produce json
// @Param id path uint true "Order ID"
// @Success 200 {object} models.Order
// @Failure 404 {object} utils.ErrorResponse
// @Router /orders/{id} [get]
func GetOrderByID(c *gin.Context) {
	orderID := c.MustGet("orderID").(uint)

	order, err := services.GetOrderByID(orderID)
	if err != nil {
		if err == utils.ErrOrderNotFound {
			utils.RespondError(c, http.StatusNotFound, "ORDER_NOT_FOUND", "Order not found")
			return
		}
		utils.RespondError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to retrieve order")
		return
	}

	c.JSON(http.StatusOK, order)
}