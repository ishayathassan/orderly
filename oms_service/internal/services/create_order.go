package services

import (
	"orderly/oms-service/internal/metrics"
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/repositories"
	"time"
)

func CreateOrder(order models.Order) (*models.Order, error) {
	now := time.Now()
	order.CreatedAt = now
	order.UpdatedAt = now

	if err := repositories.Create(&order); err != nil {
		return &models.Order{}, err
	}

	metrics.OrdersCreatedTotal.Inc()
	return &order, nil
}