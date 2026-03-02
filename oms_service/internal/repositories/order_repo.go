package repositories

import (
	"errors"
	"orderly/oms-service/internal/models"
	"time"
)

var Orders = []models.Order{
	{
		ID:        1,
		UserID:    "user_99", // Context typically injected by the API Gateway
		ItemName:  "Kubernetes in Action",
		Quantity:  2,
		Amount:    90.00,
		Status:    "completed",
		CreatedAt: time.Now().Add(-24 * time.Hour),
	},
	{
		ID:        2,
		UserID:    "user_42",
		ItemName:  "Prometheus Guide",
		Quantity:  1,
		Amount:    45.50,
		Status:    "pending",
		CreatedAt: time.Now().Add(-2 * time.Hour),
	},
	{
		ID:        3,
		UserID:    "user_101",
		ItemName:  "Mechanical Keyboard",
		Quantity:  1,
		Amount:    150.00,
		Status:    "processing",
		CreatedAt: time.Now(),
	},
}

func SearchOrder(id uint) (models.Order, error) {
	for _, order := range Orders {
		if order.ID == id {
			return order, nil
		}
	}
	return models.Order{}, errors.New("order not found")
}
