package handlers

import (
	"net/http"
	"orderly/oms-service/internal/services"
	"orderly/oms-service/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Delete an order
// @Description Delete an order by ID
// @Tags orders
// @Param id path uint true "Order ID"
// @Success 204 "No Content"
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /orders/{id} [delete]
func DeleteOrder(c *gin.Context) {
	orderID := c.MustGet("orderID").(uint)

	err := services.DeleteOrder(orderID)
	if err != nil {
		if err == utils.ErrOrderNotFound {
			utils.RespondError(c, http.StatusNotFound, "ORDER_NOT_FOUND", "Order not found")
			return
		}
		utils.RespondError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to delete order")
		return
	}

	c.Status(http.StatusNoContent)
}