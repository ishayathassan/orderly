package repositories

import (
	"orderly/oms-service/internal/database"
	"orderly/oms-service/internal/metrics"
	"orderly/oms-service/internal/models"
	"time"
)

func GetAll() ([]models.Order, error) {

	start := time.Now()

	var orders []models.Order
	err := database.DB.Find(&orders).Error

	duration := time.Since(start).Seconds()

	metrics.DatabaseQueryDuration.
		WithLabelValues("SELECT").
		Observe(duration)

	return orders, err
}