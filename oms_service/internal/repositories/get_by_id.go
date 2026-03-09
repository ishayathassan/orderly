package repositories

import (
	"orderly/oms-service/internal/database"
	"orderly/oms-service/internal/metrics"
	"orderly/oms-service/internal/models"
	"time"
)

func GetByID(id uint) (models.Order, error) {
	start := time.Now()

	var order models.Order
	err := database.DB.First(&order, id).Error

	duration := time.Since(start).Seconds()

	metrics.DatabaseQueryDuration.
		WithLabelValues("SELECT").
		Observe(duration)

	return order, err
}