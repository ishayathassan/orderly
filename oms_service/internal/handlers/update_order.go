package handlers

import (
	"net/http"
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/services"
	"orderly/oms-service/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Update an order
// @Description Update an existing order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path uint true "Order ID"
// @Param order body models.Order true "Updated order"
// @Success 200 {object} models.Order
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /orders/{id} [put]
func UpdateOrder(c *gin.Context) {
	orderID := c.MustGet("orderID").(uint)

	var updatedOrder models.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	order, err := services.UpdateOrder(orderID, updatedOrder)
	if err != nil {
		if err == services.ErrOrderNotFound {
			utils.RespondError(c, http.StatusNotFound, "ORDER_NOT_FOUND", "Order not found")
			return
		}
		utils.RespondError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Could not update order")
		return
	}

	c.JSON(http.StatusOK, order)
}