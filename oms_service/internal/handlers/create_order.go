package handlers

import (
	"net/http"
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/services"
	"orderly/oms-service/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new order
// @Description Create an order with item details
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order details"
// @Success 201 {object} models.Order
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
    var newOrder models.Order

    if err := c.ShouldBindJSON(&newOrder); err != nil {
        utils.RespondError(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
        return
    }

    createdOrder, err := services.CreateOrder(newOrder)
    if err != nil {
        utils.RespondError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Could not save order")
        return
    }

    c.JSON(http.StatusCreated, createdOrder)
}